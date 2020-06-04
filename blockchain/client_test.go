package blockchain

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/airbloc/logger"

	"github.com/gin-gonic/gin"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/klaytn/klaytn/params"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

type rpcRequest struct {
	RequestID     int           `json:"id"`
	RpcVersion    string        `json:"jsonrpc"`
	RequestMethod string        `json:"method"`
	RequestParams []interface{} `json:"params"`
}

type rpcHandlerFunc func(rpcRequest) (int, gin.H)
type rpcHandler struct {
	m string
	f rpcHandlerFunc
}

var (
	testNetworkID = "2470"
	testTxHash    = common.HexToHash("0xdeadbeef")
)

func newNetworkVersionHandler() rpcHandler {
	return rpcHandler{
		m: "net_version",
		f: func(req rpcRequest) (int, gin.H) {
			return http.StatusOK, gin.H{
				"id":      req.RequestID,
				"jsonrpc": req.RpcVersion,
				"result":  testNetworkID,
			}
		},
	}
}

func newGetTransactionReceiptHandler(status uint, contractAddress common.Address) rpcHandler {
	return rpcHandler{
		m: "klay_getTransactionReceipt",
		f: func(req rpcRequest) (int, gin.H) {
			receipt := types.NewReceipt(
				status,
				common.HexToHash(req.RequestParams[0].(string)),
				uint64(0),
			)
			receipt.ContractAddress = contractAddress
			receipt.Bloom = types.BytesToBloom(testTxHash.Bytes())
			receipt.Logs = []*types.Log{{Topics: []common.Hash{testTxHash}}}

			return http.StatusOK, gin.H{
				"id":      req.RequestID,
				"jsonrpc": req.RpcVersion,
				"result":  receipt,
			}
		},
	}
}

func newGetCodeHandler(hexcode string) rpcHandler {
	return rpcHandler{
		m: "klay_getCode",
		f: func(req rpcRequest) (int, gin.H) {
			return http.StatusOK, gin.H{
				"id":      req.RequestID,
				"jsonrpc": req.RpcVersion,
				"result":  hexcode,
			}
		},
	}
}

func NewTestRPCServer(handlers ...rpcHandler) *httptest.Server {
	handlerMap := make(map[string]rpcHandlerFunc, len(handlers))
	for _, handler := range handlers {
		handlerMap[handler.m] = handler.f
	}

	mux := gin.New()
	mux.Use(gin.ErrorLogger())
	mux.POST("/", func(c *gin.Context) {
		req := rpcRequest{}
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, errors.Wrap(err, "decode request to rpcRequest"))
			return
		}

		handler, exist := handlerMap[req.RequestMethod]
		if !exist {
			_ = c.AbortWithError(http.StatusNotFound, errors.New("unregistered method call"))
			return
		}
		c.JSON(handler(req))
	})
	return httptest.NewServer(mux)
}

type l struct{}

func (l) Init()             {}
func (l) Write(*logger.Log) {}

func TestClient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger.SetLogger(l{})

	rootContext, cancelRootContext := context.WithCancel(context.Background())
	defer cancelRootContext()

	Convey("Test Blockchain Client", t, func() {
		Convey("Test NewClient", func() {
			ctx, cancel := context.WithCancel(rootContext)
			defer cancel()

			Convey("Should return Client struct", func() {
				testServer := NewTestRPCServer(newNetworkVersionHandler())
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				So(client, ShouldNotBeNil)
				client.Close()
			})
			Convey("Should return error if endpoint is invalid", func() {
				_, err := NewClient(ctx, string([]byte{0x7f}))
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "net/url: invalid")
			})
			Convey("Should return error if failed to fetch network id", func() {
				testServer := httptest.NewServer(nil)
				_, err := NewClient(ctx, testServer.URL)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "fetch network id")
			})
		})
		Convey("Test Client.waitMined", func() {
			ctx, cancel := context.WithTimeout(rootContext, 2*time.Second)
			defer cancel()

			Convey("Should return transaction receipt", func() {
				testServer := NewTestRPCServer(
					newNetworkVersionHandler(),
					newGetTransactionReceiptHandler(types.ReceiptStatusSuccessful, common.Address{}),
				)
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				receipt, err := client.waitMined(ctx, testTxHash)
				So(err, ShouldBeNil)
				So(receipt.Status, ShouldEqual, types.ReceiptStatusSuccessful)
			})
			Convey("Should return transaction receipt even status is failed", func() {
				testServer := NewTestRPCServer(
					newNetworkVersionHandler(),
					newGetTransactionReceiptHandler(types.ReceiptStatusFailed, common.Address{}),
				)
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				receipt, err := client.waitMined(ctx, testTxHash)
				So(err, ShouldBeNil)
				So(receipt.Status, ShouldEqual, types.ReceiptStatusFailed)
			})
			Convey("Should return error if context has canceled", func() {
				testServer := NewTestRPCServer(newNetworkVersionHandler())
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				_, err = client.waitMined(ctx, testTxHash)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, context.DeadlineExceeded.Error())
			})
		})
		Convey("Test Client.WaitMinedWithHash", func() {
			ctx, cancel := context.WithTimeout(rootContext, 2*time.Second)
			defer cancel()

			Convey("Should return transaction receipt", func() {
				testServer := NewTestRPCServer(
					newNetworkVersionHandler(),
					newGetTransactionReceiptHandler(types.ReceiptStatusSuccessful, common.Address{}),
				)
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				receipt, err := client.WaitMinedWithHash(ctx, testTxHash)
				So(err, ShouldBeNil)
				So(receipt.Status, ShouldEqual, types.ReceiptStatusSuccessful)
			})
			Convey("Should return error if receive failed transaction", func() {
				testServer := NewTestRPCServer(
					newNetworkVersionHandler(),
					newGetTransactionReceiptHandler(types.ReceiptStatusFailed, common.Address{}),
				)
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				_, err = client.WaitMinedWithHash(ctx, testTxHash)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "tx failed")
			})
		})
		Convey("Test Client.WaitDeployed", func() {
			ctx, cancel := context.WithTimeout(rootContext, 2*time.Hour)
			defer cancel()

			Convey("Should return transaction receipt", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)
				testKeyAddr := crypto.PubkeyToAddress(testKey.PublicKey)

				testServer := NewTestRPCServer(
					newNetworkVersionHandler(),
					newGetTransactionReceiptHandler(types.ReceiptStatusSuccessful, testKeyAddr),
					newGetCodeHandler("0xdeadbeefdeadbeefbeefdeadbeefdead"),
				)
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				tx, err := types.NewTransactionWithMap(types.TxTypeSmartContractDeploy, map[types.TxValueKeyType]interface{}{
					types.TxValueKeyNonce:         uint64(0),
					types.TxValueKeyGasPrice:      big.NewInt(0),
					types.TxValueKeyGasLimit:      uint64(0),
					types.TxValueKeyTo:            &testKeyAddr,
					types.TxValueKeyAmount:        big.NewInt(0),
					types.TxValueKeyFrom:          testKeyAddr,
					types.TxValueKeyData:          []byte{0x0},
					types.TxValueKeyHumanReadable: false,
					types.TxValueKeyCodeFormat:    params.CodeFormatEVM,
				})
				So(err, ShouldBeNil)

				receipt, err := client.WaitDeployed(ctx, tx)
				So(err, ShouldBeNil)
				So(receipt.ContractAddress, ShouldEqual, testKeyAddr)
			})
			Convey("Should return error if transaction type is invalid", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)
				testKeyAddr := crypto.PubkeyToAddress(testKey.PublicKey)

				testServer := NewTestRPCServer(newNetworkVersionHandler())
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				tx, err := types.NewTransactionWithMap(types.TxTypeValueTransfer, map[types.TxValueKeyType]interface{}{
					types.TxValueKeyNonce:    uint64(0),
					types.TxValueKeyGasPrice: big.NewInt(0),
					types.TxValueKeyGasLimit: uint64(0),
					types.TxValueKeyTo:       testKeyAddr,
					types.TxValueKeyAmount:   big.NewInt(0),
					types.TxValueKeyFrom:     testKeyAddr,
				})
				So(err, ShouldBeNil)

				_, err = client.WaitDeployed(ctx, tx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "tx is not contract creation")
			})
			Convey("Should return error if contract address is empty", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)
				testKeyAddr := crypto.PubkeyToAddress(testKey.PublicKey)

				testServer := NewTestRPCServer(
					newNetworkVersionHandler(),
					newGetTransactionReceiptHandler(types.ReceiptStatusSuccessful, common.Address{}),
				)
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				tx, err := types.NewTransactionWithMap(types.TxTypeSmartContractDeploy, map[types.TxValueKeyType]interface{}{
					types.TxValueKeyNonce:         uint64(0),
					types.TxValueKeyGasPrice:      big.NewInt(0),
					types.TxValueKeyGasLimit:      uint64(0),
					types.TxValueKeyTo:            &testKeyAddr,
					types.TxValueKeyAmount:        big.NewInt(0),
					types.TxValueKeyFrom:          testKeyAddr,
					types.TxValueKeyData:          []byte{0x0},
					types.TxValueKeyHumanReadable: false,
					types.TxValueKeyCodeFormat:    params.CodeFormatEVM,
				})
				So(err, ShouldBeNil)

				_, err = client.WaitDeployed(ctx, tx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "zero address")
			})
			Convey("Should return error if code at contract address is empty", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)
				testKeyAddr := crypto.PubkeyToAddress(testKey.PublicKey)

				testServer := NewTestRPCServer(
					newNetworkVersionHandler(),
					newGetTransactionReceiptHandler(types.ReceiptStatusSuccessful, testKeyAddr),
					newGetCodeHandler(""),
				)
				client, err := NewClient(ctx, testServer.URL)
				So(err, ShouldBeNil)
				defer client.Close()

				tx, err := types.NewTransactionWithMap(types.TxTypeSmartContractDeploy, map[types.TxValueKeyType]interface{}{
					types.TxValueKeyNonce:         uint64(0),
					types.TxValueKeyGasPrice:      big.NewInt(0),
					types.TxValueKeyGasLimit:      uint64(0),
					types.TxValueKeyTo:            &testKeyAddr,
					types.TxValueKeyAmount:        big.NewInt(0),
					types.TxValueKeyFrom:          testKeyAddr,
					types.TxValueKeyData:          []byte{0x0},
					types.TxValueKeyHumanReadable: false,
					types.TxValueKeyCodeFormat:    params.CodeFormatEVM,
				})
				So(err, ShouldBeNil)

				receipt, err := client.WaitDeployed(ctx, tx)
				So(err, ShouldNotBeNil)
				So(err, ShouldEqual, bind.ErrNoCodeAfterDeploy)
				So(receipt.ContractAddress, ShouldEqual, testKeyAddr)
			})
		})
	})
}
