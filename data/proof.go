package data

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/merkle"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (bundle *Bundle) generateSMT() (err error) {
	leaves := make(map[common.ID][]common.RowId, len(bundle.Data))
	for userId, rowData := range bundle.Data {
		leaves[userId] = make([]common.RowId, len(rowData))
		for i, data := range rowData {
			leaves[userId][i] = data.RowId
		}
	}

	bundle.tree, err = merkle.NewMainTree(leaves)
	return
}

// SetupUserProof creates a root of 64-depth SMT (Sparse Merkle Tree),
// which can be used as an accumulator of User IDs for the bundle.
func (bundle *Bundle) SetupUserProof() (root ethCommon.Hash, _ error) {
	if bundle.tree == nil {
		if err := bundle.generateSMT(); err != nil {
			return root, errors.Wrap(err, "setup user proof")
		}
	}
	root = bundle.tree.Root()
	return
}

func (bundle *Bundle) GenerateProof(rowId common.RowId, userId common.ID) ([]byte, error) {
	if bundle.tree == nil {
		if err := bundle.generateSMT(); err != nil {
			return nil, errors.Wrap(err, "setup user proof")
		}
	}
	return bundle.tree.GenerateProof(rowId, userId)
}
