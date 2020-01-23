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

func (req UnlockRequest) ID() uuid.UUID {
	return req.MessageID
}

func (req *UnlockRequest) SetID(id uuid.UUID) {
	req.MessageID = id
}

type UnlockResponse struct {
	MessageID uuid.UUID     `json:"message_id"`
	TxHash    common.Hash   `json:"tx_hash"`
	Sign      hexutil.Bytes `json:"signature"`
}

func (UnlockResponse) Read(reader payload.Reader) (noise.Message, error) {
	var resp UnlockResponse
	err := json.NewDecoder(reader).Decode(&resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode unlock response message")
	}
	return resp, nil
}

func (resp UnlockResponse) Write() []byte {
	respBytes, _ := json.Marshal(resp)
	return respBytes
}

func (resp UnlockResponse) ID() uuid.UUID {
	return resp.MessageID
}

func (resp *UnlockResponse) SetID(id uuid.UUID) {
	resp.MessageID = id
}

func (resp UnlockResponse) Signature() hexutil.Bytes {
	return resp.Sign
}

func (resp *UnlockResponse) SetSignature(sign hexutil.Bytes) {
	resp.Sign = sign
}
