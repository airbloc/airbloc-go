package merkle

import (
	"bytes"
	"errors"
	"github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"sort"
)

const (
	Left        = 0x00
	Right       = 0x01
	ProofLength = ethCommon.HashLength + 1
)

type SubTree struct {
	leaves []common.RowId
	tree   [][]ethCommon.Hash
	root   ethCommon.Hash
}

func (st *SubTree) Leaves() []common.RowId {
	return st.leaves
}

func (st *SubTree) Root() ethCommon.Hash {
	return st.root
}

func (st *SubTree) GenerateProof(rowId common.RowId) ([]byte, error) {
	// get index of given rowId
	index := -1
	for i, leaf := range st.leaves {
		if bytes.Equal(leaf[:], rowId[:]) {
			index = i
		}
	}
	if index == -1 {
		return nil, errors.New("cannot find given rowId in leaves")
	}

	buf := new(bytes.Buffer)
	for _, lvl := range st.tree[:len(st.tree)-1] {
		if index%2 != 0 { // left
			buf.Write(append([]byte{Left}, lvl[index-1].Bytes()...))
		} else { // right
			if len(lvl) == (index + 1) {
				buf.Write(append([]byte{Right}, lvl[index].Bytes()...))
			} else {
				buf.Write(append([]byte{Right}, lvl[index+1].Bytes()...))
			}
		}

		index /= 2
	}
	return buf.Bytes(), nil
}

func (st *SubTree) Verify(rowId common.RowId, proof []byte) (bool, error) {
	if len(proof)%ProofLength != 0 {
		return false, errors.New("invalid proof length")
	}

	base := crypto.Keccak256Hash(rowId[:])

	for {
		if len(proof) == 0 {
			break
		}

		leaf := proof[:ProofLength]
		leaf, direction := leaf[1:], leaf[0]
		proof = proof[ProofLength:]

		if direction == Left {
			base = crypto.Keccak256Hash(leaf, base.Bytes())
		} else {
			base = crypto.Keccak256Hash(base.Bytes(), leaf)
		}
	}

	return bytes.Equal(st.root.Bytes(), base.Bytes()), nil
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
