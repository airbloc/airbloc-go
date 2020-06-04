package account

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/crypto"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFeePayer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	rootContext, cancelRootContext := context.WithCancel(context.Background())
	defer cancelRootContext()

	Convey("Testing fee payer client", t, func() {
		Convey("Test NewFeePayer", func() {
			Convey("Should return FeePayer struct", func() {
				testServer := httptest.NewServer(nil)
				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)
				So(feePayer, ShouldNotBeNil)
			})
			Convey("Should set client to http.DefaultClient when paramter is nil", func() {
				testServer := httptest.NewServer(nil)
				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)
				So(*feePayer.client, ShouldResemble, *http.DefaultClient)
			})
			Convey("Should return error when rawurl is invalid", func() {
				_, err := NewFeePayer(nil, string([]byte{0x7f}))
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "net/url: invalid")
			})
		})
		Convey("Test FeePayer.request", func() {
			ctx, cancel := context.WithCancel(rootContext)
			defer cancel()

			Convey("Should return response body", func() {
				mux := gin.New()
				mux.GET("/"+feePayerAPIVersion+"/", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{})
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				resp, err := feePayer.request(ctx, http.MethodGet, "/", nil, http.StatusOK)
				So(err, ShouldBeNil)
				So(string(resp), ShouldEqual, "{}")
			})
			Convey("Should return error if failed to make new request", func() {
				Convey("invalid request method", func() {
					testServer := httptest.NewServer(nil)

					feePayer, err := NewFeePayer(nil, testServer.URL)
					So(err, ShouldBeNil)

					_, err = feePayer.request(ctx, string([]byte{0x7f}), "/", nil)
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldContainSubstring, "net/http: invalid")
				})
				Convey("invalid request endpoint", func() {
					testServer := httptest.NewServer(nil)

					feePayer, err := NewFeePayer(nil, testServer.URL)
					So(err, ShouldBeNil)

					_, err = feePayer.request(ctx, http.MethodGet, string([]byte{0x7f}), nil)
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldContainSubstring, "net/url: invalid")
				})
			})
			Convey("Should return error if failed to request fee payer server", func() {
				mux := gin.New()
				mux.GET("/"+feePayerAPIVersion+"/", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{})
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				testServer.Close()

				_, err = feePayer.request(ctx, http.MethodGet, "/", nil, http.StatusOK)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "connect: connection refused")
			})
			Convey("Should return error if response code is not expected", func() {
				mux := gin.New()
				mux.GET("/"+feePayerAPIVersion+"/", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{})
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				_, err = feePayer.request(ctx, http.MethodGet, "/", nil, http.StatusCreated)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "request failed")
			})
		})
		Convey("Test FeePayer.Address", func() {
			ctx, cancel := context.WithCancel(rootContext)
			defer cancel()

			Convey("Should return fee payer address", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)

				mux := gin.New()
				mux.GET("/"+feePayerAPIVersion+"/address", func(c *gin.Context) {
					resp := gin.H{"address": crypto.PubkeyToAddress(testKey.PublicKey).Hex()}
					c.JSON(http.StatusOK, resp)
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				feePayerAddr, err := feePayer.Address(ctx)
				So(err, ShouldBeNil)
				So(feePayerAddr.Hex(), ShouldEqual, crypto.PubkeyToAddress(testKey.PublicKey).Hex())
			})
			Convey("Should return error if failed to marshal response to address", func() {
				mux := gin.New()
				mux.GET("/"+feePayerAPIVersion+"/address", func(c *gin.Context) {
					c.Data(http.StatusOK, "application/xml", []byte("address"))
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				_, err = feePayer.Address(ctx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "marshal response body")
			})
			Convey("Should return error if received address is empty", func() {
				mux := gin.New()
				mux.GET("/"+feePayerAPIVersion+"/address", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{})
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				_, err = feePayer.Address(ctx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "received fee payer address is empty")
			})
		})
		Convey("Test FeePayer.Transact", func() {
			ctx, cancel := context.WithCancel(rootContext)
			defer cancel()

			Convey("Should return transaction hash", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)

				mux := gin.New()
				mux.POST("/"+feePayerAPIVersion+"/transact", func(c *gin.Context) {
					resp := gin.H{"tx_hash": crypto.PubkeyToAddress(testKey.PublicKey).Hash().Hex()}
					c.JSON(http.StatusOK, resp)
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				tx, err := types.NewTransactionWithMap(types.TxTypeValueTransfer, map[types.TxValueKeyType]interface{}{
					types.TxValueKeyNonce:    uint64(0),
					types.TxValueKeyGasPrice: big.NewInt(0),
					types.TxValueKeyGasLimit: uint64(0),
					types.TxValueKeyTo:       crypto.PubkeyToAddress(testKey.PublicKey),
					types.TxValueKeyAmount:   big.NewInt(0),
					types.TxValueKeyFrom:     crypto.PubkeyToAddress(testKey.PublicKey),
				})
				So(err, ShouldBeNil)

				txHash, err := feePayer.Transact(ctx, tx)
				So(err, ShouldBeNil)
				So(txHash.Hex(), ShouldEqual, crypto.PubkeyToAddress(testKey.PublicKey).Hash().Hex())
			})
			Convey("Should return error if failed to marshal response to txhash", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)

				mux := gin.New()
				mux.POST("/"+feePayerAPIVersion+"/transact", func(c *gin.Context) {
					c.Data(http.StatusOK, "application/xml", []byte("txhash"))
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				tx, err := types.NewTransactionWithMap(types.TxTypeValueTransfer, map[types.TxValueKeyType]interface{}{
					types.TxValueKeyNonce:    uint64(0),
					types.TxValueKeyGasPrice: big.NewInt(0),
					types.TxValueKeyGasLimit: uint64(0),
					types.TxValueKeyTo:       crypto.PubkeyToAddress(testKey.PublicKey),
					types.TxValueKeyAmount:   big.NewInt(0),
					types.TxValueKeyFrom:     crypto.PubkeyToAddress(testKey.PublicKey),
				})
				So(err, ShouldBeNil)

				_, err = feePayer.Transact(ctx, tx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "marshal response body")
			})
			Convey("Should return error if received txhash is empty", func() {
				testKey, err := crypto.GenerateKey()
				So(err, ShouldBeNil)

				mux := gin.New()
				mux.POST("/"+feePayerAPIVersion+"/transact", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{})
				})
				testServer := httptest.NewServer(mux)

				feePayer, err := NewFeePayer(nil, testServer.URL)
				So(err, ShouldBeNil)

				tx, err := types.NewTransactionWithMap(types.TxTypeValueTransfer, map[types.TxValueKeyType]interface{}{
					types.TxValueKeyNonce:    uint64(0),
					types.TxValueKeyGasPrice: big.NewInt(0),
					types.TxValueKeyGasLimit: uint64(0),
					types.TxValueKeyTo:       crypto.PubkeyToAddress(testKey.PublicKey),
					types.TxValueKeyAmount:   big.NewInt(0),
					types.TxValueKeyFrom:     crypto.PubkeyToAddress(testKey.PublicKey),
				})
				So(err, ShouldBeNil)

				_, err = feePayer.Transact(ctx, tx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "received transaction hash is empty")
			})
		})
	})
}
