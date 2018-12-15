package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/airbloc/airbloc-go/ablclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

const numberOfUsers = 20
const numberOfBundles = 5
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
	require.NoError(t, err)
	return result.GetAppId()
}

func testCreateSchema(t *testing.T, ctx context.Context, conn *grpc.ClientConn) string {
	schemas := pb.NewSchemaClient(conn)
	result, err := schemas.Create(ctx, &pb.CreateSchemaRequest{
		Name:   fmt.Sprintf("data-test-%d", time.Now().Unix()),
		Schema: testSchema,
	})
	require.NoError(t, err)
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
	require.NoError(t, err)
	return result.GetCollectionId()
}

func testCreateUserAccount(t *testing.T, conn *grpc.ClientConn, index int) string {
	accounts := ablclient.NewClient(conn)

	priv, err := crypto.GenerateKey()
	require.NoError(t, err)

	walletAddress := crypto.PubkeyToAddress(priv.PublicKey)
	password := fmt.Sprintf("password%d", index)

	session, err := accounts.Create(walletAddress, password)
	require.NoError(t, err)
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
	require.NoError(t, err)
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

	// warehouse: store bundle data
	warehouse := pb.NewWarehouseClient(conn)
	storeResults := make([]*pb.StoreResult, numberOfBundles)
	for n := 0; n < numberOfBundles; n++ {
		log.Println("Creating Bundle #", n)
		stream, err := warehouse.StoreBundle(ctx)
		require.NoError(t, err)

		for i := 0; i < numberOfUsers; i++ {
			rawData := &pb.RawDataRequest{
				ProviderId:   appId,
				CollectionId: collectionId,
				OwnerId:      userIds[i],
				Payload:      fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userIds[i], i),
			}
			require.NoError(t, stream.Send(rawData), "datum", rawData.String())
		}

		storeResults[n], err = stream.CloseAndRecv()
		require.NoError(t, err)

		log.Println("Stored URI:", storeResults[n].Uri)
		log.Println("Stored Data Count:", storeResults[n].DataCount)
		log.Println("Bundle ID:", storeResults[n].BundleId)
	}

	// exchange: trade bundle data
	deployments := make(map[string]common.Address)
	file, err := os.OpenFile("deployment.local.json", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(file).Decode(&deployments))

	/*
		Exchange TODOs:
		- Add common api
		    - get current account
		    - get current node status
		    - etc...
		- sign/args generation (or just make input of func to abi)
	*/
	//exchange := pb.NewExchangeClient(conn)
	//for _, bundleData := range storeResults {
	//	req := &pb.OrderRequest{
	//		Contract: &pb.Contract{
	//			Type: pb.Contract_SMART,
	//			SmartEscrow: &pb.SmartContract{
	//				Address: deployments["SimpleContract"].Hex(),
	//			},
	//		},
	//	}
	//}
}
