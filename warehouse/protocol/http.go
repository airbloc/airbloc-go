package protocol

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

type HttpProtocol struct {
	timeout time.Duration
	client  *fasthttp.Client
}

func NewHttpProtocol(timeout time.Duration, maxConnsPerHost int) *HttpProtocol {
	client := &fasthttp.Client{
		ReadTimeout:     timeout,
		MaxConnsPerHost: maxConnsPerHost,
	}
	return &HttpProtocol{timeout, client}
}

func (http *HttpProtocol) Name() string {
	return "http"
}

func (http *HttpProtocol) Read(uri *url.URL) ([]byte, error) {
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(uri.String())

	response := fasthttp.AcquireResponse()
	if err := http.client.DoTimeout(request, response, http.timeout); err != nil {
		return nil, errors.Wrap(err, "failed to estabilish HTTP connection")
	}

	if response.Header.StatusCode() == 404 {
		return nil, ErrNotFound
	}
	return response.Body(), nil
}

type HttpsProtocol struct {
	HttpProtocol
}

func NewHttpsProtocol(timeout time.Duration, maxConnsPerHost int) *HttpsProtocol {
	return &HttpsProtocol{HttpProtocol: *NewHttpProtocol(timeout, maxConnsPerHost)}
}

func (https *HttpsProtocol) Name() string {
	return "https"
}
