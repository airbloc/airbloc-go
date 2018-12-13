package warehouse

import (
	"github.com/airbloc/airbloc-go/common"
)

type BundleStream struct {
	provider   common.ID
	collection common.ID
	warehouse  *DataWarehouse
	data       []*common.EncryptedData
	DataCount  int
}

func newBundleStream(warehouse *DataWarehouse, provider, collection common.ID) *BundleStream {
	return &BundleStream{
		provider:   provider,
		collection: collection,
		warehouse:  warehouse,
		data:       []*common.EncryptedData{},
		DataCount:  0,
	}
}

func (stream *BundleStream) Add(data *common.Data) (err error) {
	err = stream.warehouse.validate(stream.collection, data)
	if err != nil {
		return
	}

	encryptedData, err := stream.warehouse.encrypt(data)
	if err != nil {
		return
	}
	return stream.AddEncrypted(encryptedData)
}

func (stream *BundleStream) AddEncrypted(encryptedData *common.EncryptedData) error {
	stream.data = append(stream.data, encryptedData)
	stream.DataCount += 1
	return nil
}
