package users

import (
	"encoding/json"

	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	uuid "github.com/satori/go.uuid"
)

type SignUpRequest struct {
	MessageID    uuid.UUID   `json:"message_id"` // Ignore when write message
	IdentityHash common.Hash `json:"identity_hash"`
}

func (SignUpRequest) Read(reader payload.Reader) (noise.Message, error) {
	var req SignUpRequest
	if err := json.NewDecoder(reader).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func (req SignUpRequest) Write() []byte {
	req.MessageID = uuid.NewV4()
	reqBytes, _ := json.Marshal(req)
	return reqBytes
}

type SignUpResponse struct {
	MessageID uuid.UUID `json:"message_id"`
	Tx        struct {
		Hash    common.Hash        `json:"hash"`
		Receipt *klayTypes.Receipt `json:"receipt"`
	} `json:"tx"`
}

func (SignUpResponse) Read(reader payload.Reader) (noise.Message, error) {
	var req SignUpResponse
	if err := json.NewDecoder(reader).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func (resp SignUpResponse) Write() []byte {
	reqBytes, _ := json.Marshal(resp)
	return reqBytes
}
