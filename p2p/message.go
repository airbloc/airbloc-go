package p2p

import (
	"reflect"

	"github.com/gogo/protobuf/proto"
)

func messageType(msg proto.Message) reflect.Type {
	return reflect.ValueOf(msg).Elem().Type()
}
