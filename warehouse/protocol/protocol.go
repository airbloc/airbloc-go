package protocol

import (
	"github.com/pkg/errors"
	"net/url"
)

var (
	ErrNotFound = errors.New("given URI is not found")
)

type Protocol interface {
	Name() string
	Read(uri *url.URL) ([]byte, error)
}
