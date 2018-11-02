package data

import (
	"net"

	"github.com/airbloc/airbloc-go/common"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type Service struct {
	conn   net.Conn
	server *grpc.Server
}

func (s *Service) Close() {
	s.server.Stop()
	s.conn.Close()
}

func (s *Service) Get(ctx context.Context, dataId *DataId) (*DataResult, error) {
	return nil, nil
}

func (s *Service) BatchGet(ctx context.Context, batchId *BatchRequest) (*BatchGetResult, error) {
	return nil, nil
}

func (s *Service) SetPermission(ctx context.Context, req *SetDataPermissionRequest) (*common.Result, error) {
	return nil, nil
}

func (s *Service) SetPermissionBatch(ctx context.Context, req *SetBatchDataPermissionRequest) (*common.Results, error) {
	return nil, nil
}

func (s *Service) Delete(ctx context.Context, dataId *DataId) (*common.Result, error) {
	return nil, nil
}

func (s *Service) DeleteBatch(ctx context.Context, batchId *BatchRequest) (*common.Results, error) {
	return nil, nil
}

func (s *Service) Select(stream Data_SelectServer) error {
	return nil
}

func (s *Service) Release(ctx context.Context, batchId *BatchRequest) (*common.Result, error) {
	return nil, nil
}

func NewService(conn net.Conn) (*Service, error) {
	service := &Service{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterDataServer(service.server, service)
	return service, nil
}
