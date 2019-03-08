package storage

import (
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/json-iterator/go"
	"net/url"
)

type Storage interface {
	Save(string, *types.Bundle) (*url.URL, error)
	Update(*url.URL, *types.Bundle) error
	Delete(*url.URL) error
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary
