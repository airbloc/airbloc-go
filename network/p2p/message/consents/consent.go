package consents

import (
	ablTypes "github.com/airbloc/airbloc-go/bind/types"

	json "github.com/json-iterator/go"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var (
	_ noise.Message = (*ConsentRequest)(nil)
	_ noise.Message = (*ConsentResponse)(nil)
)

type ConsentRequest struct {
	MessageID   uuid.UUID              `json:"message_id"`
	ConsentData []ablTypes.ConsentData `json:"consent_data"`
}

func (ConsentRequest) Read(reader payload.Reader) (noise.Message, error) {
	var req ConsentRequest
	err := json.NewDecoder(reader).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode consent request message")
	}
	return req, nil
}

func (req ConsentRequest) Write() []byte {
	reqBytes, _ := json.Marshal(req)
	return reqBytes
}

type ConsentResponse struct {
	MessageID uuid.UUID     `json:"message_id"`
	TxHash    common.Hash   `json:"tx_hash"`
	Signature hexutil.Bytes `json:"signature"`
}

func (ConsentResponse) Read(reader payload.Reader) (noise.Message, error) {
	var resp ConsentResponse
	err := json.NewDecoder(reader).Decode(&resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode consent response message")
	}
	return resp, nil
}

func (resp ConsentResponse) Write() []byte {
	respBytes, _ := json.Marshal(resp)
	return respBytes
}
