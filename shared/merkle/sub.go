package merkle

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"sort"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/pkg/errors"
)

const (
	left  = 0x00
	right = 0x01
)

type SubTree struct {
	leaves [][4]byte
	tree   [][]common.Hash
	root   common.Hash
}

func (st *SubTree) Leaves() [][4]byte {
	return st.leaves
}

func (st *SubTree) Root() common.Hash {
	return st.root
}

func (st *SubTree) GenerateProof(rowId [4]byte) ([]byte, error) {
	// get index of given rowId
	index := sort.Search(len(st.leaves), func(i int) bool {
		return binary.LittleEndian.Uint32(st.leaves[i][:]) >= binary.LittleEndian.Uint32(rowId[:])
	})

	if index == len(st.leaves) {
		return nil, errors.New("cannot find leaf in sub tree")
	}

	buf := new(bytes.Buffer)
	for _, lvl := range st.tree[:len(st.tree)-1] {
		if index%2 != 0 { // left
			buf.Write(append([]byte{left}, lvl[index-1].Bytes()...))
		} else { // right
			if len(lvl) == (index + 1) {
				buf.Write(append([]byte{right}, lvl[index].Bytes()...))
			} else {
				buf.Write(append([]byte{right}, lvl[index+1].Bytes()...))
			}
		}

		index /= 2
	}
	return append(st.root.Bytes(), buf.Bytes()...), nil
}

func verifySubProof(rowId [4]byte, subRoot, proof []byte) bool {
	root := common.BytesToHash(subRoot)
	base := crypto.Keccak256Hash(rowId[:])

	for {
		if len(proof) == 0 {
			break
		}

		leaf := proof[:SubProofLength]
		leaf, direction := leaf[1:], leaf[0]
		proof = proof[SubProofLength:]

		if direction == left {
			base = crypto.Keccak256Hash(leaf, base.Bytes())
		} else {
			base = crypto.Keccak256Hash(base.Bytes(), leaf)
		}
	}

	return bytes.Equal(root.Bytes(), base.Bytes())
}

func VerifySubProof(rowId [4]byte, proof []byte) (bool, error) {
	if len(proof) < HashLength+SubProofLength {
		return false, errors.New("invalid proof length")
	}

	root := proof[:HashLength]
	proof = proof[HashLength:]

	if len(proof)%SubProofLength != 0 {
		return false, errors.New("invalid proof length")
	}

	return verifySubProof(rowId, root, proof), nil
}

func NewSubTree(input [][4]byte) (*SubTree, error) {
	// check input
	pow := new(big.Int).Lsh(big.NewInt(1), uint(32))
	if big.NewInt(int64(len(input))).Cmp(pow) > 0 {
		return nil, errors.New("too long input")
	}

	// sort
	sort.Slice(input, func(i, j int) bool {
		return binary.LittleEndian.Uint32(input[i][:]) < binary.LittleEndian.Uint32(input[j][:])
	})

	// hashing
	base := make([]common.Hash, len(input))
	for index, elem := range input {
		base[index] = crypto.Keccak256Hash(elem[:])
	}

	// create tree structure
	st := &SubTree{
		leaves: input,
		tree:   [][]common.Hash{base},
	}

	// generate tree
	for i := 0; ; i++ {
		if len(st.tree[i]) == 1 {
			st.root = st.tree[i][0]
			break
		}

		var nextLvl []common.Hash
		for j := 0; j < len(st.tree[i]); j++ {
			if j%2 != 0 {
				continue
			}

			leaf := st.tree[i][j]
			coleaf := common.Hash{}

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
