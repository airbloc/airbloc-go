package p2p

import (
	"context"
	"log"

	"github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
)

func DatasyncReponse(s Server, ctx context.Context, msg common.Message) {

}

func DatasyncRequest(s Server, ctx context.Context, msg common.Message) {

}

func Ping(s Server, ctx context.Context, message common.Message) {
	log.Println("Ping", message.Info.ID.Pretty(), message.Data.String())

	s.Send(ctx, &pb.TestPing{Message: "World!"}, "ping", message.Info.ID)
}

func Pong(s Server, ctx context.Context, message common.Message) {
	log.Println("Pong", message.Info.ID.Pretty(), message.Data.String())
}
