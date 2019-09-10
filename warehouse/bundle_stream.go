package warehouse

import (
	"sync"

	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
	"github.com/pkg/errors"
)

type BundleStream struct {
	provider  types.ID
	warehouse *Manager
	data      map[types.ID][]*types.EncryptedData
	DataCount int

	mu sync.Mutex
}

func newBundleStream(warehouse *Manager, provider types.ID, collection interface{}) *BundleStream {
	return &BundleStream{
		provider:  provider,
		warehouse: warehouse,
		data:      map[types.ID][]*types.EncryptedData{},
		DataCount: 0,
	}
}

func (stream *BundleStream) Add(data *types.Data) error {
	err := stream.warehouse.validate(nil, data)
	if errors.Cause(err) == errValidationFailed {
		stream.warehouse.log.Info("warning: %s", err.Error(), logger.Attrs{
			//"collection": stream.collection.Id.Hex(),
			"dataOwner": data.UserId.Hex(),
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
