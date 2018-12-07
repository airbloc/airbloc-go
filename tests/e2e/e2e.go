package e2e

import (
	"fmt"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"google.golang.org/grpc"
	"log"
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

func main() {
	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer func() { _ = conn.Close() }()

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
		stream, err := warehouse.StoreBundle(ctx)
		if err != nil {
			log.Fatalf("Failed to open stream: %+v", err)
		}

		for i := 0; i < numberOfUsers; i++ {
			rawData := &pb.RawDataRequest{
				CollectionId: collectionId,
				OwnerId:      userIds[i],
				Payload:      fmt.Sprintf("{\"name\":\"%s\",\"age\":%d}", userIds[i], i),
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
