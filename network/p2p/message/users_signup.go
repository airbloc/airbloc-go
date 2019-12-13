package message

import (
	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
)

type SignUpRequest struct {
	DataController common.Address
	IdentityHash   common.Hash
}

func (SignUpRequest) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (req SignUpRequest) Write() []byte {
	return nil
}

type SignUpResponse struct {
	TxHash common.Hash
}

func (SignUpResponse) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (resp SignUpResponse) Write() []byte {
	return nil
}
