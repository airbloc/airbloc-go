package collections

import (
	"net"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct {
	conn    net.Conn
	adapter *Adapter
	server  *grpc.Server
}

func (cs *Service) Close() {
	cs.server.Stop()
	cs.conn.Close()
}

func (cs *Service) Create(ctx context.Context, req *CreateCollectionRequest) (*CreateCollectionResponse, error) {
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
func (cs *Service) List(ctx context.Context, req *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, nil
}

func NewService(conn net.Conn, adapter *Adapter) (*Service, error) {
	service := &Service{
		conn:    conn,
		adapter: adapter,
		server:  grpc.NewServer(),
	}
	RegisterCollectionServer(service.server, service)
	return service, nil
}
