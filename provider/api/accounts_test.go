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

var (
	testAccountId = "deadbeefdeadbeef"
)

// happy path
func TestAccountsAPI_Create(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		Create(c).
		Return(types.HexToID(testAccountId))

	api := accountsAPI{mockManager}
	api.create(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testAccountId), w.Body.String())
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

	w, c := testutils.CreateTestRequest(t, gin.H{"identityHash": testAccountId}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		CreateTemporary(c, common.HexToHash(testAccountId)).
		Return(types.HexToID(testAccountId))

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testAccountId), w.Body.String())
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

	w, c := testutils.CreateTestRequest(t, gin.H{"identityHash": testAccountId}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		CreateTemporary(c, common.HexToHash(testAccountId)).
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
		"identityPreimage":  testAccountId,
		"newOwner":          testAccountId,
		"passwordSignature": testAccountId,
	}, binding.JSON)
	passSig, _ := hex.DecodeString(testAccountId)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		UnlockTemporary(
			c, common.HexToHash(testAccountId),
			common.HexToAddress(testAccountId),
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
		"identityPreimage":  testAccountId,
		"newOwner":          testAccountId,
		"passwordSignature": testAccountId + "z", // make invalid hex
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
		"identityPreimage":  testAccountId,
		"newOwner":          testAccountId,
		"passwordSignature": testAccountId,
	}, binding.JSON)
	passSig, _ := hex.DecodeString(testAccountId)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		UnlockTemporary(
			c, common.HexToHash(testAccountId),
			common.HexToAddress(testAccountId),
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

	w, c := testutils.CreateTestRequest(t, gin.H{"controller": testAccountId}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		SetController(c, common.HexToAddress(testAccountId)).
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

	w, c := testutils.CreateTestRequest(t, gin.H{"controller": testAccountId}, binding.JSON)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		SetController(c, common.HexToAddress(testAccountId)).
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

	w, c := testutils.CreateTestRequest(t, gin.H{"accountId": testAccountId}, nil)
	accountId, _ := types.HexToID(testAccountId)

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

	w, c := testutils.CreateTestRequest(t, gin.H{"accountId": testAccountId + "z"}, nil)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)

	api := accountsAPI{mockManager}
	api.getAccount(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestAccountsAPI_GetAccount_FailedToGetAccount(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"accountId": testAccountId}, nil)
	accountId, _ := types.HexToID(testAccountId)

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

	w, c := testutils.CreateTestRequest(t, gin.H{"owner": testAccountId}, binding.Query)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountId(common.HexToAddress(testAccountId)).
		Return(types.HexToID(testAccountId))

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testAccountId), w.Body.String())
}

func TestAccountsAPI_GetAccountId_FailedToGetAccountId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"owner": testAccountId}, binding.Query)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountId(common.HexToAddress(testAccountId)).
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
		"messageHash": testAccountId,
		"signature":   testAccountId,
	}, binding.Query)
	sig, _ := hex.DecodeString(testAccountId)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountIdFromSignature(common.HexToHash(testAccountId), sig).
		Return(types.HexToID(testAccountId))

	api := accountsAPI{mockManager}
	api.getAccountId(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testAccountId), w.Body.String())
}

func TestAccountsAPI_GetAccountIdWithSignature_InvalidSignature(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"messageHash": testAccountId,
		"signature":   testAccountId + "z",
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
		"messageHash": testAccountId,
		"signature":   testAccountId,
	}, binding.Query)
	sig, _ := hex.DecodeString(testAccountId)

	mockManager := adapterMocks.NewMockIAccountsManager(mockController)
	mockManager.EXPECT().
		GetAccountIdFromSignature(common.HexToHash(testAccountId), sig).
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
