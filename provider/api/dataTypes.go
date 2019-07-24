package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/provider/dataTypes"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// DataTypeRegistryAPI is api wrapper of contract DataTypeRegistry.sol
type DataTypeRegistryAPI struct {
	dataTypes *dataTypes.Manager
}

// NewDataTypeRegistryAPI makes new *DataTypeRegistryAPI struct
func NewDataTypeRegistryAPI(backend service.Backend) (api.API, error) {
	dt := dataTypes.NewManager(backend.Client())
	return &DataTypeRegistryAPI{dt}, nil
}

// Register is a paid mutator transaction binding the contract method 0x656afdee.
//
// Solidity: function register(string name, bytes32 schemaHash) returns()
func (api *DataTypeRegistryAPI) Register(c *gin.Context) {
	var req struct {
		Name       string
		SchemaHash string
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	var (
		name       = req.Name
		schemaHash = common.HexToHash(req.SchemaHash)
	)
	if err := api.dataTypes.Register(c, name, schemaHash); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string name) returns()
func (api *DataTypeRegistryAPI) Unregister(c *gin.Context) {
	name := c.Param("name")
	if err := api.dataTypes.Unregister(c, name); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string name) constant returns((string,address,bytes32))
func (api *DataTypeRegistryAPI) Get(c *gin.Context) {
	name := c.Param("name")
	dataType, err := api.dataTypes.Get(c, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, dataType)
}

// AttachToAPI is a registrant of an api.
func (api *DataTypeRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/data-types")
	apiMux.GET("/:name", api.Get)
	apiMux.POST("/", api.Register)
	apiMux.DELETE("/:name", api.Unregister)
}
