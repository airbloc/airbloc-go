package api

import (
	"encoding/hex"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
)

// accountsAPI is api wrapper of contract Accounts.sol
type accountsAPI struct {
	accounts adapter.IAccountsManager
	feePayer common.Address
}

// NewAccountsAPI makes new *accountsAPI struct
func NewAccountsAPI(backend service.Backend) (api.API, error) {
	ac := adapter.NewAccountsManager(backend.Client())
	ad := backend.Kms().NodeKey().EthereumAddress
	return &accountsAPI{ac, ad}, nil
}

func (api *accountsAPI) create(c *gin.Context) {
	var req struct {
		PrivateKey string `json:"private_key"`
		Controller string `json:"controller"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var txOpts *blockchain.TransactOpts = nil
	if req.PrivateKey != "" {
		privateKeyData, err := hex.DecodeString(req.PrivateKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		key, err := crypto.ToECDSA(privateKeyData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		txOpts = blockchain.NewKeyedTransactor(key)
		txOpts.FeePayer = api.feePayer
		txOpts.Context = c
	}

	accountId, err := api.accounts.Create(c, txOpts)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Controller != "" {
		err = api.accounts.SetController(c, txOpts, common.HexToAddress(req.Controller))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"account_id": accountId.Hex()})
}

func (api *accountsAPI) createTemporary(c *gin.Context) {
	var req struct {
		IdentityHash string `json:"identity_hash" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	identityHash := common.HexToHash(req.IdentityHash)
	accountId, err := api.accounts.CreateTemporary(c, nil, identityHash)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"account_id": accountId.Hex()})
}

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

	err = api.accounts.UnlockTemporary(c, nil, identityPreimage, newOwner, passwordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *accountsAPI) setController(c *gin.Context) {
	var req struct {
		Controller string `json:"controller" binding:"required"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller := common.HexToAddress(req.Controller)
	if err := api.accounts.SetController(c, nil, controller); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

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

func (api *accountsAPI) getAccountId(c *gin.Context) {
	var (
		accountIdRequest struct {
			Owner string `form:"owner" binding:"required"`
		}
		accountIdFromSigRequest struct {
			MessageHash string `form:"message_hash" binding:"required"`
			Signature   string `form:"signature" binding:"required"`
		}
		accountIdByIdentityRequest struct {
			IdentityHash string `form:"identity_hash" binding:"required"`
		}
	)

	isAccountId := c.ShouldBindWith(&accountIdRequest, binding.Query)
	isAccountIdFromSig := c.ShouldBindWith(&accountIdFromSigRequest, binding.Query)
	isAccountIdByIdentity := c.ShouldBindWith(&accountIdByIdentityRequest, binding.Query)

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

	if isAccountIdByIdentity == nil {
		identityHash := common.HexToHash(accountIdByIdentityRequest.IdentityHash)
		accountId, err := api.accounts.GetAccountIdByIdentityHash(identityHash)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account_id": accountId.Hex()})
		return
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Cannot find account id"})
}

// AttachToAPI is a registrant of an api.
func (api *accountsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/accounts")
	apiMux.GET("/", api.getAccount)
	apiMux.GET("/id", api.getAccountId)
	apiMux.POST("/", api.create)
	apiMux.POST("/temporary", api.createTemporary)
	apiMux.PUT("/controller", api.setController)
	apiMux.PUT("/temporary/unlock", api.unlockTemporary)
}
