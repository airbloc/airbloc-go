package merkle

import (
	"math/big"

	"bytes"

	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type node struct {
	cIdx *big.Int // current 'tree' index
	nIdx int      // next 'node' index
	hash common.Hash
}

// true - left
// false - right
func (n node) Direction() bool {
	return bigModInt(n.cIdx, 2).Cmp(big.NewInt(0)) == 0
}

func (n node) Cmp(y *big.Int) int {
	return n.cIdx.Cmp(y)
}

func (n node) CmpBool(y *big.Int) bool {
	return n.Cmp(y) == 0
}

func (n node) Index() *big.Int {
	return n.cIdx
}

func (n node) Hash() common.Hash {
	return n.hash
}

func (n node) Key(level int, blockNum *big.Int) []byte {
	buf := new(bytes.Buffer)
	buf.Write(blockNum.Bytes())
	buf.WriteRune(':')
	buf.Write([]byte("txs"))
	buf.WriteRune(':')
	buf.Write([]byte(strconv.Itoa(level)))
	buf.WriteRune(':')
	buf.Write(n.cIdx.Bytes())
	return buf.Bytes()
}

type level []*node

// binary search lel
func (lv *level) Search(key common.Hash) (int, error) {
	return lv.search(0, len(*lv), key.Big())
}

func (lv *level) search(x, y int, key *big.Int) (int, error) {
	if y-x == 1 {
		if (*lv)[x].CmpBool(key) {
			return x, nil
		}
		return 0, errors.New("no node found")
	}
	crit := (*lv)[(x+y)/2]
	switch crit.Cmp(key) {
	case 1: // find left
		return lv.search(x, (x+y)/2, key)
	case 0: // equals!
		return (x + y) / 2, nil
	case -1: // find right
		return lv.search((x+y)/2, y, key)
	default:
		panic("LOL")
	}
}
