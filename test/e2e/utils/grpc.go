package e2eutils

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

func ConnectGRPC() *grpc.ClientConn {
	endpoint := "localhost:9124"
	if e, ok := os.LookupEnv("ENDPOINT"); ok {
		endpoint = e
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithInsecure())
	Ω(err).ShouldNot(HaveOccurred())

	return conn
}
