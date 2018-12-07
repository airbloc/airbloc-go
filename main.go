package main

import (
	"encoding/hex"
	"log"
)

func main() {
	input := "b86041f2a688f2480000000000000001"
	bs, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	for i, b := range bs {
		log.Printf("%3d %3v %3s", i, b, input[i*2:i*2+2])
	}
}
