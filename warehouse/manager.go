package warehouse

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"net/url"
	"time"

	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/database/localdb"
	"github.com/airbloc/airbloc-go/shared/database/metadb"
	"github.com/airbloc/airbloc-go/shared/database/resdb"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/service"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/airbloc/airbloc-go/warehouse/validator"
	"github.com/airbloc/airframe/afclient"
	"github.com/airbloc/logger"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var (
	// errValidationFailed is returned when an incoming data
	// failed to met condition of the one of the validators (DAuth, Schema, ...)
	errValidationFailed = errors.New("data validation failed.")
)

type Manager struct {
	kms        key.Manager
	localCache *localdb.Model

	// for data registration
	metaDatabase metadb.Database
	ethclient    blockchain.TxClient
	//dataRegistry *adapter.DataRegistry

	// data storage layer
	protocols      map[string]protocol.Protocol
	defaultStorage storage.Storage
	resourceDB     resdb.Model

	// data validators
	dauthValidator *validator.Validator

	config service.Config
	log    *logger.Logger
}

func NewManager(
	kms key.Manager,
	localDatabase localdb.Database,
	metaDatabase metadb.Database,
	ethclient blockchain.TxClient,
	defaultStorage storage.Storage,
	supportedProtocols []protocol.Protocol,
	config service.Config,
) (*Manager, error) {
	protocols := map[string]protocol.Protocol{}
	for _, protoc := range supportedProtocols {
		protocols[protoc.Name()] = protoc
	}

	log := logger.New("warehouse")
	if config.Warehouse.Debug.DisableSchemaValidation {
		log.Error("warning: You have disabled schema validation. " +
			"It is recommended to avoid disabling the validation on production mode.")
	}
	if config.Warehouse.Debug.DisableUserAuthValidation {
		log.Error("warning: You have disabled user auth validation. \n" + "\033[31m" +
			"DO NOT DISABLE THE USER VALIDATION ON PRODUCTION MODE, " +
			"BECAUSE IT CAN CAUSE A FINANCIAL LOSS OF YOUR STAKED COLLETRALS. " + "\033[0m")
	}

	consentManager := adapter.NewConsentsManager(ethclient)
	dauthValidator := validator.NewValidator(consentManager)

	//contract := ethclient.GetContract(&adapter.DataRegistry{})

	resdbClient, err := afclient.Dial(config.ResourceDB.Endpoint, kms.NodeKey().PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial airframe")
	}

	return &Manager{
		kms:        kms,
		localCache: localdb.NewModel(localDatabase, "bundle"),

		metaDatabase: metadb.NewModel(metaDatabase, "bundles"),
		ethclient:    ethclient,
		//dataRegistry: contract.(*adapter.DataRegistry),

		protocols:      protocols,
		defaultStorage: defaultStorage,
		resourceDB:     resdb.NewModel(resdbClient, "bundle"),
		dauthValidator: dauthValidator,

		config: config,
		log:    log,
	}, nil
}

// TODO
func (dw *Manager) CreateBundle(ctx context.Context, collectionId types.ID) (*BundleStream, error) {
	return nil, nil
	//collection, err := dw.collections.Get(collectionId)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to retrieve a collection")
	//}
	//schema, err := dw.schemas.Get(ctx, collection.Schema.Id)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to retrieve a schema")
	//}
	//collection.Schema = *schema
	//return newBundleStream(dw, collection.AppId, collection), nil
}

// TODO
func (dw *Manager) validate(collection interface{}, data *types.Data) error {
	//if !dw.config.Warehouse.Debug.DisableUserAuthValidation && !dw.dauthValidator.IsCollectible(collection.Id, data) {
	//	return errors.Wrap(errValidationFailed, "user hasn't been authorized the data collection")
	//}
	//
	//if !dw.config.Warehouse.Debug.DisableSchemaValidation {
	//	isValidFormat, err := collection.Schema.IsValidFormat(data)
	//	if err != nil {
	//		return err
	//	} else if !isValidFormat {
	//		return errors.Wrap(errValidationFailed, "wrong format")
	//	}
	//}
	return nil
}

func generateBundleNameOf(bundle *Bundle) string {
	tokenBytes := make([]byte, 4)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	currentTime := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s.bundle", currentTime, bundle.Collection.Hex(), token)
}

func (dw *Manager) Store(ctx context.Context, stream *BundleStream) (*Bundle, error) {
	if stream == nil {
		return nil, errors.New("No data in the stream.")
	}
	ingestedAt := types.Time{Time: time.Now()}

	createdBundle := &Bundle{
		Provider: stream.provider,
		//Collection: stream.collection.Id,
		DataCount:  stream.DataCount,
		IngestedAt: ingestedAt,
		Data:       stream.data,
	}

	if err := createdBundle.SetupRowId(); err != nil {
		return nil, errors.Wrap(err, "failed to setup row id")
	}

	bundleName := generateBundleNameOf(createdBundle)
	bundleData, err := createdBundle.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal bundle")
	}

	uri, err := dw.defaultStorage.Save(bundleName, bundleData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save bundle to the storage")
	}
	createdBundle.Uri = uri.String()
	dw.log.Debug(uri.String())

	// register to on-chain
	bundleId, err := dw.registerBundleOnChain(createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register bundle to blockchain")
	}
	createdBundle.Id = bundleId.Hex()

	var dataIds []map[string]interface{}
	for _, rowData := range createdBundle.Data {
		for _, d := range rowData {
			dataId := make(map[string]interface{}, 4)
			dataId["bundleId"] = bundleId.Hex()
			dataId["userId"] = d.UserId.Hex()
			dataId["rowId"] = d.RowId.Hex()
			dataId["collectedAt"] = d.CollectedAt.Timestamp()
			dataIds = append(dataIds, dataId)
		}
	}

	// save metadata to make the bundle searchable
	bundleInfo := map[string]interface{}{
		"bundleId":   createdBundle.Id,
		"uri":        createdBundle.Uri,
		"provider":   createdBundle.Provider.Hex(),
		"collection": createdBundle.Collection.Hex(),
		"dataCount":  createdBundle.DataCount,
		"ingestedAt": ingestedAt.Timestamp(),
		"dataIds":    dataIds,
	}

	dw.log.Debug("Putting bundleInfo to resourceDB")
	res, err := dw.resourceDB.Put(ctx, createdBundle.Id, bundleInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save bundleInfo to resourceDB")
	}
	if !res.Created {
		return nil, errors.Errorf("")
	}
	dw.log.Debug("Bundle successfully created. Fee used : {}", res.FeeUsed)
	dw.log.Info("Bundle {} registered on", bundleName, logger.Attrs{
		"id":    bundleId.Hex(),
		"count": bundleInfo["dataCount"],
	})
	return createdBundle, nil
}

// TODO
func (dw *Manager) registerBundleOnChain(bundle *Bundle) (bundleId types.ID, _ error) {
	//bundleDataHash, err := bundle.Hash()
	//if err != nil {
	//	return bundleId, errors.Wrap(err, "failed to get hash of the bundle data")
	//}
	//
	//// for setup rowId
	//userMerkleRoot, err := bundle.SetupUserProof()
	//if err != nil {
	//	return bundleId, errors.Wrap(err, "failed to setup SMT")
	//}
	//
	//dw.log.Info("Bundle data", logger.Attrs{
	//	"hash":       bundleDataHash.Hex(),
	//	"merkleRoot": userMerkleRoot.Hex(),
	//})
	//
	//tx, err := dw.dataRegistry.RegisterBundle(
	//	dw.ethclient.Account(),
	//	bundle.Collection,
	//	userMerkleRoot,
	//	bundleDataHash,
	//	bundle.Uri)
	//if err != nil {
	//	return bundleId, errors.Wrap(err, "failed to register a bundle to DataRegistry")
	//}
	//
	//receipt, err := dw.ethclient.WaitMined(context.Background(), tx)
	//if err != nil {
	//	return bundleId, errors.Wrap(err, "failed to wait for tx to be mined")
	//}
	//
	//registerResult, err := dw.dataRegistry.ParseBundleRegisteredFromReceipt(receipt)
	//if err != nil {
	//	return bundleId, errors.Wrap(err, "failed to parse a event from the receipt")
	//}
	//
	//bundleId = types.ID(registerResult.BundleId)
	return
}

// TODO
func (dw *Manager) Get(id types.DataId) (*Bundle, error) {
	//bundle, err := dw.dataRegistry.Bundles(nil, id.BundleId)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to get uri")
	//}
	//
	//if bundle.Uri == "" {
	//	return nil, errors.Wrap(err, "failed to get bundle info")
	//}
	//
	//uri, err := url.Parse(bundle.Uri)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to get bundle data")
	//}
	//return dw.Fetch(uri)
	return nil, nil
}

func (dw *Manager) Fetch(uri *url.URL) (bundle *Bundle, _ error) {
	protoc, exists := dw.protocols[uri.Scheme]
	if !exists {
		return nil, errors.Errorf("the protocol %s is not supported", uri.Scheme)
	}

	dw.log.Info("{}", uri.String())
	bundleData, err := protoc.Read(uri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read bundle data")
	}

	bundle = new(Bundle)
	if err := bundle.UnmarshalJSON(bundleData); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal bundle")
	}
	return
}

type rawBundleData struct {
	BundleId   string  `mapstructure:"bundleId"`
	Uri        string  `mapstructure:"uri"`
	Provider   string  `mapstructure:"provider"`
	Collection string  `mapstructure:"collection"`
	DataCount  int     `mapstructure:"dataCount"`
	IngestedAt float64 `mapstructure:"ingestedAt"`
}

func (dw *Manager) List(ctx context.Context, providerId types.ID) ([]*Bundle, error) {
	bundleDataList, err := dw.resourceDB.Query(ctx, afclient.M{"provider": providerId.Hex()})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list bundles")
	}

	var bundles []*Bundle
	for _, bundleData := range bundleDataList {
		rawBundle := new(rawBundleData)
		if err := mapstructure.Decode(bundleData.Data, &rawBundle); err != nil {
			return nil, errors.Wrap(err, "failed to decode bundle")
		}

		providerId, err := types.HexToID(rawBundle.Provider)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert provider id")
		}

		collectionId, err := types.HexToID(rawBundle.Collection)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert collection id")
		}

		rawTime, _ := new(big.Float).SetFloat64(rawBundle.IngestedAt).Int64()
		bundles = append(bundles, &Bundle{
			Id:         rawBundle.BundleId,
			Uri:        rawBundle.Uri,
			Provider:   providerId,
			Collection: collectionId,
			DataCount:  rawBundle.DataCount,
			IngestedAt: types.ParseTimestamp(rawTime),
		})
	}
	return bundles, nil
}
