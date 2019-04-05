package resdb

import (
	"context"
	"github.com/airbloc/airframe/afclient"
)

type Model struct {
	typ    string
	client afclient.Client
}

func NewModel(client afclient.Client, typ string) Model {
	return Model{
		typ:    typ,
		client: client,
	}
}

func (m Model) Get(
	ctx context.Context,
	id string,
) (*afclient.Object, error) {
	return m.client.Get(ctx, m.typ, id)
}

func (m Model) Put(
	ctx context.Context,
	id string,
	data afclient.M,
) (*afclient.PutResult, error) {
	return m.client.Put(ctx, m.typ, id, data)
}

func (m Model) Query(
	ctx context.Context,
	query afclient.M,
	opts ...afclient.QueryOption,
) ([]*afclient.Object, error) {
	return m.client.Query(ctx, m.typ, query, opts...)
}
