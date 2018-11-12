package main

import (
	"context"
	"fmt"
	"log"

	"github.com/airbloc/airbloc-go/warehouse/api"
	"google.golang.org/grpc"
)

//go:generate ./build.sh

func init() {

}

func main() {
	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	warehouse := api.NewWarehouseClient(conn)
	stream, err := warehouse.StoreBundle(context.Background())
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		rawData := &api.RawDataRequest{
			Collection: "deadbeefdeadbeef",
			OwnerId:    "deadbeefdeadbeef",
			Payload:    fmt.Sprintf("{\"foo\":\"%d\"}", i),
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
