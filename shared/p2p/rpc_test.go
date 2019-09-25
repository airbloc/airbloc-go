package p2p

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func setupTestRPC(t *testing.T) (aliceAddr, bobAddr common.Address, alice, bob RPC, teardown func()) {
	keys, servers, teardown := setupTestPeers(t, 2)
	aliceAddr, bobAddr = keys[0].EthereumAddress, keys[1].EthereumAddress

	for _, server := range servers {
		require.NoError(t, StartNameServer(server))
	}
	alice, bob = NewRPC(servers[0]), NewRPC(servers[1])
	return
}

func TestRpc_Invoke(t *testing.T) {
	aliceAddr, bobAddr, alice, bob, teardown := setupTestRPC(t)
	defer teardown()

	bob.Handle("ping", pingMsg, pongMsg, func(_ context.Context, from SenderInfo, req proto.Message) (proto.Message, error) {
		require.Equal(t, aliceAddr, from.Addr)
		return pongMsg, nil
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reply, err := alice.Invoke(ctx, bobAddr, "ping", pingMsg, pongMsg)
	require.NoError(t, err)
	require.Equal(t, reply, pongMsg)
}

func TestRpc_Invoke_WithError(t *testing.T) {
	_, bobAddr, alice, bob, teardown := setupTestRPC(t)
	defer teardown()

	errFromRemote := errors.New("some error")
	bob.Handle("faultyMethod", pingMsg, pongMsg, func(_ context.Context, from SenderInfo, req proto.Message) (proto.Message, error) {
		return nil, errFromRemote
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := alice.Invoke(ctx, bobAddr, "faultyMethod", pingMsg, pongMsg)
	require.Error(t, err)
	require.EqualError(t, errors.Cause(err), errFromRemote.Error())
}
