package api

import (
	"net/http"

	"github.com/airbloc/airbloc-go/provider/data"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/warehouse"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type dataAPI struct {
	manager *data.Manager
}

func NewDataAPI(backend service.Backend) (api.API, error) {
	manager := data.NewManager(
		backend.Kms(),
		backend.P2P(),
		backend.MetaDatabase(),
		backend.LocalDatabase(),
		backend.Client(),
		backend.GetService("warehouse").(*warehouse.Service).GetManager())
	return &dataAPI{manager}, nil
}

func (api *dataAPI) GetData(c *gin.Context) {
	rawDataId := c.Param("dataId")
	dataId, err := types.NewDataIdFromStr(rawDataId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := api.manager.Get(dataId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (api *dataAPI) GetBatch(c *gin.Context) {
	batchId := c.Param("batchId")
	batchManager := api.manager.Batches()
	batchInfo, err := batchManager.Get(batchId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	res, err := api.manager.GetBatch(batchInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (api *dataAPI) GetBundle(c *gin.Context) {
	rawBundleId := c.Param("bundleId")
	bundleId, err := types.HexToID(rawBundleId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := api.manager.GetBundle(c, bundleId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (api *dataAPI) SetPermission(c *gin.Context) {
	var req struct {
		DataId     string
		ConsumerId string
		Allowed    bool
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	// TODO: implement SetPermission

	c.AbortWithStatus(http.StatusNotImplemented)
}

func (api *dataAPI) SetPermissionBatch(c *gin.Context) {
	var req struct {
		BatchId    string
		ConsumerId string
		Allowed    bool
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	// TODO: implement SetPermission

	c.AbortWithStatus(http.StatusNotImplemented)
}

func (api *dataAPI) Delete(c *gin.Context) {
	rawDataId := c.Param("dataId")
	dataId, err := types.NewDataIdFromStr(rawDataId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_ = dataId
	// TODO: implement Delete

	c.AbortWithStatus(http.StatusNotImplemented)
}

func (api *dataAPI) DeleteBatch(c *gin.Context) {
	batchId := c.Param("batchId")
	batchManager := api.manager.Batches()
	batchInfo, err := batchManager.Get(batchId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	_ = batchInfo
	// TODO: implement DeleteBatch

	c.AbortWithStatus(http.StatusNotImplemented)
}

//func (api *dataAPI) Select(stream pb.Data_SelectServer) error {
//	return status.Error(codes.Unimplemented, "unimplemented method")
//}
//
//func (api *dataAPI) Release(c *gin.Context, batchId *pb.BatchRequest) (*empty.Empty, error) {
//	return nil, status.Error(codes.Unimplemented, "unimplemented method")
//}

func (api *dataAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/data")
	apiMux.GET("/:dataId", api.GetData)
	apiMux.GET("/batch/:batchId", api.GetBatch)
	apiMux.GET("/bundle/:bundleId", api.GetBundle)
	apiMux.PUT("/permission", api.SetPermission)
	apiMux.PUT("/permission/batch", api.SetPermissionBatch)
	apiMux.DELETE("/:dataId", api.Delete)
	apiMux.DELETE("/batch/:batchId", api.DeleteBatch)
	// TODO
	// apiMux.[Method]([path], api.Select)
	// apiMux.[Method]([path], api.Release)
}
