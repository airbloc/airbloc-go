package serverapi

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/exchange"
	"github.com/airbloc/airbloc-go/node"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	googlepb "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
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

	from := common.AddresFromBytes(req.GetFrom().GetAddress())
	to := common.AddresFromBytes(req.GetFrom().GetAddress())
	escrowAddr := common.AddresFromBytes(contract.GetAddress().GetAddress())

	var escrowSign [4]byte
	copy(escrowSign[:], contract.GetSelector())

	rawDataIds := req.GetDataIds()
	dataIds := make([][16]byte, len(rawDataIds))
	for i, id := range rawDataIds {
		copy(dataIds[i][:], id)
	}

	offerId, err := api.manager.Prepare(
		ctx,
		from, to,
		escrowAddr, escrowSign, contract.GetArguments(),
		dataIds...,
	)
	return &commonpb.Hash{Hash: offerId[:]}, err
}
func (api *ExchangeAPI) AddDataIds(ctx context.Context, req *pb.DataIds) (*googlepb.Empty, error) {

	api.manager.AddDataIds(ctx)
}
func (api *ExchangeAPI) Order(ctx context.Context, req *commonpb.Hash) (*googlepb.Empty, error) {

}
func (api *ExchangeAPI) Settle(ctx context.Context, req *commonpb.Hash) (*googlepb.Empty, error) {

}
func (api *ExchangeAPI) Reject(ctx context.Context, req *commonpb.Hash) (*googlepb.Empty, error) {

}
func (api *ExchangeAPI) CloseOrder(ctx context.Context, req *commonpb.Hash) (*pb.Receipt, error) {

}
func (api *ExchangeAPI) GetOffer(ctx context.Context, req *commonpb.Hash) (*pb.Offer, error) {

}
func (api *ExchangeAPI) GetOfferCompact(ctx context.Context, req *commonpb.Hash) (*pb.OfferCompact, error) {
}
func (api *ExchangeAPI) GetReceiptsByOfferor(ctx context.Context, req *commonpb.Address) (*pb.Offers, error) {
}
func (api *ExchangeAPI) GetReceiptsByOfferee(ctx context.Context, req *commonpb.Address) (*pb.Offers, error) {
}
func (api *ExchangeAPI) GetReceiptsByEscrow(ctx context.Context, req *commonpb.Address) (*pb.Offers, error) {
}

func (api *ExchangeAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterExchangeServer(service.GrpcServer, api)
}
