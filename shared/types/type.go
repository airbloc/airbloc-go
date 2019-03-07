package types

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/json-iterator/go"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	IDLength    = 8
	IDStrLength = 16

	RowIdLength    = 4
	RowIdStrLength = 8
)

type ID [IDLength]byte

// HexToID converts heximedical string ID (ex: deadbeef1a2b3c4d) to ID instance.
func HexToID(idStr string) (ID, error) {
	var id ID
	byteId, err := hex.DecodeString(idStr)
	if err != nil {
		return id, err
	}
	if len(byteId) != IDLength {
		return id, errors.Errorf("invalid ID: %s", idStr)
	}
	copy(id[:], byteId[:IDLength])
	return BytesToID(byteId), nil
}

func BytesToID(idBytes []byte) ID {
	var id ID
	copy(id[:], idBytes[:IDLength])
	return id
}

func UintToID(i uint64) ID {
	id := ID{}
	binary.LittleEndian.PutUint64(id[:], i)
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

func (id ID) Uint64() uint64 {
	return binary.LittleEndian.Uint64(id[:])
}

func (id ID) Hex() string {
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

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Hex())
}

func IDListToByteList(ids []ID) [][8]byte {
	byteIds := make([][8]byte, len(ids))
	for i, id := range ids {
		byteIds[i] = id
	}
	return byteIds
}

// IDFilter fucks with go-ethereum/accounts/abi/bind/topics.go
func IDFilter(ids ...ID) [][32]byte {
	byteIds := make([][32]byte, len(ids))
	for i, id := range ids {
		copy(byteIds[i][:8], id[:])
	}
	return byteIds
}

type RowId [RowIdLength]byte

func HexToRowId(idStr string) (RowId, error) {
	var id RowId
	byteId, err := hex.DecodeString(idStr)
	if err != nil {
		return id, err
	}
	if len(byteId) != RowIdLength {
		return id, errors.Errorf("invalid ID: %s", idStr)
	}
	copy(id[:], byteId[:RowIdLength])
	return BytesToRowId(byteId), nil
}

func BytesToRowId(idBytes []byte) RowId {
	var id RowId
	copy(id[:], idBytes[:RowIdLength])
	return id
}

func UintToRowId(i uint32) RowId {
	id := RowId{}
	binary.LittleEndian.PutUint32(id[:], i)
	return id
}

func (id *RowId) Uint32() uint32 {
	return binary.LittleEndian.Uint32(id[:])
}

func (id *RowId) Hex() string {
	return hex.EncodeToString(id[:])
}

func (id *RowId) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tempId, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(tempId) != RowIdLength {
		return errors.Errorf("invalid ID format: %s", string(b))
	}
	copy(id[:], tempId)
	return nil
}

func (id *RowId) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Hex())
}
