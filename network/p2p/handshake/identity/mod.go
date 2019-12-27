package identity

import (
	"time"

	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/account"

	"github.com/rs/zerolog/log"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/protocol"
)

const (
	KeyAddress = "airbloc.peer.address"
)

type block struct {
	opcodePing      noise.Opcode
	timeoutDuration time.Duration

	nodeAccount account.Account
}

func New(nodeAccount account.Account) *block {
	return &block{
		timeoutDuration: 10 * time.Second,
		nodeAccount:     nodeAccount,
	}
}

func (b *block) TimeoutAfter(timeoutDuration time.Duration) *block {
	b.timeoutDuration = timeoutDuration
	return b
}

func (b *block) OnRegister(p *protocol.Protocol, node *noise.Node) {
	b.opcodePing = noise.RegisterMessage(noise.NextAvailableOpcode(), (*Ping)(nil))
}

func (b *block) OnBegin(p *protocol.Protocol, peer *noise.Peer) error {
	err := peer.SendMessage(Ping{Address: b.nodeAccount.TxOpts().From})
	if err != nil {
		return errors.Wrap(errors.Wrap(protocol.DisconnectPeer, err.Error()), "failed to send ping message to peer")
	}

	var (
		resp Ping
		ok   bool
	)
	select {
	case <-time.After(b.timeoutDuration):
		return errors.Wrap(protocol.DisconnectPeer, "timed out receiving pong message")
	case msg := <-peer.Receive(b.opcodePing):
		resp, ok = msg.(Ping)
		if !ok {
			return errors.Wrap(protocol.DisconnectPeer, "did not get a pong message")
		}
	}

	peer.Set(KeyAddress, resp.Address)

	log.Debug().
		Str("peer_address", resp.Address.Hex()).
		Msg("Successfully exchange pubkey with our peer.")

	return nil
}

func (b *block) OnEnd(p *protocol.Protocol, peer *noise.Peer) error {
	peer.Delete(KeyAddress)
	return nil
}
