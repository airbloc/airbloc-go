package common

import (
	"bufio"
	"context"
	"io"
	"reflect"

	"github.com/ethereum/go-ethereum/common"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/pkg/errors"
)

type ProtoMessage struct {
	pb.Message
}

func NewProtoMessage(msg proto.Message, topic string) (*ProtoMessage, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal message")
	}

	return &ProtoMessage{
		Message: pb.Message{
			Topic: topic,
			Data:  data,
		},
	}, nil
}

func (message ProtoMessage) ID() peer.ID {
	return peer.ID(message.GetFrom())
}

func (message ProtoMessage) Pid() (Pid, error) {
	return ParsePid(string(message.GetProtocol()))
}

func ReadMessage(stream net.Stream) (ProtoMessage, error) {
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

	msg := ProtoMessage{}
	err := proto.Unmarshal(raw, &msg)
	return msg, err
}

func (message ProtoMessage) WriteMessage(stream net.Stream) error {
	raw, err := proto.Marshal(&message)
	if err != nil {
		return errors.Wrap(err, "proto error : failed to marshal proto message")
	}

	if _, err := stream.Write(raw); err != nil {
		return errors.Wrap(err, "proto error : failed to write data to stream")
	}
	return nil
}

func (message ProtoMessage) MakeMessage(ctx context.Context, typ reflect.Type) (Message, error) {
	msg, ok := reflect.New(typ).Interface().(proto.Message)
	if !ok {
		return Message{}, errors.New("message is not protobuf message")
	}

	if err := proto.Unmarshal(message.GetData(), msg); err != nil {
		return Message{}, errors.Wrap(err, "failed to unmarshal data")
	}

	pid, err := message.Pid()
	if err != nil {
		return Message{}, errors.Wrap(err, "failed to parse protocol")
	}
	m := NewMessage(msg, peerstore.PeerInfo{ID: message.ID()}, pid)

	// recover sender's public key from the ID.
	libp2pPubKey, err := message.ID().ExtractPublicKey()
	if err != nil {
		return Message{}, errors.Wrap(err, "failed to recover sender address")
	}
	pubKeyBytes, err := libp2pPubKey.Raw()
	if err != nil {
		return Message{}, errors.Wrap(err, "failed to recover sender address")
	}
	pubKey, err := crypto.DecompressPubkey(pubKeyBytes)
	if err != nil {
		return Message{}, errors.Wrap(err, "failed to recover sender address")
	}
	m.SenderAddr = crypto.PubkeyToAddress(*pubKey)
	return m, nil
}

func MessageType(msg proto.Message) reflect.Type {
	return reflect.ValueOf(msg).Elem().Type()
}

type Message struct {
	Data       proto.Message
	SenderInfo peerstore.PeerInfo
	SenderAddr common.Address
	Protocol   Pid
}

func NewMessage(data proto.Message, info peerstore.PeerInfo, protocol Pid) Message {
	return Message{
		Data:       data,
		SenderInfo: info,
		Protocol:   protocol,
	}
}
