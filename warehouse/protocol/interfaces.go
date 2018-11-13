package protocol

import (
	"errors"
	"github.com/airbloc/airbloc-go/data"
	"net/url"
)

var (
	ErrNotFound = errors.New("given URI is not found")
)

type Protocol interface {
	Name() string
	Read(uri *url.URL) (*data.Bundle, error)
}
