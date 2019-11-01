package types

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/pkg/errors"
)

const (
	RowIdLength    = 4
	RowIdStrLength = 8
)

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

func (id RowId) Uint32() uint32 {
	return binary.LittleEndian.Uint32(id[:])
}

func (id RowId) Hex() string {
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

func (id RowId) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Hex())
}
