package serverapi

import (
	"context"
	"github.com/airbloc/airbloc-go/warehouse/service"
	"github.com/azer/logger"
	"sync"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type WarehouseAPI struct {
	warehouse *warehouse.DataWarehouse
	log       *logger.Logger
}

func NewWarehouseAPI(backend node.Backend) (_ node.API, err error) {
	service, ok := backend.GetService("warehouse").(*warehouseservice.Service)
	if !ok {
		return nil, errors.New("warehouse service is not registered")
	}
	return &WarehouseAPI{
		warehouse: service.GetManager(),
		log:       logger.New("warehouse"),
	}, nil
}

func (api *WarehouseAPI) StoreBundle(stream pb.Warehouse_StoreBundleServer) error {
	total := 0
	successful := 0
	timer := api.log.Timer()
	defer func() {
		timer.End("Successfully Ingested %d of %d data.", successful, total)
	}()

	var bundleStream *warehouse.BundleStream
	var wg sync.WaitGroup
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
			bundleStream, err = api.warehouse.CreateBundle(collectionId)
			if err != nil {
				return err
			}
		}

		userId, err := common.HexToID(request.GetUserId())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "Invalid user ANID: %s", request.GetUserId())
		}
		total += 1
		wg.Add(1)

		datum := &common.Data{
			Payload: request.GetPayload(),
			UserId:  userId,
		}
		go func() {
			if err := bundleStream.Add(datum); err != nil {
				api.log.Error("failed to add a data: %s", err.Error())
			} else {
				successful += 1
			}
			wg.Done()
		}()
	}
	wg.Wait()

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
				return status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", request.GetCollectionId())
			}
			bundleStream, err = api.warehouse.CreateBundle(collectionId)
			if err != nil {
				return err
			}
		}

		userId, err := common.HexToID(request.GetUserId())
		if err != nil {
			return errors.Wrapf(err, "failed to parse ANID %s", request.GetUserId())
		}

		datum := &common.EncryptedData{
			Payload: request.GetEncryptedPayload(),
			UserId:  userId,
			Capsule: request.GetCapsule(),
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

func (api *WarehouseAPI) DeleteBundle(ctx context.Context, request *pb.DeleteBundleRequest) (*pb.DeleteBundleResult, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (api *WarehouseAPI) ListBundle(ctx context.Context, req *pb.ListBundleRequest) (*pb.ListBundleResult, error) {
	providerId, err := common.HexToID(req.GetProviderId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid provider ID: %s", req.GetProviderId())
	}
	bundles, err := api.warehouse.List(providerId)
	if err != nil {
		return nil, err
	}

	bundleResults := make([]*pb.ListBundleResult_Bundle, len(bundles))
	for i, bundle := range bundles {
		bundleResults[i] = &pb.ListBundleResult_Bundle{
			CollectionId: bundle.Collection.Hex(),
			Index:        1010,
			CreatedAt:    uint64(bundle.IngestedAt.Unix()),
			DataCount:    uint64(bundle.DataCount),
			Uri:          bundle.Uri,
		}
	}
	return &pb.ListBundleResult{
		Bundles: bundleResults,
	}, nil
}

func (api *WarehouseAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterWarehouseServer(service.GrpcServer, api)
}
