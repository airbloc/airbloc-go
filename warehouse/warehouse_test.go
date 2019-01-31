package warehouse

import (
	"context"
	"encoding/json"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"log"
	"testing"
)

type T struct {
	*testing.T
	ctx  context.Context
	conn *grpc.ClientConn
}

func TestDataWarehouse_GetBundleInfo(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	require.NoError(t, err)
	defer func() { _ = conn.Close() }()

	c := &T{
		T:    t,
		ctx:  ctx,
		conn: conn,
	}

	warehouse := pb.NewWarehouseClient(c.conn)
	res, err := warehouse.GetBundleInfo(c.ctx, &pb.BundleInfoRequest{BundleId: "8545dd6dd9850c26"})
	require.NoError(t, err)

	d, err := json.MarshalIndent(res, "", "    ")
	require.NoError(t, err)
	log.Println(string(d))
}

func TestDataWarehouse_GetUserInfo(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	require.NoError(t, err)
	defer func() { _ = conn.Close() }()

	c := &T{
		T:    t,
		ctx:  ctx,
		conn: conn,
	}

	warehouse := pb.NewWarehouseClient(c.conn)
	res, err := warehouse.GetUserDataIds(c.ctx, &pb.UserDataIdsRequest{UserId: "2b87eacb22341daa"})
	require.NoError(t, err)

	d, err := json.MarshalIndent(res, "", "    ")
	require.NoError(t, err)
	log.Println(string(d))
}
