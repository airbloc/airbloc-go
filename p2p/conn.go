package p2p

import (
	"context"
	"fmt"
	"os"

	"bufio"

	"github.com/airbloc/airbloc-go/key"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-net"
	"github.com/multiformats/go-multiaddr"
)

func readData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		rw.WriteString(fmt.Sprintf("%s\n", sendData))
		rw.Flush()
	}

}

func handleStream(stream net.Stream) {
	defer stream.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw)
	go writeData(rw)
}

func Connect(addr multiaddr.Multiaddr, key *key.Key) error {
	ctx := context.Background()

	privKey, err := key.DeriveLibp2pKeyPair()
	if err != nil {
		return err
	}

	host, err := libp2p.New(ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr))
	if err != nil {
		return err
	}
	defer host.Close()

	node, err := dht.New(ctx, host)
	if err != nil {
		return err
	}
	defer node.Close()

	if err = node.Bootstrap(ctx); err != nil {
		return err
	}
	return nil
}
