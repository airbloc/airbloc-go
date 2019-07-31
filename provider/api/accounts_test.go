package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/airbloc/airbloc-go/shared/adapter/mocks"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const testIdHex = "deadbeefdeadbeef"
const testErrStr = "error"

// happy path
func TestNewAccountsAPI_Create(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockManager.EXPECT().Create(c).Return(types.HexToID(testIdHex))

	api := accountsAPI{mockManager}
	api.create(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testIdHex), w.Body.String())
}

func TestNewAccountAPI_Create_Fail(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockManager.EXPECT().Create(c).Return(types.ID{}, errors.New(testErrStr))

	api := accountsAPI{mockManager}
	api.create(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"message":"%s"}`, testErrStr), w.Body.String())
}

// happy path
func TestNewAccountsAPI_CreateTemporary(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockManager := mocks.NewMockIAccountsManager(mockController)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("", "", strings.NewReader(fmt.Sprintf(`{"identityHash":"%s"}`, testIdHex)))

	mockManager.EXPECT().CreateTemporary(c, common.HexToHash(testIdHex)).Return(types.HexToID(testIdHex))

	api := accountsAPI{mockManager}
	api.createTemporary(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"accountId":"%s"}`, testIdHex), w.Body.String())
}

//
//func TestNewAccountsAPI_UnlockTemporary(t *testing.T) {
//	mockController := gomock.NewController(t)
//	defer mockController.Finish()
//
//	mockManager := mocks.NewMockIAccountsManager(mockController)
//
//	resp := httptest.NewRecorder()
//	ctx, _ := gin.CreateTestContext(resp)
//
//}
//
//func TestNewAccountsAPI_SetController(t *testing.T) {
//	mockController := gomock.NewController(t)
//	defer mockController.Finish()
//
//	mockManager := mocks.NewMockIAccountsManager(mockController)
//
//	resp := httptest.NewRecorder()
//	ctx, _ := gin.CreateTestContext(resp)
//
//}
//
//func TestNewAccountsAPI_GetAccount(t *testing.T) {
//	mockController := gomock.NewController(t)
//	defer mockController.Finish()
//
//	mockManager := mocks.NewMockIAccountsManager(mockController)
//
//	resp := httptest.NewRecorder()
//	ctx, _ := gin.CreateTestContext(resp)
//
//}
//
//func TestNewAccountsAPI_GetAccountId(t *testing.T) {
//	mockController := gomock.NewController(t)
//	defer mockController.Finish()
//
//	mockManager := mocks.NewMockIAccountsManager(mockController)
//
//	resp := httptest.NewRecorder()
//	ctx, _ := gin.CreateTestContext(resp)
//
//}
