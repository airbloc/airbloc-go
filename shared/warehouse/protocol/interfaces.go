package protocol

import (
	"github.com/pkg/errors"
	"net/url"

	"github.com/airbloc/airbloc-go/shared/types"
)

var (
	ErrNotFound = errors.New("given URI is not found")
)

type Protocol interface {
	Name() string
	Read(uri *url.URL) (*types.Bundle, error)
}
