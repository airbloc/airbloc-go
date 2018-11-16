package p2p

import (
	"context"
	"log"

	"github.com/airbloc/airbloc-go/p2p/common"
	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
)

func DatasyncReponse(s Server, ctx context.Context, msg common.Message) {

}

func DatasyncRequest(s Server, ctx context.Context, msg common.Message) {

}

func Ping(s Server, ctx context.Context, message common.Message) {
	log.Println("Ping", message.Info.ID.Pretty(), message.Data.String())

	pong, _ := proto.Marshal(&p2p.TestPing{Message: "World!"})
	pongMsg := common.ProtoMessage{
		Message: p2p.Message{
			Topic: p2p.Topic_TEST_PONG,
			Data:  pong,
		},
	}

	s.Send(ctx, pongMsg, message.Info.ID, message.Protocol)
}

func Pong(s Server, ctx context.Context, message common.Message) {
	log.Println("Pong", message.Info.ID.Pretty(), message.Data.String())
}
