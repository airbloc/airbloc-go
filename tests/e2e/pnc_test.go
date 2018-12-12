package e2e

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/airbloc/airbloc-go/ablclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

const numberOfUsers = 10
const testSchema = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "http://airbloc.io/testdata.schema.json",
  "title": "Test Data",
  "type": "object",
  "properties": {
	"name": {
      "description": "Who are you lol",
      "type": "string"
	},
    "age": {
      "description": "Age is Age",
      "type": "integer"
    }
  },
  "required": [ "name" ]
}`

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

func testCreateUserAccount(t *testing.T, conn *grpc.ClientConn, index int) string {
	accounts := ablclient.NewClient(conn)

	priv, err := crypto.GenerateKey()
	assert.NoError(t, err)

	walletAddress := crypto.PubkeyToAddress(priv.PublicKey)
	password := fmt.Sprintf("password%d", index)

	session, err := accounts.Create(walletAddress, password)
	assert.NoError(t, err)
	return session.AccountId.Hex()
}

func testCreateUserAccountParallel(t *testing.T, conn *grpc.ClientConn) (userIds [numberOfUsers]string) {
	var accCreationWait sync.WaitGroup
	accCreationWait.Add(numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		go func(index int) {
			defer accCreationWait.Done()
			userIds[index] = testCreateUserAccount(t, conn, index)
			log.Printf("Created user %d : %s\n", i, userIds[index])
		}(i)
	}
	accCreationWait.Wait()
	return
}

func TestPnc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer func() { _ = conn.Close() }()

	appId := testCreateApp(t, ctx, conn)
	log.Println("Created App ID:", appId)

	schemaId := testCreateSchema(t, ctx, conn)
	log.Printf("Created Schema ID: %s\n", schemaId)

	collectionId := testCreateCollection(t, ctx, appId, schemaId, conn)
	log.Printf("Created Collection: %s\n", collectionId)

	// create 10 accounts
	var userIds [numberOfUsers]string
	for i := 0; i < numberOfUsers; i++ {
		userIds[i] = testCreateUserAccount(t, conn, i)
	}

	// make two bundles!
	warehouse := pb.NewWarehouseClient(conn)
	for n := 0; n < 2; n++ {
		log.Println("Creating Bundle #", n)
		stream, err := warehouse.StoreBundle(ctx)
		assert.NoError(t, err)

		for i := 0; i < numberOfUsers; i++ {
			rawData := &pb.RawDataRequest{
				CollectionId: collectionId,
				OwnerId:      userIds[i],
				Payload:      fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userIds[i], i),
			}
			assert.NoError(t, stream.Send(rawData), "datum", rawData.String())
		}

		result, err := stream.CloseAndRecv()
		assert.NoError(t, err)

		log.Println("Stored URI:", result.Uri)
		log.Println("Stored Data Count:", result.DataCount)
		log.Println("Bundle ID:", result.BundleId)
	}
}
