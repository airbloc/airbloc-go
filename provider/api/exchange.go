package api

import (
	"encoding/hex"
	"net/http"

	"github.com/airbloc/airbloc-go/provider/exchange"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ExchangeAPI is api wrapper of contract Exchange.sol
type ExchangeAPI struct {
	manager *exchange.Manager
}

// NewExchangeAPI makes new *ExchangeAPI struct
func NewExchangeAPI(backend service.Backend) (api.API, error) {
	ex := exchange.NewManager(backend.Client())
	return &ExchangeAPI{ex}, nil
}

// Prepare is a paid mutator transaction binding the contract method 0x77e61c33.
//
// Solidity: function prepare(string provider, address consumer, address escrow, bytes4 escrowSign, bytes escrowArgs, bytes20[] dataIds) returns(bytes8)
func (api *ExchangeAPI) Prepare(c *gin.Context) {
	var req struct {
		Provider   string   // string
		Consumer   string   // address
		Escrow     string   // address
		EscrowSign string   // bytes4
		EscrowArgs string   // bytes
		DataIds    []string // [][20]bytes
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	consumer := common.HexToAddress(req.Consumer)
	escrow := common.HexToAddress(req.Escrow)

	var escrowSign [4]byte
	if tmpBytes, err := hex.DecodeString(req.EscrowSign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	} else {
		copy(escrowSign[:], tmpBytes[:])
	}

	escrowArgs, err := hex.DecodeString(req.EscrowArgs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	dataIds := make([][20]byte, len(req.DataIds))
	for index, rawDataId := range req.DataIds {
		dataId, err := types.NewDataId(rawDataId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		dataIds[index] = dataId.Bytes()
	}

	offerId, err := api.manager.Prepare(
		c, req.Provider, consumer,
		escrow, escrowSign, escrowArgs,
		dataIds...,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"offerId": offerId})
}

// AddDataIds is a paid mutator transaction binding the contract method 0x367a9005.
//
// Solidity: function addDataIds(bytes8 offerId, bytes20[] dataIds) returns()
func (api *ExchangeAPI) AddDataIds(c *gin.Context) {
	var req struct{ DataIds []string }
	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	dataIds := make([][20]byte, len(req.DataIds))
	for index, rawDataId := range req.DataIds {
		dataId, err := types.NewDataId(rawDataId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		dataIds[index] = dataId.Bytes()
	}

	if err := api.manager.AddDataIds(c, offerId, dataIds); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Order is a paid mutator transaction binding the contract method 0x0cf833fb.
//
// Solidity: function order(bytes8 offerId) returns()
func (api *ExchangeAPI) Order(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := api.manager.Order(c, offerId); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Cancel is a paid mutator transaction binding the contract method 0xb2d9ba39.
//
// Solidity: function cancel(bytes8 offerId) returns()
func (api *ExchangeAPI) Cancel(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := api.manager.Cancel(c, offerId); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetOffer is a free data retrieval call binding the contract method 0x107f04b4.
//
// Solidity: function getOffer(bytes8 offerId) constant returns((string,address,bytes20[],uint256,uint256,(address,bytes4,bytes),uint8))
func (api *ExchangeAPI) GetOffer(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	offer, err := api.manager.GetOffer(offerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, offer)
}

// AttachToAPI is a registrant of an api.
func (api *ExchangeAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/exchange")
	apiMux.POST("/prepare", api.Prepare)
	apiMux.GET("/order/:offerId", api.GetOffer)
	apiMux.POST("/order/:offerId", api.Order)
	apiMux.PATCH("/order/:offerId", api.AddDataIds)
	apiMux.DELETE("/order/:offerId", api.Cancel)
}
