package serverapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"io"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

type WarehouseAPI struct {
	warehouse *warehouse.DataWarehouse
}

func NewWarehouseAPI(airbloc node.Backend) (_ node.API, err error) {
	config := airbloc.Config().Warehouse

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
		airbloc.Kms(),
		airbloc.LocalDatabase(),
		airbloc.MetaDatabase(),
		airbloc.Client(),
		defaultStorage,
		supportedProtocols)
	return &WarehouseAPI{dw}, nil
}

func (api *WarehouseAPI) StoreBundle(stream pb.Warehouse_StoreBundleServer) error {
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
			collectionId, err := common.HexToID(request.GetCollectionId())
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", request.GetCollectionId())
			}
			bundleStream = api.warehouse.CreateBundle(collectionId)
		}

		ownerAnid, err := common.HexToID(request.GetOwnerId())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "Invalid user ANID: %s", request.GetOwnerId())
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

	return stream.SendAndClose(&pb.StoreResult{
		BundleId:  bundle.Id,
		Uri:       bundle.Uri,
		DataCount: uint64(bundle.DataCount),
		GasUsed:   0,
	})
}

func (api *WarehouseAPI) StoreEncryptedBundle(stream pb.Warehouse_StoreEncryptedBundleServer) error {
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
			collectionId, err := common.HexToID(request.GetCollectionId())
			if err != nil {
				return errors.Wrapf(err, "failed to parse collection ID (%s)", request.GetCollectionId())
			}
			bundleStream = api.warehouse.CreateBundle(collectionId)
		}

		ownerAnid, err := common.HexToID(request.GetOwnerId())
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

	return stream.SendAndClose(&pb.StoreResult{
		BundleId:  bundle.Id,
		Uri:       bundle.Uri,
		DataCount: uint64(bundle.DataCount),
		GasUsed:   0,
	})
}

func (api *WarehouseAPI) DeleteBundle(context context.Context, request *pb.DeleteBundleRequest) (*pb.DeleteBundleResult, error) {
	return nil, nil
}

func (api *WarehouseAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterWarehouseServer(service.GrpcServer, api)
}
