package storage

import (
	"github.com/json-iterator/go"
	"net/url"

	"github.com/airbloc/airbloc-go/data"
)

type Storage interface {
	Save(string, *data.Bundle) (*url.URL, error)
	Update(*url.URL, *data.Bundle) error
	Delete(*url.URL) error
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary
