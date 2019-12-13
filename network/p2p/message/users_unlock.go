package message

import (
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
)

type UnlockRequest struct {
}

func (UnlockRequest) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (req UnlockRequest) Write() []byte {
	return nil
}

type UnlockResponse struct {
}

func (UnlockResponse) Read(reader payload.Reader) (noise.Message, error) {
	return nil, nil
}

func (resp UnlockResponse) Write() []byte {
	return nil
}
