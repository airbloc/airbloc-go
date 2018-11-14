package data

import (
	"bytes"
	"encoding/binary"
	ablCommon "github.com/airbloc/airbloc-go/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/mamamerkle"
	"github.com/pkg/errors"
)

var (
	// since we only need membership proof, we'll use 32byte left-padded 0x1 as a leaf value.
	leaveExist = append(bytes.Repeat([]byte{0}, 31), 0x1)
)

func (bundle *Bundle) generateSMT() (*mamamerkle.SparseMerkleTree, error) {
	leaves := make(map[uint64][]byte)

	for _, data := range bundle.Data {
		leafId := binary.LittleEndian.Uint64(data.OwnerAnid[:])
		leaves[leafId] = leaveExist
	}
	return mamamerkle.NewSparseMerkleTree(64, leaves)
}

// SetupUserProof creates a root of 64-depth SMT (Sparse Merkle Tree),
// which can be used as an accumulator of User IDs for the bundle.
func (bundle *Bundle) SetupUserProof() (ethCommon.Hash, error) {
	root := ethCommon.Hash{}

	smt, err := bundle.generateSMT()
	if err != nil {
		return root, errors.Wrap(err, "failed to construct SMT")
	}
	copy(root[:], smt.Root())
	return root, nil
}

func (bundle *Bundle) GenerateProofOfUser(userId ablCommon.ID) ([]byte, error) {
	smt, err := bundle.generateSMT()
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct SMT")
	}

	leafId := binary.LittleEndian.Uint64(userId[:])
	return smt.CreateMerkleProof(leafId), nil
}
