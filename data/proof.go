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

func (bundle *Bundle) generateSubSMT(rowData []*ablCommon.EncryptedData) (ethCommon.Hash, error) {
	leaves := make(map[uint64][]byte, len(rowData))
	for i, data := range rowData {
		binary.LittleEndian.PutUint32(data.RowId[:], uint32(i))
		leaves[uint64(i)] = leaveExist
	}

	root := ethCommon.Hash{}
	smt, err := mamamerkle.NewSparseMerkleTree(32, leaves)
	if err != nil {
		return root, errors.Wrap(err, "smt generation error")
	}
	copy(root[:], smt.Root())
	return root, nil
}

func (bundle *Bundle) generateSMT() (*mamamerkle.SparseMerkleTree, error) {
	leaves := make(map[uint64][]byte, len(bundle.Data))

	for userId, rowData := range bundle.Data {
		leafId := binary.LittleEndian.Uint64(userId[:])
		if _, exists := leaves[leafId]; !exists {
			leaf, err := bundle.generateSubSMT(rowData)
			if err != nil {
				return nil, err
			}
			leaves[leafId] = leaf[:]
		}
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
