package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (api *appRegistryAPI) register(c *gin.Context) {
	var req struct {
		AppName string `json:"app_name" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := api.apps.Register(c, req.AppName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (api *appRegistryAPI) unregister(c *gin.Context) {
	var req struct {
		AppName string `json:"app_name" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := api.apps.Unregister(c, req.AppName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (api *appRegistryAPI) get(c *gin.Context) {
	var req struct {
		AppName string `form:"app_name" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app, err := api.apps.Get(req.AppName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
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
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// AttachToAPI is a registrant of an api.
func (api *appRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/apps")
	apiMux.GET("/", api.get)
	apiMux.POST("/", api.register)
	apiMux.PATCH("/", api.transferAppOwner)
	apiMux.DELETE("/", api.unregister)
}
