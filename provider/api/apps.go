package api

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/gin-gonic/gin"
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
	appName := c.Param("appName")
	if appName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := api.apps.Register(c, appName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (api *appRegistryAPI) unregister(c *gin.Context) {
	appName := c.Param("appName")
	if appName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := api.apps.Unregister(c, appName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (api *appRegistryAPI) get(c *gin.Context) {
	appName := c.Param("appName")
	if appName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	app, err := api.apps.Get(appName)
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
	appName := c.Param("appName")
	newOwner := c.Param("newOwner")
	if appName == "" || newOwner == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := api.apps.TransferAppOwner(c, appName, common.HexToAddress(newOwner)); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// AttachToAPI is a registrant of an api.
func (api *appRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/apps")
	apiMux.GET("/:appName", api.register)
	apiMux.POST("/:appName", api.unregister)
	apiMux.PATCH("/:appName", api.get)
	apiMux.DELETE("/:appName/:newOwner", api.transferAppOwner)
}
