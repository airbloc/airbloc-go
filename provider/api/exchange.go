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

type ExchangeAPI struct {
	manager *exchange.Manager
}

func NewExchangeAPI(backend service.Backend) (api.API, error) {
	ex := exchange.NewManager(backend.Client())
	return &ExchangeAPI{ex}, nil
}

func (api *ExchangeAPI) Prepare(c *gin.Context) {
	var req struct {
		provider   string   // string
		consumer   string   // address
		escrow     string   // address
		escrowSign string   // bytes4
		escrowArgs string   // bytes
		dataIds    []string // [][20]bytes
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	consumer := common.HexToAddress(req.consumer)
	escrow := common.HexToAddress(req.escrow)

	var escrowSign [4]byte
	if tmpBytes, err := hex.DecodeString(req.escrowSign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	} else {
		copy(escrowSign[:], tmpBytes[:])
	}

	escrowArgs, err := hex.DecodeString(req.escrowArgs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	dataIds := make([][20]byte, len(req.dataIds))
	for index, rawDataId := range req.dataIds {
		dataId, err := types.NewDataId(rawDataId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		dataIds[index] = dataId.Bytes()
	}

	offerId, err := api.manager.Prepare(
		c, provider, consumer,
		escrow, escrowSign, escrowArgs,
		dataIds...,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"offerId": offerId})
}

func (api *ExchangeAPI) AddDataIds(c *gin.Context) {
	var req struct{ dataIds []string }
	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	dataIds := make([][20]byte, len(req.dataIds))
	for index, rawDataId := range req.dataIds {
		dataId, err := types.NewDataId(rawDataId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		dataIds[index] = dataId.Bytes()
	}

	if err := api.manager.AddDataIds(c, offerId, dataIds); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"offerId": offerId})
}

func (api *ExchangeAPI) Order(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := api.manager.Order(c, offerId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"offerId": offerId})
}

func (api *ExchangeAPI) Cancel(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := api.manager.Cancel(c, offerId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"offerId": offerId})
}

func (api *ExchangeAPI) GetOffer(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	offer, err := api.manager.GetOffer(offerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, offer)
}

func (api *ExchangeAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/exchange")
	apiMux.POST("/prepare", api.Prepare)
	apiMux.GET("/order/:offerId", api.GetOffer)
	apiMux.POST("/order/:offerId", api.Order)
	apiMux.PATCH("/order/:offerId", api.AddDataIds)
	apiMux.DELETE("/order/:offerId", api.Cancel)
}
