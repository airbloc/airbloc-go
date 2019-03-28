package api

import (
	"bytes"
	"context"
	"github.com/json-iterator/go"
	"time"

	"github.com/airbloc/logger"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	// JsonPbMarshaller is the marshaller used for serializing protobuf messages.
	JsonPbMarshaller = &jsonpb.Marshaler{}
)

const maxParameters = 2

func UnaryServerLogger() grpc.UnaryServerInterceptor {
	log := logger.New("api")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		timer := log.Timer()
		resp, err := handler(ctx, req)

		p, _ := peer.FromContext(ctx)
		reqLog := LogProtoMessage(req)

		if err != nil {
			code := codes.Internal
			if s, ok := status.FromError(err); ok {
				code = s.Code()
			} else {
				// wrap raw error
				err = status.Error(codes.Internal, err.Error())
			}
			elasped := time.Duration(time.Now().UnixNano() - timer.Time)
			log.Error("Error {} — {} ({}): {}", info.FullMethod, code, elasped, err.Error(), reqLog)

		} else {
			timer.End("{} from={}", info.FullMethod, p.Addr.String(), reqLog)
		}
		return resp, err
	}
}

// LogProtoMessage serializes proto.Message into loggable attributes.
func LogProtoMessage(msg interface{}) (attrs logger.Attrs) {
	protoMsg, ok := msg.(proto.Message)
	if !ok {
		return
	}
	rawJson := &bytes.Buffer{}
	if err := JsonPbMarshaller.Marshal(rawJson, protoMsg); err != nil {
		return
	}

	attrs = make(logger.Attrs)
	fullAttrs := make(map[string]interface{})
	json.Unmarshal(rawJson.Bytes(), &fullAttrs)

	// truncate only a few parameter
	n := 0
	for key, value := range fullAttrs {
		if n < maxParameters {
			break
		}
		attrs[key] = value
		n++
	}
	return
}
