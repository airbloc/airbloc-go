package blockchain

import (
	"fmt"
	"math/big"

	"github.com/azer/logger"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// SelectorToMethodSignature is filled by automatically generated contract binds in package `adapter`.
	SelectorToMethodSignature = make(map[Selector]string)
)

// Selector is first 4-byte slice of hashed method signatures like "transfer(uint256,address)".
type Selector [4]byte

// RegisterSelector fills SelectorToMethodSignature with given 8-chars long heximedical selector and signature,
// and called by automatically generated contract binds in package `adapter`.
func RegisterSelector(selectorHex string, signature string) {
	var selector Selector
	byteSelector := hexutil.MustDecode(selectorHex)
	copy(selector[:], byteSelector)
	SelectorToMethodSignature[selector] = signature
}

func getSignatureFromTxData(txdata []byte) string {
	var selector Selector
	copy(selector[:], txdata[:4])

	if signature, ok := SelectorToMethodSignature[selector]; ok {
		return signature
	}
	// unregistered signature. shouldn't be possible!
	return hexutil.Encode(txdata)
}

// GetTransactionDetails parses Ethereum transaction and
// returns loggable, human-readable informations.
func GetTransactionDetails(cm *ContractManager, tx *types.Transaction) (methodInfo string, attrs logger.Attrs) {
	if contractName, registered := cm.addrToName[*tx.To()]; registered {
		// get
		signature := getSignatureFromTxData(tx.Data())
		methodInfo = fmt.Sprintf("%s.%s", contractName, signature)
	} else {
		// external contract which has no ABI information.
		methodInfo = tx.To().Hex()
	}

	attrs = make(logger.Attrs)
	attrs["txid"] = tx.Hash().TerminalString()
	attrs["gas"] = tx.Gas()
	if tx.Value().Int64() > 0 {
		attrs["value"] = fmt.Sprintf("%s eth", WeiToEth(tx.Value()).Text('e', 2))
	}
	return
}

// WeiToEth converts Solidity uint256 value to Ether (1e18).
func WeiToEth(wei *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(wei),
		new(big.Float).SetInt64(params.Ether),
	)
}
