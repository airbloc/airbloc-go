package api

import (
	"fmt"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

type Service struct {
	GrpcServer *grpc.Server
	HttpServer *http.Server
	Address    string

	port int

	// for logging
	logger *logger.Logger
}

func NewService(backend service.Backend) (service.Service, error) {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			UnaryServerLogger(),
		)),
	)
	config := backend.Config()
	address := fmt.Sprintf("localhost:%d", config.Port)
	svc := &Service{
		GrpcServer: grpcServer,
		HttpServer: &http.Server{
			Addr:    address,
			Handler: restAPImux,
		},
		Address: address,
		port:    config.Port,
		logger:  logger.New("apiservice"),
	}
	return svc, nil
}

func (service *Service) Attach(name string, api API) {

}

// Start serves gRPC server and HTTP REST API Server on given TCP port.
func (service *Service) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", service.port))
	if err != nil {
		return errors.Wrapf(err, "failed to listen to TCP port %d for RPC", service.port)
	}

	// Route gRPC (HTTP2), REST (HTTP) connection accordingly.
	m := cmux.New(lis)
	grpcLis := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	restLis := m.Match(cmux.HTTP1Fast())

	var g errgroup.Group

	g.Go(func() error {
		return service.GrpcServer.Serve(grpcLis)
	})

	g.Go(func() error {
		return service.HttpServer.Serve(restLis)
	})

	if err := g.Wait(); err != nil {
		service.logger.Error("failed to run rpc server: %+v", err)
		os.Exit(1)
	}

	service.logger.Info("Server started at {}", service.Address)
	return m.Serve()
}

// Stop stops gRPC server.
func (service *Service) Stop() {
	service.GrpcServer.GracefulStop()
	service.HttpServer.Close()
}
