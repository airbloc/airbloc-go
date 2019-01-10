package main

import (
	"fmt"
	"github.com/airbloc/airbloc-go/key"
)

var files = []string{
	"private.key",
	"private.userdelegate.key",
}

func main() {
	// this code generates two private key
	for _, filename := range files {
		k, _ := key.Generate()
		k.Save(filename)
		fmt.Printf("%s generated (Address: %s)\n",
			filename, k.EthereumAddress.Hex())
	}
}
