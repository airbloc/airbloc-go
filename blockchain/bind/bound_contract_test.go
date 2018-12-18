package bind

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var testRevertData = `0x08c379a0
0000000000000000000000000000000000000000000000000000000000000020
0000000000000000000000000000000000000000000000000000000000000019
636f6c6c656374696f6e20646f6573206e6f7420657869737400000000000000`

func TestHasRevertError(t *testing.T) {
	revertData := hexutil.MustDecode(strings.Replace(testRevertData, "\n", "", -1))
	reason, ok := HasRevertError(revertData)

	assert.True(t, ok)
	assert.Equal(t, "collection does not exist", reason)

	normalData := hexutil.MustDecode("0x01")
	reason, ok = HasRevertError(normalData)
	assert.False(t, ok)
}
