package e2e

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/airbloc/airbloc-go/ablclient"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
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

func testCreateApp(conn *grpc.ClientConn) string {
	apps := pb.NewAppsClient(conn)
	result, err := apps.Register(context.Background(), &pb.RegisterRequest{
		Name: fmt.Sprintf("app-test-%d", time.Now().Unix()),
	})
	if err != nil {
		log.Fatalf("Failed to create app: %+v", err)
	}
	return result.GetAppId()
}

func testCreateSchema(conn *grpc.ClientConn) string {
	schemas := pb.NewSchemaClient(conn)
	result, err := schemas.Create(context.Background(), &pb.CreateSchemaRequest{
		Name:   fmt.Sprintf("data-test-%d", time.Now().Unix()),
		Schema: testSchema,
	})
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}
	return result.GetId()
}

func testCreateCollection(appId string, schemaId string, conn *grpc.ClientConn) string {
	collections := pb.NewCollectionClient(conn)

	result, err := collections.Create(context.Background(), &pb.CreateCollectionRequest{
		AppId:    appId,
		SchemaId: schemaId,
		Policy: &pb.Policy{
			DataOwner:    0.3,
			DataProvider: 0.7,
		},
	})
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}
	return result.GetCollectionId()
}

func testCreateUserAccount(conn *grpc.ClientConn, index int) string {
	accounts := ablclient.NewClient(conn)

	priv, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to generate a private key").Error())
	}

	walletAddress := crypto.PubkeyToAddress(priv.PublicKey)
	password := fmt.Sprintf("password%d", index)

	session, err := accounts.Create(walletAddress, password)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to create account").Error())
	}
	return session.AccountId.String()
}

func testCreateUserAccountParallel(conn *grpc.ClientConn) (userIds [numberOfUsers]string) {
	var accCreationWait sync.WaitGroup
	accCreationWait.Add(numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		go func(index int) {
			userIds[index] = testCreateUserAccount(conn, index)
			accCreationWait.Done()
			log.Printf("Created user %d : %s\n", i, userIds[index])
		}(i)
	}
	accCreationWait.Wait()
	return
}

func main() {
	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	appId := testCreateApp(conn)
	log.Println("Created App ID:", appId)

	schemaId := testCreateSchema(conn)
	log.Printf("Created Schema ID: %s\n", schemaId)

	collectionId := testCreateCollection(appId, schemaId, conn)
	log.Printf("Created Collection: %s\n", collectionId)

	// create 10 accounts
	var userIds [numberOfUsers]string
	for i := 0; i < numberOfUsers; i++ {
		userIds[i] = testCreateUserAccount(conn, i)
	}

	// make two bundles!
	warehouse := pb.NewWarehouseClient(conn)
	for n := 0; n < 2; n++ {
		log.Println("Creating Bundle #", n)
		stream, err := warehouse.StoreBundle(context.Background())
		if err != nil {
			log.Fatalf("Failed to open stream: %+v", err)
		}

		for i := 0; i < numberOfUsers; i++ {
			rawData := &pb.RawDataRequest{
				Collection: collectionId,
				OwnerId:    userIds[i],
				Payload:    fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userIds[i], i),
			}
			if err := stream.Send(rawData); err != nil {
				log.Fatalf("Failed to send datum %v: %+v", rawData, err)
			}
		}
		result, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("Error occurred after RPC call: %+v", err)
		}

		log.Println("Stored URI:", result.Uri)
		log.Println("Stored Data Count:", result.DataCount)
		log.Println("Bundle ID:", result.BundleId)
	}
}
