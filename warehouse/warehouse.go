package warehouse

import (
	"net/url"
	"time"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse/bundle"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

type DataWarehouse struct {
	kms            *key.Manager
	protocols      map[string]protocol.Protocol
	localCache     *localdb.Model
	metaDatabase   *metadb.Model
	DefaultStorage storage.Storage
}

func New(
	kms *key.Manager,
	localDatabase localdb.Database,
	metaDatabase metadb.Database,
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

func (warehouse *DataWarehouse) Store(stream *BundleStream) (*bundle.Bundle, error) {
	if stream == nil {
		return nil, errors.New("No data in the stream.")
	}
	ingestedAt := time.Now()

	// TODO: hash collision proof / generate on contract
	bundleId := common.GenerateID(
		warehouse.kms.OwnerKey.EthereumAddress,
		time.Now(),
		stream.collection[:])

	createdBundle := &bundle.Bundle{
		Id:         bundleId,
		Provider:   common.ID{}, /* TODO: implement appID */
		Collection: stream.collection,
		DataCount:  stream.DataCount,
		IngestedAt: ingestedAt,
		Data:       stream.data,
	}

	uri, err := warehouse.DefaultStorage.Save(createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save bundle to the storage")
	}
	createdBundle.Uri = uri.String()

	// save metadata to make the bundle searchable
	bundleInfo := map[string]interface{}{
		"bundleId":   bundleId.String(),
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

func (warehouse *DataWarehouse) Get(bundleId string) (*bundle.Bundle, error) {
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

func (warehouse *DataWarehouse) Fetch(uri *url.URL) (*bundle.Bundle, error) {
	protoc, exists := warehouse.protocols[uri.Scheme]
	if !exists {
		return nil, errors.Errorf("the protocol %s is not supported", uri.Scheme)
	}
	return protoc.Read(uri)
}
