package p2p

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const numOfLookupPeers = 5

func TestStartNameServer(t *testing.T) {
	_, servers, teardown := setupTestPeers(t, 1)
	defer teardown()

	err := StartNameServer(servers[0])
	require.NoError(t, err)
}

func TestLookup(t *testing.T) {
	keys, servers, teardown := setupTestPeers(t, numOfLookupPeers)
	defer teardown()

	for _, server := range servers {
		err := StartNameServer(server)
		require.NoError(t, err)
	}

	// every peer should identify each other
	for i := 0; i < numOfLookupPeers; i++ {
		for j := 0; j < numOfLookupPeers; j++ {
			if i == j {
				continue
			}
			alice := servers[i]
			bob := keys[j].EthereumAddress
			bobId := servers[j].getHost().ID()

			log.Printf("Alice: %s, Bob: %s, BobID: %s\n", keys[i].EthereumAddress.Hex(), bob.Hex(), bobId.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			bobIdFound, err := Lookup(ctx, alice, bob, 5*time.Second)
			require.NoError(t, err)

			require.Equal(t, bobId, bobIdFound)
		}
	}
}
