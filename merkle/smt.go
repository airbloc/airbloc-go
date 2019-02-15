package merkle

import (
	"bytes"
	"github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"math/big"
	"sort"
)

const depth = 64

var (
	empty []ethCommon.Hash
	hash  = func(b ...[]byte) ethCommon.Hash { return crypto.Keccak256Hash(b...) }
)

func init() {
	empty = make([]ethCommon.Hash, depth+1)
	base := crypto.Keccak256Hash(bytes.Repeat([]byte{0x00}, 32))
	empty[0] = base

	for lvl := 1; lvl < depth+1; lvl++ {
		prev := empty[lvl-1]
		next := crypto.Keccak256Hash(prev.Bytes(), prev.Bytes())
		empty[lvl] = next
	}
}

// n is smallest element type of sparse merkle tree
type n struct {
	k uint64
	v ethCommon.Hash
	n uint64 // points to leaf that placed in next level
}
type ns []*n

func (ns ns) Len() int           { return len(ns) }
func (ns ns) Less(i, j int) bool { return ns[i].k < ns[j].k }
func (ns ns) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }

type MainTree struct {
	leaves []struct {
		userId common.ID
		*SubTree
	}
	tree  []ns
	root  ethCommon.Hash
	cache map[uint64]*SubTree
}

func (mt *MainTree) Leaves() map[common.ID][]common.RowId {
	leaves := make(map[common.ID][]common.RowId, len(mt.leaves))
	for _, subTree := range mt.leaves {
		leaves[subTree.userId] = subTree.Leaves()
	}
	return leaves
}

func (mt *MainTree) Root() ethCommon.Hash {
	return mt.root
}

func (mt *MainTree) GenerateProof(rowId common.RowId, userId common.ID) ([]byte, error) {
	leafIndex := sort.Search(len(mt.tree[0]), func(i int) bool {
		return mt.tree[0][i].k >= userId.Uint64()
	})

	if leafIndex == len(mt.tree[0]) {
		return nil, errors.New("cannot find leaf in main tree")
	}

	leaf := mt.leaves[leafIndex]
	subProof, err := leaf.GenerateProof(rowId)
	if err != nil {
		return nil, err
	}
	subProof = append(leaf.Root().Bytes(), subProof...)

	var proofBits = make([]byte, depth/8)
	var proofBytes []byte

	for pos, lvl := range mt.tree[:len(mt.tree)-1] { // remove root
		leaf := lvl[leafIndex]

		if leafIndex%2 == 0 {
			// right
			coKey := leaf.k + 1
			coIndex := leafIndex + 1

			switch {
			case len(lvl) == coIndex: // avoid panic
			case lvl[coIndex].k == coKey:
				setBit(proofBits, uint64(depth-pos))
				proofBytes = append(proofBytes, leaf.v.Bytes()...)
			}
		}

		leafIndex /= 2
	}

	mainProof := append(proofBits, proofBytes...)

	return append(mainProof, subProof...), nil
}

func verifyMainProof(userId common.ID, subRoot, mainRoot, proofBits, proofBytes []byte) bool {
	k := userId.Uint64()
	v := ethCommon.BytesToHash(subRoot)

	for i := 0; i < len(proofBits)*8; i++ {
		h := empty[i].Bytes()
		if hasBit(proofBits, uint64(i)) {
			h, proofBytes = proofBytes[:HashLength], proofBytes[HashLength:]
		}

		if k%2 != 0 {
			v = hash(h, v.Bytes())
		} else {
			v = hash(v.Bytes(), h)
		}
		k /= 2
	}

	return bytes.Equal(mainRoot, v.Bytes())
}

func VerifyMainProof(rowId common.RowId, userId common.ID, mainRoot, proof []byte) bool {
	// split proofs
	// main
	proofBits, proof := proof[:common.IDLength], proof[common.IDLength:]
	count := 0
	for i := 0; i < len(proofBits)*8; i++ {
		if hasBit(proofBits, uint64(i)) {
			count++
		}
	}
	proofBytes, proof := proof[:HashLength*count], proof[HashLength*count:]

	// sub
	subRoot := proof[:HashLength]
	subProof := proof[HashLength:]

	// verify main proof
	if !verifyMainProof(userId, subRoot, mainRoot, proofBits, proofBytes) {
		return false
	}

	// verify sub proof
	if !verifySubProof(rowId, subRoot, subProof) {
		return false
	}

	return true
}

func (mt *MainTree) createEmptyHash() {

}

func (mt *MainTree) createTree() {
	for _, leaf := range mt.leaves {
		mt.tree[0] = append(mt.tree[0], &n{k: leaf.userId.Uint64(), v: leaf.Root()})
	}

	for lvl := 0; lvl < depth; lvl++ {
		treeLvl := mt.tree[lvl]
		nextLvl := ns{}

		for i, v := range treeLvl {
			if v.k%2 != 0 {
				// left
				coKey := v.k - 1
				coIndex := i - 1

				switch {
				case i == 0: // avoid panic
					fallthrough
				case treeLvl[coIndex].k != coKey:
					nextLvl = append(nextLvl, &n{
						k: v.k / 2,
						v: hash(empty[lvl].Bytes(), v.v.Bytes()),
						n: uint64(len(nextLvl)),
					})
				}
			} else {
				// right
				coKey := v.k + 1
				coIndex := i + 1

				switch {
				case len(treeLvl) == coIndex: // avoid panic
					fallthrough
				case treeLvl[coIndex].k != coKey:
					nextLvl = append(nextLvl, &n{
						k: v.k / 2,
						v: hash(v.v.Bytes(), empty[lvl].Bytes()),
						n: uint64(len(nextLvl)),
					})
				case treeLvl[coIndex].k == coKey:
					nextLvl = append(nextLvl, &n{
						k: v.k / 2,
						v: hash(v.v.Bytes(), treeLvl[coIndex].v.Bytes()),
						n: uint64(len(nextLvl)),
					})
				}
			}
		}

		mt.tree[lvl+1] = nextLvl
	}
}

func NewMainTree(input map[common.ID][]common.RowId) (*MainTree, error) {
	// check input
	pow := new(big.Int).Lsh(big.NewInt(1), uint(64))
	if big.NewInt(int64(len(input))).Cmp(pow) > 0 {
		return nil, errors.New("too long input")
	}

	// create SubTrees
	leaves := make([]struct {
		userId common.ID
		*SubTree
	}, len(input))
	i := 0

	cache := make(map[uint64]*SubTree)
	for k, v := range input {
		leaf := struct {
			userId common.ID
			*SubTree
		}{}
		leaf.userId = k

		if subTree, exists := cache[uint64(len(v))]; exists {
			leaf.SubTree = subTree
		} else {
			subTree, err := NewSubTree(v)
			if err != nil {
				return nil, err
			}
			leaf.SubTree = subTree
			cache[uint64(len(v))] = subTree
		}
		leaves[i] = leaf
		i++
	}

	// sort by userId (key)
	sort.Slice(leaves, func(i, j int) bool {
		return leaves[i].userId.Uint64() < leaves[j].userId.Uint64()
	})

	// initialize struct & create empty hash
	mt := &MainTree{
		leaves: leaves,
		tree:   make([]ns, depth+1),
		cache:  cache,
	}
	mt.createEmptyHash()

	if len(mt.leaves) == 0 {
		mt.root = empty[depth]
	} else {
		mt.createTree()
		root := mt.tree[len(mt.tree)-1]
		if len(root) > 1 {
			return nil, errors.Errorf("root array should have one element : %v", root)
		}
		mt.root = root[0].v
	}
	return mt, nil
}
