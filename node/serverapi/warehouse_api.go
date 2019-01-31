package serverapi

import (
	"context"
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
}

func NewWarehouseAPI(backend node.Backend) (_ node.API, err error) {
	service, ok := backend.GetService("warehouse").(*warehouse.Service)
	if !ok {
		return nil, errors.New("warehouse service is not registered")
	}

	return &WarehouseAPI{service.GetManager()}, nil
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
			bundleStream, err = api.warehouse.CreateBundle(collectionId)
			if err != nil {
				return err
			}
		}

		userId, err := common.HexToID(request.GetUserId())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "Invalid user ANID: %s", request.GetUserId())
		}

		datum := &common.Data{
			Payload: request.GetPayload(),
			UserId:  userId,
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

func (api *WarehouseAPI) GetBundleInfo(ctx context.Context, request *pb.BundleInfoRequest) (*pb.BundleInfoResponse, error) {
	bundleId, err := common.HexToID(request.GetBundleId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert bundleId to common.ID format : %v", err)
	}

	bundleInfo, err := api.warehouse.GetBundleInfo(ctx, bundleId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get bundle info : %v", err)
	}

	return &pb.BundleInfoResponse{
		BundleId:   bundleInfo.Id,
		Uri:        bundleInfo.Uri,
		Provider:   bundleInfo.Provider,
		Collection: bundleInfo.Collection,
		IngestedAt: bundleInfo.IngestedAt,
		DataIds:    bundleInfo.DataIds,
	}, nil
}

func (api *WarehouseAPI) GetUserDataIds(ctx context.Context, request *pb.UserDataIdsRequest) (*pb.UserDataIdsResponse, error) {
	userId, err := common.HexToID(request.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert userId to common.ID format : %v", err)
	}

	userInfoes, err := api.warehouse.GetUserInfo(ctx, userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info : %v", err)
	}

	collections := make([]*pb.UserDataIdsResponse_Collection, len(userInfoes))
	for i, collection := range userInfoes {
		collections[i] = &pb.UserDataIdsResponse_Collection{
			AppId:        collection.AppId,
			SchemaId:     collection.SchemaId,
			CollectionId: collection.CollectionId,
			DataIds:      make([]*pb.UserDataIdsResponse_DataInfo, len(collection.DataIds)),
		}

		for j, dataId := range collection.DataIds {
			collections[i].DataIds[j] = &pb.UserDataIdsResponse_DataInfo{
				Id:         dataId.Id,
				IngestedAt: dataId.IngestedAt,
			}
		}
	}
	return &pb.UserDataIdsResponse{Collections: collections}, nil
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
