package merkle

import (
	"encoding/hex"
	"github.com/airbloc/airbloc-go/common"
	"github.com/loomnetwork/mamamerkle"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestNewMainTree(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	runtime.GOMAXPROCS(runtime.NumCPU())

	// prepare
	sInput := make(map[common.ID][]common.RowId)
	for i := uint64(0); i < 1000000; i++ {
		userId := common.UintToID(i)
		for j := uint32(0); j < rand.Uint32()%20; j++ {
			sInput[userId] = append(sInput[userId], common.UintToRowId(j))
		}
	}

	s := time.Now()
	sTree, err := NewMainTree(sInput)
	require.NoError(t, err)
	log.Println(time.Since(s).String())

	mInput := make(map[uint64][]byte)
	for k, v := range sInput {
		mInput[k.Uint64()] = sTree.cache[uint64(len(v))].Root().Bytes()
	}

	m := time.Now()
	mTree, err := mamamerkle.NewSparseMerkleTree(64, mInput)
	require.NoError(t, err)
	log.Println(time.Since(m).String())

	log.Println(sTree.root.Hex())
	log.Println("0x" + hex.EncodeToString(mTree.Root()))

	log.Println(len(sTree.tree))
}
