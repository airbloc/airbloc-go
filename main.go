package main

import (
	"airbloc/go-producer/blockchain"
	"log"
)

const (
	ABIPath  = "./abi"
	BindPath = "./blockchain"
)

func init() {
	if err := blockchain.GenerateBind(ABIPath, BindPath); err != nil {
		log.Fatal(err)
	}
}

func main() {

}
