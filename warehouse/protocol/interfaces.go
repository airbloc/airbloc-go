package protocol

import (
	"errors"
	"net/url"

	"github.com/airbloc/airbloc-go/warehouse/bundle"
)

var (
	ErrNotFound = errors.New("given URI is not found")
)

type Protocol interface {
	Name() string
	Read(uri url.URL) (*bundle.Bundle, error)
}
