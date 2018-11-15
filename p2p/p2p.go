package p2p

import p2p "github.com/airbloc/airbloc-go/proto/p2p"

type Adapter func(TopicHandler) TopicHandler
type ProtocolHandler func(p2p.Message)
type TopicHandler func(Message)
