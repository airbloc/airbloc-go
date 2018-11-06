package p2p

import (
	"testing"

	"github.com/airbloc/airbloc-go/key"
	multiaddr "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

func launchClient() {

}

func launchBootstrap() {

}

func TestConnect(t *testing.T) {
	tcpRawAddr := "/ip4/127.0.0.1/tcp/3000"
	tcpAddr, err := multiaddr.NewMultiaddr(tcpRawAddr)
	assert.NoError(t, err)

	k, err := key.Generate()
	assert.NoError(t, err)

	for i := 0; i < 10; i++ {
		assert.NoError(t, Connect(tcpAddr, k))
	}
}
