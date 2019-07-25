package api

import (
	"encoding/hex"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// accountsAPI is api wrapper of contract Accounts.sol
type accountsAPI struct {
	accounts adapter.IAccountsManager
}

// NewAccountsAPI makes new *accountsAPI struct
func NewAccountsAPI(backend service.Backend) (api.API, error) {
	ac := adapter.NewAccountsManager(backend.Client())
	return &accountsAPI{ac}, nil
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (api *accountsAPI) create(c *gin.Context) {
	accountId, err := api.accounts.Create(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns()
func (api *accountsAPI) createTemporary(c *gin.Context) {
	var req struct {
		IdentityHash string
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	identityHash := common.HexToHash(req.IdentityHash)
	accountId, err := api.accounts.CreateTemporary(c, identityHash)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (api *accountsAPI) unlockTemporary(c *gin.Context) {
	var req struct {
		IdentityPreimage  string
		NewOwner          string
		PasswordSignature string
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	identityPreimage := common.HexToHash(req.IdentityPreimage)
	newOwner := common.HexToAddress(req.NewOwner)
	passwordSignature, err := hex.DecodeString(req.PasswordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = api.accounts.UnlockTemporary(c, identityPreimage, newOwner, passwordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (api *accountsAPI) setController(c *gin.Context) {
	var req struct {
		Controller string
	}

	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		return
	}

	controller := common.HexToAddress(req.Controller)
	if err := api.accounts.SetController(c, controller); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (api *accountsAPI) getAccount(c *gin.Context) {
	var req struct {
		AccountId string
	}

	if err := c.MustBindWith(&req, binding.Query); err != nil {
		return
	}

	accountId, err := types.HexToID(req.AccountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	account, err := api.accounts.GetAccount(accountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, account)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
// GetAccountIdFromSignature is a free data retrieval call binding the contract method 0x23d0601d.
//
// Solidity: function getAccountId(address sender) constant returns(bytes8)
// Solidity: function getAccountIdFromSignature(bytes32 messageHash, bytes signature) constant returns(bytes8)
func (api *accountsAPI) getAccountId(c *gin.Context) {
	var req struct {
		Owner       string `form:"owner"`
		MessageHash string `form:"messageHash"`
		Signature   string `form:"signature"`
	}

	if err := c.MustBindWith(&req, binding.Query); err != nil {
		return
	}

	if req.Owner != "" {
		owner := common.HexToAddress(req.Owner)
		accountId, err := api.accounts.GetAccountId(owner)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
		return
	}

	if req.MessageHash != "" && req.Signature != "" {
		messageHash := common.HexToHash(req.MessageHash)
		signature, err := hex.DecodeString(req.Signature)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		accountId, err := api.accounts.GetAccountIdFromSignature(messageHash, signature)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
}

// AttachToAPI is a registrant of an api.
func (api *accountsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/accounts")
	apiMux.POST("/", api.create)
	apiMux.GET("/:accountId", api.getAccount)
	apiMux.GET("/id", api.getAccountId)
	apiMux.PATCH("/controller", api.setController)
	apiMux.POST("/temporary", api.createTemporary)
	apiMux.PATCH("/temporary/unlock", api.unlockTemporary)
}
