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
	apps        adapter.IAppRegistryManager
	consents    adapter.IConsentsManager
	dataTypes   adapter.IDataTypeRegistryManager
}

// NewDAuthAPI makes new *dAuthAPI struct
func NewDAuthAPI(backend service.Backend) (api.API, error) {
	dauthClient := dauth.NewProviderClient(backend.Kms(), backend.Client(), backend.P2P())
	return &dAuthAPI{
		dauthClient: dauthClient,
		apps:        adapter.NewAppRegistryManager(backend.Client()),
		consents:    adapter.NewConsentsManager(backend.Client()),
		dataTypes:   adapter.NewDataTypeRegistryManager(backend.Client()),
	}, nil
}

func (api *dAuthAPI) signUp(c *gin.Context) {
	var req struct {
		IdentityHash string `json:"identity_hash" binding:"required"`
		Controller   string `json:"controller" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller := common.HexToAddress(req.Controller)
	accountId, err := api.dauthClient.SignIn(c, req.IdentityHash, controller)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
}

func (api *dAuthAPI) getAuthorizations(c *gin.Context) {
	var req struct {
		AccountId string `form:"account_id" binding:"required"`
		AppName   string `form:"app_name" binding:"required"`
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

	if exists, err := api.apps.Exists(req.AppName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Cannot find app information."})
		return
	}

	app, err := api.apps.Get(req.AppName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp struct {
		HasAuthorizedBefore bool                `json:"has_authorized_before"`
		Authorizations      []types.ConsentData `json:"authorizations"`
	}

	consentEventIter, err := api.consents.FilterConsented(&bind.FilterOpts{
		Start:   api.consents.CreatedAt().Uint64(),
		End:     nil,
		Context: c,
	}, []uint8{0, 1}, []types.ID{accountId}, []common.Address{app.Addr})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp.Authorizations = []types.ConsentData{}

	for consentEventIter.Next() {
		event := consentEventIter.Event
		if event == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse event from data."})
			return
		}

		resp.Authorizations = append(resp.Authorizations, types.ConsentData{
			Action:   event.Action,
			DataType: event.DataType,
			Allow:    event.Allowed,
		})
	}

	resp.HasAuthorizedBefore = len(resp.Authorizations) == 0
	c.JSON(http.StatusOK, resp)
}

func (api *dAuthAPI) allow(c *gin.Context) {
	var req struct {
		AccountId string `json:"account_id" binding:"required"`
		DataType  string `json:"data_type" binding:"required"`
		Action    uint8  `json:"action" binding:"required"`
		AppName   string `json:"app_name" binding:"required"`
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

	c.JSON(http.StatusOK, gin.H{})
}

func (api *dAuthAPI) deny(c *gin.Context) {
	var req struct {
		AccountId string `json:"account_id" binding:"required"`
		DataType  string `json:"data_type" binding:"required"`
		Action    uint8  `json:"action" binding:"required"`
		AppName   string `json:"app_name" binding:"required"`
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

	c.JSON(http.StatusOK, gin.H{})
}

func (api *dAuthAPI) many(c *gin.Context) {
	var req struct {
		AccountId   string              `json:"account_id" binding:"required"`
		AppName     string              `json:"app_name" binding:"required"`
		ConsentData []types.ConsentData `json:"consent_data" binding:"required"`
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

	err = api.dauthClient.Many(
		c, accountId,
		req.AppName,
		req.ConsentData,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// AttachToAPI is a registrant of an api.
func (api *dAuthAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/dauth")
	apiMux.GET("/auth", api.getAuthorizations)
	apiMux.POST("/signin", api.signUp)
	apiMux.PUT("/allow", api.allow)
	apiMux.PUT("/deny", api.deny)
	apiMux.PUT("/many", api.many)
}
