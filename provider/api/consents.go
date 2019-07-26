package api

import (
	"encoding/hex"
	"math/big"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/consents"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ConsentsAPI is api wrapper of contract Consents.sol
type consentsAPI struct {
	consents adapter.IConsentsManager
}

// NewConsentsAPI makes new *ConsentsAPI struct
func NewConsentsAPI(backend service.Backend) (api.API, error) {
	cs := consents.NewManager(backend.Client())
	return &consentsAPI{cs}, nil
}

// Consent is a paid mutator transaction binding the contract method 0xbecae241.
//
// Solidity: function consent(uint8 action, string appName, string dataType, bool allowed) returns()
func (api *consentsAPI) consent(c *gin.Context) {
	var req struct {
		Action   uint8
		AppName  string
		DataType string
		Allowed  bool
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	err := api.consents.Consent(
		c, req.Action,
		req.AppName, req.DataType, req.Allowed,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// ConsentByController is a paid mutator transaction binding the contract method 0xf92519d8.
//
// Solidity: function consentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed) returns()
func (api *consentsAPI) consentByController(c *gin.Context) {
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

// ModifyConsentByController is a paid mutator transaction binding the contract method 0xedf2ef20.
//
// Solidity: function modifyConsentByController(uint8 action, bytes8 userId, string appName, string dataType, bool allowed, bytes passwordSignature) returns()
func (api *consentsAPI) modifyConsentByController(c *gin.Context) {
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

// IsAllowed is a free data retrieval call binding the contract method 0xa1d2bbf5.
// IsAllowedAt is a free data retrieval call binding the contract method 0x118642e1.
//
// Solidity: function isAllowed(uint8 action, bytes8 userId, string appName, string dataType) constant returns(bool)
// Solidity: function isAllowedAt(uint8 action, bytes8 userId, string appName, string dataType, uint256 blockNumber) constant returns(bool)
func (api *consentsAPI) isAllowed(c *gin.Context) {
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
		allowed, rerr := api.consents.IsAllowed(req.Action, userId, req.AppName, req.DataType)
		if rerr != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": rerr})
			return
		}
		c.JSON(http.StatusOK, gin.H{"allowed": allowed})
	}

	blockNumber, ok := new(big.Int).SetString(req.BlockNumber, 10)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	allowed, err := api.consents.IsAllowedAt(req.Action, userId, req.AppName, req.DataType, blockNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"allowed": allowed})
}

// AttachToAPI is a registrant of an api.
func (api *consentsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/consents")
	apiMux.POST("/", api.consent)
	apiMux.GET("/allowed", api.isAllowed)
	apiMux.PUT("/controller", api.modifyConsentByController)
	apiMux.POST("/controller", api.consentByController)
}
