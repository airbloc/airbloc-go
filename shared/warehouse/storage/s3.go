package storage

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"

	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

var S3ProtocolFmt = "%s.s3.%s.amazonaws.com"
var S3ProtocolExp = regexp.MustCompile(`https://([^.]+).s3.([^.]+).amazonaws.com/([^ ]+)`)

type S3Storage struct {
	client *s3.S3
	region string
	bucket string
	prefix string
}

func NewS3Storage(bucket, prefix string, sess *session.Session) Storage {
	return &S3Storage{
		client: s3.New(sess),
		region: *sess.Config.Region,
		bucket: bucket,
		prefix: prefix,
	}
}

func ExtractS3ObjectInfo(rawUrl string) (region, bucket, key string) {
	s := S3ProtocolExp.FindStringSubmatch(rawUrl)[1:]
	return s[0], s[1], s[2]
}

func (ss *S3Storage) Save(bundleId string, bundle *types.Bundle) (*url.URL, error) {
	bundleUri := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf(S3ProtocolFmt, ss.bucket, ss.region),
		Path:   filepath.Join(ss.prefix, bundleId),
	}
	err := ss.Update(bundleUri, bundle)
	if err != nil {
		return nil, err
	}
	return bundleUri, nil
}

func (ss *S3Storage) Update(bundlePath *url.URL, bundle *types.Bundle) error {
	bundleData, err := json.Marshal(bundle)
	if err != nil {
		return errors.Wrap(err, "s3 update: marshal error")
	}

	reqObj := &s3.PutObjectInput{
		Key:             aws.String(bundlePath.Path),
		Body:            bytes.NewReader(bundleData),
		Bucket:          aws.String(ss.bucket),
		ACL:             aws.String(s3.BucketCannedACLPublicRead),
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
		Key:    aws.String(bundlePath.Path),
		Bucket: aws.String(ss.bucket),
	}

	_, err := ss.client.DeleteObject(reqObj)
	if err != nil {
		return errors.Wrap(err, "s3 delete: request error")
	}
	return nil
}
