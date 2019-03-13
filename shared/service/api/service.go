package api

import (
	"fmt"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

type Service struct {
	GrpcServer *grpc.Server
	HttpServer *http.Server
	RestAPIMux *runtime.ServeMux
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
	restAPImux := runtime.NewServeMux()
	config := backend.Config()
	address := fmt.Sprintf("localhost:%d", config.Port)
	svc := &Service{
		GrpcServer: grpcServer,
		RestAPIMux: restAPImux,
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

	go service.withErrorHandler(service.GrpcServer.Serve, grpcLis)
	go service.withErrorHandler(service.HttpServer.Serve, restLis)

	service.logger.Info("Server started at {}", service.Address)
	return m.Serve()
}

func (service *Service) withErrorHandler(serveMethod func(net.Listener) error, lis net.Listener) {
	if err := serveMethod(lis); err != http.ErrServerClosed {
		service.logger.Error("failed to run rpc server: %+v", err)
		os.Exit(1)
	}
}

// Stop stops gRPC server.
func (service *Service) Stop() {
	service.GrpcServer.GracefulStop()
	service.HttpServer.Close()
}
