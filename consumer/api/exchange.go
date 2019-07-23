package api

import (
	"context"
	"github.com/airbloc/airbloc-go/consumer/exchange"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExchangeAPI struct {
	manager *exchange.Manager
}

func NewExchangeAPI(backend service.Backend) (api.API, error) {
	manager := exchange.NewManager(backend.Client())
	return &ExchangeAPI{manager}, nil
}

func (api *ExchangeAPI) Settle(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := api.manager.Settle(c, offerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (api *ExchangeAPI) Reject(c *gin.Context) {
	rawOfferId := c.Param("offerId")
	offerId, err := types.HexToID(rawOfferId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := api.manager.Reject(c, offerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (api *ExchangeAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/exchange")
	apiMux.GET("/settle/:offerId", api.Settle)
	apiMux.GET("/reject/:offerId", api.Reject)
}
