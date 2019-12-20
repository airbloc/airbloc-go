package consents

import (
	ablTypes "github.com/airbloc/airbloc-go/bind/types"
	klayTypes "github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	uuid "github.com/satori/go.uuid"
)

var (
	_ noise.Message = (*ConsentRequest)(nil)
	_ noise.Message = (*ConsentResponse)(nil)
)

type ConsentRequest struct {
	MessageID   uuid.UUID
	ConsentData []ablTypes.ConsentData
}

func (ConsentRequest) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (req ConsentRequest) Write() []byte {
	//messageID := uuid.NewV4()
	return nil
}

type ConsentResponse struct {
	MessageID uuid.UUID
	Tx        struct {
		Hash    common.Hash
		Receipt klayTypes.Receipt
	}
}

func (ConsentResponse) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (resp ConsentResponse) Write() []byte {
	return nil
}
