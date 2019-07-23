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
	rawAccountId := c.Param("accountId")
	accountId, err := types.HexToID(rawAccountId)
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
	if rawOwner, ok := c.GetQuery("owner"); ok {
		owner := common.HexToAddress(rawOwner)
		accountId, err := api.accounts.GetAccountId(owner)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"accountId": accountId.Hex()})
		return
	}

	rawMessageHash, ok1 := c.GetQuery("messageHash")
	rawSignature, ok2 := c.GetQuery("signature")
	if ok1 && ok2 {
		messageHash := common.HexToHash(rawMessageHash)
		signature, err := hex.DecodeString(rawSignature)
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

func (api *AccountsAPI) IsTemporary(c *gin.Context) {
	if rawAccountId, ok := c.GetQuery("accountId"); ok {
		accountId, err := types.HexToID(rawAccountId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		isTemporary, err := api.accounts.IsTemporary(accountId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"isTemporary": isTemporary})
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
}

func (api *AccountsAPI) IsControllerOf(c *gin.Context) {
	rawController, ok1 := c.GetQuery("controller")
	rawAccountId, ok2 := c.GetQuery("accountId")
	if ok1 && ok2 {
		controller := common.HexToAddress(rawController)
		accountId, err := types.HexToID(rawAccountId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		isControllerOf, err := api.accounts.IsControllerOf(controller, accountId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"isControllerOf": isControllerOf})
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
}

func (api *AccountsAPI) Exists(c *gin.Context) {
	rawAccountId := c.Param("accountId")
	accountId, err := types.HexToID(rawAccountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	exist, err := api.accounts.Exists(accountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exist": exist})
}

func (api *AccountsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.RestAPIMux.Group("/accounts")
	apiMux.POST("/", api.Create)

	apiMux.GET("/:accountId", api.GetAccount)
	apiMux.GET("/:accountId/check", api.Exists)

	apiMux.GET("/id", api.GetAccountId)

	apiMux.PATCH("/controller", api.SetController)
	apiMux.GET("/controller/check", api.IsControllerOf)

	apiMux.POST("/temporary", api.CreateTemporary)
	apiMux.GET("/temporary/check", api.IsTemporary)
	apiMux.PATCH("/temporary/unlock", api.UnlockTemporary)
}
