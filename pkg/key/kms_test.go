package key

import (
	"testing"

	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/stretchr/testify/require"
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
	srcData := &types.Data{Payload: "Hello"}
	edata, err := TestManager.EncryptData(srcData)
	require.NoError(t, err)

	data, err := TestManager.DecryptData(edata)
	require.NoError(t, err)

	require.Equal(t, srcData.Payload, data.Payload)
}
