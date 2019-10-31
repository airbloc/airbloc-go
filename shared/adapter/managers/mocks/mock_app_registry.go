// Code generated by MockGen. DO NOT EDIT.
// Source: app_registry.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	bind "github.com/airbloc/airbloc-go/shared/adapter"
	contracts "github.com/airbloc/airbloc-go/shared/adapter/contracts"
	types "github.com/airbloc/airbloc-go/shared/adapter/types"
	gomock "github.com/golang/mock/gomock"
	bind0 "github.com/klaytn/klaytn/accounts/abi/bind"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// MockIAppRegistryManager is a mock of IAppRegistryManager interface
type MockIAppRegistryManager struct {
	ctrl     *gomock.Controller
	recorder *MockIAppRegistryManagerMockRecorder
}

// MockIAppRegistryManagerMockRecorder is the mock recorder for MockIAppRegistryManager
type MockIAppRegistryManagerMockRecorder struct {
	mock *MockIAppRegistryManager
}

// NewMockIAppRegistryManager creates a new mock instance
func NewMockIAppRegistryManager(ctrl *gomock.Controller) *MockIAppRegistryManager {
	mock := &MockIAppRegistryManager{ctrl: ctrl}
	mock.recorder = &MockIAppRegistryManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAppRegistryManager) EXPECT() *MockIAppRegistryManagerMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockIAppRegistryManager) Address() common.Address {
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockIAppRegistryManagerMockRecorder) Address() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockIAppRegistryManager)(nil).Address))
}

// TxHash mocks base method
func (m *MockIAppRegistryManager) TxHash() common.Hash {
	ret := m.ctrl.Call(m, "TxHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// TxHash indicates an expected call of TxHash
func (mr *MockIAppRegistryManagerMockRecorder) TxHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxHash", reflect.TypeOf((*MockIAppRegistryManager)(nil).TxHash))
}

// CreatedAt mocks base method
func (m *MockIAppRegistryManager) CreatedAt() *big.Int {
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CreatedAt indicates an expected call of CreatedAt
func (mr *MockIAppRegistryManagerMockRecorder) CreatedAt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockIAppRegistryManager)(nil).CreatedAt))
}

// Exists mocks base method
func (m *MockIAppRegistryManager) Exists(appName string) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", appName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockIAppRegistryManagerMockRecorder) Exists(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIAppRegistryManager)(nil).Exists), appName)
}

// Get mocks base method
func (m *MockIAppRegistryManager) Get(appName string) (types.App, error) {
	ret := m.ctrl.Call(m, "Get", appName)
	ret0, _ := ret[0].(types.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIAppRegistryManagerMockRecorder) Get(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAppRegistryManager)(nil).Get), appName)
}

// IsOwner mocks base method
func (m *MockIAppRegistryManager) IsOwner(appName string, owner common.Address) (bool, error) {
	ret := m.ctrl.Call(m, "IsOwner", appName, owner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockIAppRegistryManagerMockRecorder) IsOwner(appName, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockIAppRegistryManager)(nil).IsOwner), appName, owner)
}

// Register mocks base method
func (m *MockIAppRegistryManager) Register(ctx context.Context, opts *bind.TransactOpts, appName string) error {
	ret := m.ctrl.Call(m, "Register", ctx, opts, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockIAppRegistryManagerMockRecorder) Register(ctx, opts, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAppRegistryManager)(nil).Register), ctx, opts, appName)
}

// TransferAppOwner mocks base method
func (m *MockIAppRegistryManager) TransferAppOwner(ctx context.Context, opts *bind.TransactOpts, appName string, newOwner common.Address) error {
	ret := m.ctrl.Call(m, "TransferAppOwner", ctx, opts, appName, newOwner)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferAppOwner indicates an expected call of TransferAppOwner
func (mr *MockIAppRegistryManagerMockRecorder) TransferAppOwner(ctx, opts, appName, newOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferAppOwner", reflect.TypeOf((*MockIAppRegistryManager)(nil).TransferAppOwner), ctx, opts, appName, newOwner)
}

// Unregister mocks base method
func (m *MockIAppRegistryManager) Unregister(ctx context.Context, opts *bind.TransactOpts, appName string) error {
	ret := m.ctrl.Call(m, "Unregister", ctx, opts, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister
func (mr *MockIAppRegistryManagerMockRecorder) Unregister(ctx, opts, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockIAppRegistryManager)(nil).Unregister), ctx, opts, appName)
}

// FilterAppOwnerTransferred mocks base method
func (m *MockIAppRegistryManager) FilterAppOwnerTransferred(opts *bind0.FilterOpts, appAddr, oldOwner []common.Address) (bind.EventIterator, error) {
	ret := m.ctrl.Call(m, "FilterAppOwnerTransferred", opts, appAddr, oldOwner)
	ret0, _ := ret[0].(bind.EventIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterAppOwnerTransferred indicates an expected call of FilterAppOwnerTransferred
func (mr *MockIAppRegistryManagerMockRecorder) FilterAppOwnerTransferred(opts, appAddr, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryManager)(nil).FilterAppOwnerTransferred), opts, appAddr, oldOwner)
}

// FilterRegistration mocks base method
func (m *MockIAppRegistryManager) FilterRegistration(opts *bind0.FilterOpts, appAddr []common.Address) (bind.EventIterator, error) {
	ret := m.ctrl.Call(m, "FilterRegistration", opts, appAddr)
	ret0, _ := ret[0].(bind.EventIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterRegistration indicates an expected call of FilterRegistration
func (mr *MockIAppRegistryManagerMockRecorder) FilterRegistration(opts, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterRegistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).FilterRegistration), opts, appAddr)
}

// FilterUnregistration mocks base method
func (m *MockIAppRegistryManager) FilterUnregistration(opts *bind0.FilterOpts, appAddr []common.Address) (bind.EventIterator, error) {
	ret := m.ctrl.Call(m, "FilterUnregistration", opts, appAddr)
	ret0, _ := ret[0].(bind.EventIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterUnregistration indicates an expected call of FilterUnregistration
func (mr *MockIAppRegistryManagerMockRecorder) FilterUnregistration(opts, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterUnregistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).FilterUnregistration), opts, appAddr)
}

// WatchAppOwnerTransferred mocks base method
func (m *MockIAppRegistryManager) WatchAppOwnerTransferred(opts *bind0.WatchOpts, sink chan<- *contracts.AppRegistryAppOwnerTransferred, appAddr, oldOwner []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchAppOwnerTransferred", opts, sink, appAddr, oldOwner)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAppOwnerTransferred indicates an expected call of WatchAppOwnerTransferred
func (mr *MockIAppRegistryManagerMockRecorder) WatchAppOwnerTransferred(opts, sink, appAddr, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryManager)(nil).WatchAppOwnerTransferred), opts, sink, appAddr, oldOwner)
}

// WatchRegistration mocks base method
func (m *MockIAppRegistryManager) WatchRegistration(opts *bind0.WatchOpts, sink chan<- *contracts.AppRegistryRegistration, appAddr []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchRegistration", opts, sink, appAddr)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRegistration indicates an expected call of WatchRegistration
func (mr *MockIAppRegistryManagerMockRecorder) WatchRegistration(opts, sink, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRegistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).WatchRegistration), opts, sink, appAddr)
}

// WatchUnregistration mocks base method
func (m *MockIAppRegistryManager) WatchUnregistration(opts *bind0.WatchOpts, sink chan<- *contracts.AppRegistryUnregistration, appAddr []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchUnregistration", opts, sink, appAddr)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUnregistration indicates an expected call of WatchUnregistration
func (mr *MockIAppRegistryManagerMockRecorder) WatchUnregistration(opts, sink, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUnregistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).WatchUnregistration), opts, sink, appAddr)
}
