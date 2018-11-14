package merkle

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type proof struct {
	bits []byte
	hash []common.Hash
}

func encodeProof(p proof) []byte {
	var pb []byte
	pb = append(pb, p.bits...)
	for _, h := range p.hash {
		pb = append(pb, h.Bytes()...)
	}
	return pb
}

func decodeProof(b []byte) (proof, error) {
	if len(b)%common.HashLength != 0 {
		return proof{}, errors.New("wrong proof length")
	}

	p := proof{}
	bits, hash := b[:common.HashLength], b[common.HashLength:]
	p.bits = bits
	for i := common.HashLength; i < len(b); i += common.HashLength {
		var tmp []byte
		tmp, hash = hash[:common.HashLength], hash[common.HashLength:]
		p.hash = append(p.hash, common.BytesToHash(tmp))
	}
	return p, nil
}
