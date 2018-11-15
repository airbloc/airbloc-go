package p2p

type Service interface {
	Topic()
	MessageType()

	RegisterAdapter(Adapter)
	UnregisterAdapter(Adapter)

	Handle(Message)
}
