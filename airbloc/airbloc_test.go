package airbloc

import (
	"context"
	"testing"

	"github.com/airbloc/airbloc-go/blockchain"

	"github.com/stretchr/testify/assert"
)

func TestAirbloc(t *testing.T) {
	//testKey, err := crypto.GenerateKey()
	//assert.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	abl, err := NewAirbloc(ctx, blockchain.KlaytnBaobab)

	assert.NoError(t, err)
	assert.NotNil(t, abl)
}
