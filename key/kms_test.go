package key

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/stretchr/testify/require"
	"testing"
)

var TestManager *manager

func init() {
	var err error
	TestManager = new(manager)
	TestManager.ownerKey, err = Generate()
	if err != nil {
		panic(err)
	}
}

func TestManager_Decrypt(t *testing.T) {
	srcData := "Hello"
	edata, err := TestManager.Encrypt(srcData)
	require.NoError(t, err)

	data, err := TestManager.Decrypt(edata)
	require.NoError(t, err)

	require.Equal(t, srcData, data)
}

func TestManager_DecryptData(t *testing.T) {
	srcData := &common.Data{Payload: "Hello"}
	edata, err := TestManager.EncryptData(srcData)
	require.NoError(t, err)

	data, err := TestManager.DecryptData(edata)
	require.NoError(t, err)

	require.Equal(t, srcData.Payload, data.Payload)
}
