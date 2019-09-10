package e2eutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type payloadType uint8
type httpMethod string

const (
	RequestJSON = payloadType(iota)
	RequestQuery
	RequestParam

	MethodGet     = httpMethod(http.MethodGet)
	MethodHead    = httpMethod(http.MethodHead)
	MethodPost    = httpMethod(http.MethodPost)
	MethodPut     = httpMethod(http.MethodPut)
	MethodPatch   = httpMethod(http.MethodPatch)
	MethodDelete  = httpMethod(http.MethodDelete)
	MethodConnect = httpMethod(http.MethodConnect)
	MethodOptions = httpMethod(http.MethodOptions)
	MethodTrace   = httpMethod(http.MethodTrace)
)

var (
	ErrUnsupportedPayloadType = errors.New("unsupported payload type")
)

type request struct {
	*http.Request
}

func (req request) Do() (resp *http.Response, body []byte, rerr error) {
	var err error
	client := http.DefaultClient
	resp, err = client.Do(req.Request)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			rerr = err
		}
	}()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}

func CreateRequest(method httpMethod, url string, payload map[string]interface{}, typ payloadType) (request, error) {
	req, err := http.NewRequest(string(method), url, nil)
	if err != nil {
		return request{}, err
	}

	switch typ {
	case RequestJSON:
		var d []byte
		d, err = json.Marshal(payload)
		if err != nil {
			return request{}, err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(d))
	case RequestQuery:
		q := req.URL.Query()
		for key, val := range payload {
			q.Add(key, val.(string))
		}
		req.URL.RawQuery = q.Encode()
	case RequestParam:
		for key, val := range payload {
			req.Form.Add(key, val.(string))
		}
	default:
		return request{}, ErrUnsupportedPayloadType
	}

	return request{req}, nil
}
