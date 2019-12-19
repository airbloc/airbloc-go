package identity

import (
	"crypto/ecdsa"

	"github.com/klaytn/klaytn/crypto"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	"github.com/pkg/errors"
)

var (
	_ noise.Message = (*HandshakeRequest)(nil)
	_ noise.Message = (*HandshakeResponse)(nil)
)

type HandshakeRequest struct {
	Payload []byte
}

func (HandshakeRequest) Read(reader payload.Reader) (noise.Message, error) {
	verifyPayload, err := reader.ReadBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read payload")
	}
	return HandshakeRequest{Payload: verifyPayload}, nil
}

func (m HandshakeRequest) Write() []byte {
	return payload.NewWriter(nil).WriteBytes(m.Payload).Bytes()
}

type HandshakeResponse struct {
	PubKey    *ecdsa.PublicKey
	Signature []byte
}

func (HandshakeResponse) Read(reader payload.Reader) (noise.Message, error) {
	pubKeyBytes, err := reader.ReadBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read address")
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal pubkey")
	}

	signature, err := reader.ReadBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read signature")
	}

	return HandshakeResponse{
		PubKey:    pubKey,
		Signature: signature,
	}, nil
}

func (m HandshakeResponse) Write() []byte {
	return payload.NewWriter(nil).
		WriteBytes(crypto.FromECDSAPub(m.PubKey)).
		WriteBytes(m.Signature).Bytes()
}
