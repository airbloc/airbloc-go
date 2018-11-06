package p2p

import "github.com/airbloc/airbloc-go/api"

type Sender interface {
	Send()
}

type Receiver interface {
	Receive()
}

type Client interface {
	Find()
	Connect()
	Sender
	Receiver
	Disconnect()
}

type Node interface {
	api.Service
	Client
}
