package data

import (
	"net"

	"github.com/airbloc/airbloc-go/api/common"
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

// TODO
func (s *Service) PreEncrypt(stream DataService_PreEncryptServer) error {
	return nil
}

func (s *Service) Register(stream DataService_RegisterServer) error {
	return nil
}

func (s *Service) Get(ctx context.Context, dataId *DataId) (*Data, error) {
	return nil, nil
}

func (s *Service) BatchGet(ctx context.Context, batchId *Batch) (*BatchGetResult, error) {
	return nil, nil
}

func (s *Service) SetPermission(ctx context.Context, req *SetDataPermissionRequest) (*common.Result, error) {
	return nil, nil
}

func (s *Service) SetPermissionBatch(ctx context.Context, req *SetBatchDataPermissionRequest) (*common.Results, error) {
	return nil, nil
}

func (s *Service) Unregister(ctx context.Context, dataId *DataId) (*common.Result, error) {
	return nil, nil
}

func (s *Service) UnregisterBatch(ctx context.Context, batchId *Batch) (*common.Results, error) {
	return nil, nil
}

func (s *Service) Select(stream DataService_SelectServer) error {
	return nil
}

func (s *Service) Release(ctx context.Context, batchId *Batch) (*common.Result, error) {
	return nil, nil
}

func NewService(conn net.Conn) (*Service, error) {
	service := &Service{
		conn:   conn,
		server: grpc.NewServer(),
	}
	RegisterDataServiceServer(service.server, service)
	return service, nil
}
