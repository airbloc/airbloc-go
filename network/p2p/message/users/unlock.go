package users

import (
	json "github.com/json-iterator/go"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var (
	_ noise.Message = (*UnlockRequest)(nil)
	_ noise.Message = (*UnlockResponse)(nil)
)

type UnlockRequest struct {
	MessageID        uuid.UUID      `json:"message_id"` // Ignore when write message
	IdentityPreimage common.Hash    `json:"identity_preimage"`
	NewOwner         common.Address `json:"new_owner"`
}

func (UnlockRequest) Read(reader payload.Reader) (noise.Message, error) {
	var req UnlockRequest
	err := json.NewDecoder(reader).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode unlock request message")
	}
	return req, nil
}

func (req UnlockRequest) Write() []byte {
	reqBytes, _ := json.Marshal(req)
	return reqBytes
}

type UnlockResponse struct {
	MessageID uuid.UUID     `json:"message_id"`
	TxHash    common.Hash   `json:"tx_hash"`
	Signature hexutil.Bytes `json:"signature"`
}

func (UnlockResponse) Read(reader payload.Reader) (noise.Message, error) {
	var resp UnlockResponse
	err := json.NewDecoder(reader).Decode(&resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode unlock response message")
	}
	return nil, nil
}

func (resp UnlockResponse) Write() []byte {
	respBytes, _ := json.Marshal(resp)
	return respBytes
}
