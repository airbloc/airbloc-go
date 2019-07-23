package api

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin/binding"

	"github.com/airbloc/airbloc-go/provider/apps"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/gin-gonic/gin"
)

type AppRegistryAPI struct {
	apps *apps.Manager
}

func NewAppRegistryAPI(backend service.Backend) (api.API, error) {
	ar := apps.NewManager(backend.Client())
	return &AppRegistryAPI{ar}, nil
}

func (api *AppRegistryAPI) Register(c *gin.Context) {
	appName := c.Param(":appName")
	if err := api.apps.Register(appName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *AppRegistryAPI) Unregister(c *gin.Context) {
	appName := c.Param(":appName")
	if err := api.apps.Unregister(appName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *AppRegistryAPI) Get(c *gin.Context) {
	appName := c.Param(":appName")

	app, err := api.apps.Get(appName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, app)
}

func (api *AppRegistryAPI) TransferAppOwner(c *gin.Context) {
	appName := c.Param(":appName")

	var req struct{ NewOwner string }
	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	newOwner := common.HexToAddress(req.NewOwner)
	if err := api.apps.TransferAppOwner(appName, newOwner); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *AppRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/apps")
	apiMux.GET("/:appName", api.Register)
	apiMux.POST("/:appName", api.Unregister)
	apiMux.PATCH("/:appName", api.Get)
	apiMux.DELETE("/:appName", api.TransferAppOwner)
}
