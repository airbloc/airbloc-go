package testutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

var (
	TestErr              = errors.New("error")
	TestErrStr           = `{"error":"error"}`
	TestErrBadRequestStr = `{"error":"Bad Request"}`
	TestIdHex            = "deadbeefdeadbeef"
	TestSuccessStr       = `{"message":"success"}`
)

// CreateTestRequest creates gin TestContext and inject request data in context and return.
func CreateTestRequest(t *testing.T, msg gin.H, b binding.Binding) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var err error
	c.Request, err = http.NewRequest("", "", nil)
	assert.NoError(t, err)

	switch b {
	case binding.JSON:
		var d []byte
		d, err = json.Marshal(msg)
		assert.NoError(t, err)

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(d))
	case binding.Query:
		q := c.Request.URL.Query()
		for key, val := range msg {
			q.Add(key, val.(string))
		}
		c.Request.URL.RawQuery = q.Encode()
	default:
		for key, val := range msg {
			c.Params = append(c.Params, gin.Param{
				Key:   key,
				Value: val.(string),
			})
		}
	}

	return w, c
}
