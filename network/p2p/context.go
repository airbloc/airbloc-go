package p2p

import (
	"context"
)

type Context struct {
	context.Context
	Cancel context.CancelFunc
}

func newContext(context context.Context, cancel context.CancelFunc) Context {
	return Context{
		Context: context,
		Cancel:  cancel,
	}
}
