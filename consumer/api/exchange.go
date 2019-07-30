package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/exchange"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
)

type ExchangeAPI struct {
	manager adapter.IExchangeManager
}

func NewExchangeAPI(backend service.Backend) (api.API, error) {
	manager := exchange.NewManager(backend.Client())
	return &ExchangeAPI{manager}, nil
}

func (api *ExchangeAPI) Settle(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := api.manager.Settle(c, offerId); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *ExchangeAPI) Reject(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := api.manager.Reject(c, offerId); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, offer)
}

func (api *ExchangeAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/exchange")
	apiMux.GET("/settle/:offerId", api.Settle)
	apiMux.GET("/reject/:offerId", api.Reject)
	apiMux.GET("/order/:offerId", api.GetOffer)
}
