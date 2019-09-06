package api

import (
	"encoding/hex"
	"net/http"

	"github.com/airbloc/airbloc-go/shared/types"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/service/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// consentsAPI is api wrapper of contract Consents.sol
type consentsAPI struct {
	apps     adapter.IAppRegistryManager
	consents adapter.IConsentsManager
}

// NewConsentsAPI makes new *consentsAPI struct
func NewConsentsAPI(backend service.Backend) (api.API, error) {
	am := adapter.NewAppRegistryManager(backend.Client())
	cm := adapter.NewConsentsManager(backend.Client())
	return &consentsAPI{am, cm}, nil
}

type consentRequest struct {
	UserId            string                   `json:"user_id"`
	ConsentData       types.ConsentRequestData `json:"consent_data" binding:"required"`
	PasswordSignature string                   `json:"password_signature"`
}

func (api *consentsAPI) consentByController(c *gin.Context, appName string, req consentRequest) {
	userId, err := types.HexToID(req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = api.consents.ConsentByController(
		c, userId, appName,
		req.ConsentData.ToConsentData(),
	); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) modifyConsentByController(c *gin.Context, appName string, req consentRequest) {
	userId, err := types.HexToID(req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordSignature, err := hex.DecodeString(req.PasswordSignature)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = api.consents.ModifyConsentByController(
		c, userId, appName,
		req.ConsentData.ToConsentData(),
		passwordSignature,
	); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) consent(c *gin.Context) {
	req := consentRequest{}
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appName := c.Param("appName")
	if appName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
		return
	}

	exists, err := api.apps.Exists(appName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot find app information"})
		return
	}

	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
	}

	if req.ConsentData == (types.ConsentRequestData{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid consent data"})
		return
	}

	if req.PasswordSignature != "" {
		api.modifyConsentByController(c, appName, req)
		return
	}

	if req.UserId != "" {
		api.consentByController(c, appName, req)
		return
	}

	if err := api.consents.Consent(c, appName, req.ConsentData.ToConsentData()); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type consentManyRequest struct {
	UserId            string                     `json:"user_id"`
	ConsentData       []types.ConsentRequestData `json:"consent_data" binding:"required"`
	PasswordSignature string                     `json:"password_signature"`
}

func (api *consentsAPI) consentManyByController(c *gin.Context, appName string, consentData []types.ConsentData, req consentManyRequest) {
	userId, err := types.HexToID(req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = api.consents.ConsentManyByController(c, userId, appName, consentData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) modifyConsentManyByController(c *gin.Context, appName string, consentData []types.ConsentData, req consentManyRequest) {
	userId, err := types.HexToID(req.UserId)
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
		c, userId,
		appName,
		consentData,
		passwordSignature,
	); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (api *consentsAPI) consentMany(c *gin.Context) {
	req := consentManyRequest{}
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appName := c.Param("appName")
	if appName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
		return
	}

	exists, err := api.apps.Exists(appName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot find app information"})
		return
	}

	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid app name"})
	}

	if len(req.ConsentData) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid consent data"})
		return
	}

	consentData := make([]types.ConsentData, len(req.ConsentData))
	for index, data := range req.ConsentData {
		consentData[index] = data.ToConsentData()
	}

	if req.PasswordSignature != "" {
		api.modifyConsentManyByController(c, appName, consentData, req)
		return
	}

	if req.UserId != "" {
		api.consentManyByController(c, appName, consentData, req)
		return
	}

	if err = api.consents.ConsentMany(c, appName, consentData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// AttachToAPI is a registrant of an api.
func (api *consentsAPI) AttachToAPI(service *api.Service) {
	apiMux := service.HttpServer.Group("/consents")
	apiMux.PUT("/:appName", api.consent)
	apiMux.PUT("/:appName/many", api.consentMany)
}
