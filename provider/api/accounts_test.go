package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	adapterMocks "github.com/airbloc/airbloc-go/shared/adapter/mocks"
	"github.com/airbloc/airbloc-go/shared/types"
	testutils "github.com/airbloc/airbloc-go/test/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// happy path
func TestAccountsAPI_Create(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		Create(c).
		Return(types.HexToID(testutils.TestIdHex))

	api := accountsAPI{mockManager}
	api.create(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testutils.TestIdHex), w.Body.String())
}

func TestAccountsAPI_Create_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		Create(c).
		Return(types.ID{}, testutils.TestErr)

	api := accountsAPI{mockManager}
	api.create(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_CreateTemporary(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"identityHash": testutils.TestIdHex}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		CreateTemporary(c, common.HexToHash(testutils.TestIdHex)).
		Return(types.HexToID(testutils.TestIdHex))

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testutils.TestIdHex), w.Body.String())
}

func TestAccountsAPI_CreateTemporary_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_CreateTemporary_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"identityHash": testutils.TestIdHex}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		CreateTemporary(c, common.HexToHash(testutils.TestIdHex)).
		Return(types.ID{}, testutils.TestErr)

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_UnlockTemporary(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"identityPreimage":  testutils.TestIdHex,
		"newOwner":          testutils.TestIdHex,
		"passwordSignature": testutils.TestIdHex,
	}, binding.JSON)
	passSig, _ := hex.DecodeString(testutils.TestIdHex)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		UnlockTemporary(
			c, common.HexToHash(testutils.TestIdHex),
			common.HexToAddress(testutils.TestIdHex),
			passSig,
		).Return(nil)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

func TestAccountsAPI_UnlockTemporary_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_UnlockTemporary_InvalidPassordSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"identityPreimage":  testutils.TestIdHex,
		"newOwner":          testutils.TestIdHex,
		"passwordSignature": testutils.TestIdHex + "z", // make invalid hex
	}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_UnlockTemporary_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"identityPreimage":  testutils.TestIdHex,
		"newOwner":          testutils.TestIdHex,
		"passwordSignature": testutils.TestIdHex,
	}, binding.JSON)
	passSig, _ := hex.DecodeString(testutils.TestIdHex)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		UnlockTemporary(
			c, common.HexToHash(testutils.TestIdHex),
			common.HexToAddress(testutils.TestIdHex),
			passSig,
		).Return(testutils.TestErr)

	api := accountsAPI{mockManager}
	api.unlockTemporary(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_SetController(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"controller": testutils.TestIdHex}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		SetController(c, common.HexToAddress(testutils.TestIdHex)).
		Return(nil)

	api := accountsAPI{mockManager}
	api.setController(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

func TestAccountsAPI_SetController_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.setController(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_SetController_Conflict(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"controller": testutils.TestIdHex}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		SetController(c, common.HexToAddress(testutils.TestIdHex)).
		Return(testutils.TestErr)

	api := accountsAPI{mockManager}
	api.setController(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_GetAccount(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"accountId": testutils.TestIdHex}, nil)
	accountId, _ := types.HexToID(testutils.TestIdHex)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccount(accountId).
		Return(types.Account{}, nil)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusOK, w.Code)
	d, _ := json.Marshal(types.Account{})
	assert.Equal(t, string(d), w.Body.String())
}

func TestAccountsAPI_GetAccount_InvalidParam(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_GetAccount_InvalidAccountId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"accountId": testutils.TestIdHex + "z"}, nil)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_GetAccount_FailedToGetAccount(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"accountId": testutils.TestIdHex}, nil)
	accountId, _ := types.HexToID(testutils.TestIdHex)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccount(accountId).
		Return(types.Account{}, testutils.TestErr)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_GetAccountId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"owner": testutils.TestIdHex}, binding.Query)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountId(common.HexToAddress(testutils.TestIdHex)).
		Return(types.HexToID(testutils.TestIdHex))

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testutils.TestIdHex), w.Body.String())
}

func TestAccountsAPI_GetAccountId_FailedToGetAccountId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"owner": testutils.TestIdHex}, binding.Query)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountId(common.HexToAddress(testutils.TestIdHex)).
		Return(types.ID{}, testutils.TestErr)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

// happy path
func TestAccountsAPI_GetAccountIdWithSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"messageHash": testutils.TestIdHex,
		"signature":   testutils.TestIdHex,
	}, binding.Query)
	sig, _ := hex.DecodeString(testutils.TestIdHex)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountIdFromSignature(common.HexToHash(testutils.TestIdHex), sig).
		Return(types.HexToID(testutils.TestIdHex))

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testutils.TestIdHex), w.Body.String())
}

func TestAccountsAPI_GetAccountIdWithSignature_InvalidSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"messageHash": testutils.TestIdHex,
		"signature":   testutils.TestIdHex + "z",
	}, binding.Query)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_GetAccountIdWithSignature_FailedToGetAccountIdFromSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"messageHash": testutils.TestIdHex,
		"signature":   testutils.TestIdHex,
	}, binding.Query)
	sig, _ := hex.DecodeString(testutils.TestIdHex)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountIdFromSignature(common.HexToHash(testutils.TestIdHex), sig).
		Return(types.ID{}, testutils.TestErr)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

func TestAccountsAPI_GetAccount_BadRequest(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Bad Request"}`, w.Body.String())
}
