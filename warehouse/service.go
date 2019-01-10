package warehouse

import (
	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
)

type Service struct {
	manager *DataWarehouse
}

func NewService(backend node.Backend) (node.Service, error) {
	var err error
	config := backend.Config().Warehouse

	supportedProtocols := []protocol.Protocol{
		protocol.NewHttpProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
		protocol.NewHttpsProtocol(config.Http.Timeout, config.Http.MaxConnsPerHost),
	}

	var defaultStorage storage.Storage
	switch storage.Type_value[config.DefaultStorage] {
	case storage.Local:
		cfg := config.LocalStorage
		defaultStorage, err = storage.NewLocalStorage(
			cfg.SavePath,
			cfg.Endpoint)

		if err != nil {
			return nil, err
		}
	case storage.CloudS3:
		cfg := config.S3

		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKey,
				cfg.SecretKey,
				cfg.Token,
			),
			Region: aws.String(cfg.Region),
		})
		if err != nil {
			return nil, err
		}

		defaultStorage = storage.NewS3Storage(cfg.Bucket, cfg.PathPrefix, sess)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("unknown storage type: %s", config.DefaultStorage)
	}

	dw := New(
		backend.Kms(),
		backend.LocalDatabase(),
		backend.MetaDatabase(),
		backend.Client(),
		defaultStorage,
		supportedProtocols,
	)

	return &Service{manager: dw}, nil
}

func (service *Service) GetManager() *DataWarehouse { return service.manager }
func (service *Service) Start() error               { return nil }
func (service *Service) Stop()                      {}
