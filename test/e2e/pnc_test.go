package e2e

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/json-iterator/go"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const numberOfUsers = 3
const numberOfCollections = 5
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

type T struct {
	*testing.T
	ctx    context.Context
	conn   *grpc.ClientConn
	config *testConfig
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func generateUniqueName() string {
	return time.Now().Format("20060102-150405")
}

func (t *T) testCreateApp() string {
	apps := pb.NewAppsClient(t.conn)
	result, err := apps.Register(t.ctx, &pb.RegisterRequest{
		Name: fmt.Sprintf("app-test-%s-%d", generateUniqueName(), rand.Int()),
	})
	require.NoError(t, err)
	return result.GetAppId()
}

func (t *T) testCreateSchema() string {
	schemas := pb.NewSchemaClient(t.conn)
	result, err := schemas.Create(t.ctx, &pb.CreateSchemaRequest{
		Name:   fmt.Sprintf("data-test-%s-%d", generateUniqueName(), rand.Int()),
		Schema: testSchema,
	})
	require.NoError(t, err)
	return result.GetId()
}

func (t *T) testCreateCollection(appId string, schemaId string) string {
	collections := pb.NewCollectionClient(t.conn)
	result, err := collections.Create(t.ctx, &pb.CreateCollectionRequest{
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

func (t *T) testCreateUserAccount(index int) string {
	dauth := pb.NewDAuthClient(t.conn)
	response, err := dauth.SignIn(t.ctx, &pb.SignInRequest{
		Identity:     fmt.Sprintf("test-user-%s-%d@airbloc.org", generateUniqueName(), index),
		UserDelegate: t.config.UserDelegateAddress.Hex(),
	})
	require.NoError(t, err)
	return response.GetAccountId()
}

func (t *T) testStoreBundleData(userIds [numberOfUsers]string, collectionId string) []*pb.StoreResult {
	warehouse := pb.NewWarehouseClient(t.conn)
	storeResults := make([]*pb.StoreResult, numberOfBundles)
	for n := 0; n < numberOfBundles; n++ {
		log.Println("Creating Bundle #", n+1)
		stream, err := warehouse.StoreBundle(t.ctx)
		require.NoError(t, err)

		for index, userId := range userIds {
			rawData := &pb.RawDataRequest{
				CollectionId: collectionId,
				UserId:       userId,
				Payload:      fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userId, index),
				CollectedAt:  types.Time{Time: time.Now()}.Timestamp(),
			}
			require.NoError(t, stream.Send(rawData), "datum", rawData.String())
		}

		for index, userId := range userIds {
			rawData := &pb.RawDataRequest{
				CollectionId: collectionId,
				UserId:       userId,
				Payload:      fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userId, index+1),
				CollectedAt:  types.Time{Time: time.Now()}.Timestamp(),
			}
			require.NoError(t, stream.Send(rawData), "datum", rawData.String())
		}

		storeResults[n], err = stream.CloseAndRecv()
		require.NoError(t, err)

		log.Println("Stored URI:", storeResults[n].Uri)
		log.Println("Stored Data Count:", storeResults[n].DataCount)
		log.Println("Bundle ID:", storeResults[n].BundleId)
	}
	return storeResults
}

func TestPnc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http.Handle("/", http.FileServer(http.Dir("../../local/warehouse")))
	testServer := &http.Server{Addr: ":9126"}
	defer testServer.Shutdown(ctx)

	go func() {
		require.NoError(t, testServer.ListenAndServe())
	}()

	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	c := &T{
		T:    t,
		ctx:  ctx,
		conn: conn,
	}
	c.loadConfig()

	appId := c.testCreateApp()
	log.Println("Created App ID:", appId)

	schemaId := c.testCreateSchema()
	log.Printf("Created Schema ID: %s\n", schemaId)

	var collectionIds [numberOfCollections]string
	for i := 0; i < numberOfCollections; i++ {
		collectionIds[i] = c.testCreateCollection(appId, schemaId)
		log.Printf("Created Collection: %s\n", collectionIds[i])
	}

	// create 10 accounts
	var userIds [numberOfUsers]string
	for i := 0; i < numberOfUsers; i++ {
		//userIds[i] = c.testCreateUserAccount(i)
		b := make([]byte, 8)
		rand.Read(b)
		userIds[i] = hex.EncodeToString(b)
		log.Printf("Created user %d : %s\n", i, userIds[i])
	}

	// DAuth: allow data collection of these users
	//for _, userId := range userIds {
	//	for _, collectionId := range collectionIds {
	//		testDAuth(conn, collectionId, userId, true)
	//	}
	//	log.Printf("Allowed collection of user %s's data\n", userId)
	//}

	var storeResults [len(collectionIds)][]*pb.StoreResult
	for i, collectionId := range collectionIds {
		storeResults[i] = c.testStoreBundleData(userIds, collectionId)
	}

	warehouse := pb.NewWarehouseClient(c.conn)
	listRes, err := warehouse.ListBundle(c.ctx, &pb.ListBundleRequest{ProviderId: appId})
	require.NoError(t, err)

	log.Println(listRes.Bundles[0])
	// exchange: Test exchanging uploaded data
	//log.Println("Start exchanging", storeResults[0][0].DataCount, "data")
	//
	//user := pb.NewUserClient(c.conn)
	//userData, err := user.GetData(c.ctx, &pb.DataRequest{
	//	UserId: userIds[0],
	//	From:   0,
	//})
	//require.NoError(t, err)
	//
	//d, _ := json.MarshalIndent(userData, "", "    ")
	//log.Println(string(d))
}
