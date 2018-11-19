package protocol

import (
	"errors"
	"net/url"

	"github.com/airbloc/airbloc-go/data"
)

var (
	ErrNotFound = errors.New("given URI is not found")
)

type Protocol interface {
	Name() string
	Read(uri *url.URL) (*data.Bundle, error)
}
