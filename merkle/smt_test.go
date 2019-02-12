package merkle

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/stretchr/testify/require"
	"log"
	"runtime"
	"testing"
	"time"
)

func TestNewMainTree(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	runtime.GOMAXPROCS(runtime.NumCPU())

	// prepare
	sInput := make(map[common.ID][]common.RowId)
	for i := uint64(0); i < 100000; i++ {
		userId := common.UintToID(i)
		for j := uint32(0); j < uint32(i%20); j++ {
			sInput[userId] = append(sInput[userId], common.UintToRowId(j))
		}
	}

	s := time.Now()
	sTree, err := NewMainTree(sInput)
	require.NoError(t, err)
	log.Println(time.Since(s).String())

	proof, err := sTree.GenerateProof(common.UintToRowId(2), common.UintToID(10))
	require.NoError(t, err)

	require.True(t, VerifyMainProof(common.UintToRowId(2), common.UintToID(10), sTree.Root().Bytes(), proof))
}
