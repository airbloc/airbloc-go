package p2p

type Adapter func(Handler) Handler
type Handler func(Message)
