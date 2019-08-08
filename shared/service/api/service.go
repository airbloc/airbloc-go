package api

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/logger"
	"github.com/gin-gonic/gin"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type Service struct {
	GrpcServer *grpc.Server
	HttpServer *gin.Engine
	Address    string

	port     int
	listener net.Listener

	// for logging
	logger *logger.Logger
}

func NewService(backend service.Backend) (service.Service, error) {
	config := backend.Config()
	address := fmt.Sprintf("localhost:%d", config.Port)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			UnaryServerLogger(),
		)),
	)

	svc := &Service{
		GrpcServer: grpcServer,
		HttpServer: gin.New(),
		Address:    address,
		port:       config.Port,
		logger:     logger.New("apiservice"),
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
	if service.port == 0 {
		service.logger.Debug("Port randomly chosed. listening address is : %s", lis.Addr().String())
	}

	// Route gRPC (HTTP2), REST (HTTP) connection accordingly.
	m := cmux.New(lis)

	grpcLis := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	restLis := m.Match(cmux.HTTP1Fast())

	// grpc server
	go func() {
		if err = service.GrpcServer.Serve(grpcLis); err != http.ErrServerClosed {
			service.logger.Error("failed to run grpc api server: %+v", err)
			os.Exit(1)
		}
	}()
	// rest server
	go func() {
		if err = http.Serve(restLis, service.HttpServer); err != nil {
			service.logger.Error("failed to run rest api server: %+v", err)
			os.Exit(1)
		}
	}()

	service.logger.Info("Server started at {}", service.Address)
	service.listener = lis

	return m.Serve()
}

// Stop stops gRPC server.
func (service *Service) Stop() {
	service.GrpcServer.GracefulStop()
	service.listener.Close()
}
