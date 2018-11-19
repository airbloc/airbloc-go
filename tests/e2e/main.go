package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/airbloc/airbloc-go/account"
	collectionApi "github.com/airbloc/airbloc-go/collections/api"
	schemaApi "github.com/airbloc/airbloc-go/schemas/api"
	warehouseApi "github.com/airbloc/airbloc-go/warehouse/api"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

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

func testCreateSchema(conn *grpc.ClientConn) string {
	schemas := schemaApi.NewSchemaClient(conn)
	result, err := schemas.Create(context.Background(), &schemaApi.CreateSchemaRequest{
		Name:   fmt.Sprintf("data-test-%d", time.Now().Unix()),
		Schema: testSchema,
	})
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}
	return result.GetId()
}

func testCreateCollection(appId string, schemaId string, conn *grpc.ClientConn) string {
	collections := collectionApi.NewCollectionClient(conn)

	result, err := collections.Create(context.Background(), &collectionApi.CreateCollectionRequest{
		AppId:    appId,
		SchemaId: schemaId,
		Policy: &collectionApi.Policy{
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
	accounts := account.NewClient(conn)

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

func main() {
	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	appId := "786d674f9f07fa21"

	schemaId := testCreateSchema(conn)
	log.Printf("Created Schema ID: %s\n", schemaId)

	collectionId := testCreateCollection(appId, schemaId, conn)
	log.Printf("Created Collection: %s\n", collectionId)

	warehouse := warehouseApi.NewWarehouseClient(conn)
	stream, err := warehouse.StoreBundle(context.Background())
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		userId := testCreateUserAccount(conn, i)
		log.Printf("Created user %d : %s\n", i, userId)

		rawData := &warehouseApi.RawDataRequest{
			Collection: collectionId,
			OwnerId:    userId,
			Payload:    fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userId, i),
		}
		if err := stream.Send(rawData); err != nil {
			log.Fatalf("Failed to send datum %v: %v", rawData, err)
		}
	}

	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error occurred after RPC call: %v", err)
	}

	log.Printf("Stored URI : %s\n", result.Uri)
	log.Printf("Stored Data Count : %d\n", result.DataCount)
	log.Printf("Bundle ID : %s\n", result.BundleId)
}
