package p2p

import (
	"bufio"
	"github.com/golang/protobuf/proto"
	"github.com/klaytn/klaytn/common"
	"github.com/libp2p/go-libp2p-net"
	"io"
	"reflect"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/klaytn/klaytn/crypto"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/pkg/errors"
)

type directMessage struct {
	pb.Message
}

// readDirectMessageFrom reads raw protobuf message from stream.
func readDirectMessageFrom(stream net.Stream) (directMessage, error) {
	var raw []byte
	reader := bufio.NewReader(stream)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		raw = append(raw, p[:n]...)
	}
	msg := directMessage{}
	err := proto.Unmarshal(raw, &msg)
	return msg, err
}

func marshalDirectMessage(server Server, payload proto.Message) ([]byte, error) {
	payloadData, err := proto.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}
	return proto.Marshal(&directMessage{
		Message: pb.Message{
			Payload: payloadData,
			From:    []byte(server.getHost().ID()),
		},
	})
}

type IncomingMessage struct {
	Payload    proto.Message
	Sender     peer.ID
	SenderAddr common.Address
}

func newIncomingMessage(msg []byte, payloadType reflect.Type, sender peer.ID) (*IncomingMessage, error) {
	// unmarshal payload as given type
	payload, ok := reflect.New(payloadType).Interface().(proto.Message)
	if !ok {
		return nil, errors.New("payload of given message is not protobuf")
	}
	if err := proto.Unmarshal(msg, payload); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal payload")
	}
	addr, err := addrFromID(sender)
	if err != nil {
		return nil, errors.Wrap(err, "failed to recover sender address from ID")
	}
	return &IncomingMessage{
		Payload:    payload,
		SenderAddr: addr,
		Sender:     sender,
	}, nil
}

// AddrFromID converts a peer.ID (made from SECP256k1 public) into the Ethereum address.
func addrFromID(id peer.ID) (addr common.Address, err error) {
	libp2pPubKey, err := id.ExtractPublicKey()
	if err != nil {
		return
	}
	pubKeyBytes, err := libp2pPubKey.Raw()
	if err != nil {
		return
	}
	pubKey, err := crypto.DecompressPubkey(pubKeyBytes)
	if err != nil {
		return
	}
	addr = crypto.PubkeyToAddress(*pubKey)
	return
}
