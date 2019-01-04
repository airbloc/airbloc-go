package warehouse

import (
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/azer/logger"
	"github.com/pkg/errors"
)

type BundleStream struct {
	provider   common.ID
	collection *collections.Collection
	warehouse  *DataWarehouse
	data       []*common.EncryptedData
	DataCount  int
}

func newBundleStream(warehouse *DataWarehouse, provider common.ID, collection *collections.Collection) *BundleStream {
	return &BundleStream{
		provider:   provider,
		collection: collection,
		warehouse:  warehouse,
		data:       []*common.EncryptedData{},
		DataCount:  0,
	}
}

func (stream *BundleStream) Add(data *common.Data) error {
	err := stream.warehouse.validate(stream.collection, data)
	if errors.Cause(err) == errValidationFailed {
		stream.warehouse.log.Info("warning: %s", err.Error(), logger.Attrs{
			"collection": stream.collection.Id.Hex(),
			"dataOwner":  data.OwnerAnid.Hex(),
		})
		return err
	}
	if err != nil {
		return err
	}

	encryptedData, err := stream.warehouse.encrypt(data)
	if err != nil {
		return err
	}
	return stream.AddEncrypted(encryptedData)
}

func (stream *BundleStream) AddEncrypted(encryptedData *common.EncryptedData) error {
	stream.data = append(stream.data, encryptedData)
	stream.DataCount += 1
	return nil
}
