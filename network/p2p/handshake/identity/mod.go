package identity

import (
	"bytes"
	"crypto/rand"
	"sync"
	"time"

	"github.com/airbloc/airbloc-go/account"

	"github.com/klaytn/klaytn/crypto"
	"github.com/rs/zerolog/log"

	"github.com/pkg/errors"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/protocol"
)

const (
	KeyAddress = "airbloc.node.address"
)

type block struct {
	opcodeHandshakeRequest  noise.Opcode
	opcodeHandshakeResponse noise.Opcode
	timeoutDuration         time.Duration

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
	b.opcodeHandshakeRequest = noise.RegisterMessage(noise.NextAvailableOpcode(), (*HandshakeRequest)(nil))
	b.opcodeHandshakeResponse = noise.RegisterMessage(noise.NextAvailableOpcode(), (*HandshakeResponse)(nil))
}

func (b *block) sendHandshakeRequest(peer *noise.Peer) (HandshakeResponse, error) {
	var verifyPayload [32]byte
	_, err := rand.Read(verifyPayload[:])
	if err != nil {
		return HandshakeResponse{}, errors.Wrap(errors.Wrap(protocol.DisconnectPeer, err.Error()), "failed to generate request payload")
	}

	req := HandshakeRequest{Payload: verifyPayload[:]}

	err = peer.SendMessage(req)
	if err != nil {
		return HandshakeResponse{}, errors.Wrap(errors.Wrap(protocol.DisconnectPeer, err.Error()), "failed to send verify request to peer")
	}

	var (
		res HandshakeResponse
		ok  bool
	)

	select {
	case <-time.After(b.timeoutDuration):
		return HandshakeResponse{}, errors.Wrap(protocol.DisconnectPeer, "timed out receiving handshake response")
	case msg := <-peer.Receive(b.opcodeHandshakeResponse):
		res, ok = msg.(HandshakeResponse)
		if !ok {
			return HandshakeResponse{}, errors.Wrap(protocol.DisconnectPeer, "did not get a handshake response back")
		}
	}

	pubKey, err := crypto.SigToPub(verifyPayload[:], res.Signature)
	if err != nil {
		return HandshakeResponse{}, errors.Wrap(errors.Wrap(protocol.DisconnectPeer, err.Error()), "failed to derive pubkey from signature")
	}

	if bytes.Compare(crypto.FromECDSAPub(pubKey), crypto.FromECDSAPub(res.PubKey)) != 0 {
		return HandshakeResponse{}, errors.Wrap(protocol.DisconnectPeer, "failed to verify signature")
	}
	return res, nil
}

func (b *block) handleHandshakeRequest(peer *noise.Peer) error {
	var (
		req HandshakeRequest
		ok  bool
	)

	select {
	case <-time.After(b.timeoutDuration):
		return errors.Wrap(protocol.DisconnectPeer, "timed out receiving handshake request")
	case msg := <-peer.Receive(b.opcodeHandshakeRequest):
		req, ok = msg.(HandshakeRequest)
		if !ok {
			return errors.Wrap(protocol.DisconnectPeer, "did not get a handshake request")
		}
	}

	signature, err := b.nodeAccount.SignMessage(req.Payload)
	if err != nil {
		return errors.Wrap(errors.Wrap(protocol.DisconnectPeer, err.Error()), "failed to sign payload")
	}

	pubKey := b.nodeAccount.PublicKey()
	err = peer.SendMessage(HandshakeResponse{
		PubKey:    &pubKey,
		Signature: signature,
	})

	resp := HandshakeResponse{
		PubKey:    &pubKey,
		Signature: signature,
	}

	err = peer.SendMessage(resp)
	if err != nil {
		return errors.Wrap(errors.Wrap(protocol.DisconnectPeer, err.Error()), "failed to send response to peer")
	}

	return nil
}

func (b *block) OnBegin(p *protocol.Protocol, peer *noise.Peer) error {
	var (
		res       HandshakeResponse
		errChan   = make(chan error, 2)
		waitGroup = new(sync.WaitGroup)
	)

	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		var err error
		res, err = b.sendHandshakeRequest(peer)
		errChan <- err
	}()

	go func() {
		defer waitGroup.Done()
		err := b.handleHandshakeRequest(peer)
		if err != nil {
			log.Error().Err(err)
		}
		errChan <- err
	}()

	waitGroup.Wait()

	for len(errChan) > 0 {
		if err := <-errChan; err != nil {
			return err
		}
	}

	peer.Set(KeyAddress, crypto.PubkeyToAddress(*res.PubKey))

	log.Debug().
		Hex("peer_ecdsa_public_key", crypto.FromECDSAPub(res.PubKey)).
		Msg("Successfully exchange pubkey with our peer.")

	return nil
}

func (b *block) OnEnd(p *protocol.Protocol, peer *noise.Peer) error {
	peer.Delete(KeyAddress)
	return nil
}
