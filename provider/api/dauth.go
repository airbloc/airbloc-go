package api

import (
	"context"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/collections"
	"github.com/airbloc/airbloc-go/shared/dauth"
	"github.com/airbloc/airbloc-go/shared/node"
	"github.com/airbloc/airbloc-go/shared/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DAuthAPI struct {
	dauthClient *dauth.ProviderClient
	collections *collections.Collections
}

func NewDAuthAPI(backend node.Backend) (node.API, error) {
	dauthClient := dauth.NewProviderClient(backend.Kms(), backend.Client(), backend.P2P())
	return &DAuthAPI{
		dauthClient: dauthClient,
		collections: collections.New(backend.Client()),
	}, nil
}

func (api *DAuthAPI) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	userDelegateAddr := ethCommon.HexToAddress(req.GetUserDelegate())
	accountId, err := api.dauthClient.SignIn(ctx, req.GetIdentity(), userDelegateAddr)
	if err != nil {
		return nil, err
	}
	return &pb.SignInResponse{
		AccountId: accountId.Hex(),
	}, nil
}

func (api *DAuthAPI) GetAuthorizations(ctx context.Context, req *pb.GetAuthorizationsRequest) (*pb.GetAuthorizationsResponse, error) {
	appId, err := types.HexToID(req.GetAppId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid app ID: %s", req.GetAppId())
	}
	accountId, err := types.HexToID(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid account ID: %s", req.GetAccountId())
	}

	collectionIds, err := api.collections.ListID(ctx, appId)
	if err != nil {
		return nil, err
	}

	// check that user have been DAuthed at least once
	hasAuthorizedBefore, err := api.hasAuthedBefore(collectionIds, accountId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAuthorizationsResponse{
		HasAuthorizedBefore: hasAuthorizedBefore,
	}
	for _, collectionId := range collectionIds {
		// collectionId
		authorized, err := api.collections.IsCollectionAllowed(collectionId, accountId)
		if err != nil {
			return nil, err
		}
		response.Authorizations = append(response.Authorizations, &pb.GetAuthorizationsResponse_Authorization{
			CollectionId: collectionId.Hex(),
			Authorized:   authorized,
		})
	}
	return response, nil
}

func (api *DAuthAPI) Allow(ctx context.Context, req *pb.DAuthRequest) (*empty.Empty, error) {
	collectionId, err := types.HexToID(req.GetCollectionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", req.GetCollectionId())
	}
	accountId, err := types.HexToID(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid account ID: %s", req.GetAccountId())
	}
	err = api.dauthClient.Allow(ctx, collectionId, accountId)
	return &empty.Empty{}, err
}

func (api *DAuthAPI) Deny(ctx context.Context, req *pb.DAuthRequest) (*empty.Empty, error) {
	collectionId, err := types.HexToID(req.GetCollectionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collection ID: %s", req.GetCollectionId())
	}
	accountId, err := types.HexToID(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid account ID: %s", req.GetAccountId())
	}
	err = api.dauthClient.Deny(ctx, collectionId, accountId)
	return &empty.Empty{}, err
}

func (api *DAuthAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDAuthServer(service.GrpcServer, api)
	pb.RegisterDAuthHandlerFromEndpoint(context.Background(), service.RestAPIMux, service.Address, []grpc.DialOption{grpc.WithInsecure()})
}

func (api *DAuthAPI) hasAuthedBefore(collectionIds []types.ID, accountId types.ID) (bool, error) {
	collectionIdBytes := types.IDListToByteList(collectionIds)

	allowEvents, err := api.collections.GetContract().FilterAllowed(&bind.FilterOpts{}, collectionIdBytes, [][8]byte{accountId})
	if err != nil {
		return false, errors.Wrap(err, "failed to scan allow DAuth events of the user")
	}
	defer allowEvents.Close()
	if allowEvents.Next() {
		return true, nil
	}
	if allowEvents.Error() != nil {
		return false, allowEvents.Error()
	}

	denyEvents, err := api.collections.GetContract().FilterAllowed(&bind.FilterOpts{}, collectionIdBytes, [][8]byte{accountId})
	if err != nil {
		return false, errors.Wrap(err, "failed to scan deny DAuth events of the user")
	}
	defer denyEvents.Close()
	if denyEvents.Next() {
		return true, nil
	}
	return false, denyEvents.Error()
}
