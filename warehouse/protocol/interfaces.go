package protocol

import (
	"errors"
	"github.com/airbloc/airbloc-go/warehouse/bundle"
	"net/url"
)

var (
	ErrNotFound = errors.New("given URI is not found")
)

type Protocol interface {
	Name() string
	Read(uri url.URL) (*bundle.Bundle, error)
}
