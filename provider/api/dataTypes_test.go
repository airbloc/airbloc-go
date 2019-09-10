package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	adapterMock "github.com/airbloc/airbloc-go/shared/adapter/mocks"
	"github.com/airbloc/airbloc-go/shared/types"
	testutils "github.com/airbloc/airbloc-go/test/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	testDataType   = "test"
	testSchemaHash = "deadbeefdeadbeefdeadbeef"
)

func TestDataTypeRegistryAPI_Register(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"name":       testDataType,
		"schemaHash": testSchemaHash,
	}, binding.JSON)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)
	mockManager.EXPECT().
		Register(c, testDataType, common.HexToHash(testSchemaHash)).
		Return(nil)

	api := &dataTypeRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

func TestDataTypeRegistryAPI_Register_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)

	api := &dataTypeRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestDataTypeRegistryAPI_Register_FailedToRegister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"name":       testDataType,
		"schemaHash": testSchemaHash,
	}, binding.JSON)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)
	mockManager.EXPECT().
		Register(c, testDataType, common.HexToHash(testSchemaHash)).
		Return(testutils.TestErr)

	api := &dataTypeRegistryAPI{mockManager}
	api.register(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

func TestDataTypeRegistryAPI_Unregister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"dataType": testDataType}, nil)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)
	mockManager.EXPECT().
		Unregister(c, testDataType).
		Return(nil)

	api := &dataTypeRegistryAPI{mockManager}
	api.unregister(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

func TestDataTypeRegistryAPI_Unregister_InvalidDataType(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)

	api := &dataTypeRegistryAPI{mockManager}
	api.unregister(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestDataTypeRegistryAPI_Unregister_FailedToUnregister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"dataType": testDataType}, nil)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)
	mockManager.EXPECT().
		Unregister(c, testDataType).
		Return(testutils.TestErr)

	api := &dataTypeRegistryAPI{mockManager}
	api.unregister(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

func TestDataTypeRegistryAPI_Get(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"dataType": testDataType}, nil)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)
	mockManager.EXPECT().
		Get(testDataType).
		Return(types.DataType{}, nil)

	api := &dataTypeRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusOK, w.Code)
	d, _ := json.Marshal(types.DataType{})
	assert.Equal(t, string(d), w.Body.String())
}

func TestDataTypeRegistryAPI_Get_InvalidDataType(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, nil)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)

	api := &dataTypeRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, testutils.TestErrBadRequestStr, w.Body.String())
}

func TestDataTypeRegistryAPI_Get_FailedToGet(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"dataType": testDataType}, nil)

	mockManager := adapterMock.NewMockIDataTypeRegistryManager(mockController)
	mockManager.EXPECT().
		Get(testDataType).
		Return(types.DataType{}, testutils.TestErr)

	api := &dataTypeRegistryAPI{mockManager}
	api.get(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}
