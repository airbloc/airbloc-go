package identity

import (
	"github.com/klaytn/klaytn/common"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	"github.com/pkg/errors"
)

var (
	_ noise.Message = (*Ping)(nil)
)

type Ping struct {
	Address common.Address
}

func (Ping) Read(reader payload.Reader) (noise.Message, error) {
	address, err := reader.ReadBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read address")
	}
	return Ping{Address: common.BytesToAddress(address)}, nil
}

func (p Ping) Write() []byte {
	return payload.NewWriter(nil).WriteBytes(p.Address.Bytes()).Bytes()
}
