package e2e

import (
	"context"
	"fmt"
	"github.com/airbloc/airbloc-go/ablclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"log"
	"sync"
	"testing"
	"time"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

func testCreateApp(t *testing.T, ctx context.Context, conn *grpc.ClientConn) string {
	apps := pb.NewAppsClient(conn)
	result, err := apps.Register(ctx, &pb.RegisterRequest{
		Name: fmt.Sprintf("app-test-%d", time.Now().Unix()),
	})
	assert.NoError(t, err)
	return result.GetAppId()
}

func testCreateSchema(t *testing.T, ctx context.Context, conn *grpc.ClientConn) string {
	schemas := pb.NewSchemaClient(conn)
	result, err := schemas.Create(ctx, &pb.CreateSchemaRequest{
		Name:   fmt.Sprintf("data-test-%d", time.Now().Unix()),
		Schema: testSchema,
	})
	assert.NoError(t, err)
	return result.GetId()
}

func testCreateCollection(t *testing.T, ctx context.Context, appId string, schemaId string, conn *grpc.ClientConn) string {
	collections := pb.NewCollectionClient(conn)
	result, err := collections.Create(ctx, &pb.CreateCollectionRequest{
		AppId:    appId,
		SchemaId: schemaId,
		Policy: &pb.Policy{
			DataOwner:    0.3,
			DataProvider: 0.7,
		},
	})
	assert.NoError(t, err)
	return result.GetCollectionId()
}

func testCreateUserAccount(t *testing.T, ctx context.Context, conn *grpc.ClientConn, index int) string {
	accounts := ablclient.NewClient(conn)

	priv, err := crypto.GenerateKey()
	assert.NoError(t, err)

	walletAddress := crypto.PubkeyToAddress(priv.PublicKey)
	password := fmt.Sprintf("password%d", index)

	session, err := accounts.Create(ctx, walletAddress, password)
	assert.NoError(t, err)
	return session.AccountId.Hex()
}

func testCreateUserAccountParallel(t *testing.T, ctx context.Context, conn *grpc.ClientConn) (userIds [numberOfUsers]string) {
	var accCreationWait sync.WaitGroup
	accCreationWait.Add(numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		go func(index int) {
			defer accCreationWait.Done()
			userIds[index] = testCreateUserAccount(t, ctx, conn, index)
			log.Printf("Created user %d : %s\n", i, userIds[index])
		}(i)
	}
	accCreationWait.Wait()
	return
}

func TestStoreBundleData(t *testing.T) {

}

func TestSellData(t *testing.T) {

}

func TestPurchaseData(t *testing.T) {

}

func TestPnc(t *testing.T) {
	main()
}
