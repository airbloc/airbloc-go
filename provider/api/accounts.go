package api

import (
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/ethereum/go-ethereum/common"

	"github.com/airbloc/airbloc-go/provider/accounts"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"

	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
)

type AccountsAPI struct {
	accounts *accounts.Manager
}

func NewAccountsAPI(backend service.Backend) (api.API, error) {
	ac := accounts.NewManager(backend.Client())
	return &AccountsAPI{ac}, nil
}

func (api *AccountsAPI) Create(c *gin.Context) {
	accountId, err := api.accounts.Create(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
}

func (api *AccountsAPI) CreateTemporary(c *gin.Context) {
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

func (api *AccountsAPI) UnlockTemporary(c *gin.Context) {
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

	if err := api.accounts.UnlockTemporary(c, identityPreimage, newOwner, passwordSignature); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api *AccountsAPI) SetController(c *gin.Context) {
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

func (api *AccountsAPI) GetAccount(c *gin.Context) {
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

func (api *AccountsAPI) GetAccountId(c *gin.Context) {
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

func (api *AccountsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/accounts")
	apiMux.POST("/", api.Create)
	apiMux.GET("/:accountId", api.GetAccount)
	apiMux.GET("/id", api.GetAccountId)
	apiMux.PATCH("/controller", api.SetController)
	apiMux.POST("/temporary", api.CreateTemporary)
	apiMux.PATCH("/temporary/unlock", api.UnlockTemporary)
}
