package dauth

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/azer/logger"
)

// Validator ensures data providers to only register data to Airbloc
// from users who have allowed collection of the data through DAuth.
type Validator struct {
	dauth *Manager
	log   *logger.Logger
}

// NewValidator creates validator instance.
func NewValidator(manager *Manager) *Validator {
	return &Validator{
		dauth: manager,
		log:   logger.New("dauth-validator"),
	}
}

// IsCollectible returns true if the owner of the given data
// has authorized data collection of the given collection (data type).
func (validator *Validator) IsCollectible(collectionId common.ID, data *common.Data) bool {
	allowed, err := validator.dauth.IsCollectionAllowed(collectionId, data.OwnerAnID)
	if err != nil {
		validator.log.Error("error: %s", err.Error(), logger.Attrs{
			"collectionId": collectionId.Hex(),
			"user":         data.OwnerAnID.Hex(),
		})
	}
	return allowed
}
