package merkle

import (
	"errors"
	"github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"sort"
)

type SubTree struct {
	leaves []common.RowId
	tree   [][]ethCommon.Hash
	root   ethCommon.Hash
}

func (st *SubTree) Leaves() []common.RowId {
	return st.leaves
}

func NewSubTree(input []common.RowId) (*SubTree, error) {
	// check input
	pow := new(big.Int).Lsh(big.NewInt(1), uint(32))
	if big.NewInt(int64(len(input))).Cmp(pow) > 0 {
		return nil, errors.New("too long input")
	}

	// sort
	sort.Slice(input, func(i, j int) bool {
		return input[i].Uint32() < input[j].Uint32()
	})

	// hashing
	base := make([]ethCommon.Hash, len(input))
	for index, elem := range input {
		base[index] = crypto.Keccak256Hash(elem[:])
	}

	// create tree structure
	st := &SubTree{
		leaves: input,
		tree:   [][]ethCommon.Hash{base},
	}

	// generate tree
	for i := 0; ; i++ {
		if len(st.tree[i]) == 1 {
			st.root = st.tree[i][0]
			break
		}

		var nextLvl []ethCommon.Hash
		for j := 0; j < len(st.tree[i]); j++ {
			if j%2 != 0 {
				continue
			}

			leaf := st.tree[i][j]
			coleaf := ethCommon.Hash{}

			if len(st.tree[i]) == j+1 {
				coleaf = leaf
			} else {
				coleaf = st.tree[i][j+1]
			}

			nextLvl = append(nextLvl, crypto.Keccak256Hash(leaf.Bytes(), coleaf.Bytes()))
		}
		st.tree = append(st.tree, nextLvl)
	}
	return st, nil
}
