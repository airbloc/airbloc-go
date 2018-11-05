package exchange

import (
	"github.com/airbloc/airbloc-go/common"
)

type Status int8

const (
	StatusNeutral = Status(iota)
	StatusPending
	StatusSettled
	StatusRejected
	StatusOpened
	StatusClosed
)

type Offer struct {
	Offeror  common.Address
	Offeree  common.Address
	Contract common.Address
	Status   Status
}
