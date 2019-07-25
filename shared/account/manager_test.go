package account

import (
	"testing"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestManager_CreateTemporary(t *testing.T) {

}

func TestManager_HashIdentity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// create stub
	stub := mocks.NewMockTxClient(ctrl)
	stub.EXPECT().GetContract(gomock.Any()).Return(&adapter.Accounts{})

	manager := NewManager(stub)

	// the result should not equal to identity string
	hash := manager.HashIdentity("foo@bar.io")
	assert.NotEqual(t, hash.String(), "foo@bar.io")
}
