package storage

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/airbloc/airbloc-go/data"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type GoogleCloudStorage struct {
	ctx       context.Context
	client    *storage.Client
	projectId string
	bucket    string
	prefix    string
}

func NewGoogleCloudStorage(ctx context.Context, bucket, prefix, projectId string, opts ...option.ClientOption) (Storage, error) {
	client, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "gs init: client error")
	}

	bk := client.Bucket(bucket)
	_, err = bk.Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		err = bk.Create(ctx, projectId, nil)
		if err != nil {
			return nil, errors.Wrap(err, "gs init: create bucket error")
		}
	} else if err != nil {
		return nil, errors.Wrap(err, "gs init: get bucket attr error")
	}

	return &GoogleCloudStorage{
		ctx:       ctx,
		client:    client,
		projectId: projectId,
		bucket:    bucket,
		prefix:    prefix,
	}, nil
}

func (gs *GoogleCloudStorage) getBucket() *storage.BucketHandle {
	return gs.client.Bucket(gs.bucket)
}

//gs://[prefix]/[bundleId].bundle
func (gs *GoogleCloudStorage) Save(bundleId string, bundle *data.Bundle) (*url.URL, error) {
	bundlePath := filepath.Join(gs.prefix, bundleId+".bundle")
	bundleUrl := &url.URL{
		Scheme: "gs",
		Path:   bundlePath,
	}
	return bundleUrl, gs.Update(bundleUrl, bundle)
}

func (gs *GoogleCloudStorage) Update(bundlePath *url.URL, bundle *data.Bundle) error {
	bundleData, err := bundle.Marshal()
	if err != nil {
		return errors.Wrap(err, "gs update: marshal error")
	}
	bucket := gs.getBucket()
	writer := bucket.Object(bundlePath.Path).NewWriter(gs.ctx)
	if _, err = io.Copy(writer, bytes.NewReader(bundleData)); err != nil {
		return errors.Wrap(err, "gs update: write error")
	}
	if err = writer.Close(); err != nil {
		return errors.Wrap(err, "gs update: cannot close writer")
	}
	return nil
}

func (gs *GoogleCloudStorage) Delete(bundlePath *url.URL) error {
	if bundlePath.Scheme != "gs" {
		return errors.New("gs delete: invalid bundle url")
	}
	bucket := gs.getBucket()
	if err := bucket.Object(bundlePath.Path).Delete(gs.ctx); err != nil {
		return errors.Wrap(err, "gs delete: cannot delete object")
	}
	return nil
}
