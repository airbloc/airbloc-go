package node

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type APIService struct {
	GrpcServer *grpc.Server
	HttpServer *http.Server
	RestAPIMux *http.ServeMux
	Address    string

	port int
}

func NewAPIService(airbloc Backend) (Service, error) {
	grpcServer := grpc.NewServer()
	restAPImux := http.NewServeMux()

	proxyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			// redirect other traffics (e.g. Swagger, Prometheus, Custom REST Endpoint, ...)
			restAPImux.ServeHTTP(w, r)
		}
	})

	config := airbloc.Config()
	address := fmt.Sprintf("localhost:%d", config.Port)
	service := &APIService{
		GrpcServer: grpcServer,
		RestAPIMux: restAPImux,
		HttpServer: &http.Server{
			Addr:    address,
			Handler: proxyHandler,
		},
		Address: address,
		port:    config.Port,
	}
	return service, nil
}

func (service *APIService) Attach(name string, api API) {

}

// Start serves gRPC server on given TCP port.
func (service *APIService) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", service.port))
	if err != nil {
		return errors.Wrapf(err, "failed to listen to TCP port %d for RPC", service.port)
	}

	log.Info("Server started", "address", service.Address)
	if err := service.GrpcServer.Serve(lis); err != http.ErrServerClosed {
		return errors.Wrapf(err, "failed to start HTTP server")
	}
	return nil
}

// Stop stops gRPC server.
func (service *APIService) Stop() {
	service.GrpcServer.GracefulStop()
	service.HttpServer.Close()
}
