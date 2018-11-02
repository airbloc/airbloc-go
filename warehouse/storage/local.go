package storage

import (
	"fmt"
	"github.com/airbloc/airbloc-go/warehouse/bundle"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

type LocalStorage struct {
	SavePath string
	Endpoint string
}

func NewLocalStorage(savePath string, endpoint string) *LocalStorage {
	return &LocalStorage{
		SavePath: savePath,
		Endpoint: endpoint,
	}
}

func (localStorage *LocalStorage) Save(bundle *bundle.Bundle) (*url.URL, error) {
	if _, err := os.Stat(localStorage.SavePath); os.IsNotExist(err) {
		return nil, errors.Errorf("the path %s does not exist", localStorage.SavePath)
	}

	fileName := fmt.Sprintf("%s.json", bundle.Id.String())
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

func (localStorage *LocalStorage) Update(url *url.URL, bundle *bundle.Bundle) error {
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
