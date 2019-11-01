package blockchain

import (
	"math/big"

	"github.com/klaytn/klaytn/params"
)

// WeiToKlay converts Solidity uint256 value to Klay (1e18).
func WeiToKlay(peb *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(peb),
		new(big.Float).SetInt64(params.KLAY),
	)
}

// getChainName returns chain name by chain ID (network ID), according to EIP-155.
func getChainName(cid *big.Int) string {
	switch cid.String() {
	case "1":
		return "Ethereum Main"
	case "3":
		return "Ethereum Ropsten Test"
	case "4":
		return "Ethereum Rinkeby Test"
	case "1000":
		return "Klaytn Aspen Test"
	case "1001":
		return "Klaytn Baobab Test"
	case "8217":
		return "Klaytn Main"
	}
	return "EVM Private"
}
