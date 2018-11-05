package main

import (
	"math/big"

	"log"

	"github.com/ethereum/go-ethereum/params"
)

//go:generate ./build.sh

func init() {

}

func main() {
	bigFloat := new(big.Float).Mul(
		big.NewFloat(37.8),
		big.NewFloat(params.Ether),
	)
	bigInt := new(big.Int).Set
	log.Println(bigFloat.Int())
}
