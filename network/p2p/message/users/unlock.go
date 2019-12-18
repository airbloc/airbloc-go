package users

import (
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	uuid "github.com/satori/go.uuid"
)

type UnlockRequest struct {
	MessageID        uuid.UUID // Ignore when write message
	IdentityPreimage common.Hash
	NewOwner         common.Address
}

func (UnlockRequest) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (req UnlockRequest) Write() []byte {
	return nil
}

type UnlockResponse struct {
	MessageID uuid.UUID
	Tx        struct {
		Hash    common.Hash
		Receipt klayTypes.Receipt
	}
}

func (UnlockResponse) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (resp UnlockResponse) Write() []byte {
	return nil
}
