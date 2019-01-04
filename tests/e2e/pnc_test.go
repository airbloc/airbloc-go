package e2e

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

const numberOfUsers = 10
const numberOfBundles = 2
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

func generateUniqueName() string {
	return time.Now().Format("20060102-150405")
}

func testCreateApp(t *testing.T, ctx context.Context, conn *grpc.ClientConn) string {
	apps := pb.NewAppsClient(conn)
	result, err := apps.Register(ctx, &pb.RegisterRequest{
		Name: fmt.Sprintf("app-test-%s", generateUniqueName()),
	})
	require.NoError(t, err)
	return result.GetAppId()
}

func testCreateSchema(t *testing.T, ctx context.Context, conn *grpc.ClientConn) string {
	schemas := pb.NewSchemaClient(conn)
	result, err := schemas.Create(ctx, &pb.CreateSchemaRequest{
		Name:   fmt.Sprintf("data-test-%s", generateUniqueName()),
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

func testCreateUserAccount(t *testing.T, ctx context.Context, conn *grpc.ClientConn, config *testConfig, index int) string {
	dauth := pb.NewDAuthClient(conn)
	response, err := dauth.SignIn(ctx, &pb.SignInRequest{
		Identity:     fmt.Sprintf("test-user-%s-%d@airbloc.org", generateUniqueName(), index),
		UserDelegate: config.UserDelegateAddress.Hex(),
	})
	require.NoError(t, err)
	return response.GetAccountId()
}

// func testCreateUserAccountParallel(t *testing.T, ctx context.Context, conn *grpc.ClientConn) (userIds [numberOfUsers]string) {
// 	var accCreationWait sync.WaitGroup
// 	accCreationWait.Add(numberOfUsers)
// 	for i := 0; i < numberOfUsers; i++ {
// 		go func(index int) {
// 			defer accCreationWait.Done()
// 			userIds[index] = testCreateUserAccount(t, ctx, conn, index)
// 			log.Printf("Created user %d : %s\n", i, userIds[index])
// 		}(i)
// 	}
// 	accCreationWait.Wait()
// 	return
// }

func TestPnc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := loadConfig(t)

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
		userIds[i] = testCreateUserAccount(t, ctx, conn, config, i)
		log.Printf("Created user %d : %s\n", i, userIds[i])
	}

	// DAuth: allow data collection of these users
	for _, userId := range userIds {
		testDAuth(conn, collectionId, userId, true)
		log.Printf("Allowed collection of user %s's data\n", userId)
	}

	// warehouse: store bundle data
	warehouse := pb.NewWarehouseClient(conn)
	storeResults := make([]*pb.StoreResult, numberOfBundles)
	bundles := make([][]string, numberOfBundles)
	for n := 0; n < numberOfBundles; n++ {
		log.Println("Creating Bundle #", n+1)
		stream, err := warehouse.StoreBundle(ctx)
		require.NoError(t, err)

		for index, userId := range userIds {
			rawData := &pb.RawDataRequest{
				CollectionId: collectionId,
				OwnerId:      userId,
				Payload:      fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userId, index),
			}
			require.NoError(t, stream.Send(rawData), "datum", rawData.String())
		}

		storeResults[n], err = stream.CloseAndRecv()
		require.NoError(t, err)

		log.Println("Stored URI:", storeResults[n].Uri)
		log.Println("Stored Data Count:", storeResults[n].DataCount)
		log.Println("Bundle ID:", storeResults[n].BundleId)

		// collectionId		bundleNumber		ownerId
		// deadbeefdeadbeef	0000000000000001	deadbeefdeadbeef
		bundles[n] = make([]string, numberOfUsers)
		for index, userId := range userIds {
			dataId, _ := toDataId(storeResults[n].BundleId, userId)
			bundles[n][index] = hex.EncodeToString(dataId[:])
		}
	}

	// exchange: Test exchanging uploaded data
	log.Println("Start exchanging", len(bundles[0]), "data")
	testExchange(t, ctx, conn, config, bundles[0])
}
