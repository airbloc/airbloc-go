package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/provider/controllers"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ControllerRegistryAPI is api wrapper of contract ControllerRegistry.sol
type ControllerRegistryAPI struct {
	controllers *controllers.Manager
}

// NewControllerRegistryAPI makes new *ControllerRegistryAPI struct
func NewControllerRegistryAPI(backend service.Backend) (api.API, error) {
	cr := controllers.NewManager(backend.Client())
	return &ControllerRegistryAPI{cr}, nil
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address controllerAddr) returns()
func (api *ControllerRegistryAPI) Register(c *gin.Context) {
	rawControllerAddr := c.Param("controllerAddr")
	controllerAddr := common.HexToAddress(rawControllerAddr)

	if err := api.controllers.Register(c, controllerAddr); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address controller) constant returns((address,uint256))
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

// AttachToAPI is a registrant of an api.
func (api *ControllerRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/controllers")
	apiMux.GET("/:controllerAddr", api.Get)
	apiMux.POST("/:controllerAddr", api.Register)
}
