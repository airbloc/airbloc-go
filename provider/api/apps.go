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

// AppRegistryAPI is api wrapper of contract AppRegistry.sol
type AppRegistryAPI struct {
	apps *apps.Manager
}

// NewAppRegistryAPI makes new *AppRegistryAPI struct
func NewAppRegistryAPI(backend service.Backend) (api.API, error) {
	ar := apps.NewManager(backend.Client())
	return &AppRegistryAPI{ar}, nil
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string appName) returns()
func (api *AppRegistryAPI) Register(c *gin.Context) {
	appName := c.Param(":appName")
	if err := api.apps.Register(c, appName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string appName) returns()
func (api *AppRegistryAPI) Unregister(c *gin.Context) {
	appName := c.Param(":appName")
	if err := api.apps.Unregister(c, appName); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string appName) constant returns((string,address,bytes32))
func (api *AppRegistryAPI) Get(c *gin.Context) {
	appName := c.Param(":appName")

	app, err := api.apps.Get(appName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, app)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0x1a9dff9f.
//
// Solidity: function transferAppOwner(string appName, address newOwner) returns()
func (api *AppRegistryAPI) TransferAppOwner(c *gin.Context) {
	appName := c.Param(":appName")

	var req struct{ NewOwner string }
	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	newOwner := common.HexToAddress(req.NewOwner)
	if err := api.apps.TransferAppOwner(c, appName, newOwner); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// AttachToAPI is a registrant of an api.
func (api *AppRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/apps")
	apiMux.GET("/:appName", api.Register)
	apiMux.POST("/:appName", api.Unregister)
	apiMux.PATCH("/:appName", api.Get)
	apiMux.DELETE("/:appName", api.TransferAppOwner)
}
