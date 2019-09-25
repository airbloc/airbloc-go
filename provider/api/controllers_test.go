package api

import (
	"encoding/json"
	"net/http"
	"testing"

	adapterMock "github.com/airbloc/airbloc-go/shared/adapter/mocks"
	"github.com/airbloc/airbloc-go/shared/types"
	testutils "github.com/airbloc/airbloc-go/test/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/klaytn/klaytn/common"
	"github.com/stretchr/testify/assert"
)

var (
	testControllerAddr = "deadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
)

func TestControllerRegistryAPI_Register(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"controllerAddr": testControllerAddr}, nil)

	mockManager := adapterMock.NewMockIControllerRegistryManager(mockController)
	mockManager.EXPECT().
		Register(c, common.HexToAddress(testControllerAddr)).
		Return(nil)

	api := &controllerRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

func TestControllerRegistryAPI_Register_InvalidControllerAddr(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMock.NewMockIControllerRegistryManager(mockController)

	api := &controllerRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestControllerRegistryAPI_Register_FailedToRegister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"controllerAddr": testControllerAddr}, nil)

	mockManager := adapterMock.NewMockIControllerRegistryManager(mockController)
	mockManager.EXPECT().
		Register(c, common.HexToAddress(testControllerAddr)).
		Return(testutils.TestErr)

	api := &controllerRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

func TestControllerRegistryAPI_Get(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"controllerAddr": testControllerAddr}, nil)

	mockManager := adapterMock.NewMockIControllerRegistryManager(mockController)
	mockManager.EXPECT().
		Get(common.HexToAddress(testControllerAddr)).
		Return(types.DataController{}, nil)

	api := &controllerRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusOK, w.Code)
	d, _ := json.Marshal(types.DataController{})
	assert.Equal(t, string(d), w.Body.String())
}

func TestControllerRegistryAPI_Get_InvalidControllerAddr(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMock.NewMockIControllerRegistryManager(mockController)

	api := &controllerRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestControllerRegistryAPI_Get_FailedToGet(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"controllerAddr": testControllerAddr}, nil)

	mockManager := adapterMock.NewMockIControllerRegistryManager(mockController)
	mockManager.EXPECT().
		Get(common.HexToAddress(testControllerAddr)).
		Return(types.DataController{}, testutils.TestErr)

	api := &controllerRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}
