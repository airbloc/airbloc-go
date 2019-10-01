package api

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"testing"

	adapterMock "github.com/airbloc/airbloc-go/shared/adapter/mocks"
	"github.com/airbloc/airbloc-go/shared/types"
	testutils "github.com/airbloc/airbloc-go/test/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/klaytn/klaytn/common"
	"github.com/stretchr/testify/assert"
)

var (
	testProvider   = "deadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
	testConsumer   = "beefdeadbeefdeadbeefdeadbeefdeadbeefdead"
	testEscrow     = "deaddeaddeaddeaddeaddeaddeaddeaddeaddead"
	testEscrowSign = "deadbeef"
	testEscrowArgs = "deadbeefdeadbeef"
	testDataIds    = []string{
		"beefdeadbeefdeadbeefdeadbeefdeadbeefdead",
		"deadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
		"beefdeadbeefdeadbeefdeadbeefdeadbeefdead",
		"deadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
		"beefdeadbeefdeadbeefdeadbeefdeadbeefdead",
		"deadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
	}
	testOfferId = "deadbeefdeadbeef"
)

// happy path
func TestExchangeAPI_Prepare(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"provider":   testProvider,
		"consumer":   testConsumer,
		"escrow":     testEscrow,
		"escrowSign": testEscrowSign,
		"escrowArgs": testEscrowArgs,
		"dataIds":    testDataIds,
	}, binding.JSON)

	var escrowSign [4]byte
	b, _ := hex.DecodeString(testEscrowSign)
	copy(escrowSign[:], b[:])
	escrowArgs, _ := hex.DecodeString(testEscrowArgs)

	dataIds := make([]types.DataId, len(testDataIds))
	for i, raw := range testDataIds {
		dataIds[i], _ = types.NewDataIdFromStr(raw)
	}

	mockManager := adapterMock.NewMockIExchangeManager(mockController)
	mockManager.EXPECT().
		Prepare(c, nil,
			testProvider,
			common.HexToAddress(testConsumer),
			common.HexToAddress(testEscrow),
			escrowSign,
			escrowArgs,
			dataIds,
		).Return(types.HexToID(testOfferId))

	api := &exchangeAPI{mockManager}
	api.prepare(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"offer_id":"%s"}`, testOfferId), w.Body.String())
}

func TestExchangeAPI_Prepare_InvalidJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{}, binding.JSON)

	mockManager := adapterMock.NewMockIExchangeManager(mockController)

	api := &exchangeAPI{mockManager}
	api.prepare(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestExchangeAPI_Prepare_InvalidEscrowSign(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"provider":   testProvider,
		"consumer":   testConsumer,
		"escrow":     testEscrow,
		"escrowSign": testEscrowSign + "z",
		"escrowArgs": testEscrowArgs,
		"dataIds":    testDataIds,
	}, binding.JSON)

	mockManager := adapterMock.NewMockIExchangeManager(mockController)

	api := &exchangeAPI{mockManager}
	api.prepare(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestExchangeAPI_Prepare_InvalidEscrowArgs(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"provider":   testProvider,
		"consumer":   testConsumer,
		"escrow":     testEscrow,
		"escrowSign": testEscrowSign,
		"escrowArgs": testEscrowArgs + "z",
		"dataIds":    testDataIds,
	}, binding.JSON)

	mockManager := adapterMock.NewMockIExchangeManager(mockController)

	api := &exchangeAPI{mockManager}
	api.prepare(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestExchangeAPI_Prepare_InvalidDataIds(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"provider":   testProvider,
		"consumer":   testConsumer,
		"escrow":     testEscrow,
		"escrowSign": testEscrowSign + "z",
		"escrowArgs": testEscrowArgs,
		"dataIds":    append([]string{testDataIds[0] + "z"}, testDataIds[1:]...),
	}, binding.JSON)

	mockManager := adapterMock.NewMockIExchangeManager(mockController)

	api := &exchangeAPI{mockManager}
	api.prepare(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), `{"error":`))
}

func TestExchangeAPI_Prepare_FailedToPrepare(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{
		"provider":   testProvider,
		"consumer":   testConsumer,
		"escrow":     testEscrow,
		"escrowSign": testEscrowSign,
		"escrowArgs": testEscrowArgs,
		"dataIds":    testDataIds,
	}, binding.JSON)

	var escrowSign [4]byte
	b, _ := hex.DecodeString(testEscrowSign)
	copy(escrowSign[:], b[:])
	escrowArgs, _ := hex.DecodeString(testEscrowArgs)

	dataIds := make([]types.DataId, len(testDataIds))
	for i, raw := range testDataIds {
		dataIds[i], _ = types.NewDataIdFromStr(raw)
	}

	mockManager := adapterMock.NewMockIExchangeManager(mockController)
	mockManager.EXPECT().
		Prepare(c, nil,
			testProvider,
			common.HexToAddress(testConsumer),
			common.HexToAddress(testEscrow),
			escrowSign,
			escrowArgs,
			dataIds,
		).Return(types.ID{}, testutils.TestErr)

	api := &exchangeAPI{mockManager}
	api.prepare(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, testutils.TestErrStr, w.Body.String())
}

// happy path
func TestExchangeAPI_AddDataIds(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"dataIds": testDataIds}, binding.JSON)
	c.Params = append(c.Params, gin.Param{Key: "offer_id", Value: testOfferId})

	offerId, _ := types.HexToID(testOfferId)
	dataIds := make([]types.DataId, len(testDataIds))
	for i, raw := range testDataIds {
		dataIds[i], _ = types.NewDataIdFromStr(raw)
	}

	mockManager := adapterMock.NewMockIExchangeManager(mockController)
	mockManager.EXPECT().
		AddDataIds(c, nil, offerId, dataIds).
		Return(nil)

	api := &exchangeAPI{mockManager}
	api.addDataIds(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

// happy path
func TestExchangeAPI_Order(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"offer_id": testOfferId}, nil)

	offerId, _ := types.HexToID(testOfferId)

	mockManager := adapterMock.NewMockIExchangeManager(mockController)
	mockManager.EXPECT().
		Order(c, nil, offerId).
		Return(nil)

	api := &exchangeAPI{mockManager}
	api.order(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}

// happy path
func TestExchangeAPI_Cancel(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	w, c := testutils.CreateTestRequest(t, gin.H{"offer_id": testOfferId}, nil)

	offerId, _ := types.HexToID(testOfferId)

	mockManager := adapterMock.NewMockIExchangeManager(mockController)
	mockManager.EXPECT().
		Cancel(c, nil, offerId).
		Return(nil)

	api := &exchangeAPI{mockManager}
	api.cancel(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testutils.TestSuccessStr, w.Body.String())
}
