package collections

import "github.com/airbloc/airbloc-go/adapter"

const (
	EventRegistered   = "Registered"
	EventUnregistered = "Unregistered"
	EventAllowed      = "Allowed"
	EventDenied       = "Denied"
)

func (s *Service) ParseRegisteredEvent(logData []byte) (
	event *adapter.CollectionRegistryRegistered,
	err error,
) {
	err = s.contractABI.Unpack(&event, EventRegistered, logData)
	return
}

func (s *Service) ParseUnregsiteredEvent(logData []byte) (
	event *adapter.CollectionRegistryUnregistered,
	err error,
) {
	err = s.contractABI.Unpack(&event, EventUnregistered, logData)
	return
}

func (s *Service) ParseAllowedEvent(logData []byte) (
	event *adapter.CollectionRegistryAllowed,
	err error,
) {
	err = s.contractABI.Unpack(&event, EventAllowed, logData)
	return
}

func (s *Service) ParseDenideEvent(logData []byte) (
	event *adapter.CollectionRegistryDenied,
	err error,
) {
	err = s.contractABI.Unpack(&event, EventDenied, logData)
	return
}
