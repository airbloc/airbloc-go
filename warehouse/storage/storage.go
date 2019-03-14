package storage

import (
	"github.com/json-iterator/go"
	"net/url"
)

type Storage interface {
	Save(string, []byte) (*url.URL, error)
	Update(*url.URL, []byte) error
	Delete(*url.URL) error
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary
