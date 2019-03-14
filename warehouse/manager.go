package warehouse

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/airbloc/airbloc-go/shared/service"
	"math/rand"
	"net/url"
	"time"

	"github.com/airbloc/airbloc-go/provider/collections"
	"github.com/airbloc/airbloc-go/provider/schemas"
	"github.com/airbloc/airbloc-go/shared/adapter"
	"github.com/airbloc/airbloc-go/shared/blockchain"
	"github.com/airbloc/airbloc-go/shared/database/localdb"
	"github.com/airbloc/airbloc-go/shared/database/metadb"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/airbloc-go/shared/types"
	"github.com/airbloc/airbloc-go/warehouse/dauth"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/airbloc/logger"
	"github.com/mongodb/mongo-go-driver/bson"
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
	metaDatabase *metadb.Model
	ethclient    blockchain.TxClient
	dataRegistry *adapter.DataRegistry
	schemas      *schemas.Schemas
	collections  *collections.Manager

	// data storage layer
	protocols      map[string]protocol.Protocol
	DefaultStorage storage.Storage

	// data validators
	dauthValidator *dauth.Validator

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
) *Manager {
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

	dauthManager := dauth.NewManager(ethclient)
	dauthValidator := dauth.NewValidator(dauthManager)

	contract := ethclient.GetContract(&adapter.DataRegistry{})
	return &Manager{
		kms:        kms,
		localCache: localdb.NewModel(localDatabase, "bundle"),

		metaDatabase: metadb.NewModel(metaDatabase, "bundles"),
		ethclient:    ethclient,
		dataRegistry: contract.(*adapter.DataRegistry),
		collections:  collections.NewManager(ethclient),
		schemas:      schemas.New(metaDatabase, ethclient),

		protocols:      protocols,
		DefaultStorage: defaultStorage,
		dauthValidator: dauthValidator,

		config: config,
		log:    log,
	}
}

func (dw *Manager) CreateBundle(collectionId types.ID) (*BundleStream, error) {
	collection, err := dw.collections.Get(collectionId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve a collection")
	}
	schema, err := dw.schemas.Get(collection.Schema.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve a schema")
	}
	collection.Schema = *schema
	return newBundleStream(dw, collection.AppId, collection), nil
}

func (dw *Manager) validate(collection *collections.Collection, data *types.Data) error {
	if !dw.config.Warehouse.Debug.DisableUserAuthValidation && !dw.dauthValidator.IsCollectible(collection.Id, data) {
		return errors.Wrap(errValidationFailed, "user hasn't been authorized the data collection")
	}

	if !dw.config.Warehouse.Debug.DisableSchemaValidation {
		isValidFormat, err := collection.Schema.IsValidFormat(data)
		if err != nil {
			return err
		} else if !isValidFormat {
			return errors.Wrap(errValidationFailed, "wrong format")
		}
	}
	return nil
}

func generateBundleNameOf(bundle *Bundle) string {
	tokenBytes := make([]byte, 4)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	currentTime := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s.bundle", currentTime, bundle.Collection.Hex(), token)
}

func (dw *Manager) Store(stream *BundleStream) (*Bundle, error) {
	if stream == nil {
		return nil, errors.New("No data in the stream.")
	}
	ingestedAt := types.Time{Time: time.Now()}

	createdBundle := &Bundle{
		Provider:   stream.provider,
		Collection: stream.collection.Id,
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

	uri, err := dw.DefaultStorage.Save(bundleName, bundleData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save bundle to the storage")
	}
	createdBundle.Uri = uri.String()

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

	_, err = dw.metaDatabase.Create(bundleInfo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save metadata")
	}

	dw.log.Info("Bundle %s registered on", bundleName, logger.Attrs{
		"id":    bundleId.Hex(),
		"count": bundleInfo["dataCount"],
	})
	return createdBundle, nil
}

func (dw *Manager) registerBundleOnChain(bundle *Bundle) (bundleId types.ID, _ error) {
	bundleDataHash, err := bundle.Hash()
	if err != nil {
		return bundleId, errors.Wrap(err, "failed to get hash of the bundle data")
	}

	// for setup rowId
	userMerkleRoot, err := bundle.SetupUserProof()
	if err != nil {
		return bundleId, errors.Wrap(err, "failed to setup SMT")
	}

	dw.log.Info("Bundle data", logger.Attrs{
		"hash":       bundleDataHash.Hex(),
		"merkleRoot": userMerkleRoot.Hex(),
	})

	tx, err := dw.dataRegistry.RegisterBundle(
		dw.ethclient.Account(),
		bundle.Collection,
		userMerkleRoot,
		bundleDataHash,
		bundle.Uri)
	if err != nil {
		return bundleId, errors.Wrap(err, "failed to register a bundle to DataRegistry")
	}

	receipt, err := dw.ethclient.WaitMined(context.Background(), tx)
	if err != nil {
		return bundleId, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	registerResult, err := dw.dataRegistry.ParseBundleRegisteredFromReceipt(receipt)
	if err != nil {
		return bundleId, errors.Wrap(err, "failed to parse a event from the receipt")
	}

	bundleId = types.ID(registerResult.BundleId)
	return
}

func (dw *Manager) Get(id *types.DataId) (*Bundle, error) {
	bundle, err := dw.dataRegistry.Bundles(nil, id.BundleId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get uri")
	}

	if bundle.Uri == "" {
		return nil, errors.Wrap(err, "failed to get bundle info")
	}

	uri, err := url.Parse(bundle.Uri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bundle data")
	}
	return dw.Fetch(uri)
}

func (dw *Manager) Fetch(uri *url.URL) (bundle *Bundle, _ error) {
	protoc, exists := dw.protocols[uri.Scheme]
	if !exists {
		return nil, errors.Errorf("the protocol %s is not supported", uri.Scheme)
	}

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

func (dw *Manager) List(providerId types.ID) ([]*Bundle, error) {
	bundleDataList, err := dw.metaDatabase.RetrieveMany(context.TODO(), bson.M{"provider": providerId.Hex()})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list bundles")
	}

	var bundles []*Bundle
	for _, bundleData := range bundleDataList {
		collectionId, _ := types.HexToID(bundleData["collection"].(string))
		ingestedAt, _ := time.Parse(time.RFC3339Nano, bundleData["ingestedAt"].(string))
		bundles = append(bundles, &Bundle{
			Id:         bundleData["bundleId"].(string),
			Uri:        bundleData["uri"].(string),
			DataCount:  bundleData["dataCount"].(int),
			IngestedAt: types.Time{Time: ingestedAt},
			Provider:   providerId,
			Collection: collectionId,
		})
	}
	return bundles, nil
}
