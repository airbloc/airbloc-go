package merkle

import (
	"github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestNewSubTree(t *testing.T) {
	for i := 0; i < 5; i++ {
		var sInput []common.RowId
		for j := uint32(0); j < rand.Uint32()%500; j++ {
			sInput = append(sInput, common.UintToRowId(j))
		}

		lvl := make([]ethCommon.Hash, len(sInput))
		for index, rowId := range sInput {
			lvl[index] = crypto.Keccak256Hash(rowId[:])
		}

		for {
			if len(lvl) == 1 {
				break
			}

			var nlvl []ethCommon.Hash
			for index := 0; index < len(lvl); index++ {
				if index%2 != 0 {
					continue
				}

				if len(lvl) == index+1 {
					nlvl = append(nlvl, crypto.Keccak256Hash(lvl[index].Bytes(), lvl[index].Bytes()))
				} else {
					nlvl = append(nlvl, crypto.Keccak256Hash(lvl[index].Bytes(), lvl[index+1].Bytes()))
				}
			}

			lvl = nlvl
		}

		st, err := NewSubTree(sInput)
		require.NoError(t, err)
		assert.Equal(t, lvl[0], st.root)
	}
}

func TestSubTree_GenerateProof(t *testing.T) {
	var sInput []common.RowId
	for i := uint32(0); i < 500; i++ {
		sInput = append(sInput, common.UintToRowId(i))
	}
	st, err := NewSubTree(sInput)
	require.NoError(t, err)

	proof, err := st.GenerateProof(sInput[250])
	require.NoError(t, err)

	res, err := st.Verify(sInput[250], proof)
	require.NoError(t, err)

	assert.True(t, res)
}
