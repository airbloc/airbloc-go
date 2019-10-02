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

// dataTypeRegistryAPI is api wrapper of contract DataTypeRegistry.sol
type dataTypeRegistryAPI struct {
	dataTypes adapter.IDataTypeRegistryManager
}

// NewDataTypeRegistryAPI makes new *dataTypeRegistryAPI struct
func NewDataTypeRegistryAPI(backend service.Backend) (api.API, error) {
	dt := adapter.NewDataTypeRegistryManager(backend.Client())
	return &dataTypeRegistryAPI{dt}, nil
}

func (api *dataTypeRegistryAPI) register(c *gin.Context) {
	var req struct {
		Name       string `json:"name" binding:"required"`
		SchemaHash string `json:"schema_hash" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var (
		name       = req.Name
		schemaHash = common.HexToHash(req.SchemaHash)
	)

	if exists, err := api.dataTypes.Exists(name); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if exists {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{})
		return
	}

	if err := api.dataTypes.Register(c, nil, name, schemaHash); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (api *dataTypeRegistryAPI) unregister(c *gin.Context) {
	dataType := c.Param("data_type")
	if dataType == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := api.dataTypes.Unregister(c, nil, dataType); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (api *dataTypeRegistryAPI) get(c *gin.Context) {
	dataType := c.Param("data_type")
	if dataType == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	typ, err := api.dataTypes.Get(dataType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, typ)
}

// AttachToAPI is a registrant of an api.
func (api *dataTypeRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/data-types")
	apiMux.GET("/:data_type", api.get)
	apiMux.POST("/", api.register)
	apiMux.DELETE("/:data_type", api.unregister)
}
