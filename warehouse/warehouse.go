package warehouse

import (
	"time"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse/bundle"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

type DataWarehouse struct {
	kms            *key.Manager
	protocols      map[string]protocol.Protocol
	DefaultStorage storage.Storage
}

func New(
	kms *key.Manager,
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
		DefaultStorage: defaultStorage,
	}
}

func (warehouse *DataWarehouse) CreateBundle(collection common.ID) *BundleStream {
	return newBundleStream(warehouse, collection)
}

func (warehouse *DataWarehouse) validate(collection common.ID, data *data.Data) error {
	return nil
}

func (warehouse *DataWarehouse) encrypt(d *data.Data) (*data.EncryptedData, error) {
	encryptedPayload, err := warehouse.kms.Encrypt(d.Payload)
	if err != nil {
		return nil, err
	}
	return &data.EncryptedData{
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

	// TODO: save metadata
	return createdBundle, nil
}
