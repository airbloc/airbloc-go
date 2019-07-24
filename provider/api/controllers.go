package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/provider/controllers"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ControllerRegistryAPI struct {
	controllers *controllers.Manager
}

func NewControllerRegistryAPI(backend service.Backend) (api.API, error) {
	cr := controllers.NewManager(backend.Client())
	return &ControllerRegistryAPI{cr}, nil
}

func (api *ControllerRegistryAPI) Register(c *gin.Context) {
	rawControllerAddr := c.Param("controllerAddr")
	controllerAddr := common.HexToAddress(rawControllerAddr)

	if err := api.controllers.Register(c, controllerAddr); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *ControllerRegistryAPI) Get(c *gin.Context) {
	rawControllerAddr := c.Param("controllerAddr")
	controllerAddr := common.HexToAddress(rawControllerAddr)

	controller, err := api.controllers.Get(c, controllerAddr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, controller)
}

func (api *ControllerRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/controllers")
	apiMux.GET("/:controllerAddr", api.Get)
	apiMux.POST("/:controllerAddr", api.Register)
}
