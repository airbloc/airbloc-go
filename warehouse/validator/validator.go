package validator

import (
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/logger"
)

// Validator ensures data providers to only register data to Airbloc
// from users who have allowed collection of the data through DAuth.
type Validator struct {
	consents adapter.IConsentsManager
	log      *logger.Logger
}

// NewValidator creates validator instance.
func NewValidator(manager adapter.IConsentsManager) *Validator {
	return &Validator{
		consents: manager,
		log:      logger.New("dauth-validator"),
	}
}

// IsCollectible returns true if the owner of the given data
// has authorized data collection of the given collection (data type).
func (validator *Validator) IsCollectible(appName string, action types.ConsentActionTypes, dataType string, data *types.Data) bool {
	allowed, err := validator.consents.IsAllowed(data.UserId, appName, uint8(action), dataType)
	if err != nil {
		validator.log.Error("error: {} {}", err.Error(), logger.Attrs{
			"app-name":  appName,
			"action":    action,
			"data-type": dataType,
			"user":      data.UserId.Hex(),
		})
	}
	return allowed
}
