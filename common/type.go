package common

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	IDLength = 8
)

type ID [IDLength]byte

// HexToID converts heximedical string ID (ex: deadbeef1a2b3c4d) to ID instance.
func HexToID(idStr string) (ID, error) {
	var id ID
	byteId, err := hex.DecodeString(idStr)
	if err != nil {
		return id, err
	}
	copy(id[:], byteId[:IDLength])
	return BytesToID(byteId), nil
}

func BytesToID(idBytes []byte) ID {
	var id ID
	copy(id[:], idBytes[:IDLength])
	return id
}

func GenerateID(issuer common.Address, time time.Time, seed []byte) (id ID) {
	// use ABI-compatible 32-byte padding,
	// to be compatible with keccak256 on Ethereum
	preimageWithPadding := make([]byte, 64)

	copy(preimageWithPadding[12:32], issuer.Bytes())
	binary.LittleEndian.PutUint64(preimageWithPadding[56:64], uint64(time.Unix()))

	// TODO: pad seed
	hash := crypto.Keccak256(preimageWithPadding, seed)

	// use only first 8 byte in 32-byte hash
	copy(id[:], hash[:8])
	return
}

func (id *ID) Hex() string {
	return hex.EncodeToString(id[:])
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tempId, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(tempId) != IDLength {
		return errors.Errorf("invalid ID format: %s", string(b))
	}
	copy(id[:], tempId)
	return nil
}

func (id *ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Hex())
}
