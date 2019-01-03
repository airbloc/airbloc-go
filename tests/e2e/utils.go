package e2e

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func toDataId(bundleId, userId string) ([20]byte, error) {
	dataId := [20]byte{}

	// bundleId: deadbeefdeadbeef/1
	// userId: cafebabe12345678
	// turn into [8byte collectionId][4byte bundleNumber][8byte userId] = 20 bytes
	idSegs := strings.Split(bundleId, "/")
	if len(idSegs) != 2 {
		return [20]byte{}, errors.Errorf("Invalid bundle Id: %s", bundleId)
	}
	collectionId, err := hex.DecodeString(idSegs[0])
	if err != nil {
		return dataId, errors.Wrapf(err, "Wrong collection Id in bundleId %s", bundleId)
	}

	bundleNumberInt, err := strconv.Atoi(idSegs[1])
	if err != nil {
		return dataId, errors.Wrapf(err, "Invalid bundle number in bundleId %s", bundleId)
	}
	bundleNumber := new(bytes.Buffer)
	binary.Write(bundleNumber, binary.LittleEndian, bundleNumberInt)

	userIdRaw, err := hex.DecodeString(userId)
	if err != nil {
		return [20]byte{}, errors.Errorf("Invalid user ID: %s", userId)
	}
	copy(dataId[:8], collectionId)
	copy(dataId[8:12], bundleNumber.Bytes())
	copy(dataId[12:20], userIdRaw)
	return dataId, nil
}
