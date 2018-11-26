package userdelegate

import (
	"context"
	"fmt"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/apps"
	"github.com/airbloc/airbloc-go/collections"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"log"

	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p"
	p2pcommon "github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

func newServer(memdb localdb.Database, k *key.Key, port int, bootnode bool, bootinfos ...peerstore.PeerInfo) p2p.Server {
	addr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	server, err := p2p.NewAirblocServer(memdb, k, addr, bootnode, bootinfos)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	return server
}

func Main(backend node.Backend) {
	memdb := localdb.NewMemDB()
	k, err := key.Generate()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// setup node
	addr, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/9100")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	info, err := peerstore.InfoFromP2pAddr(addr)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	server := newServer(memdb, k, 9100, false, *info)
	server.Start()
	defer server.Stop()

	addUser(backend, server, "deadbeefdeadbeef")
}

func addUser(backend node.Backend, server p2p.Server, accountId string) {
	server.SubscribeTopic("dauth-allow-"+accountId, &pb.DAuthRequest{}, authHandler(backend, true))
	server.SubscribeTopic("dauth-deny-"+accountId, &pb.DAuthRequest{}, authHandler(backend, false))

	accountMan, err := account.NewManager(backend.Client())
	if err != nil {
		log.Panicln(err.Error())
	}

	server.SubscribeTopic("dauth-signup", &pb.DAuthSignUpRequest{}, func(server p2p.Server, ctx context.Context, message p2pcommon.Message) {
		request, ok := message.Data.(*pb.DAuthSignUpRequest)
		if !ok {
			log.Println("error: Invalid topic.")
			return
		}

		identityHash := ethCommon.HexToHash(request.GetIdentityHash())
		accountId, err := accountMan.CreateTemporary(identityHash)
		if err != nil {
			log.Println("error: Failed to create temporary account:", err.Error())
		}

		log.Println(
			"Created account", accountId.String(),
			"by request from the data provider", crypto.PubkeyToAddress(*message.Sender).Hex())

		response := &pb.DAuthSignUpResponse{
			UserId: accountId.String(),
		}
		if err = server.Send(context.Background(), response, "dauth-signup-response", message.Info.ID); err != nil {
			log.Println("error: Failed to send response to data provider:", err.Error())
		}
	})
}

func authHandler(backend node.Backend, allow bool) p2p.TopicHandler {
	dauthMan := dauth.NewManager(backend.Client())

	return func(server p2p.Server, ctx context.Context, message p2pcommon.Message) {
		request, ok := message.Data.(*pb.DAuthRequest)
		if !ok {
			log.Println("error: Invalid topic.")
			return
		}

		collectionId, err := ablCommon.IDFromString(request.CollectionId)
		if err != nil {
			log.Println("error: Invalid Collection ID", collectionId, err.Error())
			return
		}

		// the message sender should be the data provider (the collection's owner)
		senderAddr := crypto.PubkeyToAddress(*message.Sender)
		if ok, err := isCollectionOwner(ctx, backend, collectionId, senderAddr); err != nil {
			log.Println("error: Failed to retrieve collection owner", err.Error())
			return
		} else if !ok {
			log.Println("error: The address", senderAddr.Hex(), "is not a data provider.")
			return
		}

		if allow {
			err = dauthMan.Allow(collectionId, request.GetIdentityHash())
		} else {
			err = dauthMan.Deny(collectionId, request.GetIdentityHash())
		}
		if err != nil {
			log.Println("error: Failed to modify DAuth settings: ", err.Error())
		}

		// TODO: reply to the data provider :D
	}
}

// isCollectionOwner checks that the P2P message sender is
// same with the owner of the collection (data provider, app owner).
func isCollectionOwner(ctx context.Context, backend node.Backend, collectionId ablCommon.ID, senderAddr ethCommon.Address) (bool, error) {
	collectionMan, err := collections.New(backend.LocalDatabase(), backend.MetaDatabase(), backend.Client())
	if err != nil {
		return false, errors.Wrap(err, "failed to initiate collection manager")
	}

	appsMan, err := apps.NewManager(backend.Client())
	if err != nil {
		return false, errors.Wrap(err, "failed to initiate apps manager")
	}

	collection, err := collectionMan.Get(collectionId)
	if err != nil {
		return false, errors.Wrap(err, "unable to retrieve collection")
	}
	return appsMan.CheckOwner(ctx, collection.AppId, senderAddr)
}
