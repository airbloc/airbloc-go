package metadb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mode int

const (
	DBName = "airbloc-core"
)

type Database interface {
	Insert(ctx context.Context, docs []interface{}, opt *options.InsertManyOptions) error
	Find(ctx context.Context, query interface{}, opt *options.FindOptions) ([]bson.M, error)
	Aggregate(ctx context.Context, pipeline mongo.Pipeline, opt *options.AggregateOptions) ([]bson.M, error)
	Update(ctx context.Context, filter, update interface{}, opt *options.UpdateOptions) error
	Delete(ctx context.Context, filter interface{}, opt *options.DeleteOptions) error
	Close() error
}
