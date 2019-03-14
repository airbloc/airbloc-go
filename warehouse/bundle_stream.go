package warehouse

import (
	"github.com/airbloc/airbloc-go/provider/collections"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/pkg/errors"
	"sync"
)

type BundleStream struct {
	provider   types.ID
	collection *collections.Collection
	warehouse  *Manager
	data       map[types.ID][]*types.EncryptedData
	DataCount  int

	mu sync.Mutex
}

func newBundleStream(warehouse *Manager, provider types.ID, collection *collections.Collection) *BundleStream {
	return &BundleStream{
		provider:   provider,
		collection: collection,
		warehouse:  warehouse,
		data:       map[types.ID][]*types.EncryptedData{},
		DataCount:  0,
	}
}

func (stream *BundleStream) Add(data *types.Data) error {
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

func (stream *BundleStream) AddEncrypted(encryptedData *types.EncryptedData) {
	stream.mu.Lock()
	stream.data[encryptedData.UserId] = append(stream.data[encryptedData.UserId], encryptedData)
	stream.DataCount += 1
	stream.mu.Unlock()
}
