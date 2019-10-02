package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/gin-gonic/gin"
	"github.com/klaytn/klaytn/common"
)

// controllerRegistryAPI is api wrapper of contract ControllerRegistry.sol
type controllerRegistryAPI struct {
	controllers adapter.IControllerRegistryManager
}

// NewControllerRegistryAPI makes new *controllerRegistryAPI struct
func NewControllerRegistryAPI(backend service.Backend) (api.API, error) {
	cr := adapter.NewControllerRegistryManager(backend.Client())
	return &controllerRegistryAPI{cr}, nil
}

func (api *controllerRegistryAPI) register(c *gin.Context) {
	controllerAddrHex := c.Param("controller_addr")
	if controllerAddrHex == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	controllerAddr := common.HexToAddress(controllerAddrHex)

	if exists, err := api.controllers.Exists(controllerAddr); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if exists {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{})
		return
	}

	if err := api.controllers.Register(c, nil, controllerAddr); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (api *controllerRegistryAPI) get(c *gin.Context) {
	controllerAddr := c.Param("controller_addr")
	if controllerAddr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	controller, err := api.controllers.Get(common.HexToAddress(controllerAddr))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, controller)
}

// AttachToAPI is a registrant of an api.
func (api *controllerRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/controllers")
	apiMux.GET("/:controller_addr", api.get)
	apiMux.POST("/:controller_addr", api.register)
}
