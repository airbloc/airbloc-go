package user

import (
	"context"
	"encoding/json"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestManager_GetData(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	user := pb.NewUserClient(conn)
	resp, err := user.GetData(ctx, &pb.DataRequest{
		UserId: "b64e0af23bb9ae56",
		From:   1550827549155,
	})
	require.NoError(t, err)

	d, _ := json.MarshalIndent(resp, "", "    ")
	log.Println(string(d))
}

func TestManager_GetDataIds(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	user := pb.NewUserClient(conn)
	resp, err := user.GetDataIds(ctx, &pb.DataIdRequest{
		UserId: "b64e0af23bb9ae56",
	})
	require.NoError(t, err)

	d, _ := json.MarshalIndent(resp, "", "    ")
	log.Println(string(d))
}
