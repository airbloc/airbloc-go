package api

import (
	"encoding/hex"
	"math/big"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/provider/consents"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ConsentsAPI struct {
	consents *consents.Manager
}

func NewConsentsAPI(backend service.Backend) (api.API, error) {
	cs := consents.NewManager(backend.Client())
	return &ConsentsAPI{cs}, nil
}

func (api *ConsentsAPI) Consent(c *gin.Context) {
	var req struct {
		Action   uint8
		AppName  string
		DataType string
		Allowed  bool
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	err = api.consents.Consent(
		c, req.Action,
		req.AppName, req.DataType, req.Allowed,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *ConsentsAPI) ConsentByController(c *gin.Context) {
	var req struct {
		Action   uint8
		UserId   string
		AppName  string
		DataType string
		Allowed  bool
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	userId, err := types.HexToID(req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = api.consents.ConsentByController(
		c, req.Action,
		userId, req.AppName, req.DataType, req.Allowed,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *ConsentsAPI) ModifyConsentByController(c *gin.Context) {
	var req struct {
		Action            uint8
		UserId            string
		AppName           string
		DataType          string
		Allowed           bool
		PasswordSignature string
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	userId, err := types.HexToID(req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	passwordSignature, err := hex.DecodeString(req.PasswordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = api.consents.ModifyConsentByController(
		c, req.Action,
		userId, req.AppName, req.DataType, req.Allowed, passwordSignature,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *ConsentsAPI) IsAllowed(c *gin.Context) {
	var req struct {
		Action      uint8
		UserId      string
		AppName     string
		DataType    string
		BlockNumber string
	}

	if err := c.MustBindWith(&req, binding.Query); err != nil {
		return
	}

	userId, err := types.HexToID(req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if req.BlockNumber == "" {
		allowed, err := api.consents.IsAllowed(
			c, req.Action,
			userId, req.AppName, req.DataType,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"allowed": allowed})
	}

	blockNumber, ok := new(big.Int).SetString(req.BlockNumber, 10)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	allowed, err := api.consents.IsAllowedAt(
		c, req.Action,
		userId, req.AppName, req.DataType, blockNumber,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"allowed": allowed})
}

func (api *ConsentsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/consents")
	apiMux.POST("/", api.Consent)
	apiMux.GET("/allowed", api.IsAllowed)
	apiMux.PUT("/controller", api.ModifyConsentByController)
	apiMux.POST("/controller", api.ConsentByController)
}
