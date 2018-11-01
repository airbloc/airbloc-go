package collections

import (
	"net"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type API struct {
	conn    net.Conn
	adapter *Service
	server  *grpc.Server
}

func (cs *API) Close() {
	cs.server.Stop()
	cs.conn.Close()
}

func (cs *API) Create(ctx context.Context, req *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	hash, err := cs.adapter.Register(ctx, &Collection{
		AppId:    common.HexToHash(req.AppId),
		SchemaId: common.HexToHash(req.SchemaId),
		Policy: &IncentivePolicy{
			DataProducer:  req.Policy.DataProducer,
			DataProcessor: req.Policy.DataProcessor,
			DataRelayer:   req.Policy.DataRelayer,
			DataSource:    req.Policy.DataSource,
		},
	})
	return &CreateCollectionResponse{
		CollectionId: hash.Hex(),
	}, err
}

// TODO after localdb integration
func (cs *API) List(ctx context.Context, req *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, nil
}

func NewAPI(conn net.Conn, adapter *Service) (*API, error) {
	api := &API{
		conn:    conn,
		adapter: adapter,
		server:  grpc.NewServer(),
	}
	RegisterCollectionServer(api.server, api)
	return api, nil
}
