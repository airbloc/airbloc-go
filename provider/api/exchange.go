package api

import (
	"encoding/hex"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/klaytn/klaytn/common"
)

// exchangeAPI is api wrapper of contract Exchange.sol
type exchangeAPI struct {
	manager adapter.IExchangeManager
}

// NewExchangeAPI makes new *exchangeAPI struct
func NewExchangeAPI(backend service.Backend) (api.API, error) {
	ex := adapter.NewExchangeManager(backend.Client())
	return &exchangeAPI{ex}, nil
}

func (api *exchangeAPI) prepare(c *gin.Context) {
	var req struct {
		Provider   string   `binding:"required"` // string
		Consumer   string   `binding:"required"` // address
		Escrow     string   `binding:"required"` // address
		EscrowSign string   `binding:"required"` // bytes4
		EscrowArgs string   `binding:"required"` // bytes
		DataIds    []string `binding:"required"` // [][20]bytes
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consumer := common.HexToAddress(req.Consumer)
	escrow := common.HexToAddress(req.Escrow)

	var escrowSign [4]byte
	if tmpBytes, err := hex.DecodeString(req.EscrowSign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		copy(escrowSign[:], tmpBytes[:])
	}

	escrowArgs, err := hex.DecodeString(req.EscrowArgs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataIds := make([]types.DataId, len(req.DataIds))
	for index, rawDataId := range req.DataIds {
		dataIds[index], err = types.NewDataIdFromStr(rawDataId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	offerId, err := api.manager.Prepare(
		c, req.Provider, consumer,
		escrow, escrowSign, escrowArgs, dataIds,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"offer_id": offerId})
}

func (api *exchangeAPI) addDataIds(c *gin.Context) {
	var req struct {
		DataIds []string `binding:"required"`
	}
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rawOfferId := c.Param("offer_id")
	if rawOfferId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataIds := make([]types.DataId, len(req.DataIds))
	for index, rawDataId := range req.DataIds {
		dataIds[index], err = types.NewDataIdFromStr(rawDataId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	err = api.manager.AddDataIds(c, offerId, dataIds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *exchangeAPI) order(c *gin.Context) {
	rawOfferId := c.Param("offer_id")
	if rawOfferId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.manager.Order(c, offerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *exchangeAPI) cancel(c *gin.Context) {
	rawOfferId := c.Param("offer_id")
	if rawOfferId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.manager.Cancel(c, offerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *exchangeAPI) getOffer(c *gin.Context) {
	rawOfferId := c.Param("offer_id")
	if rawOfferId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := api.manager.GetOffer(offerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, offer)
}

// AttachToAPI is a registrant of an api.
func (api *exchangeAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/exchange")
	apiMux.POST("/prepare", api.prepare)
	apiMux.GET("/order/:offer_id", api.getOffer)
	apiMux.POST("/order/:offer_id", api.order)
	apiMux.PATCH("/order/:offer_id", api.addDataIds)
	apiMux.DELETE("/order/:offer_id", api.cancel)
}
