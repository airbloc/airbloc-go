package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/signal"

	logger2 "github.com/airbloc/airbloc-go/logger"
	"github.com/azer/logger"
	"github.com/pkg/errors"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"gopkg.in/urfave/cli.v1"
)

var log = logger.New("bootnode")

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
	logger2.Setup(os.Stdout, "*", "*")

	var nodekey *key.Key
	if ctx.IsSet("nodekeyhex") {
		priv, err := crypto.HexToECDSA(ctx.String("nodekeyhex"))
		if err != nil {
			return errors.Wrap(err, "wrong node key")
		}
		nodekey = key.FromECDSA(priv)

	} else if ctx.IsSet("nodekey") {
		nodekey, err = key.Load(ctx.String("nodekey"))
		if err != nil {
			return errors.Wrap(err, "failed to load node key")
		}
	} else {
		log.Info("No node key was given. Generating new key...")
		nodekey, err = key.Generate()
		if err != nil {
			return errors.Wrap(err, "failed to generate node key")
		}
	}

	keypair, err := nodekey.DeriveLibp2pKeyPair()
	public, err := keypair.GetPublic().Bytes()
	log.Info("Node public key: %s", base64.StdEncoding.EncodeToString(public))
	log.Info("Node ID: %s", nodekey.EthereumAddress.Hex())

	addrStr := fmt.Sprintf("/ip4/%s/tcp/%d", ctx.String("bind"), ctx.Int("port"))
	addr, err := multiaddr.NewMultiaddr(addrStr)
	if err != nil {
		return errors.Wrap(err, "failed to create multiaddr")
	}

	server, err := p2p.NewAirblocServer(localdb.NewMemDB(), nodekey, addr, true, []peerstore.PeerInfo{})
	if err != nil {
		return errors.Wrap(err, "unable to start bootnode p2p server")
	}
	defer server.Stop()

	bootInfo, err := server.BootInfo()
	if err != nil {
		return errors.Wrap(err, "unable to get bootnode address")
	}
	log.Info("Address: %s", multiaddr.Join(bootInfo.Addrs...).String()+"/ipfs/"+bootInfo.ID.Pretty())
	log.Info("You can put the address to p2p.bootNodes in config.yml.")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	log.Info("Bye 👋")
	return
}

func main() {
	app := newApp()
	if err := app.Run(os.Args); err != nil {
		log.Error("Error: %+v", err)
		os.Exit(1)
	}
}
