package api

import (
	"context"

	"io"

	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

type API struct {
	warehouse *warehouse.DataWarehouse
}

func New(airbloc *api.AirblocBackend) (_ api.API, err error) {
	config := airbloc.Config.Warehouse

	supportedProtocols := []protocol.Protocol{
		protocol.NewHttpProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
		protocol.NewHttpsProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
	}

	var defaultStorage storage.Storage
	if config.DefaultStorage == "local" {
		defaultStorage, err = storage.NewLocalStorage(
			config.LocalStorage.SavePath,
			config.LocalStorage.Endpoint)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.Errorf("unknown storage type: %s", config.DefaultStorage)
	}

	dw := warehouse.New(
		airbloc.Kms,
		airbloc.LocalDatabase,
		airbloc.MetaDatabase,
		airbloc.Ethclient,
		defaultStorage,
		supportedProtocols,
	)
	return &API{dw}, nil
}

func (api *API) StoreBundle(stream Warehouse_StoreBundleServer) error {
	var bundleStream *warehouse.BundleStream
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if bundleStream == nil {
			collectionId, err := common.IDFromString(request.GetCollection())
			if err != nil {
				return errors.Wrapf(err, "failed to parse collection ID (%s)", request.GetCollection())
			}
			bundleStream = api.warehouse.CreateBundle(collectionId)
		}

		ownerAnid, err := common.IDFromString(request.GetOwnerId())
		if err != nil {
			return errors.Wrapf(err, "failed to parse ANID %s", request.GetOwnerId())
		}

		datum := &common.Data{
			Payload:   request.GetPayload(),
			OwnerAnid: ownerAnid,
		}
		bundleStream.Add(datum)
	}

	bundle, err := api.warehouse.Store(bundleStream)
	if err != nil {
		return errors.Wrap(err, "failed to store a bundle")
	}

	return stream.SendAndClose(&StoreResult{
		BundleId:  bundle.Id.String(),
		Uri:       bundle.Uri,
		DataCount: uint64(bundle.DataCount),
		GasUsed:   0,
	})
}

func (api *API) StoreEncryptedBundle(stream Warehouse_StoreEncryptedBundleServer) error {
	var bundleStream *warehouse.BundleStream
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if bundleStream == nil {
			collectionId, err := common.IDFromString(request.GetCollection())
			if err != nil {
				return errors.Wrapf(err, "failed to parse collection ID (%s)", request.Collection)
			}
			bundleStream = api.warehouse.CreateBundle(collectionId)
		}

		ownerAnid, err := common.IDFromString(request.GetOwnerId())
		if err != nil {
			return errors.Wrapf(err, "failed to parse ANID %s", request.GetOwnerId())
		}

		datum := &common.EncryptedData{
			Payload:   request.GetEncryptedPayload(),
			OwnerAnid: ownerAnid,
			Capsule:   request.GetCapsule(),
		}
		bundleStream.AddEncrypted(datum)
	}

	bundle, err := api.warehouse.Store(bundleStream)
	if err != nil {
		return errors.Wrap(err, "failed to store a bundle")
	}

	return stream.SendAndClose(&StoreResult{
		BundleId:  bundle.Id.String(),
		Uri:       bundle.Uri,
		DataCount: uint64(bundle.DataCount),
		GasUsed:   0,
	})
}

func (api *API) DeleteBundle(context context.Context, request *DeleteBundleRequest) (*DeleteBundleResult, error) {
	return nil, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterWarehouseServer(service.GrpcServer, api)
}
