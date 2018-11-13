package merkle

import (
	"bytes"
	"math/big"

	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type Key []common.Hash

func (k Key) Len() int           { return len(k) }
func (k Key) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k Key) Less(i, j int) bool { return bytes.Compare(k[i].Bytes(), k[j].Bytes()) == -1 }

type SMT struct {
	db     *badger.DB
	tree   tree          // height/index/hash
	empty  []common.Hash // height/hash
	height int
	hash   func(...[]byte) common.Hash
}

func NewSMT(db *badger.DB, hash func(...[]byte) common.Hash) *SMT {
	smt := new(SMT)
	smt.db = db
	smt.hash = hash
	smt.height = common.HashLength * 8
	smt.empty = append(smt.empty, smt.hash([]byte{}))
	for i := 0; i < smt.height-1; i++ {
		smt.empty = append(smt.empty, smt.hash(
			smt.empty[i].Bytes(),
			smt.empty[i].Bytes(),
		))
	}
	return smt
}

func (s *SMT) Close() error {
	return s.db.Close()
}

func (s *SMT) Make(keys Key) common.Hash {
	var base level
	for _, key := range keys {
		base = append(base, &node{cIdx: key.Big(), hash: key})
	}
	s.tree = append(s.tree, &base)
	return s.make(keys)
}

func (s *SMT) make(keys Key) common.Hash {
	for i := 0; i < s.height; i++ {
		nLv := level{}
		pLv := s.tree[i]
		pIdx := big.NewInt(-1)
		for j, n := range *pLv {
			direction := n.Direction()

			if !direction && n.CmpBool(bigAddInt(pIdx, 1)) {
				//log.Println(i, direction, node.hash.Hex(), nLv[len(nLv)-1].hash.Hex(), "*")
				// get last node and override it
				nLv[len(nLv)-1].hash = s.makeHash(
					direction,
					n.hash,
					(*pLv)[j-1].hash,
				)
			} else {
				//if idx == 0 {
				//	log.Println(i, direction, node.hash.Hex(), s.empty[i].Hex())
				//}
				nLv = append(nLv, &node{
					cIdx: bigDivInt(n.cIdx, 2),
					hash: s.makeHash(
						direction,
						n.hash,
						s.empty[i],
					),
				})
			}

			n.nIdx = len(nLv) - 1
			pIdx = n.cIdx
		}
		s.tree = append(s.tree, &nLv)
	}
	return (*s.tree[len(s.tree)-1])[0].hash
}

func (s *SMT) Flush(blockNum *big.Int, data []byte) error {
	return s.tree.flush(s.db, s.height, blockNum, data)
}

// TODO: generate proof without any node data - only transaction set
// func (s *SMT) Proof(base Key, key common.Hash) ([]byte, error) {
func (s *SMT) Proof(key common.Hash) ([]byte, error) {
	if s.tree == nil {
		return nil, errors.New("tree is empty")
	}

	n, err := (*s.tree[0]).Search(key)
	if err != nil {
		return nil, err
	}

	var p = proof{
		bits: make([]byte, common.HashLength),
		hash: []common.Hash{},
	}
	for i := 0; i < s.height; i++ {
		cLv := *s.tree[i] // max(s.height) = 255
		node := cLv[n]    // node

		sbkIdx := new(big.Int) // sibling key index
		sbnIdx := n            // sibling node index
		if node.Direction() {
			sbnIdx++
			sbkIdx = bigAddInt(node.cIdx, 1)
		} else {
			sbnIdx--
			sbkIdx = bigSubInt(node.cIdx, 1)
		}

		if sbnIdx <= len(cLv)-1 && sbnIdx >= 0 &&
			cLv[sbnIdx].CmpBool(sbkIdx) {
			p.hash = append(p.hash, cLv[sbnIdx].hash)
		} else {
			bitSet(p.bits, uint64(i))
		}
		n = node.nIdx
	}

	return encodeProof(p), nil
}

func (s *SMT) Verify(key, root common.Hash, p []byte) bool {
	pr, err := decodeProof(p)
	if err != nil {
		return false
	}

	proofBits := pr.bits
	proofHash := pr.hash

	node := node{cIdx: key.Big()}
	for i, j := 0, 0; i < s.height; i++ {
		direction := node.Direction()
		node.cIdx = bigDivInt(node.cIdx, 2)

		if bitIsSet(proofBits, uint64(i)) {
			//log.Println(i, direction, key.Hex(), s.empty[i].Hex())
			key = s.makeHash(
				direction,
				key, s.empty[i],
			)
		} else {
			//log.Println(i, direction, key.Hex(), proofHash[j].Hex(), "*")
			key = s.makeHash(
				direction,
				key, proofHash[j],
			)
			j++
		}
	}

	return bytes.Equal(key.Bytes(), root.Bytes())
}

func (s *SMT) makeHash(
	direction bool,
	x, y common.Hash,
) common.Hash {
	if direction {
		return s.hash(x.Bytes(), y.Bytes())
	} else {
		return s.hash(y.Bytes(), x.Bytes())
	}
}
