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

// n is smallest element type of sparse merkle tree
type n struct {
	k uint64
	v ethCommon.Hash
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
	empty []ethCommon.Hash // length of this array is depth of this tree
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

func (mt *MainTree) GenerateProof() ([]byte, error) {
	return nil, nil
}

func (mt *MainTree) Verify(rowId common.RowId, userId common.ID, proof []byte) (bool, error) {
	return false, nil
}

func (mt *MainTree) hash(b ...[]byte) ethCommon.Hash {
	return crypto.Keccak256Hash(b...)
}

func (mt *MainTree) createEmptyHash() {
	base := mt.hash(bytes.Repeat([]byte{0x00}, 32))
	mt.empty[0] = base

	for lvl := 1; lvl < depth+1; lvl++ {
		prev := mt.empty[lvl-1]
		next := mt.hash(prev.Bytes(), prev.Bytes())
		mt.empty[lvl] = next
	}
}

func (mt *MainTree) createTree() {
	for _, leaf := range mt.leaves {
		mt.tree[0] = append(mt.tree[0], &n{leaf.userId.Uint64(), leaf.Root()})
	}

	for lvl := 0; lvl < depth; lvl++ {
		treeLvl := mt.tree[lvl]
		nextLvl := make(map[uint64]ethCommon.Hash)

		for i, v := range treeLvl {
			if v.k%2 != 0 {
				coKey := v.k - 1
				coIndex := i - 1

				if i == 0 && treeLvl[coIndex].k != coKey {
					nextLvl[v.k/2] = mt.hash(mt.empty[lvl].Bytes(), v.v.Bytes())
				}
			} else {
				coKey := v.k + 1
				coIndex := i + 1

				switch {
				case len(treeLvl) == coIndex:
					fallthrough
				case treeLvl[coIndex].k != coKey:
					nextLvl[v.k/2] = mt.hash(v.v.Bytes(), mt.empty[lvl].Bytes())
				case treeLvl[coIndex].k == coKey:
					nextLvl[v.k/2] = mt.hash(v.v.Bytes(), treeLvl[coIndex].v.Bytes())
				}
			}
		}

		for k, v := range nextLvl {
			mt.tree[lvl+1] = append(mt.tree[lvl+1], &n{k, v})
		}
		sort.Sort(mt.tree[lvl+1])
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

	cache := make(map[int]*SubTree)
	for k, v := range input {
		leaf := struct {
			userId common.ID
			*SubTree
		}{}
		leaf.userId = k

		if subTree, exists := cache[len(v)]; exists {
			leaf.SubTree = subTree
		} else {
			subTree, err := NewSubTree(v)
			if err != nil {
				return nil, err
			}
			leaf.SubTree = subTree
			cache[len(v)] = subTree
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
		empty:  make([]ethCommon.Hash, depth+1),
	}
	mt.createEmptyHash()

	if len(mt.leaves) == 0 {
		mt.root = mt.empty[depth]
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
