package stream

import (
	"airbloc/go-producer/airbloc_producer"
	"context"
	"net"

	"google.golang.org/grpc"
)

type Stream struct {
	watcher
	gs *grpc.Server
	ls net.Listener
}

func (s *Stream) serve() error {
	return s.gs.Serve(s.ls)
}

func (s *Stream) Close() {
	s.gs.Stop()
	s.ls.Close()
}

func Run(
	context context.Context,
	network, address string,
	opts ...grpc.ServerOption,
) (stream *Stream, err error) {
	stream = new(Stream)
	stream.gs = grpc.NewServer(opts...)
	stream.watcher = watcher{
		Msg: make(chan *airbloc_producer.RawData),
		Err: make(chan error),
	}
	airbloc_producer.RegisterProducerServer(stream.gs, newServer(context, stream.watcher))
	stream.ls, err = net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return stream, stream.serve()
}
