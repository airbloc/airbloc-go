package serverapi

import (
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/exchange"
	"github.com/airbloc/airbloc-go/node"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExchangeAPI struct {
	manager *exchange.Manager
}

func NewExchangeAPI(backend node.Backend) (node.API, error) {
	ex := exchange.NewManager(backend.Client())
	return &ExchangeAPI{ex}, nil
}

func (api *ExchangeAPI) Prepare(ctx context.Context, req *pb.OrderRequest) (*commonpb.Hash, error) {
	contract := req.GetContract().GetSmartEscrow()

	to := common.BytesToAddress(req.GetTo().GetAddress())
	escrowAddr := common.BytesToAddress(contract.GetAddress().GetAddress())

	var escrowOpenSign, escrowCloseSign [4]byte
	copy(escrowOpenSign[:], contract.GetOpenSign())
	copy(escrowCloseSign[:], contract.GetCloseSign())

	rawDataIds := req.GetDataIds()
	dataIds := make([][16]byte, len(rawDataIds))
	for i, id := range rawDataIds {
		copy(dataIds[i][:], id)
	}

	offerId, err := api.manager.Prepare(
		ctx, to, escrowAddr,
		escrowOpenSign, contract.GetOpenArgs(),
		escrowCloseSign, contract.GetCloseSign(),
		dataIds...,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to prepare order request")
	}
	return &commonpb.Hash{Hash: offerId[:]}, err
}

func (api *ExchangeAPI) AddDataIds(ctx context.Context, req *pb.DataIds) (*commonpb.Result, error) {
	offerId := ablCommon.IDFromBytes(req.GetOfferId().GetHash())
	rawDataIds := req.GetDataIds()
	dataIds := make([][16]byte, len(rawDataIds))
	for i, id := range rawDataIds {
		copy(dataIds[i][:], id)
	}

	err := api.manager.AddDataIds(ctx, offerId, dataIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add data ids")
	}
	return &commonpb.Result{}, nil
}

func (api *ExchangeAPI) Order(ctx context.Context, req *commonpb.Hash) (*commonpb.Result, error) {
	offerId := ablCommon.IDFromBytes(req.GetHash())
	err := api.manager.Order(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to order")
	}
	return &commonpb.Result{}, nil
}

func (api *ExchangeAPI) Settle(ctx context.Context, req *commonpb.Hash) (*commonpb.Result, error) {
	offerId := ablCommon.IDFromBytes(req.GetHash())
	err := api.manager.Settle(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to settle")
	}
	return &commonpb.Result{}, nil
}

func (api *ExchangeAPI) Reject(ctx context.Context, req *commonpb.Hash) (*commonpb.Result, error) {
	offerId := ablCommon.IDFromBytes(req.GetHash())
	err := api.manager.Reject(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to reject")
	}
	return &commonpb.Result{}, nil
}

func (api *ExchangeAPI) CloseOrder(ctx context.Context, req *commonpb.Hash) (*pb.Receipt, error) {
	offerId := ablCommon.IDFromBytes(req.GetHash())

	err := api.manager.CloseOrder(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to close offer")
	}
	return &pb.Receipt{}, nil
}

// TODO: hard-coded chainID
func (api *ExchangeAPI) GetOffer(ctx context.Context, req *commonpb.Hash) (*pb.Offer, error) {
	offerId := ablCommon.IDFromBytes(req.GetHash())

	offer, err := api.manager.GetOffer(offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get offer")
	}
	escrow := offer.Escrow

	rawDataIds := make([][]byte, len(offer.DataIds))
	for i, id := range offer.DataIds {
		copy(rawDataIds[i], id[:])
	}

	return &pb.Offer{
		From:    &commonpb.Address{ChainId: 1337, Address: offer.From.Bytes()},
		To:      &commonpb.Address{ChainId: 1337, Address: offer.To.Bytes()},
		DataIds: rawDataIds,
		Contract: &pb.Contract{
			Type: pb.Contract_SMART,
			SmartEscrow: &pb.SmartContract{
				Address:   &commonpb.Address{ChainId: 1337, Address: escrow.Addr.Bytes()},
				OpenSign:  escrow.OpenSign[:],
				OpenArgs:  escrow.OpenArgs,
				CloseSign: escrow.CloseSign[:],
				CloseArgs: escrow.CloseArgs,
			},
		},
		Status:   pb.Status(offer.Status),
		Reverted: offer.Reverted,
	}, nil
}

// TODO: hard-coded chainID
func (api *ExchangeAPI) GetOfferCompact(ctx context.Context, req *commonpb.Hash) (*pb.OfferCompact, error) {
	offerId := ablCommon.IDFromBytes(req.GetHash())

	offer, err := api.manager.GetOfferCompact(offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get offer")
	}

	return &pb.OfferCompact{
		From:     &commonpb.Address{ChainId: 1337, Address: offer.From.Bytes()},
		To:       &commonpb.Address{ChainId: 1337, Address: offer.To.Bytes()},
		Escrow:   &commonpb.Address{ChainId: 1337, Address: offer.Escrow.Bytes()},
		Reverted: offer.Reverted,
	}, nil
}

// TODO: ignored chainID
func (api *ExchangeAPI) GetReceiptsByOfferor(ctx context.Context, req *commonpb.Address) (*pb.Offers, error) {
	offeror := common.BytesToAddress(req.GetAddress())

	offers, err := api.manager.GetReceiptsByOfferor(offeror)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get receipts")
	}

	rawOffers := make([]*commonpb.Hash, len(offers))
	for i, offer := range offers {
		rawOffers[i] = &commonpb.Hash{Hash: offer[:]}
	}

	return &pb.Offers{OfferIds: rawOffers}, nil
}

func (api *ExchangeAPI) GetReceiptsByOfferee(ctx context.Context, req *commonpb.Address) (*pb.Offers, error) {
	offeree := common.BytesToAddress(req.GetAddress())

	offers, err := api.manager.GetReceiptsByOfferee(offeree)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get receipts")
	}

	rawOffers := make([]*commonpb.Hash, len(offers))
	for i, offer := range offers {
		rawOffers[i] = &commonpb.Hash{Hash: offer[:]}
	}

	return &pb.Offers{OfferIds: rawOffers}, nil
}

func (api *ExchangeAPI) GetReceiptsByEscrow(ctx context.Context, req *commonpb.Address) (*pb.Offers, error) {
	escrow := common.BytesToAddress(req.GetAddress())

	offers, err := api.manager.GetReceiptsByEscrow(escrow)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get receipts")
	}

	rawOffers := make([]*commonpb.Hash, len(offers))
	for i, offer := range offers {
		rawOffers[i] = &commonpb.Hash{Hash: offer[:]}
	}

	return &pb.Offers{OfferIds: rawOffers}, nil
}

func (api *ExchangeAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterExchangeServer(service.GrpcServer, api)
}
