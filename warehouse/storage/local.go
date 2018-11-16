package storage

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"

	"github.com/airbloc/airbloc-go/data"
	"github.com/pkg/errors"
)

type LocalStorage struct {
	SavePath string
	Endpoint string
}

func NewLocalStorage(savePath string, endpoint string) (*LocalStorage, error) {
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err := os.MkdirAll(savePath, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create save path")
		}
	}

	return &LocalStorage{
		SavePath: savePath,
		Endpoint: endpoint,
	}, nil
}

func (localStorage *LocalStorage) Save(bundleId string, bundle *data.Bundle) (*url.URL, error) {
	fileName := fmt.Sprintf("%s.json", bundleId)
	savePath := path.Join(localStorage.SavePath, fileName)

	bundleData, err := bundle.Marshal()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse bundle")
	}

	if err := ioutil.WriteFile(savePath, bundleData, 0644); err != nil {
		return nil, errors.Wrap(err, "failed to write bundle to the path")
	}

	return url.Parse(path.Join(localStorage.Endpoint, fileName))
}

func (localStorage *LocalStorage) Update(url *url.URL, bundle *data.Bundle) error {
	_, fileName := path.Split(url.Path)
	savedPath := path.Join(localStorage.SavePath, fileName)

	if _, err := os.Stat(savedPath); os.IsNotExist(err) {
		return errors.Errorf("the given URI %s does not exist", url.String())
	}

	bundleData, err := bundle.Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to parse bundle")
	}
	if err := ioutil.WriteFile(savedPath, bundleData, 0644); err != nil {
		return errors.Wrap(err, "failed to write bundle to the path")
	}
	return nil
}

func (localStorage *LocalStorage) Delete(url *url.URL) error {
	_, fileName := path.Split(url.Path)
	savedPath := path.Join(localStorage.SavePath, fileName)

	if _, err := os.Stat(savedPath); os.IsNotExist(err) {
		return errors.Errorf("the given URI %s does not exist", url.String())
	}

	if err := os.Remove(savedPath); err != nil {
		return errors.Errorf("failed to delete the bundle at %s", url.String())
	}
	return nil
}
