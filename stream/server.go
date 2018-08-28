package stream

import (
	"airbloc/go-producer/airbloc_producer"
	"context"
	"log"
)

type watcher struct {
	Msg chan *airbloc_producer.RawData
	Err chan error
}

type server struct {
	ctx context.Context
	watcher
}

func newServer(ctx context.Context, watcher watcher) *server {
	return &server{ctx, watcher}
}

func (s *server) AddData(stream airbloc_producer.Producer_AddDataServer) error {
	log.Println("Started stream")
	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			if in, err := stream.Recv(); err != nil {
				s.Err <- err
			} else {
				s.Msg <- in
			}
		}
	}
}
