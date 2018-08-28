package main

import (
	"airbloc/go-producer/stream"
	"context"
)

func init() {

}

func main() {
	ctx, cncl := context.WithCancel(context.Background())
	strm, err := stream.Run(ctx, "tcp", ":6000")
	if err != nil {
		panic(err)
	}
	defer strm.Close()
	cncl()
}
