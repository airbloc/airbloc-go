package p2p

import (
	"bufio"
	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-protocol"
	"io"
	"reflect"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/pkg/errors"
)

type RawMessage struct {
	pb.Message
}

func (m *RawMessage) ID() peer.ID {
	return peer.ID(m.GetFrom())
}

func (m *RawMessage) Pid() (Pid, error) {
	return ParsePid(string(m.GetProtocol()))
}

func (m *RawMessage) WriteTo(stream net.Stream) error {
	raw, err := proto.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "failed to marshal message")
	}

	if _, err := stream.Write(raw); err != nil {
		return errors.Wrap(err, "failed to write RawMessage to stream")
	}
	return nil
}

func ReadRawMessage(stream net.Stream) (RawMessage, error) {
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

	msg := RawMessage{}
	err := proto.Unmarshal(raw, &msg)
	return msg, err
}

func MarshalOutgoingMessage(payload proto.Message, topic string, id peer.ID, pid protocol.ID) (*RawMessage, error) {
	data, err := proto.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	return &RawMessage{
		Message: pb.Message{
			Topic: topic,
			Data:  data,

			From:     []byte(id),
			Protocol: []byte(pid),
		},
	}, nil
}

type IncomingMessage struct {
	Payload    proto.Message
	SenderInfo peerstore.PeerInfo
	SenderAddr common.Address
	Protocol   Pid
}

func NewIncomingMessage(msg RawMessage, payloadType reflect.Type) (*IncomingMessage, error) {
	// unmarshal payload as given type
	payload, ok := reflect.New(payloadType).Interface().(proto.Message)
	if !ok {
		return nil, errors.New("payload of given message is not protobuf")
	}

	if err := proto.Unmarshal(msg.GetData(), payload); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal payload")
	}

	pid, err := msg.Pid()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse protocol")
	}
	m := &IncomingMessage{
		Payload:    payload,
		SenderInfo: peerstore.PeerInfo{ID: msg.ID()},
		Protocol:   pid,
	}

	// recover sender's public key from the ID.
	m.SenderAddr, err = AddrFromID(m.SenderInfo.ID)
	if err != nil {
		return m, errors.Wrap(err, "failed to recover sender address")
	}
	return m, nil
}

// AddrFromID converts a peer.ID (made from SECP256k1 public) into the Ethereum address.
func AddrFromID(id peer.ID) (addr common.Address, err error) {
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
