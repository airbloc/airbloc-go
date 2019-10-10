package api

import (
	"encoding/hex"
	"net/http"

	"github.com/klaytn/klaytn/common"

	"github.com/klaytn/klaytn/crypto"

	"github.com/airbloc/airbloc-go/shared/blockchain"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// consentsAPI is api wrapper of contract Consents.sol
type consentsAPI struct {
	apps     adapter.IAppRegistryManager
	consents adapter.IConsentsManager
	feePayer common.Address
}

// NewConsentsAPI makes new *consentsAPI struct
func NewConsentsAPI(backend service.Backend) (api.API, error) {
	am := adapter.NewAppRegistryManager(backend.Client())
	cm := adapter.NewConsentsManager(backend.Client())
	ad := backend.Kms().NodeKey().EthereumAddress
	return &consentsAPI{am, cm, ad}, nil
}

// Consent scenario
// 1. basic
// 2. delegated
// 2-1. use nodekey
// 2-2.
// 3. by controller
// 3-1. initial consent (first time only)
// 3-2. modify consent

type consentRequest struct {
	// delegated
	PrivateKey string `json:"private_key"`

	// by controller
	AccountID         string            `json:"account_id"`
	PasswordSignature string            `json:"password_signature"`
	ConsentData       types.ConsentData `json:"consent_data" binding:"required"`
}

func (api *consentsAPI) consentHandler(c *gin.Context) {
	var req consentRequest
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appName := c.Param("app_name")
	if appName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
		return
	}

	if exists, err := api.apps.Exists(appName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot find app information"})
		return
	} else if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
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

	switch {
	case req.PasswordSignature != "":
		api.modifyConsentByController(c, txOpts, appName, req)
	case req.AccountID != "":
		api.consentByController(c, txOpts, appName, req)
	default:
		api.consent(c, txOpts, appName, req)
	}
}

func (api *consentsAPI) consent(c *gin.Context, txOpts *blockchain.TransactOpts, appName string, req consentRequest) {
	if err := api.consents.Consent(c, nil, appName, req.ConsentData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) consentByController(c *gin.Context, txOpts *blockchain.TransactOpts, appName string, req consentRequest) {
	accountID, err := types.HexToID(req.AccountID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = api.consents.ConsentByController(c, txOpts, accountID, appName, req.ConsentData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) modifyConsentByController(c *gin.Context, txOpts *blockchain.TransactOpts, appName string, req consentRequest) {
	accountID, err := types.HexToID(req.AccountID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordSignature, err := hex.DecodeString(req.PasswordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = api.consents.ModifyConsentByController(c, txOpts, accountID, appName, req.ConsentData, passwordSignature); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type consentManyRequest struct {
	// delegated
	PrivateKey string `json:"private_key"`

	// by controller
	AccountID         string              `json:"account_id"`
	PasswordSignature string              `json:"password_signature"`
	ConsentData       []types.ConsentData `json:"consent_data" binding:"required"`
}

func (api *consentsAPI) consentManyHandler(c *gin.Context) {
	var req consentManyRequest
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appName := c.Param("app_name")
	if appName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
		return
	}

	if exists, err := api.apps.Exists(appName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot find app information"})
		return
	} else if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
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

	switch {
	case req.PasswordSignature != "":
		api.modifyConsentManyByController(c, txOpts, appName, req.ConsentData, req)
	case req.AccountID != "":
		api.consentManyByController(c, txOpts, appName, req.ConsentData, req)
	default:
		api.consentMany(c, txOpts, appName, req.ConsentData, req)
	}
}

func (api *consentsAPI) consentMany(
	c *gin.Context,
	txOpts *blockchain.TransactOpts,
	appName string,
	consentData []types.ConsentData,
	req consentManyRequest,
) {
	if err := api.consents.ConsentMany(c, txOpts, appName, req.ConsentData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) consentManyByController(
	c *gin.Context,
	txOpts *blockchain.TransactOpts,
	appName string,
	consentData []types.ConsentData,
	req consentManyRequest,
) {
	accountID, err := types.HexToID(req.AccountID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = api.consents.ConsentManyByController(c, txOpts, accountID, appName, consentData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) modifyConsentManyByController(
	c *gin.Context,
	txOpts *blockchain.TransactOpts,
	appName string,
	consentData []types.ConsentData,
	req consentManyRequest,
) {
	accountID, err := types.HexToID(req.AccountID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordSignature, err := hex.DecodeString(req.PasswordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = api.consents.ModifyConsentManyByController(
		c, txOpts,
		accountID,
		appName,
		consentData,
		passwordSignature,
	); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// AttachToAPI is a registrant of an api.
func (api *consentsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/consents")
	apiMux.PUT("/:app_name", api.consentHandler)
	apiMux.PUT("/:app_name/many", api.consentManyHandler)
}
