package warehouse

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/data"
)

type BundleStream struct {
	collection common.ID
	warehouse  *DataWarehouse
	data       []*data.EncryptedData
	DataCount  int
}

func newBundleStream(warehouse *DataWarehouse, collection common.ID) *BundleStream {
	return &BundleStream{
		collection: collection,
		warehouse:  warehouse,
		data:       []*data.EncryptedData{},
		DataCount:  0,
	}
}

func (stream *BundleStream) Add(data *data.Data) (err error) {
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

func (stream *BundleStream) AddEncrypted(encryptedData *data.EncryptedData) error {
	stream.data = append(stream.data, encryptedData)
	stream.DataCount += 1
	return nil
}
