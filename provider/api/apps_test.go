package api

import (
	"net/http"
	"testing"

	adapterMock "github.com/airbloc/airbloc-go/shared/adapter/mocks"
	"github.com/airbloc/airbloc-go/shared/types"
	testutils "github.com/airbloc/airbloc-go/test/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	testAppName = "test"
)

func TestAppRegistryAPI_Register(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		Register(c, testAppName).
		Return(nil)

	api := &appRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

func TestAppRegistryAPI_Register_InvalidAppName(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)

	api := &appRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestAppRegistryAPI_Register_FailedToRegister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		Register(c, testAppName).
		Return(testutils.TestErr)

	api := &appRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())

}

func TestAppRegistryAPI_Unregister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		Unregister(c, testAppName).
		Return(nil)

	api := &appRegistryAPI{mockManager}
	api.unregister(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())

}

func TestAppRegistryAPI_Unregister_InvalidAppName(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)

	api := &appRegistryAPI{mockManager}
	api.unregister(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestAppRegistryAPI_Unregister_FailedToUnregister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		Unregister(c, testAppName).
		Return(testutils.TestErr)

	api := &appRegistryAPI{mockManager}
	api.unregister(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

func TestAppRegistryAPI_Get(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		Get(testAppName).
		Return(types.App{}, nil)

	api := &appRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusOK, w.Code)
	d, _ := json.Marshal(types.App{})
	assert.Equal(t, string(d), w.Body.String())
}

func TestAppRegistryAPI_Get_InvalidAppName(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)

	api := &appRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestAppRegistryAPI_Get_FailedToGet(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		Get(testAppName).
		Return(types.App{}, testutils.TestErr)

	api := &appRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

func TestAppRegistryAPI_TransferAppOwner(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName, "newOwner": testAccountId}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		TransferAppOwner(c, testAppName, common.HexToAddress(testAccountId)).
		Return(nil)

	api := &appRegistryAPI{mockManager}
	api.transferAppOwner(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

func TestAppRegistryAPI_TransferAppOwner_InvalidAppName(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"newOwner": testAccountId}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)

	api := &appRegistryAPI{mockManager}
	api.transferAppOwner(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestAppRegistryAPI_TransferAppOwner_InvalidNewOwner(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)

	api := &appRegistryAPI{mockManager}
	api.transferAppOwner(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestAppRegistryAPI_TransferAppOwner_FailedToTransferAppOwner(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"appName": testAppName, "newOwner": testAccountId}, nil)

	mockManager := adapterMock.NewMockIAppRegistryManager(mockController)
	mockManager.EXPECT().
		TransferAppOwner(c, testAppName, common.HexToAddress(testAccountId)).
		Return(testutils.TestErr)

	api := &appRegistryAPI{mockManager}
	api.transferAppOwner(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}
