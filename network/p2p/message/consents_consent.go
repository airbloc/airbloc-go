package message

import (
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
)

type ConsentRequest struct {
}

func (ConsentRequest) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (req ConsentRequest) Write() []byte {
	return nil
}

type ConsentResponse struct {
}

func (ConsentResponse) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (resp ConsentResponse) Write() []byte {
	return nil
}
