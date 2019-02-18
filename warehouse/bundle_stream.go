package warehouse

import (
	"github.com/airbloc/airbloc-go/collections"
	"github.com/airbloc/airbloc-go/common"
	"github.com/azer/logger"
	"github.com/pkg/errors"
	"sync"
)

type BundleStream struct {
	provider   common.ID
	collection *collections.Collection
	warehouse  *DataWarehouse
	data       map[common.ID][]*common.EncryptedData
	DataCount  int

	mu sync.Mutex
}

func newBundleStream(warehouse *DataWarehouse, provider common.ID, collection *collections.Collection) *BundleStream {
	return &BundleStream{
		provider:   provider,
		collection: collection,
		warehouse:  warehouse,
		data:       map[common.ID][]*common.EncryptedData{},
		DataCount:  0,
	}
}

func (stream *BundleStream) Add(data *common.Data) error {
	err := stream.warehouse.validate(stream.collection, data)
	if errors.Cause(err) == errValidationFailed {
		stream.warehouse.log.Info("warning: %s", err.Error(), logger.Attrs{
			"collection": stream.collection.Id.Hex(),
			"dataOwner":  data.UserId.Hex(),
		})
		return err
	}
	if err != nil {
		return errors.Wrap(err, "error occured whilte validating data")
	}

	encryptedData, err := stream.warehouse.kms.EncryptData(data)
	if err != nil {
		return err
	}

	stream.AddEncrypted(encryptedData)
	return nil
}

func (stream *BundleStream) AddEncrypted(encryptedData *common.EncryptedData) {
	stream.mu.Lock()
	stream.data[encryptedData.UserId] = append(stream.data[encryptedData.UserId], encryptedData)
	stream.DataCount += 1
	stream.mu.Unlock()
}
