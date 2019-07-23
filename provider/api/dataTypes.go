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

type DataTypeRegistryAPI struct {
	dataTypes *dataTypes.Manager
}

func NewDataTypeRegistryAPI(backend service.Backend) (api.API, error) {
	dt := dataTypes.NewManager(backend.Client())
	return &DataTypeRegistryAPI{dt}, nil
}

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

func (api *DataTypeRegistryAPI) Unregister(c *gin.Context) {
	name := c.Param("name")
	if err := api.dataTypes.Unregister(c, name); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *DataTypeRegistryAPI) Get(c *gin.Context) {
	name := c.Param("name")
	dataType, err := api.dataTypes.Get(c, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, dataType)
}

func (api *DataTypeRegistryAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/data-types")
	apiMux.GET("/:name", api.Get)
	apiMux.POST("/", api.Register)
	apiMux.DELETE("/:name", api.Unregister)
}
