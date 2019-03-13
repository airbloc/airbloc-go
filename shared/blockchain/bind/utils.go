package bind

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"reflect"
)

var (
	// if there was REVERT on contract call, the revert reason will be returned as a call result.
	// For details, please see EIP-838: https://github.com/ethereum/EIPs/issues/838
	revertErrorSelector = crypto.Keccak256([]byte("Error(string)"))[:4]

	// revertErrorABI is used to decode revert reason.
	revertErrorABI = abi.Arguments{
		abi.Argument{
			Name: "reason",
			Type: abi.Type{
				Kind: reflect.String,
				Type: reflect.TypeOf(""),
				T:    abi.StringTy,
			},
		},
	}
)

// hasRevertError returns revert reason if there was a revert error message
// according to EIP-838 (https://github.com/ethereum/EIPs/issues/838) on given call result.
func HasRevertError(callResult []byte) (string, bool) {
	if !bytes.Equal(callResult[:4], revertErrorSelector) {
		// revert message should be started with 0x08c379a0.
		return "", false
	}
	var parsedMsg struct {
		Reason string
	}
	err := revertErrorABI.Unpack(&parsedMsg, callResult[4:])
	if err != nil {
		return err.Error(), false
	}
	return parsedMsg.Reason, true
}
