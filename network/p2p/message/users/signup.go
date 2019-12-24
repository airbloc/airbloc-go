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
	_ noise.Message = (*SignUpRequest)(nil)
	_ noise.Message = (*SignUpResponse)(nil)
)

type SignUpRequest struct {
	MessageID    uuid.UUID   `json:"message_id"` // Ignore when write message
	IdentityHash common.Hash `json:"identity_hash"`
}

func (SignUpRequest) Read(reader payload.Reader) (noise.Message, error) {
	var req SignUpRequest
	err := json.NewDecoder(reader).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode signup request message")
	}
	return req, nil
}

func (req SignUpRequest) Write() []byte {
	req.MessageID = uuid.NewV4()
	reqBytes, _ := json.Marshal(req)
	return reqBytes
}

type SignUpResponse struct {
	MessageID uuid.UUID     `json:"message_id"`
	TxHash    common.Hash   `json:"tx_hash"`
	Signature hexutil.Bytes `json:"signature"`
}

func (SignUpResponse) Read(reader payload.Reader) (noise.Message, error) {
	var resp SignUpResponse
	err := json.NewDecoder(reader).Decode(&resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode signup response message")
	}
	return resp, nil
}

func (resp SignUpResponse) Write() []byte {
	respBytes, _ := json.Marshal(resp)
	return respBytes
}
