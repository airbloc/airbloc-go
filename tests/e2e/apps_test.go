package e2e

import (
	"context"
	"fmt"
	"github.com/airbloc/airbloc-go/tests/e2e/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"log"
	"math/rand"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

func createApp(name string, conn *grpc.ClientConn) string {
	apps := pb.NewAppsClient(conn)
	result, err := apps.Register(context.Background(), &pb.RegisterRequest{
		Name: fmt.Sprintf("app-test-%s-%d", generateUniqueName(), rand.Int()),
	})
	Ω(err).ShouldNot(HaveOccurred())
	return result.GetAppId()
}

var _ = Describe("Apps", func() {
	var conn *grpc.ClientConn

	BeforeEach(func() {
		conn = e2eutils.ConnectGRPC()
	})

	It("should create a new app", func() {
		id := createApp(fmt.Sprintf("app-test-%s-%d", generateUniqueName(), rand.Int()), conn)
		log.Println("Created new app", id)
	})
})
