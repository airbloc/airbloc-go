package p2p

import (
	"context"
	"sync"

	"github.com/perlin-network/noise"
	"github.com/pkg/errors"
)

type contextGroup struct {
	context.Context
	Cancel context.CancelFunc
}

func newContextGroup(context context.Context, cancel context.CancelFunc) contextGroup {
	return contextGroup{
		Context: context,
		Cancel:  cancel,
	}
}

type contextMap struct{ sync.Map }

// Cancel cancels context with given key
func (cm contextMap) Cancel(key *noise.Peer) error {
	peerContext, ok := cm.Load(key)
	if !ok {
		return errors.New("failed to get peer's context")
	}
	peerContext.Cancel()
	return nil
}

// Delete is contextGroup version of wrapper of sync.Map.Delete
func (cm contextMap) Delete(key *noise.Peer) {
	cm.Map.Delete(key)
}

// Load is contextGroup version of wrapper of sync.Map.Load
func (cm contextMap) Load(key *noise.Peer) (contextGroup, bool) {
	v, ok := cm.Map.Load(key)
	return v.(contextGroup), ok
}

// LoadOrStore is contextGroup version of wrapper of sync.Map.LoadOrStore
func (cm contextMap) LoadOrStore(key *noise.Peer, value contextGroup) (contextGroup, bool) {
	v, loaded := cm.Map.LoadOrStore(key, value)
	return v.(contextGroup), loaded
}

// Range is contextGroup version of wrapper of sync.Map.Range
func (cm contextMap) Range(f func(key *noise.Peer, value contextGroup) bool) {
	cm.Map.Range(func(key, value interface{}) bool {
		return f(key.(*noise.Peer), value.(contextGroup))
	})
}

// Store is contextGroup version of wrapper of sync.Map.Store
func (cm contextMap) Store(key *noise.Peer, value contextGroup) {
	cm.Map.Store(key, value)
}
