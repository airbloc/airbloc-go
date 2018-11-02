package storage

import (
	"github.com/airbloc/airbloc-go/warehouse/bundle"
	"net/url"
)

type Storage interface {
	Save(*bundle.Bundle) (*url.URL, error)
	Update(*url.URL, *bundle.Bundle) error
	Delete(*url.URL) error
}
