package merkle

import (
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/loomnetwork/mamamerkle"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestNewMainTree(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// prepare
	sInput := make(map[types.ID][]types.RowId)
	for i := uint64(0); i < 100000; i++ {
		userId := types.UintToID(i)
		for j := uint32(0); j < 10; j++ {
			sInput[userId] = append(sInput[userId], types.UintToRowId(j))
		}
	}

	//s := time.Now()
	sTree, err := NewMainTree(sInput)
	require.NoError(t, err)

	mInput := make(map[uint64][]byte)
	for _, v := range sTree.leaves {
		mInput[v.userId.Uint64()] = v.root.Bytes()
	}

	//m := time.Now()
	mTree, err := mamamerkle.NewSparseMerkleTree(64, mInput)
	require.NoError(t, err)

	require.Equal(t, mTree.Root(), sTree.Root().Bytes())

	proof, err := sTree.GenerateProof(types.UintToRowId(2), types.UintToID(10))
	require.NoError(t, err)

	require.True(t, VerifyMainProof(types.UintToRowId(2), types.UintToID(10), sTree.Root().Bytes(), proof))
}
