package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/klaytn/klaytn/common"
)

// appRegistryAPI is api wrapper of contract AppRegistry.sol
type appRegistryAPI struct {
	apps adapter.IAppRegistryManager
}

// NewAppRegistryAPI makes new *appRegistryAPI struct
func NewAppRegistryAPI(backend service.Backend) (api.API, error) {
	ar := adapter.NewAppRegistryManager(backend.Client())
	return &appRegistryAPI{ar}, nil
}

func (api *appRegistryAPI) register(c *gin.Context) {
	var req struct {
		AppName string `json:"app_name" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if exists, err := api.apps.Exists(req.AppName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if exists {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{})
		return
	}

	if err := api.apps.Register(c, req.AppName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (api *appRegistryAPI) unregister(c *gin.Context) {
	var req struct {
		AppName string `json:"app_name" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := api.apps.Unregister(c, req.AppName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *appRegistryAPI) get(c *gin.Context) {
	var req struct {
		AppName string `form:"app_name" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if exists, err := api.apps.Exists(req.AppName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "cannot find app information"})
		return
	}

	app, err := api.apps.Get(req.AppName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (api *appRegistryAPI) transferAppOwner(c *gin.Context) {
	var req struct {
		AppName  string `json:"app_name" binding:"required"`
		NewOwner string `json:"new_owner" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := api.apps.TransferAppOwner(c, req.AppName, common.HexToAddress(req.NewOwner)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// AttachToAPI is a registrant of an api.
func (api *appRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/apps")
	apiMux.GET("/", api.get)
	apiMux.POST("/", api.register)
	apiMux.PATCH("/", api.transferAppOwner)
	apiMux.DELETE("/", api.unregister)
}
