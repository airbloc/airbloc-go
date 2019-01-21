package warehouse

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/schemas"
	"math/rand"
	"net/url"
	"time"

	"github.com/azer/logger"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

var (
	// errValidationFailed is returned when an incoming data
	// failed to met condition of the one of the validators (DAuth, Schema, ...)
	errValidationFailed = errors.New("data validation failed.")
)

type DataWarehouse struct {
	kms        key.Manager
	localCache *localdb.Model

	// for data registration
	metaDatabase *metadb.Model
	ethclient    blockchain.TxClient
	dataRegistry *adapter.DataRegistry
	schemas      *schemas.Schemas
	collections  *collections.Collections

	// data storage layer
	protocols      map[string]protocol.Protocol
	DefaultStorage storage.Storage

	// data validators
	dauthValidator *dauth.Validator

	log *logger.Logger
}

func New(
	kms key.Manager,
	localDatabase localdb.Database,
	metaDatabase metadb.Database,
	ethclient blockchain.TxClient,
	defaultStorage storage.Storage,
	supportedProtocols []protocol.Protocol,
) *DataWarehouse {
	protocols := map[string]protocol.Protocol{}
	for _, protoc := range supportedProtocols {
		protocols[protoc.Name()] = protoc
	}

	dauthManager := dauth.NewManager(ethclient)
	dauthValidator := dauth.NewValidator(dauthManager)

	contract := ethclient.GetContract(&adapter.DataRegistry{})
	return &DataWarehouse{
		kms:        kms,
		localCache: localdb.NewModel(localDatabase, "bundle"),

		metaDatabase: metadb.NewModel(metaDatabase, "bundles"),
		ethclient:    ethclient,
		dataRegistry: contract.(*adapter.DataRegistry),
		collections:  collections.New(ethclient),
		schemas:      schemas.New(metaDatabase, ethclient),

		protocols:      protocols,
		DefaultStorage: defaultStorage,
		dauthValidator: dauthValidator,

		log: logger.New("warehouse"),
	}
}

func (warehouse *DataWarehouse) CreateBundle(collectionId common.ID) (*BundleStream, error) {
	collection, err := warehouse.collections.Get(collectionId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve a collection")
	}
	schema, err := warehouse.schemas.Get(collection.Schema.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve a schema")
	}
	collection.Schema = *schema
	return newBundleStream(warehouse, collection.AppId, collection), nil
}

func (warehouse *DataWarehouse) validate(collection *collections.Collection, data *common.Data) error {
	if !warehouse.dauthValidator.IsCollectible(collection.Id, data) {
		return errors.Wrap(errValidationFailed, "user hasn't been authorized the data collection")
	}

	isValidFormat, err := collection.Schema.IsValidFormat(data)
	if err != nil {
		return err
	} else if !isValidFormat {
		return errors.Wrap(errValidationFailed, "wrong format")
	}
	return nil
}

func (warehouse *DataWarehouse) encrypt(d *common.Data) (*common.EncryptedData, error) {
	encryptedPayload, err := warehouse.kms.Encrypt(d.Payload)
	if err != nil {
		return nil, err
	}
	return &common.EncryptedData{
		OwnerAnID: d.OwnerAnID,
		Payload:   encryptedPayload,
		Capsule:   nil,
	}, nil
}

func generateBundleNameOf(bundle *data.Bundle) string {
	tokenBytes := make([]byte, 4)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	currentTime := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s.bundle", currentTime, bundle.Collection.Hex(), token)
}

func (warehouse *DataWarehouse) Store(stream *BundleStream) (*data.Bundle, error) {
	if stream == nil {
		return nil, errors.New("No data in the stream.")
	}
	ingestedAt := time.Now()

	createdBundle := &data.Bundle{
		Provider:   stream.provider,
		Collection: stream.collection.Id,
		DataCount:  stream.DataCount,
		IngestedAt: ingestedAt,
		Data:       stream.data,
	}
	bundleName := generateBundleNameOf(createdBundle)
	uri, err := warehouse.DefaultStorage.Save(bundleName, createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save bundle to the storage")
	}
	createdBundle.Uri = uri.String()

	// register to on-chainㄴ
	bundleId, err := warehouse.registerBundleOnChain(createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register bundle to blockchain")
	}
	empty := common.ID{}
	createdBundle.Id = fmt.Sprintf("%s%s", empty.Hex(), bundleId.Hex())

	// save metadata to make the bundle searchable
	bundleInfo := map[string]interface{}{
		"bundleId":   createdBundle.Id,
		"uri":        createdBundle.Uri,
		"provider":   createdBundle.Provider.Hex(),
		"collection": createdBundle.Collection.Hex(),
		"dataCount":  createdBundle.DataCount,
		"ingestedAt": ingestedAt,
	}
	_, err = warehouse.metaDatabase.Create(bundleInfo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save metadata")
	}
	warehouse.log.Info("Bundle %s registered on", bundleName, logger.Attrs{
		"id":    bundleId.Hex(),
		"count": bundleInfo["dataCount"],
	})
	return createdBundle, nil
}

func (warehouse *DataWarehouse) registerBundleOnChain(bundle *data.Bundle) (common.ID, error) {
	bundleDataHash, err := bundle.Hash()
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to get hash of the bundle data")
	}

	userMerkleRoot, err := bundle.SetupUserProof()
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to setup SMT")
	}

	warehouse.log.Info("Bundle data", logger.Attrs{
		"hash":       bundleDataHash.Hex(),
		"merkleRoot": userMerkleRoot.Hex(),
	})

	tx, err := warehouse.dataRegistry.RegisterBundle(
		warehouse.ethclient.Account(),
		bundle.Collection,
		userMerkleRoot,
		bundleDataHash,
		bundle.Uri)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to register a bundle to DataRegistry")
	}

	receipt, err := warehouse.ethclient.WaitMined(context.Background(), tx)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	registerResult, err := warehouse.dataRegistry.ParseBundleRegisteredFromReceipt(receipt)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	return common.ID(registerResult.BundleId), nil
}

func (warehouse *DataWarehouse) Get(id *common.DataID) (*data.Bundle, error) {
	bundle, err := warehouse.dataRegistry.Bundles(nil, id.Empty, id.BundleID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get uri")
	}
	uri, err := url.Parse(bundle.Uri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bundle data")
	}
	return warehouse.Fetch(uri)
}

func (warehouse *DataWarehouse) Fetch(uri *url.URL) (*data.Bundle, error) {
	protoc, exists := warehouse.protocols[uri.Scheme]
	if !exists {
		return nil, errors.Errorf("the protocol %s is not supported", uri.Scheme)
	}
	return protoc.Read(uri)
}
