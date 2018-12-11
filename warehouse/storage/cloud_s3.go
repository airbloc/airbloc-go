package storage

import (
	"bytes"
	"github.com/airbloc/airbloc-go/data"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"net/url"
	"path"
	"path/filepath"
)

type S3Storage struct {
	client *s3.S3
	bucket string
	prefix string
}

func NewS3Storage(bucket, prefix string, sess *session.Session, opts ...*aws.Config) Storage {
	return &S3Storage{
		client: s3.New(sess, opts...),
		bucket: bucket,
		prefix: prefix,
	}
}

func (ss *S3Storage) Save(bundleId string, bundle *data.Bundle) (*url.URL, error) {
	bundlePath := filepath.Join(ss.prefix, bundleId+".bundle")
	bundleUri := &url.URL{
		Scheme: "s3",
		Path:   path.Join(ss.bucket, bundlePath),
	}
	return bundleUri, ss.Update(bundleUri, bundle)
}

func (ss *S3Storage) Update(bundlePath *url.URL, bundle *data.Bundle) error {
	if bundlePath.Scheme != "s3" {
		return errors.New("s3 update: invalid bundle url")
	}
	bundleData, err := bundle.Marshal()
	if err != nil {
		return errors.Wrap(err, "s3 update: marshal error")
	}
	reqObj := &s3.PutObjectInput{
		Key:             aws.String(bundlePath.String()),
		Body:            bytes.NewReader(bundleData),
		Bucket:          aws.String(ss.bucket),
		ContentType:     aws.String("application/json"),
		ContentEncoding: aws.String("utf-8"),
	}

	_, err = ss.client.PutObject(reqObj)
	if err != nil {
		return errors.Wrap(err, "s3 update: request error")
	}
	return nil
}

func (ss *S3Storage) Delete(bundlePath *url.URL) error {
	reqObj := &s3.DeleteObjectInput{
		Key:    aws.String(bundlePath.String()),
		Bucket: aws.String(ss.bucket),
	}

	_, err := ss.client.DeleteObject(reqObj)
	if err != nil {
		return errors.Wrap(err, "s3 delete: request error")
	}
	return nil
}
