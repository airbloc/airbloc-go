package api

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/airbloc/airbloc-go/shared/adapter/mocks"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	testErr           = errors.New("error")
	testErrStr        = `{"error":"error"}`
	testIdHex         = "deadbeefdeadbeef"
	testSuccessMsg    = gin.H{"message": "success"}
	testSuccessMsgStr = `{"message":"success"}`
)

func createTestRequest(t *testing.T, msg gin.H, b binding.Binding) (*httptest.ResponseRecorder, *gin.Context) {
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
		assert.NoError(t, err)
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

// happy path
func TestAccountsAPI_Create(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{}, binding.JSON)
	mockManager.EXPECT().Create(c).Return(types.HexToID(testIdHex))

	api := accountsAPI{mockManager}
	api.create(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testIdHex), w.Body.String())
}

func TestAccountsAPI_Create_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{}, binding.JSON)
	mockManager.EXPECT().Create(c).Return(types.ID{}, testErr)

	api := accountsAPI{mockManager}
	api.create(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_CreateTemporary(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"identityHash": testIdHex}, binding.JSON)
	mockManager.EXPECT().CreateTemporary(c, common.HexToHash(testIdHex)).Return(types.HexToID(testIdHex))

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testIdHex), w.Body.String())
}

func TestAccountsAPI_CreateTemporary_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{}, binding.JSON)

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_CreateTemporary_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"identityHash": testIdHex}, binding.JSON)
	mockManager.EXPECT().CreateTemporary(c, common.HexToHash(testIdHex)).Return(types.ID{}, testErr)

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_UnlockTemporary(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{
		"identityPreimage":  testIdHex,
		"newOwner":          testIdHex,
		"passwordSignature": testIdHex,
	}, binding.JSON)
	passSig, _ := hex.DecodeString(testIdHex)

	mockManager.EXPECT().UnlockTemporary(
		c, common.HexToHash(testIdHex),
		common.HexToAddress(testIdHex),
		passSig,
	).Return(nil)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testSuccessMsgStr, w.Body.String())
}

func TestAccountsAPI_UnlockTemporary_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{}, binding.JSON)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_UnlockTemporary_InvalidPassordSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{
		"identityPreimage":  testIdHex,
		"newOwner":          testIdHex,
		"passwordSignature": testIdHex + "z", // make invalid hex
	}, binding.JSON)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_UnlockTemporary_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{
		"identityPreimage":  testIdHex,
		"newOwner":          testIdHex,
		"passwordSignature": testIdHex,
	}, binding.JSON)
	passSig, _ := hex.DecodeString(testIdHex)

	mockManager.EXPECT().UnlockTemporary(
		c, common.HexToHash(testIdHex),
		common.HexToAddress(testIdHex),
		passSig,
	).Return(testErr)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_SetController(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"controller": testIdHex}, binding.JSON)
	mockManager.EXPECT().SetController(c, common.HexToAddress(testIdHex)).Return(nil)

	api := accountsAPI{mockManager}
	api.setController(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testSuccessMsgStr, w.Body.String())
}

func TestAccountsAPI_SetController_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{}, binding.JSON)

	api := accountsAPI{mockManager}
	api.setController(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_SetController_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"controller": testIdHex}, binding.JSON)
	mockManager.EXPECT().SetController(c, common.HexToAddress(testIdHex)).Return(testErr)

	api := accountsAPI{mockManager}
	api.setController(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_GetAccount(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"accountId": testIdHex}, nil)
	accountId, _ := types.HexToID(testIdHex)
	mockManager.EXPECT().GetAccount(accountId).Return(types.Account{}, nil)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusOK, w.Code)
	d, _ := json.Marshal(types.Account{})
	assert.Equal(t, string(d), w.Body.String())
}

func TestAccountsAPI_GetAccount_InvalidParam(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{}, nil)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_GetAccount_InvalidAccountId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"accountId": testIdHex + "z"}, nil)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_GetAccount_FailedToGetAccount(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"accountId": testIdHex}, nil)
	accountId, _ := types.HexToID(testIdHex)
	mockManager.EXPECT().GetAccount(accountId).Return(types.Account{}, testErr)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_GetAccountId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"owner": testIdHex}, binding.Query)
	mockManager.EXPECT().GetAccountId(common.HexToAddress(testIdHex)).Return(types.HexToID(testIdHex))

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testIdHex), w.Body.String())
}

func TestAccountsAPI_GetAccountId_FailedToGetAccountId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{"owner": testIdHex}, binding.Query)
	mockManager.EXPECT().GetAccountId(common.HexToAddress(testIdHex)).Return(types.ID{}, testErr)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_GetAccountIdWithSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{
		"messageHash": testIdHex,
		"signature":   testIdHex,
	}, binding.Query)
	sig, _ := hex.DecodeString(testIdHex)

	mockManager.EXPECT().GetAccountIdFromSignature(common.HexToHash(testIdHex), sig).Return(types.HexToID(testIdHex))

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testIdHex), w.Body.String())
}

func TestAccountsAPI_GetAccountIdWithSignature_InvalidSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{
		"messageHash": testIdHex,
		"signature":   testIdHex + "z",
	}, binding.Query)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_GetAccountIdWithSignature_FailedToGetAccountIdFromSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{
		"messageHash": testIdHex,
		"signature":   testIdHex,
	}, binding.Query)
	sig, _ := hex.DecodeString(testIdHex)

	mockManager.EXPECT().GetAccountIdFromSignature(common.HexToHash(testIdHex), sig).Return(types.ID{}, testErr)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testErrStr, w.Body.String())
}

func TestAccountsAPI_GetAccount_BadRequest(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w, c := createTestRequest(t, gin.H{}, nil)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Bad Request"}`, w.Body.String())
}
