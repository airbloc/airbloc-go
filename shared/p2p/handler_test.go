package p2p

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/p2p/common"
	"log"

	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
)

func testPingHandler(s Server, ctx context.Context, message *IncomingMessage) {
	log.Println("Ping", message.SenderInfo.ID.Pretty(), message.Payload.String())

	s.Send(ctx, &pb.TestPing{Message: "World!"}, "ping", message.SenderInfo.ID)
}

func testPongHandler(s Server, ctx context.Context, message *IncomingMessage) {
	log.Println("Pong", message.SenderInfo.ID.Pretty(), message.Payload.String())
}
