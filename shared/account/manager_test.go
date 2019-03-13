package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager_CreateTemporary(t *testing.T) {

}

func TestManager_HashIdentity(t *testing.T) {
	manager := NewManager(nil)

	// the result should not equal to identity string
	hash := manager.HashIdentity("foo@bar.io")
	assert.NotEqual(t, hash.String(), "foo@bar.io")
}
