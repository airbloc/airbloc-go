package serverapi

import (
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/exchange"
	"github.com/airbloc/airbloc-go/node"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/ptypes/empty"
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

func (api *ExchangeAPI) Prepare(ctx context.Context, req *pb.OrderRequest) (*pb.OfferId, error) {
	contract := req.GetContract().GetSmartEscrow()

	to := common.HexToAddress(req.GetTo())
	escrowAddr := common.HexToAddress(req.GetContract().GetSmartEscrow().GetAddress())

	var escrowOpenSign, escrowCloseSign [4]byte
	copy(escrowOpenSign[:], contract.GetOpenSign())
	copy(escrowCloseSign[:], contract.GetCloseSign())

	rawDataIds := req.GetDataIds()
	dataIds := make([][20]byte, len(rawDataIds))
	for i, idStr := range rawDataIds {
		idBytes, err := hexutil.Decode(idStr)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Failed to decode dataId")
		}
		if len(idBytes) != 16 {
			return nil, status.Errorf(codes.InvalidArgument, "Wrong id length (expected 16)")
		}
		copy(dataIds[i][:], idBytes)
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
	return &pb.OfferId{OfferId: offerId.Hex()}, err
}

func (api *ExchangeAPI) AddDataIds(ctx context.Context, req *pb.DataIds) (*empty.Empty, error) {
	offerId, err := ablCommon.HexToID(req.GetOfferId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to decode offerId")
	}

	rawDataIds := req.GetDataIds()
	dataIds := make([][20]byte, len(rawDataIds))
	for i, idStr := range rawDataIds {
		idBytes, err := hexutil.Decode(idStr)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Failed to decode dataId")
		}
		if len(idBytes) != 16 {
			return nil, status.Errorf(codes.InvalidArgument, "Wrong id length (expected 16)")
		}
		copy(dataIds[i][:], idBytes)
	}

	err = api.manager.AddDataIds(ctx, offerId, dataIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add data ids")
	}
	return &empty.Empty{}, nil
}

func (api *ExchangeAPI) Order(ctx context.Context, req *pb.OfferId) (*empty.Empty, error) {
	offerId, err := ablCommon.HexToID(req.GetOfferId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to decode offerId")
	}

	err = api.manager.Order(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to order")
	}
	return &empty.Empty{}, nil
}

func (api *ExchangeAPI) Settle(ctx context.Context, req *pb.OfferId) (*empty.Empty, error) {
	offerId, err := ablCommon.HexToID(req.GetOfferId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to decode offerId")
	}

	err = api.manager.Settle(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to settle")
	}
	return &empty.Empty{}, nil
}

func (api *ExchangeAPI) Reject(ctx context.Context, req *pb.OfferId) (*empty.Empty, error) {
	offerId, err := ablCommon.HexToID(req.GetOfferId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to decode offerId")
	}

	err = api.manager.Reject(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to reject")
	}
	return &empty.Empty{}, nil
}

func (api *ExchangeAPI) CloseOrder(ctx context.Context, req *pb.OfferId) (*pb.Receipt, error) {
	offerId, err := ablCommon.HexToID(req.GetOfferId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to decode offerId")
	}

	err = api.manager.CloseOrder(ctx, offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to close offer")
	}
	return &pb.Receipt{}, nil
}

func (api *ExchangeAPI) GetOffer(ctx context.Context, req *pb.OfferId) (*pb.Offer, error) {
	offerId, err := ablCommon.HexToID(req.GetOfferId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to decode offerId")
	}

	offer, err := api.manager.GetOffer(offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get offer")
	}
	escrow := offer.Escrow

	rawDataIds := make([]string, len(offer.DataIds))
	for i, id := range offer.DataIds {
		rawDataIds[i] = hexutil.Encode(id[:])
	}

	return &pb.Offer{
		From:    offer.From.Hex(),
		To:      offer.To.Hex(),
		DataIds: rawDataIds,
		Contract: &pb.Contract{
			Type: pb.Contract_SMART,
			SmartEscrow: &pb.SmartContract{
				Address:   escrow.Addr.Hex(),
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

func (api *ExchangeAPI) GetOfferCompact(ctx context.Context, req *pb.OfferId) (*pb.OfferCompact, error) {
	offerId, err := ablCommon.HexToID(req.GetOfferId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to decode offerId")
	}

	offer, err := api.manager.GetOfferCompact(offerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get offer")
	}

	return &pb.OfferCompact{
		From:     offer.From.Hex(),
		To:       offer.To.Hex(),
		Escrow:   offer.Escrow.Hex(),
		Reverted: offer.Reverted,
	}, nil
}

// TODO: ignored chainID
func (api *ExchangeAPI) GetReceiptsByOfferor(ctx context.Context, req *pb.ReceiptRequest) (*pb.Offers, error) {
	offeror := common.HexToAddress(req.GetAddress())

	offers, err := api.manager.GetReceiptsByOfferor(offeror)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get receipts")
	}

	rawOffers := make([]string, len(offers))
	for i, offer := range offers {
		rawOffers[i] = hexutil.Encode(offer[:])
	}

	return &pb.Offers{OfferIds: rawOffers}, nil
}

func (api *ExchangeAPI) GetReceiptsByOfferee(ctx context.Context, req *pb.ReceiptRequest) (*pb.Offers, error) {
	offeree := common.HexToAddress(req.GetAddress())

	offers, err := api.manager.GetReceiptsByOfferee(offeree)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get receipts")
	}

	rawOffers := make([]string, len(offers))
	for i, offer := range offers {
		rawOffers[i] = hexutil.Encode(offer[:])
	}

	return &pb.Offers{OfferIds: rawOffers}, nil
}

func (api *ExchangeAPI) GetReceiptsByEscrow(ctx context.Context, req *pb.ReceiptRequest) (*pb.Offers, error) {
	escrow := common.HexToAddress(req.GetAddress())

	offers, err := api.manager.GetReceiptsByEscrow(escrow)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get receipts")
	}

	rawOffers := make([]string, len(offers))
	for i, offer := range offers {
		rawOffers[i] = hexutil.Encode(offer[:])
	}

	return &pb.Offers{OfferIds: rawOffers}, nil
}

func (api *ExchangeAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterExchangeServer(service.GrpcServer, api)
}
