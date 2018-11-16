package merkle

import (
	"crypto/rand"
	"sort"
	"testing"

	"log"

	"math/big"

	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestNewSMT(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	length := int64(1111)

	dbOpt := badger.DefaultOptions
	dbOpt.Dir = "testdata"
	dbOpt.ValueDir = "testdata"
	db, err := badger.Open(dbOpt)
	assert.NoError(t, err)

	// make test data
	var k Key
	for i := int64(0); i < length; i++ {
		tmp := make([]byte, common.HashLength)
		rand.Read(tmp)
		k = append(k, crypto.Keccak256Hash(tmp))
	}
	sort.Sort(k)

	smt := NewSMT(db, crypto.Keccak256Hash)
	defer smt.Close()

	root := smt.Make(k)
	smt.Flush(big.NewInt(1), []byte("frostornge"))

	tmp := make([]byte, common.HashLength)
	rand.Read(tmp)
	index := new(big.Int).SetBytes(tmp)
	index.Mod(index, big.NewInt(length))

	proof, err := smt.Proof(k[index.Int64()])
	assert.NoError(t, err)
	assert.True(t, smt.Verify(k[index.Int64()], root, proof))
}
