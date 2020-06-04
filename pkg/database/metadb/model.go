package metadb

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/airbloc/logger"
)

type Model struct {
	database Database
	Name     string
	log      *logger.Logger
}

// NewModel creates a collection for given metadatabase
// with the name of the data type.
func NewModel(database Database, name string) Database {
	loggerName := fmt.Sprintf("metadb (%s)", strings.ToLower(name))
	return &Model{
		Name:     name,
		database: database,
		log:      logger.New(loggerName),
	}
}

func checkType(i interface{}) (bson.M, error) {
	switch v := i.(type) {
	case bson.M:
		return v, nil
	case map[string]interface{}:
		return primitive.M(v), nil
	default:
		return nil, errors.Errorf(
			"unwrap failed : invalid type (expected: %s, actual: %s)",
			"bson.M", reflect.TypeOf(i).String())
	}
}

func (m *Model) unwrap(rawDoc interface{}) (interface{}, error) {
	doc, err := checkType(rawDoc)
	if err != nil {
		return nil, err
	}

	typ, ok := doc["type"]
	if !ok {
		return nil, errors.Errorf("unwrap failed : type does not exists")
	}

	if typ != m.Name {
		return nil, errors.Errorf(
			"unwrap failed : invalid type (expected: %s, actual: %s)",
			m.Name, typ)
	}

	data, ok := doc["data"]
	if !ok {
		return nil, errors.Errorf("unwrap failed : data does not exists")
	}
	return data, nil
}

func (m *Model) wrapDoc(rawDoc interface{}) (interface{}, error) {
	doc, err := checkType(rawDoc)
	if err != nil {
		return nil, err
	}

	wrappedDoc := bson.M{
		"type": m.Name,
		"data": doc,
	}
	return wrappedDoc, nil
}

func (m *Model) wrapQuery(rawQuery interface{}) (interface{}, error) {
	query, err := checkType(rawQuery)
	if err != nil {
		return nil, err
	}

	wrappedQuery := bson.M{"type": m.Name}
	for name, payload := range query {
		wrappedQuery["data."+name] = payload
	}
	return wrappedQuery, nil
}

func (m *Model) Insert(
	ctx context.Context,
	docs []interface{},
	opt *options.InsertManyOptions,
) (err error) {
	for i, doc := range docs {
		docs[i], err = m.wrapDoc(doc)
		if err != nil {
			return
		}
	}
	return m.database.Insert(ctx, docs, opt)
}

func (m *Model) Find(
	ctx context.Context,
	query interface{},
	opt *options.FindOptions,
) ([]bson.M, error) {
	q, err := m.wrapQuery(query)
	if err != nil {
		return nil, err
	}

	wrappedDocs, err := m.database.Find(ctx, q, opt)
	if err != nil {
		return nil, err
	}

	var docs = make([]bson.M, len(wrappedDocs))
	for i, wrappedDoc := range wrappedDocs {
		rawDoc, err := m.unwrap(wrappedDoc)
		if err != nil {
			return nil, err
		}

		doc, ok := rawDoc.(primitive.D)
		if !ok {
			return nil, errors.Errorf(
				"find failed : invalid type of data (expected: %s, actual: %s)",
				"primitive.D", reflect.TypeOf(rawDoc).String())
		}
		docs[i] = doc.Map()
	}
	return docs, nil
}

func (m *Model) Aggregate(
	ctx context.Context,
	pipeline mongo.Pipeline,
	opt *options.AggregateOptions,
) ([]bson.M, error) {
	return m.database.Aggregate(ctx, pipeline, opt)
}

func (m *Model) Update(
	ctx context.Context,
	filter, update interface{},
	opt *options.UpdateOptions,
) error {
	q, err := m.wrapQuery(filter)
	if err != nil {
		return err
	}
	return m.database.Update(ctx, q, update, opt)
}

func (m *Model) Delete(
	ctx context.Context,
	filter interface{},
	opt *options.DeleteOptions,
) error {
	q, err := m.wrapQuery(filter)
	if err != nil {
		return err
	}
	return m.database.Delete(ctx, q, opt)
}

func (m *Model) Close() error {
	return m.database.Close()
}
