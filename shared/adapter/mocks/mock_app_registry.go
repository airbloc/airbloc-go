// Code generated by MockGen. DO NOT EDIT.
// Source: app_registry.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	adapter "github.com/airbloc/airbloc-go/shared/adapter"
	bind "github.com/airbloc/airbloc-go/shared/blockchain/bind"
	types "github.com/airbloc/airbloc-go/shared/types"
	common "github.com/ethereum/go-ethereum/common"
	types0 "github.com/ethereum/go-ethereum/core/types"
	event "github.com/ethereum/go-ethereum/event"
	gomock "github.com/golang/mock/gomock"
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

// Account mocks base method
func (m *MockIAppRegistryManager) Account() common.Address {
	ret := m.ctrl.Call(m, "Account")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Account indicates an expected call of Account
func (mr *MockIAppRegistryManagerMockRecorder) Account() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Account", reflect.TypeOf((*MockIAppRegistryManager)(nil).Account))
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
func (m *MockIAppRegistryManager) Register(ctx context.Context, appName string) error {
	ret := m.ctrl.Call(m, "Register", ctx, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockIAppRegistryManagerMockRecorder) Register(ctx, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAppRegistryManager)(nil).Register), ctx, appName)
}

// TransferAppOwner mocks base method
func (m *MockIAppRegistryManager) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) error {
	ret := m.ctrl.Call(m, "TransferAppOwner", ctx, appName, newOwner)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferAppOwner indicates an expected call of TransferAppOwner
func (mr *MockIAppRegistryManagerMockRecorder) TransferAppOwner(ctx, appName, newOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferAppOwner", reflect.TypeOf((*MockIAppRegistryManager)(nil).TransferAppOwner), ctx, appName, newOwner)
}

// Unregister mocks base method
func (m *MockIAppRegistryManager) Unregister(ctx context.Context, appName string) error {
	ret := m.ctrl.Call(m, "Unregister", ctx, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister
func (mr *MockIAppRegistryManagerMockRecorder) Unregister(ctx, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockIAppRegistryManager)(nil).Unregister), ctx, appName)
}

// FilterAppOwnerTransferred mocks base method
func (m *MockIAppRegistryManager) FilterAppOwnerTransferred(opts *bind.FilterOpts, hashedAppName []common.Hash, oldOwner []common.Address) (*adapter.AppRegistryAppOwnerTransferredIterator, error) {
	ret := m.ctrl.Call(m, "FilterAppOwnerTransferred", opts, hashedAppName, oldOwner)
	ret0, _ := ret[0].(*adapter.AppRegistryAppOwnerTransferredIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterAppOwnerTransferred indicates an expected call of FilterAppOwnerTransferred
func (mr *MockIAppRegistryManagerMockRecorder) FilterAppOwnerTransferred(opts, hashedAppName, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryManager)(nil).FilterAppOwnerTransferred), opts, hashedAppName, oldOwner)
}

// WatchAppOwnerTransferred mocks base method
func (m *MockIAppRegistryManager) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryAppOwnerTransferred, hashedAppName []common.Hash, oldOwner []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchAppOwnerTransferred", opts, sink, hashedAppName, oldOwner)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAppOwnerTransferred indicates an expected call of WatchAppOwnerTransferred
func (mr *MockIAppRegistryManagerMockRecorder) WatchAppOwnerTransferred(opts, sink, hashedAppName, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryManager)(nil).WatchAppOwnerTransferred), opts, sink, hashedAppName, oldOwner)
}

// FilterRegistration mocks base method
func (m *MockIAppRegistryManager) FilterRegistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryRegistrationIterator, error) {
	ret := m.ctrl.Call(m, "FilterRegistration", opts, hashedAppName)
	ret0, _ := ret[0].(*adapter.AppRegistryRegistrationIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterRegistration indicates an expected call of FilterRegistration
func (mr *MockIAppRegistryManagerMockRecorder) FilterRegistration(opts, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterRegistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).FilterRegistration), opts, hashedAppName)
}

// WatchRegistration mocks base method
func (m *MockIAppRegistryManager) WatchRegistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryRegistration, hashedAppName []common.Hash) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchRegistration", opts, sink, hashedAppName)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRegistration indicates an expected call of WatchRegistration
func (mr *MockIAppRegistryManagerMockRecorder) WatchRegistration(opts, sink, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRegistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).WatchRegistration), opts, sink, hashedAppName)
}

// FilterUnregistration mocks base method
func (m *MockIAppRegistryManager) FilterUnregistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryUnregistrationIterator, error) {
	ret := m.ctrl.Call(m, "FilterUnregistration", opts, hashedAppName)
	ret0, _ := ret[0].(*adapter.AppRegistryUnregistrationIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterUnregistration indicates an expected call of FilterUnregistration
func (mr *MockIAppRegistryManagerMockRecorder) FilterUnregistration(opts, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterUnregistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).FilterUnregistration), opts, hashedAppName)
}

// WatchUnregistration mocks base method
func (m *MockIAppRegistryManager) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryUnregistration, hashedAppName []common.Hash) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchUnregistration", opts, sink, hashedAppName)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUnregistration indicates an expected call of WatchUnregistration
func (mr *MockIAppRegistryManagerMockRecorder) WatchUnregistration(opts, sink, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUnregistration", reflect.TypeOf((*MockIAppRegistryManager)(nil).WatchUnregistration), opts, sink, hashedAppName)
}

// MockIAppRegistryContract is a mock of IAppRegistryContract interface
type MockIAppRegistryContract struct {
	ctrl     *gomock.Controller
	recorder *MockIAppRegistryContractMockRecorder
}

// MockIAppRegistryContractMockRecorder is the mock recorder for MockIAppRegistryContract
type MockIAppRegistryContractMockRecorder struct {
	mock *MockIAppRegistryContract
}

// NewMockIAppRegistryContract creates a new mock instance
func NewMockIAppRegistryContract(ctrl *gomock.Controller) *MockIAppRegistryContract {
	mock := &MockIAppRegistryContract{ctrl: ctrl}
	mock.recorder = &MockIAppRegistryContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAppRegistryContract) EXPECT() *MockIAppRegistryContractMockRecorder {
	return m.recorder
}

// Account mocks base method
func (m *MockIAppRegistryContract) Account() common.Address {
	ret := m.ctrl.Call(m, "Account")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Account indicates an expected call of Account
func (mr *MockIAppRegistryContractMockRecorder) Account() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Account", reflect.TypeOf((*MockIAppRegistryContract)(nil).Account))
}

// TxHash mocks base method
func (m *MockIAppRegistryContract) TxHash() common.Hash {
	ret := m.ctrl.Call(m, "TxHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// TxHash indicates an expected call of TxHash
func (mr *MockIAppRegistryContractMockRecorder) TxHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxHash", reflect.TypeOf((*MockIAppRegistryContract)(nil).TxHash))
}

// CreatedAt mocks base method
func (m *MockIAppRegistryContract) CreatedAt() *big.Int {
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CreatedAt indicates an expected call of CreatedAt
func (mr *MockIAppRegistryContractMockRecorder) CreatedAt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockIAppRegistryContract)(nil).CreatedAt))
}

// Exists mocks base method
func (m *MockIAppRegistryContract) Exists(appName string) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", appName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockIAppRegistryContractMockRecorder) Exists(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIAppRegistryContract)(nil).Exists), appName)
}

// Get mocks base method
func (m *MockIAppRegistryContract) Get(appName string) (types.App, error) {
	ret := m.ctrl.Call(m, "Get", appName)
	ret0, _ := ret[0].(types.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIAppRegistryContractMockRecorder) Get(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAppRegistryContract)(nil).Get), appName)
}

// IsOwner mocks base method
func (m *MockIAppRegistryContract) IsOwner(appName string, owner common.Address) (bool, error) {
	ret := m.ctrl.Call(m, "IsOwner", appName, owner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockIAppRegistryContractMockRecorder) IsOwner(appName, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockIAppRegistryContract)(nil).IsOwner), appName, owner)
}

// Register mocks base method
func (m *MockIAppRegistryContract) Register(ctx context.Context, appName string) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Register", ctx, appName)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockIAppRegistryContractMockRecorder) Register(ctx, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAppRegistryContract)(nil).Register), ctx, appName)
}

// TransferAppOwner mocks base method
func (m *MockIAppRegistryContract) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "TransferAppOwner", ctx, appName, newOwner)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferAppOwner indicates an expected call of TransferAppOwner
func (mr *MockIAppRegistryContractMockRecorder) TransferAppOwner(ctx, appName, newOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferAppOwner", reflect.TypeOf((*MockIAppRegistryContract)(nil).TransferAppOwner), ctx, appName, newOwner)
}

// Unregister mocks base method
func (m *MockIAppRegistryContract) Unregister(ctx context.Context, appName string) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Unregister", ctx, appName)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister
func (mr *MockIAppRegistryContractMockRecorder) Unregister(ctx, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockIAppRegistryContract)(nil).Unregister), ctx, appName)
}

// FilterAppOwnerTransferred mocks base method
func (m *MockIAppRegistryContract) FilterAppOwnerTransferred(opts *bind.FilterOpts, hashedAppName []common.Hash, oldOwner []common.Address) (*adapter.AppRegistryAppOwnerTransferredIterator, error) {
	ret := m.ctrl.Call(m, "FilterAppOwnerTransferred", opts, hashedAppName, oldOwner)
	ret0, _ := ret[0].(*adapter.AppRegistryAppOwnerTransferredIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterAppOwnerTransferred indicates an expected call of FilterAppOwnerTransferred
func (mr *MockIAppRegistryContractMockRecorder) FilterAppOwnerTransferred(opts, hashedAppName, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryContract)(nil).FilterAppOwnerTransferred), opts, hashedAppName, oldOwner)
}

// ParseAppOwnerTransferredFromReceipt mocks base method
func (m *MockIAppRegistryContract) ParseAppOwnerTransferredFromReceipt(receipt *types0.Receipt) (*adapter.AppRegistryAppOwnerTransferred, error) {
	ret := m.ctrl.Call(m, "ParseAppOwnerTransferredFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.AppRegistryAppOwnerTransferred)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseAppOwnerTransferredFromReceipt indicates an expected call of ParseAppOwnerTransferredFromReceipt
func (mr *MockIAppRegistryContractMockRecorder) ParseAppOwnerTransferredFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseAppOwnerTransferredFromReceipt", reflect.TypeOf((*MockIAppRegistryContract)(nil).ParseAppOwnerTransferredFromReceipt), receipt)
}

// WatchAppOwnerTransferred mocks base method
func (m *MockIAppRegistryContract) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryAppOwnerTransferred, hashedAppName []common.Hash, oldOwner []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchAppOwnerTransferred", opts, sink, hashedAppName, oldOwner)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAppOwnerTransferred indicates an expected call of WatchAppOwnerTransferred
func (mr *MockIAppRegistryContractMockRecorder) WatchAppOwnerTransferred(opts, sink, hashedAppName, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryContract)(nil).WatchAppOwnerTransferred), opts, sink, hashedAppName, oldOwner)
}

// FilterRegistration mocks base method
func (m *MockIAppRegistryContract) FilterRegistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryRegistrationIterator, error) {
	ret := m.ctrl.Call(m, "FilterRegistration", opts, hashedAppName)
	ret0, _ := ret[0].(*adapter.AppRegistryRegistrationIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterRegistration indicates an expected call of FilterRegistration
func (mr *MockIAppRegistryContractMockRecorder) FilterRegistration(opts, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterRegistration", reflect.TypeOf((*MockIAppRegistryContract)(nil).FilterRegistration), opts, hashedAppName)
}

// ParseRegistrationFromReceipt mocks base method
func (m *MockIAppRegistryContract) ParseRegistrationFromReceipt(receipt *types0.Receipt) (*adapter.AppRegistryRegistration, error) {
	ret := m.ctrl.Call(m, "ParseRegistrationFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.AppRegistryRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseRegistrationFromReceipt indicates an expected call of ParseRegistrationFromReceipt
func (mr *MockIAppRegistryContractMockRecorder) ParseRegistrationFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRegistrationFromReceipt", reflect.TypeOf((*MockIAppRegistryContract)(nil).ParseRegistrationFromReceipt), receipt)
}

// WatchRegistration mocks base method
func (m *MockIAppRegistryContract) WatchRegistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryRegistration, hashedAppName []common.Hash) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchRegistration", opts, sink, hashedAppName)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRegistration indicates an expected call of WatchRegistration
func (mr *MockIAppRegistryContractMockRecorder) WatchRegistration(opts, sink, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRegistration", reflect.TypeOf((*MockIAppRegistryContract)(nil).WatchRegistration), opts, sink, hashedAppName)
}

// FilterUnregistration mocks base method
func (m *MockIAppRegistryContract) FilterUnregistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryUnregistrationIterator, error) {
	ret := m.ctrl.Call(m, "FilterUnregistration", opts, hashedAppName)
	ret0, _ := ret[0].(*adapter.AppRegistryUnregistrationIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterUnregistration indicates an expected call of FilterUnregistration
func (mr *MockIAppRegistryContractMockRecorder) FilterUnregistration(opts, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterUnregistration", reflect.TypeOf((*MockIAppRegistryContract)(nil).FilterUnregistration), opts, hashedAppName)
}

// ParseUnregistrationFromReceipt mocks base method
func (m *MockIAppRegistryContract) ParseUnregistrationFromReceipt(receipt *types0.Receipt) (*adapter.AppRegistryUnregistration, error) {
	ret := m.ctrl.Call(m, "ParseUnregistrationFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.AppRegistryUnregistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseUnregistrationFromReceipt indicates an expected call of ParseUnregistrationFromReceipt
func (mr *MockIAppRegistryContractMockRecorder) ParseUnregistrationFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUnregistrationFromReceipt", reflect.TypeOf((*MockIAppRegistryContract)(nil).ParseUnregistrationFromReceipt), receipt)
}

// WatchUnregistration mocks base method
func (m *MockIAppRegistryContract) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryUnregistration, hashedAppName []common.Hash) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchUnregistration", opts, sink, hashedAppName)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUnregistration indicates an expected call of WatchUnregistration
func (mr *MockIAppRegistryContractMockRecorder) WatchUnregistration(opts, sink, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUnregistration", reflect.TypeOf((*MockIAppRegistryContract)(nil).WatchUnregistration), opts, sink, hashedAppName)
}

// MockIAppRegistryCalls is a mock of IAppRegistryCalls interface
type MockIAppRegistryCalls struct {
	ctrl     *gomock.Controller
	recorder *MockIAppRegistryCallsMockRecorder
}

// MockIAppRegistryCallsMockRecorder is the mock recorder for MockIAppRegistryCalls
type MockIAppRegistryCallsMockRecorder struct {
	mock *MockIAppRegistryCalls
}

// NewMockIAppRegistryCalls creates a new mock instance
func NewMockIAppRegistryCalls(ctrl *gomock.Controller) *MockIAppRegistryCalls {
	mock := &MockIAppRegistryCalls{ctrl: ctrl}
	mock.recorder = &MockIAppRegistryCallsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAppRegistryCalls) EXPECT() *MockIAppRegistryCallsMockRecorder {
	return m.recorder
}

// Exists mocks base method
func (m *MockIAppRegistryCalls) Exists(appName string) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", appName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockIAppRegistryCallsMockRecorder) Exists(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIAppRegistryCalls)(nil).Exists), appName)
}

// Get mocks base method
func (m *MockIAppRegistryCalls) Get(appName string) (types.App, error) {
	ret := m.ctrl.Call(m, "Get", appName)
	ret0, _ := ret[0].(types.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIAppRegistryCallsMockRecorder) Get(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAppRegistryCalls)(nil).Get), appName)
}

// IsOwner mocks base method
func (m *MockIAppRegistryCalls) IsOwner(appName string, owner common.Address) (bool, error) {
	ret := m.ctrl.Call(m, "IsOwner", appName, owner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockIAppRegistryCallsMockRecorder) IsOwner(appName, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockIAppRegistryCalls)(nil).IsOwner), appName, owner)
}

// MockIAppRegistryTransacts is a mock of IAppRegistryTransacts interface
type MockIAppRegistryTransacts struct {
	ctrl     *gomock.Controller
	recorder *MockIAppRegistryTransactsMockRecorder
}

// MockIAppRegistryTransactsMockRecorder is the mock recorder for MockIAppRegistryTransacts
type MockIAppRegistryTransactsMockRecorder struct {
	mock *MockIAppRegistryTransacts
}

// NewMockIAppRegistryTransacts creates a new mock instance
func NewMockIAppRegistryTransacts(ctrl *gomock.Controller) *MockIAppRegistryTransacts {
	mock := &MockIAppRegistryTransacts{ctrl: ctrl}
	mock.recorder = &MockIAppRegistryTransactsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAppRegistryTransacts) EXPECT() *MockIAppRegistryTransactsMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockIAppRegistryTransacts) Register(ctx context.Context, appName string) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Register", ctx, appName)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockIAppRegistryTransactsMockRecorder) Register(ctx, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAppRegistryTransacts)(nil).Register), ctx, appName)
}

// TransferAppOwner mocks base method
func (m *MockIAppRegistryTransacts) TransferAppOwner(ctx context.Context, appName string, newOwner common.Address) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "TransferAppOwner", ctx, appName, newOwner)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferAppOwner indicates an expected call of TransferAppOwner
func (mr *MockIAppRegistryTransactsMockRecorder) TransferAppOwner(ctx, appName, newOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferAppOwner", reflect.TypeOf((*MockIAppRegistryTransacts)(nil).TransferAppOwner), ctx, appName, newOwner)
}

// Unregister mocks base method
func (m *MockIAppRegistryTransacts) Unregister(ctx context.Context, appName string) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Unregister", ctx, appName)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister
func (mr *MockIAppRegistryTransactsMockRecorder) Unregister(ctx, appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockIAppRegistryTransacts)(nil).Unregister), ctx, appName)
}

// MockIAppRegistryEvents is a mock of IAppRegistryEvents interface
type MockIAppRegistryEvents struct {
	ctrl     *gomock.Controller
	recorder *MockIAppRegistryEventsMockRecorder
}

// MockIAppRegistryEventsMockRecorder is the mock recorder for MockIAppRegistryEvents
type MockIAppRegistryEventsMockRecorder struct {
	mock *MockIAppRegistryEvents
}

// NewMockIAppRegistryEvents creates a new mock instance
func NewMockIAppRegistryEvents(ctrl *gomock.Controller) *MockIAppRegistryEvents {
	mock := &MockIAppRegistryEvents{ctrl: ctrl}
	mock.recorder = &MockIAppRegistryEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAppRegistryEvents) EXPECT() *MockIAppRegistryEventsMockRecorder {
	return m.recorder
}

// FilterAppOwnerTransferred mocks base method
func (m *MockIAppRegistryEvents) FilterAppOwnerTransferred(opts *bind.FilterOpts, hashedAppName []common.Hash, oldOwner []common.Address) (*adapter.AppRegistryAppOwnerTransferredIterator, error) {
	ret := m.ctrl.Call(m, "FilterAppOwnerTransferred", opts, hashedAppName, oldOwner)
	ret0, _ := ret[0].(*adapter.AppRegistryAppOwnerTransferredIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterAppOwnerTransferred indicates an expected call of FilterAppOwnerTransferred
func (mr *MockIAppRegistryEventsMockRecorder) FilterAppOwnerTransferred(opts, hashedAppName, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryEvents)(nil).FilterAppOwnerTransferred), opts, hashedAppName, oldOwner)
}

// ParseAppOwnerTransferredFromReceipt mocks base method
func (m *MockIAppRegistryEvents) ParseAppOwnerTransferredFromReceipt(receipt *types0.Receipt) (*adapter.AppRegistryAppOwnerTransferred, error) {
	ret := m.ctrl.Call(m, "ParseAppOwnerTransferredFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.AppRegistryAppOwnerTransferred)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseAppOwnerTransferredFromReceipt indicates an expected call of ParseAppOwnerTransferredFromReceipt
func (mr *MockIAppRegistryEventsMockRecorder) ParseAppOwnerTransferredFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseAppOwnerTransferredFromReceipt", reflect.TypeOf((*MockIAppRegistryEvents)(nil).ParseAppOwnerTransferredFromReceipt), receipt)
}

// WatchAppOwnerTransferred mocks base method
func (m *MockIAppRegistryEvents) WatchAppOwnerTransferred(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryAppOwnerTransferred, hashedAppName []common.Hash, oldOwner []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchAppOwnerTransferred", opts, sink, hashedAppName, oldOwner)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAppOwnerTransferred indicates an expected call of WatchAppOwnerTransferred
func (mr *MockIAppRegistryEventsMockRecorder) WatchAppOwnerTransferred(opts, sink, hashedAppName, oldOwner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAppOwnerTransferred", reflect.TypeOf((*MockIAppRegistryEvents)(nil).WatchAppOwnerTransferred), opts, sink, hashedAppName, oldOwner)
}

// FilterRegistration mocks base method
func (m *MockIAppRegistryEvents) FilterRegistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryRegistrationIterator, error) {
	ret := m.ctrl.Call(m, "FilterRegistration", opts, hashedAppName)
	ret0, _ := ret[0].(*adapter.AppRegistryRegistrationIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterRegistration indicates an expected call of FilterRegistration
func (mr *MockIAppRegistryEventsMockRecorder) FilterRegistration(opts, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterRegistration", reflect.TypeOf((*MockIAppRegistryEvents)(nil).FilterRegistration), opts, hashedAppName)
}

// ParseRegistrationFromReceipt mocks base method
func (m *MockIAppRegistryEvents) ParseRegistrationFromReceipt(receipt *types0.Receipt) (*adapter.AppRegistryRegistration, error) {
	ret := m.ctrl.Call(m, "ParseRegistrationFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.AppRegistryRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseRegistrationFromReceipt indicates an expected call of ParseRegistrationFromReceipt
func (mr *MockIAppRegistryEventsMockRecorder) ParseRegistrationFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRegistrationFromReceipt", reflect.TypeOf((*MockIAppRegistryEvents)(nil).ParseRegistrationFromReceipt), receipt)
}

// WatchRegistration mocks base method
func (m *MockIAppRegistryEvents) WatchRegistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryRegistration, hashedAppName []common.Hash) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchRegistration", opts, sink, hashedAppName)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRegistration indicates an expected call of WatchRegistration
func (mr *MockIAppRegistryEventsMockRecorder) WatchRegistration(opts, sink, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRegistration", reflect.TypeOf((*MockIAppRegistryEvents)(nil).WatchRegistration), opts, sink, hashedAppName)
}

// FilterUnregistration mocks base method
func (m *MockIAppRegistryEvents) FilterUnregistration(opts *bind.FilterOpts, hashedAppName []common.Hash) (*adapter.AppRegistryUnregistrationIterator, error) {
	ret := m.ctrl.Call(m, "FilterUnregistration", opts, hashedAppName)
	ret0, _ := ret[0].(*adapter.AppRegistryUnregistrationIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterUnregistration indicates an expected call of FilterUnregistration
func (mr *MockIAppRegistryEventsMockRecorder) FilterUnregistration(opts, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterUnregistration", reflect.TypeOf((*MockIAppRegistryEvents)(nil).FilterUnregistration), opts, hashedAppName)
}

// ParseUnregistrationFromReceipt mocks base method
func (m *MockIAppRegistryEvents) ParseUnregistrationFromReceipt(receipt *types0.Receipt) (*adapter.AppRegistryUnregistration, error) {
	ret := m.ctrl.Call(m, "ParseUnregistrationFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.AppRegistryUnregistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseUnregistrationFromReceipt indicates an expected call of ParseUnregistrationFromReceipt
func (mr *MockIAppRegistryEventsMockRecorder) ParseUnregistrationFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUnregistrationFromReceipt", reflect.TypeOf((*MockIAppRegistryEvents)(nil).ParseUnregistrationFromReceipt), receipt)
}

// WatchUnregistration mocks base method
func (m *MockIAppRegistryEvents) WatchUnregistration(opts *bind.WatchOpts, sink chan<- *adapter.AppRegistryUnregistration, hashedAppName []common.Hash) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchUnregistration", opts, sink, hashedAppName)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUnregistration indicates an expected call of WatchUnregistration
func (mr *MockIAppRegistryEventsMockRecorder) WatchUnregistration(opts, sink, hashedAppName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUnregistration", reflect.TypeOf((*MockIAppRegistryEvents)(nil).WatchUnregistration), opts, sink, hashedAppName)
}
