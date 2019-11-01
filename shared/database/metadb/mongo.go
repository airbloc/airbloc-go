package metadb

import (
	"context"

	"github.com/airbloc/logger"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	db  *mongo.Database
	col string
	log *logger.Logger
}

func NewMongoDB(ctx context.Context, url, col string) (_ Database, err error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return
	}

	db := &Mongo{
		db:  client.Database(DBName),
		col: col,
		log: logger.New("mongodb"),
	}

	return db, err
}

func (mg *Mongo) Insert(
	ctx context.Context,
	docs []interface{},
	opt *options.InsertManyOptions,
) error {
	db := mg.db.Collection(mg.col)

	if docs == nil {
		return errors.New("no document")
	}

	if len(docs) == 1 {
		oneOpt := &options.InsertOneOptions{}
		if opt != nil {
			oneOpt.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		res, err := db.InsertOne(ctx, docs[0], oneOpt)
		if err != nil {
			return errors.Wrap(err, "insert failed")
		}
		mg.log.Info("Data inserted ID:{}", res.InsertedID)
	} else {
		res, err := db.InsertMany(ctx, docs, opt)
		if err != nil {
			return errors.Wrap(err, "insert failed")
		}
		mg.log.Info("Data inserted ID:{}...{}", res.InsertedIDs[:3], len(res.InsertedIDs))
	}
	return nil
}

func (mg *Mongo) Find(
	ctx context.Context,
	query interface{},
	opt *options.FindOptions,
) (docs []bson.M, _ error) {
	db := mg.db.Collection(mg.col)
	cursor, err := db.Find(ctx, query, opt)
	if err != nil {
		return nil, errors.Wrap(err, "find failed")
	}

	for cursor.Next(ctx) {
		elem := new(bson.D)
		if err := cursor.Decode(elem); err != nil {
			return nil, errors.Wrap(err, "decode failed")
		}
		docs = append(docs, elem.Map())
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor failed")
	}

	mg.log.Info("Found {} data", len(docs))
	return
}

func (mg *Mongo) Aggregate(
	ctx context.Context,
	pipeline mongo.Pipeline,
	opt *options.AggregateOptions,
) (docs []bson.M, _ error) {
	db := mg.db.Collection(mg.col)
	cursor, err := db.Aggregate(ctx, pipeline, opt)
	if err != nil {
		return nil, errors.Wrap(err, "aggregate failed")
	}

	for cursor.Next(ctx) {
		elem := new(bson.D)
		if err := cursor.Decode(elem); err != nil {
			return nil, errors.Wrap(err, "decode failed")
		}
		docs = append(docs, elem.Map())
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor failed")
	}

	mg.log.Info("Aggregated {} data", len(docs))
	return
}

func (mg *Mongo) Update(
	ctx context.Context,
	filter, update interface{},
	opt *options.UpdateOptions,
) error {
	db := mg.db.Collection(mg.col)
	res, err := db.UpdateMany(ctx, filter, update, opt)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}
	mg.log.Info(
		"find {} data and update {} data",
		res.MatchedCount, res.ModifiedCount)
	return nil
}

func (mg *Mongo) Delete(
	ctx context.Context,
	filter interface{},
	opt *options.DeleteOptions,
) error {
	db := mg.db.Collection(mg.col)
	res, err := db.DeleteMany(ctx, filter, opt)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	mg.log.Info("delete {} data", res.DeletedCount)
	return nil
}

func (mg *Mongo) Close() error {
	return mg.db.Client().Disconnect(context.Background())
}
