package warehouse

import (
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
)

type Service struct {
	manager *Manager
}

func NewService(backend service.Backend) (service.Service, error) {
	var err error
	config := backend.Config().Warehouse

	supportedProtocols := []protocol.Protocol{
		protocol.NewHttpProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
		protocol.NewHttpsProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
	}

	var defaultStorage storage.Storage
	switch config.DefaultStorage {
	case "local":
		defaultStorage, err = storage.NewLocalStorage(
			config.LocalStorage.SavePath,
			config.LocalStorage.Endpoint,
		)
		if err != nil {
			return nil, err
		}
	case "s3":
		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewEnvCredentials(),
			Region:      aws.String(config.S3.Region),
		})
		if err != nil {
			return nil, err
		}
		defaultStorage = storage.NewS3Storage(config.S3.Bucket, config.S3.PathPrefix, sess)
	default:
		return nil, errors.Errorf("unknown storage type: %s", config.DefaultStorage)
	}

	dw, err := NewManager(
		backend.Kms(),
		backend.LocalDatabase(),
		backend.MetaDatabase(),
		backend.Client(),
		defaultStorage,
		supportedProtocols,
		*backend.Config(),
	)
	return &Service{manager: dw}, err
}

func (service *Service) GetManager() *Manager { return service.manager }
func (service *Service) Start() error         { return nil }
func (service *Service) Stop()                {}
