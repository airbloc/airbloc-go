package merkle

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestNewMainTree(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// prepare
	sInput := make(map[common.ID][]common.RowId)
	for i := uint64(0); i < 10000000; i++ {
		userId := common.UintToID(i)
		for j := uint32(0); j < rand.Uint32()%20; j++ {
			sInput[userId] = append(sInput[userId], common.UintToRowId(j))
		}
	}

	s := time.Now()
	sTree, err := NewMainTree(sInput)
	require.NoError(t, err)
	log.Println(time.Since(s).String())

	log.Println(sTree.root.Hex())
}
