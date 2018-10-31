package collections

import "github.com/airbloc/airbloc-go/adapter"

const (
	EventRegistered   = "Registered"
	EventUnregistered = "Unregistered"
	EventAllowed      = "Allowed"
	EventDenied       = "Denied"
)

func (adt *Adapter) ParseRegisteredEvent(logData []byte) (
	event *adapter.CollectionRegistryRegistered,
	err error,
) {
	err = adt.contractABI.Unpack(&event, EventRegistered, logData)
	return
}

func (adt *Adapter) ParseUnregsiteredEvent(logData []byte) (
	event *adapter.CollectionRegistryUnregistered,
	err error,
) {
	err = adt.contractABI.Unpack(&event, EventUnregistered, logData)
	return
}

func (adt *Adapter) ParseAllowedEvent(logData []byte) (
	event *adapter.CollectionRegistryAllowed,
	err error,
) {
	err = adt.contractABI.Unpack(&event, EventAllowed, logData)
	return
}

func (adt *Adapter) ParseDenideEvent(logData []byte) (
	event *adapter.CollectionRegistryDenied,
	err error,
) {
	err = adt.contractABI.Unpack(&event, EventDenied, logData)
	return
}
