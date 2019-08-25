package api

import (
	"encoding/hex"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/klaytn/klaytn/common"
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account_id": accountId.Hex()})
}

// CreateTemporary is a paid mutator transaction binding the contract method 0x56003f0f.
//
// Solidity: function createTemporary(bytes32 identityHash) returns()
func (api *accountsAPI) createTemporary(c *gin.Context) {
	var req struct {
		IdentityHash string `json:"identity_hash" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	identityHash := common.HexToHash(req.IdentityHash)
	accountId, err := api.accounts.CreateTemporary(c, identityHash)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account_id": accountId.Hex()})
}

// UnlockTemporary is a paid mutator transaction binding the contract method 0x2299219d.
//
// Solidity: function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) returns()
func (api *accountsAPI) unlockTemporary(c *gin.Context) {
	var req struct {
		IdentityPreimage  string `json:"identity_preimage" binding:"required"`
		NewOwner          string `json:"new_owner" binding:"required"`
		PasswordSignature string `json:"password_signature" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	identityPreimage := common.HexToHash(req.IdentityPreimage)
	newOwner := common.HexToAddress(req.NewOwner)
	passwordSignature, err := hex.DecodeString(req.PasswordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.accounts.UnlockTemporary(c, identityPreimage, newOwner, passwordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address controller) returns()
func (api *accountsAPI) setController(c *gin.Context) {
	var req struct {
		Controller string `json:"controller" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller := common.HexToAddress(req.Controller)
	if err := api.accounts.SetController(c, controller); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetAccount is a free data retrieval call binding the contract method 0xf9292ddb.
//
// Solidity: function getAccount(bytes8 accountId) constant returns((address,uint8,address,address))
func (api *accountsAPI) getAccount(c *gin.Context) {
	var req struct {
		AccountId string `form:"account_id" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountId, err := types.HexToID(req.AccountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := api.accounts.GetAccount(accountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	var (
		accountIdRequest struct {
			Owner string `form:"owner" binding:"required"`
		}
		accountIdFromSigRequest struct {
			MessageHash string `form:"message_hash" binding:"required"`
			Signature   string `form:"signature" binding:"required"`
		}
	)

	isAccountId := c.ShouldBindWith(&accountIdRequest, binding.Query)
	isAccountIdFromSig := c.ShouldBindWith(&accountIdFromSigRequest, binding.Query)

	if isAccountId == nil {
		owner := common.HexToAddress(accountIdRequest.Owner)
		accountId, err := api.accounts.GetAccountId(owner)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account_id": accountId.Hex()})
		return
	}

	if isAccountIdFromSig == nil {
		messageHash := common.HexToHash(accountIdFromSigRequest.MessageHash)
		signature, err := hex.DecodeString(accountIdFromSigRequest.Signature)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		accountId, err := api.accounts.GetAccountIdFromSignature(messageHash, signature)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account_id": accountId.Hex()})
		return
	}

	accountId, err := api.accounts.GetAccountId(common.Address{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account_id": accountId.Hex()})
	return
}

// AttachToAPI is a registrant of an api.
func (api *accountsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/accounts")
	apiMux.GET("/", api.getAccount)
	apiMux.GET("/id", api.getAccountId)
	apiMux.POST("/", api.create)
	apiMux.POST("/temporary", api.createTemporary)
	apiMux.PATCH("/controller", api.setController)
	apiMux.PATCH("/temporary/unlock", api.unlockTemporary)
}
