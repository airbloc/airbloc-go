package warehouse

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mongodb/mongo-go-driver/bson"
	"math/rand"
	"net/url"
	"time"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

type DataWarehouse struct {
	kms            *key.Manager
	protocols      map[string]protocol.Protocol
	localCache     *localdb.Model
	metaDatabase   *metadb.Model
	ethclient      *blockchain.Client
	dataRegistry   *adapter.DataRegistry
	DefaultStorage storage.Storage
}

func New(
	kms *key.Manager,
	localDatabase localdb.Database,
	metaDatabase metadb.Database,
	ethclient *blockchain.Client,
	defaultStorage storage.Storage,
	supportedProtocols []protocol.Protocol,
) *DataWarehouse {

	protocols := map[string]protocol.Protocol{}
	for _, protoc := range supportedProtocols {
		protocols[protoc.Name()] = protoc
	}

	return &DataWarehouse{
		kms:            kms,
		protocols:      protocols,
		localCache:     localdb.NewModel(localDatabase, "bundle"),
		metaDatabase:   metadb.NewModel(metaDatabase, "bundles"),
		ethclient:      ethclient,
		dataRegistry:   ethclient.Contracts.DataRegistry,
		DefaultStorage: defaultStorage,
	}
}

func (warehouse *DataWarehouse) CreateBundle(collection common.ID) *BundleStream {
	return newBundleStream(warehouse, collection)
}

func (warehouse *DataWarehouse) validate(collection common.ID, data *common.Data) error {
	return nil
}

func (warehouse *DataWarehouse) encrypt(d *common.Data) (*common.EncryptedData, error) {
	encryptedPayload, err := warehouse.kms.Encrypt(d.Payload)
	if err != nil {
		return nil, err
	}
	return &common.EncryptedData{
		OwnerAnid: d.OwnerAnid,
		Payload:   encryptedPayload,
		Capsule:   nil,
	}, nil
}

func generateBundleNameOf(bundle *data.Bundle) string {
	tokenBytes := make([]byte, 4)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	currentTime := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s.bundle", currentTime, bundle.Collection.String(), token)
}

func (warehouse *DataWarehouse) Store(stream *BundleStream) (*data.Bundle, error) {
	if stream == nil {
		return nil, errors.New("No data in the stream.")
	}
	ingestedAt := time.Now()

	createdBundle := &data.Bundle{
		Provider:   common.ID{}, /* TODO: implement appID */
		Collection: stream.collection,
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

	// register to on-chain
	bundleIndex, err := warehouse.registerBundleOnChain(createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register bundle to blockchain")
	}

	// save metadata to make the bundle searchable
	bundleInfo := map[string]interface{}{
		"bundleId":   fmt.Sprintf("%s/%d", createdBundle.Collection, bundleIndex),
		"uri":        createdBundle.Uri,
		"provider":   createdBundle.Provider.String(),
		"collection": createdBundle.Collection.String(),
		"dataCount":  createdBundle.DataCount,
		"ingestedAt": ingestedAt,
	}
	txn, err := warehouse.metaDatabase.Create(bundleInfo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save metadata")
	}
	log.Debug("Metadata Stored", "transactionId", txn.ID)

	return createdBundle, nil
}

func (warehouse *DataWarehouse) registerBundleOnChain(createdBundle *data.Bundle) (int, error) {
	bundleDataHash, err := createdBundle.Hash()
	if bundleDataHash, err != nil {
		return 0, errors.Wrap(err, "failed to get hash of the bundle data")
	}

	tx, err := warehouse.dataRegistry.RegisterBundle(warehouse.ethclient.Account(),
		createdBundle.Collection,
		[32]byte{'T', 'O', 'D', 'O'},
		bundleDataHash,
		createdBundle.Uri)
	if err != nil {
		return 0, err
	}

	receipt, err := warehouse.ethclient.WaitMined(context.Background(), tx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	registerResult := adapter.DataRegistryBundleRegistered{}
	err = warehouse.ethclient.GetEventFromReceipt("DataRegistry", "BundleRegistered", &registerResult, receipt)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	return int(registerResult.Index), nil
}

func (warehouse *DataWarehouse) Get(bundleId string) (*data.Bundle, error) {
	// try to fetch URI from cache. TODO: TTL of the bundle cache
	uri, err := warehouse.localCache.Get(bundleId)
	if err != nil {
		log.Warn("Failed to access local DB", "uri", uri, "error", err)
	}

	if uri != nil {
		if uri, err := url.Parse(string(uri)); err == nil {
			return warehouse.Fetch(uri)
		}
	}

	// search URI from metadatabase
	query := bson.NewDocument(bson.EC.String("data.data.bundleId", bundleId))
	metadata, err := warehouse.metaDatabase.RetrieveAsset(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query on metadatabase")
	}

	parsedUri, err := url.Parse(metadata.Lookup("uri").StringValue())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse URI")
	}
	return warehouse.Fetch(parsedUri)
}

func (warehouse *DataWarehouse) Fetch(uri *url.URL) (*data.Bundle, error) {
	protoc, exists := warehouse.protocols[uri.Scheme]
	if !exists {
		return nil, errors.Errorf("the protocol %s is not supported", uri.Scheme)
	}
	return protoc.Read(uri)
}
