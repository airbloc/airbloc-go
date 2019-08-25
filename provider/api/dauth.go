package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/provider/dauth"
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/common"
)

type dAuthAPI struct {
	dauthClient *dauth.Client
	consents    adapter.IConsentsManager
	dataTypes   adapter.IDataTypeRegistryManager
}

// NewDAuthAPI makes new *dAuthAPI struct
func NewDAuthAPI(backend service.Backend) (api.API, error) {
	dauthClient := dauth.NewProviderClient(backend.Kms(), backend.Client(), backend.P2P())
	return &dAuthAPI{
		dauthClient: dauthClient,
		consents:    adapter.NewConsentsManager(backend.Client()),
		dataTypes:   adapter.NewDataTypeRegistryManager(backend.Client()),
	}, nil
}

func (api *dAuthAPI) signIn(c *gin.Context) {
	var req struct {
		Identity   string `binding:"required"`
		Controller string `binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller := common.HexToAddress(req.Controller)
	accountId, err := api.dauthClient.SignIn(c, req.Identity, controller)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
}

func (api *dAuthAPI) getAuthorizations(c *gin.Context) {
	var req struct {
		AccountId string `binding:"required"`
		AppName   string `binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountId, err := types.HexToID(req.AccountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp struct {
		HasAuthorizedBefore bool
		Authorizations      []struct {
			Action     types.ConsentActionTypes
			DataType   string
			Authorized bool
		}
	}

	consentEventIter, err := api.consents.FilterConsented(&bind.FilterOpts{
		Context: c,
		Start:   api.consents.CreatedAt().Uint64(),
	}, nil, []types.ID{accountId}, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if !consentEventIter.Next() {
		resp.HasAuthorizedBefore = false
		c.JSON(http.StatusOK, resp)
		return
	}

	dataTypeEventIter, err := api.dataTypes.FilterRegistration(&bind.FilterOpts{
		Context: c,
		Start:   api.consents.CreatedAt().Uint64(),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var (
		actions = []types.ConsentActionTypes{
			types.ConsentActionCollection,
			types.ConsentActionExchange,
		}
		dataTypeRegisterEvent *adapter.DataTypeRegistryRegistration
	)

	// data type
	for ; dataTypeEventIter.Next(); dataTypeRegisterEvent = dataTypeEventIter.Event {
		if dataTypeRegisterEvent == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// action
		for _, action := range actions {
			dataType := dataTypeRegisterEvent.Name
			allowed, err := api.consents.IsAllowed(accountId, dataType, uint8(action), req.AppName)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			resp.Authorizations = append(resp.Authorizations, struct {
				Action     types.ConsentActionTypes
				DataType   string
				Authorized bool
			}{
				Action:     action,
				DataType:   dataType,
				Authorized: allowed,
			})
		}
	}

	c.JSON(http.StatusOK, resp)
}

func (api *dAuthAPI) allow(c *gin.Context) {
	var req struct {
		AccountId string                   `binding:"required"`
		DataType  string                   `binding:"required"`
		Action    types.ConsentActionTypes `binding:"required"`
		AppName   string                   `binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountId, err := types.HexToID(req.AccountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.dauthClient.Allow(
		c, accountId,
		req.DataType,
		req.Action,
		req.AppName,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *dAuthAPI) deny(c *gin.Context) {
	var req struct {
		AccountId string                   `binding:"required"`
		DataType  string                   `binding:"required"`
		Action    types.ConsentActionTypes `binding:"required"`
		AppName   string                   `binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountId, err := types.HexToID(req.AccountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.dauthClient.Deny(
		c, accountId,
		req.DataType,
		req.Action,
		req.AppName,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// AttachToAPI is a registrant of an api.
func (api *dAuthAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/dauth")
	apiMux.GET("/auth", api.getAuthorizations)
	apiMux.POST("/signin", api.signIn)
	apiMux.PUT("/allow", api.allow)
	apiMux.PUT("/deny", api.deny)
}
