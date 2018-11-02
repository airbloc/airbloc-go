package storage

import (
	"net/url"

	"github.com/airbloc/airbloc-go/warehouse/bundle"
)

type Storage interface {
	Save(*bundle.Bundle) (*url.URL, error)
	Update(*url.URL, *bundle.Bundle) error
	Delete(*url.URL) error
}
