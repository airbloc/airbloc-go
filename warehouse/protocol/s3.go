package protocol

import (
	"bufio"
	"github.com/airbloc/airbloc-go/data"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/url"
	"path"
	"strings"
)

type S3Protocol struct {
	client *s3.S3
}

func NewS3Protocol(sess *session.Session) Protocol {
	return &S3Protocol{client: s3.New(sess)}
}

func (s3p *S3Protocol) Name() string { return "s3" }
func (s3p *S3Protocol) Read(uri *url.URL) (*data.Bundle, error) {
	tmp := strings.Split(uri.Path, "/")
	region, bucket := tmp[0], tmp[1]
	bundlePath := strings.TrimPrefix(uri.Path, path.Join(region, bucket)+"/")

	reqObj := &s3.GetObjectInput{
		Key:    aws.String(bundlePath),
		Bucket: aws.String(bucket),
	}

	resObj, err := s3p.client.GetObject(reqObj)
	if err != nil {
		return nil, errors.Wrap(err, "s3 read: request error")
	}
	defer resObj.Body.Close()

	reader := bufio.NewReader(resObj.Body)

	var buf []byte
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "s3 read: read error")
		}
		buf = append(buf, b)
	}

	log.Println(string(buf))
	bundle, err := data.UnmarshalBundle(buf)
	if err != nil {
		return nil, errors.Wrap(err, "s3 read: unmarshal error")
	}
	return bundle, nil
}
