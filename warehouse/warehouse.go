package warehouse

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/schemas"
	"github.com/mongodb/mongo-go-driver/bson"
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

	config Config
	log    *logger.Logger
}

func New(
	kms key.Manager,
	localDatabase localdb.Database,
	metaDatabase metadb.Database,
	ethclient blockchain.TxClient,
	defaultStorage storage.Storage,
	supportedProtocols []protocol.Protocol,
	config Config,
) *DataWarehouse {
	protocols := map[string]protocol.Protocol{}
	for _, protoc := range supportedProtocols {
		protocols[protoc.Name()] = protoc
	}

	log := logger.New("warehouse")
	if config.Debug.DisableSchemaValidation {
		log.Error("warning: You have disabled schema validation. " +
			"It is recommended to avoid disabling the validation on production mode.")
	}
	if config.Debug.DisableUserAuthValidation {
		log.Error("warning: You have disabled user auth validation. \n" + "\033[31m" +
			"DO NOT DISABLE THE USER VALIDATION ON PRODUCTION MODE, " +
			"BECAUSE IT CAN CAUSE A FINANCIAL LOSS OF YOUR STAKED COLLETRALS. " + "\033[0m")
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

		config: config,
		log:    log,
	}
}

func (dw *DataWarehouse) CreateBundle(collectionId common.ID) (*BundleStream, error) {
	collection, err := dw.collections.Get(collectionId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve a collection")
	}
	schema, err := dw.schemas.Get(collection.Schema.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve a schema")
	}
	collection.Schema = *schema
	return newBundleStream(dw, collection.AppId, collection), nil
}

func (dw *DataWarehouse) validate(collection *collections.Collection, data *common.Data) error {
	if !dw.dauthValidator.IsCollectible(collection.Id, data) &&
		!dw.config.Debug.DisableUserAuthValidation {
		return errors.Wrap(errValidationFailed, "user hasn't been authorized the data collection")
	}

	if !dw.config.Debug.DisableSchemaValidation {
		isValidFormat, err := collection.Schema.IsValidFormat(data)
		if err != nil {
			return err
		} else if !isValidFormat {
			return errors.Wrap(errValidationFailed, "wrong format")
		}
	}
	return nil
}

func generateBundleNameOf(bundle *data.Bundle) string {
	tokenBytes := make([]byte, 4)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	currentTime := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s.bundle", currentTime, bundle.Collection.Hex(), token)
}

func (dw *DataWarehouse) Store(stream *BundleStream) (*data.Bundle, error) {
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
	uri, err := dw.DefaultStorage.Save(bundleName, createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save bundle to the storage")
	}
	createdBundle.Uri = uri.String()

	// register to on-chain
	bundleId, err := dw.registerBundleOnChain(createdBundle)
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
	_, err = dw.metaDatabase.Create(bundleInfo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save metadata")
	}
	dw.log.Info("Bundle %s registered on", bundleName, logger.Attrs{
		"id":    bundleId.Hex(),
		"count": bundleInfo["dataCount"],
	})
	return createdBundle, nil
}

func (dw *DataWarehouse) registerBundleOnChain(bundle *data.Bundle) (common.ID, error) {
	bundleDataHash, err := bundle.Hash()
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to get hash of the bundle data")
	}

	userMerkleRoot, err := bundle.SetupUserProof()
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to setup SMT")
	}

	dw.log.Info("Bundle data", logger.Attrs{
		"hash":       bundleDataHash.Hex(),
		"merkleRoot": userMerkleRoot.Hex(),
	})

	tx, err := dw.dataRegistry.RegisterBundle(
		dw.ethclient.Account(),
		bundle.Collection,
		userMerkleRoot,
		bundleDataHash,
		bundle.Uri)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to register a bundle to DataRegistry")
	}

	receipt, err := dw.ethclient.WaitMined(context.Background(), tx)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	registerResult, err := dw.dataRegistry.ParseBundleRegisteredFromReceipt(receipt)
	if err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	return common.ID(registerResult.BundleId), nil
}

func (dw *DataWarehouse) Get(id *common.DataID) (*data.Bundle, error) {
	bundle, err := dw.dataRegistry.Bundles(nil, id.Padding, id.BundleID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get uri")
	}
	uri, err := url.Parse(bundle.Uri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bundle data")
	}
	return dw.Fetch(uri)
}

func (dw *DataWarehouse) Fetch(uri *url.URL) (*data.Bundle, error) {
	protoc, exists := dw.protocols[uri.Scheme]
	if !exists {
		return nil, errors.Errorf("the protocol %s is not supported", uri.Scheme)
	}
	return protoc.Read(uri)
}

func (dw *DataWarehouse) List(providerId common.ID) ([]*data.Bundle, error) {
	bundleDataList, err := dw.metaDatabase.RetrieveMany(context.TODO(), bson.M{"provider": providerId.Hex()})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list bundles")
	}

	var bundles []*data.Bundle
	for _, bundleData := range bundleDataList {
		collectionId, _ := common.HexToID(bundleData["collection"].(string))
		ingestedAt, _ := time.Parse(time.RFC3339Nano, bundleData["ingestedAt"].(string))
		bundles = append(bundles, &data.Bundle{
			Id:         bundleData["bundleId"].(string),
			Uri:        bundleData["uri"].(string),
			DataCount:  bundleData["dataCount"].(int),
			IngestedAt: ingestedAt,
			Provider:   providerId,
			Collection: collectionId,
		})
	}
	return bundles, nil
}
