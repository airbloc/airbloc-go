package main

import (
	"encoding/base64"
	"fmt"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"os/signal"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "bootnode"
	app.Version = "0.1.0"
	app.Usage = "Bootstrap Node for Airbloc Network."
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "nodekey",
			Value: "",
			Usage: "Private key (ECDSA) file path. Generates new one if not specified.",
		},
		cli.StringFlag{
			Name:  "nodekeyhex",
			Value: "",
			Usage: "Private key (ECDSA) hex. Testing purpose only.",
		},
		cli.StringFlag{
			Name:  "bind",
			Value: "0.0.0.0",
			Usage: "Bind address. (0.0.0.0 for public access)",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 9100,
			Usage: "Listen address for P2P connection.",
		},
	}
	return app
}

func run(ctx *cli.Context) (err error) {
	var nodekey *key.Key
	if ctx.IsSet("nodekeyhex") {
		priv, err := crypto.HexToECDSA(ctx.String("nodekeyhex"))
		if err != nil {
			log.Fatalf("wrong node key: %+v", err)
		}
		nodekey = key.FromECDSA(priv)

	} else if ctx.IsSet("nodekey") {
		nodekey, err = key.Load(ctx.String("nodekey"))
		if err != nil {
			log.Fatalf("failed to load node key: %+v", err)
		}
	} else {
		log.Println("No node key was given. Generating new key...")
		nodekey, err = key.Generate()
		if err != nil {
			log.Fatalf("failed to generate node key: %+v", err)
		}
	}

	keypair, err := nodekey.DeriveLibp2pKeyPair()
	public, err := keypair.GetPublic().Bytes()
	log.Printf("Node public key: %s\n", base64.StdEncoding.EncodeToString(public))

	addrStr := fmt.Sprintf("/ip4/%s/tcp/%d", ctx.String("bind"), ctx.Int("port"))
	addr, err := multiaddr.NewMultiaddr(addrStr)
	if err != nil {
		log.Fatalf("failed to parse address: %+v", err)
	}

	server, err := p2p.NewServer(localdb.NewMemDB(), nodekey, addr, true, []peerstore.PeerInfo{})
	if err != nil {
		log.Fatalf("unable to create bootstrap server: %+v", err)
	}
	defer server.Stop()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	log.Println("Bye 👋")
	return
}

func main() {
	app := newApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
