// Code generated by MockGen. DO NOT EDIT.
// Source: consents_wrapper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	adapter "github.com/airbloc/airbloc-go/shared/adapter"
	blockchain "github.com/airbloc/airbloc-go/shared/blockchain"
	types "github.com/airbloc/airbloc-go/shared/types"
	gomock "github.com/golang/mock/gomock"
	bind "github.com/klaytn/klaytn/accounts/abi/bind"
	types0 "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
)

// MockIConsentsManager is a mock of IConsentsManager interface
type MockIConsentsManager struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsManagerMockRecorder
}

// MockIConsentsManagerMockRecorder is the mock recorder for MockIConsentsManager
type MockIConsentsManagerMockRecorder struct {
	mock *MockIConsentsManager
}

// NewMockIConsentsManager creates a new mock instance
func NewMockIConsentsManager(ctrl *gomock.Controller) *MockIConsentsManager {
	mock := &MockIConsentsManager{ctrl: ctrl}
	mock.recorder = &MockIConsentsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsManager) EXPECT() *MockIConsentsManagerMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockIConsentsManager) Address() common.Address {
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockIConsentsManagerMockRecorder) Address() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockIConsentsManager)(nil).Address))
}

// TxHash mocks base method
func (m *MockIConsentsManager) TxHash() common.Hash {
	ret := m.ctrl.Call(m, "TxHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// TxHash indicates an expected call of TxHash
func (mr *MockIConsentsManagerMockRecorder) TxHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxHash", reflect.TypeOf((*MockIConsentsManager)(nil).TxHash))
}

// CreatedAt mocks base method
func (m *MockIConsentsManager) CreatedAt() *big.Int {
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CreatedAt indicates an expected call of CreatedAt
func (mr *MockIConsentsManagerMockRecorder) CreatedAt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockIConsentsManager)(nil).CreatedAt))
}

// IsAllowed mocks base method
func (m *MockIConsentsManager) IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	ret := m.ctrl.Call(m, "IsAllowed", userId, appName, action, dataType)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAllowed indicates an expected call of IsAllowed
func (mr *MockIConsentsManagerMockRecorder) IsAllowed(userId, appName, action, dataType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAllowed", reflect.TypeOf((*MockIConsentsManager)(nil).IsAllowed), userId, appName, action, dataType)
}

// IsAllowedAt mocks base method
func (m *MockIConsentsManager) IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	ret := m.ctrl.Call(m, "IsAllowedAt", userId, appName, action, dataType, blockNumber)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAllowedAt indicates an expected call of IsAllowedAt
func (mr *MockIConsentsManagerMockRecorder) IsAllowedAt(userId, appName, action, dataType, blockNumber interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAllowedAt", reflect.TypeOf((*MockIConsentsManager)(nil).IsAllowedAt), userId, appName, action, dataType, blockNumber)
}

// Consent mocks base method
func (m *MockIConsentsManager) Consent(ctx context.Context, opts *blockchain.TransactOpts, appName string, consentData types.ConsentData) error {
	ret := m.ctrl.Call(m, "Consent", ctx, opts, appName, consentData)
	ret0, _ := ret[0].(error)
	return ret0
}

// Consent indicates an expected call of Consent
func (mr *MockIConsentsManagerMockRecorder) Consent(ctx, opts, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consent", reflect.TypeOf((*MockIConsentsManager)(nil).Consent), ctx, opts, appName, consentData)
}

// ConsentByController mocks base method
func (m *MockIConsentsManager) ConsentByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData) error {
	ret := m.ctrl.Call(m, "ConsentByController", ctx, opts, userId, appName, consentData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsentByController indicates an expected call of ConsentByController
func (mr *MockIConsentsManagerMockRecorder) ConsentByController(ctx, opts, userId, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentByController", reflect.TypeOf((*MockIConsentsManager)(nil).ConsentByController), ctx, opts, userId, appName, consentData)
}

// ConsentMany mocks base method
func (m *MockIConsentsManager) ConsentMany(ctx context.Context, opts *blockchain.TransactOpts, appName string, consentData []types.ConsentData) error {
	ret := m.ctrl.Call(m, "ConsentMany", ctx, opts, appName, consentData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsentMany indicates an expected call of ConsentMany
func (mr *MockIConsentsManagerMockRecorder) ConsentMany(ctx, opts, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentMany", reflect.TypeOf((*MockIConsentsManager)(nil).ConsentMany), ctx, opts, appName, consentData)
}

// ConsentManyByController mocks base method
func (m *MockIConsentsManager) ConsentManyByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData) error {
	ret := m.ctrl.Call(m, "ConsentManyByController", ctx, opts, userId, appName, consentData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsentManyByController indicates an expected call of ConsentManyByController
func (mr *MockIConsentsManagerMockRecorder) ConsentManyByController(ctx, opts, userId, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentManyByController", reflect.TypeOf((*MockIConsentsManager)(nil).ConsentManyByController), ctx, opts, userId, appName, consentData)
}

// ModifyConsentByController mocks base method
func (m *MockIConsentsManager) ModifyConsentByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) error {
	ret := m.ctrl.Call(m, "ModifyConsentByController", ctx, opts, userId, appName, consentData, passwordSignature)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyConsentByController indicates an expected call of ModifyConsentByController
func (mr *MockIConsentsManagerMockRecorder) ModifyConsentByController(ctx, opts, userId, appName, consentData, passwordSignature interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyConsentByController", reflect.TypeOf((*MockIConsentsManager)(nil).ModifyConsentByController), ctx, opts, userId, appName, consentData, passwordSignature)
}

// ModifyConsentManyByController mocks base method
func (m *MockIConsentsManager) ModifyConsentManyByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData, passwordSignature []byte) error {
	ret := m.ctrl.Call(m, "ModifyConsentManyByController", ctx, opts, userId, appName, consentData, passwordSignature)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyConsentManyByController indicates an expected call of ModifyConsentManyByController
func (mr *MockIConsentsManagerMockRecorder) ModifyConsentManyByController(ctx, opts, userId, appName, consentData, passwordSignature interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyConsentManyByController", reflect.TypeOf((*MockIConsentsManager)(nil).ModifyConsentManyByController), ctx, opts, userId, appName, consentData, passwordSignature)
}

// FilterConsented mocks base method
func (m *MockIConsentsManager) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*adapter.ConsentsConsentedIterator, error) {
	ret := m.ctrl.Call(m, "FilterConsented", opts, action, userId, appAddr)
	ret0, _ := ret[0].(*adapter.ConsentsConsentedIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterConsented indicates an expected call of FilterConsented
func (mr *MockIConsentsManagerMockRecorder) FilterConsented(opts, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterConsented", reflect.TypeOf((*MockIConsentsManager)(nil).FilterConsented), opts, action, userId, appAddr)
}

// WatchConsented mocks base method
func (m *MockIConsentsManager) WatchConsented(opts *bind.WatchOpts, sink chan<- *adapter.ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchConsented", opts, sink, action, userId, appAddr)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchConsented indicates an expected call of WatchConsented
func (mr *MockIConsentsManagerMockRecorder) WatchConsented(opts, sink, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchConsented", reflect.TypeOf((*MockIConsentsManager)(nil).WatchConsented), opts, sink, action, userId, appAddr)
}

// MockIConsentsCalls is a mock of IConsentsCalls interface
type MockIConsentsCalls struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsCallsMockRecorder
}

// MockIConsentsCallsMockRecorder is the mock recorder for MockIConsentsCalls
type MockIConsentsCallsMockRecorder struct {
	mock *MockIConsentsCalls
}

// NewMockIConsentsCalls creates a new mock instance
func NewMockIConsentsCalls(ctrl *gomock.Controller) *MockIConsentsCalls {
	mock := &MockIConsentsCalls{ctrl: ctrl}
	mock.recorder = &MockIConsentsCallsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsCalls) EXPECT() *MockIConsentsCallsMockRecorder {
	return m.recorder
}

// IsAllowed mocks base method
func (m *MockIConsentsCalls) IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	ret := m.ctrl.Call(m, "IsAllowed", userId, appName, action, dataType)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAllowed indicates an expected call of IsAllowed
func (mr *MockIConsentsCallsMockRecorder) IsAllowed(userId, appName, action, dataType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAllowed", reflect.TypeOf((*MockIConsentsCalls)(nil).IsAllowed), userId, appName, action, dataType)
}

// IsAllowedAt mocks base method
func (m *MockIConsentsCalls) IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	ret := m.ctrl.Call(m, "IsAllowedAt", userId, appName, action, dataType, blockNumber)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAllowedAt indicates an expected call of IsAllowedAt
func (mr *MockIConsentsCallsMockRecorder) IsAllowedAt(userId, appName, action, dataType, blockNumber interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAllowedAt", reflect.TypeOf((*MockIConsentsCalls)(nil).IsAllowedAt), userId, appName, action, dataType, blockNumber)
}

// MockIConsentsTransacts is a mock of IConsentsTransacts interface
type MockIConsentsTransacts struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsTransactsMockRecorder
}

// MockIConsentsTransactsMockRecorder is the mock recorder for MockIConsentsTransacts
type MockIConsentsTransactsMockRecorder struct {
	mock *MockIConsentsTransacts
}

// NewMockIConsentsTransacts creates a new mock instance
func NewMockIConsentsTransacts(ctrl *gomock.Controller) *MockIConsentsTransacts {
	mock := &MockIConsentsTransacts{ctrl: ctrl}
	mock.recorder = &MockIConsentsTransactsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsTransacts) EXPECT() *MockIConsentsTransactsMockRecorder {
	return m.recorder
}

// Consent mocks base method
func (m *MockIConsentsTransacts) Consent(ctx context.Context, opts *blockchain.TransactOpts, appName string, consentData types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Consent", ctx, opts, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consent indicates an expected call of Consent
func (mr *MockIConsentsTransactsMockRecorder) Consent(ctx, opts, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consent", reflect.TypeOf((*MockIConsentsTransacts)(nil).Consent), ctx, opts, appName, consentData)
}

// ConsentByController mocks base method
func (m *MockIConsentsTransacts) ConsentByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ConsentByController", ctx, opts, userId, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsentByController indicates an expected call of ConsentByController
func (mr *MockIConsentsTransactsMockRecorder) ConsentByController(ctx, opts, userId, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentByController", reflect.TypeOf((*MockIConsentsTransacts)(nil).ConsentByController), ctx, opts, userId, appName, consentData)
}

// ConsentMany mocks base method
func (m *MockIConsentsTransacts) ConsentMany(ctx context.Context, opts *blockchain.TransactOpts, appName string, consentData []types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ConsentMany", ctx, opts, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsentMany indicates an expected call of ConsentMany
func (mr *MockIConsentsTransactsMockRecorder) ConsentMany(ctx, opts, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentMany", reflect.TypeOf((*MockIConsentsTransacts)(nil).ConsentMany), ctx, opts, appName, consentData)
}

// ConsentManyByController mocks base method
func (m *MockIConsentsTransacts) ConsentManyByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ConsentManyByController", ctx, opts, userId, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsentManyByController indicates an expected call of ConsentManyByController
func (mr *MockIConsentsTransactsMockRecorder) ConsentManyByController(ctx, opts, userId, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentManyByController", reflect.TypeOf((*MockIConsentsTransacts)(nil).ConsentManyByController), ctx, opts, userId, appName, consentData)
}

// ModifyConsentByController mocks base method
func (m *MockIConsentsTransacts) ModifyConsentByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ModifyConsentByController", ctx, opts, userId, appName, consentData, passwordSignature)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyConsentByController indicates an expected call of ModifyConsentByController
func (mr *MockIConsentsTransactsMockRecorder) ModifyConsentByController(ctx, opts, userId, appName, consentData, passwordSignature interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyConsentByController", reflect.TypeOf((*MockIConsentsTransacts)(nil).ModifyConsentByController), ctx, opts, userId, appName, consentData, passwordSignature)
}

// ModifyConsentManyByController mocks base method
func (m *MockIConsentsTransacts) ModifyConsentManyByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData, passwordSignature []byte) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ModifyConsentManyByController", ctx, opts, userId, appName, consentData, passwordSignature)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyConsentManyByController indicates an expected call of ModifyConsentManyByController
func (mr *MockIConsentsTransactsMockRecorder) ModifyConsentManyByController(ctx, opts, userId, appName, consentData, passwordSignature interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyConsentManyByController", reflect.TypeOf((*MockIConsentsTransacts)(nil).ModifyConsentManyByController), ctx, opts, userId, appName, consentData, passwordSignature)
}

// MockIConsentsEvents is a mock of IConsentsEvents interface
type MockIConsentsEvents struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsEventsMockRecorder
}

// MockIConsentsEventsMockRecorder is the mock recorder for MockIConsentsEvents
type MockIConsentsEventsMockRecorder struct {
	mock *MockIConsentsEvents
}

// NewMockIConsentsEvents creates a new mock instance
func NewMockIConsentsEvents(ctrl *gomock.Controller) *MockIConsentsEvents {
	mock := &MockIConsentsEvents{ctrl: ctrl}
	mock.recorder = &MockIConsentsEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsEvents) EXPECT() *MockIConsentsEventsMockRecorder {
	return m.recorder
}

// FilterConsented mocks base method
func (m *MockIConsentsEvents) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*adapter.ConsentsConsentedIterator, error) {
	ret := m.ctrl.Call(m, "FilterConsented", opts, action, userId, appAddr)
	ret0, _ := ret[0].(*adapter.ConsentsConsentedIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterConsented indicates an expected call of FilterConsented
func (mr *MockIConsentsEventsMockRecorder) FilterConsented(opts, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterConsented", reflect.TypeOf((*MockIConsentsEvents)(nil).FilterConsented), opts, action, userId, appAddr)
}

// ParseConsented mocks base method
func (m *MockIConsentsEvents) ParseConsented(log types0.Log) (*adapter.ConsentsConsented, error) {
	ret := m.ctrl.Call(m, "ParseConsented", log)
	ret0, _ := ret[0].(*adapter.ConsentsConsented)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseConsented indicates an expected call of ParseConsented
func (mr *MockIConsentsEventsMockRecorder) ParseConsented(log interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseConsented", reflect.TypeOf((*MockIConsentsEvents)(nil).ParseConsented), log)
}

// ParseConsentedFromReceipt mocks base method
func (m *MockIConsentsEvents) ParseConsentedFromReceipt(receipt *types0.Receipt) ([]*adapter.ConsentsConsented, error) {
	ret := m.ctrl.Call(m, "ParseConsentedFromReceipt", receipt)
	ret0, _ := ret[0].([]*adapter.ConsentsConsented)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseConsentedFromReceipt indicates an expected call of ParseConsentedFromReceipt
func (mr *MockIConsentsEventsMockRecorder) ParseConsentedFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseConsentedFromReceipt", reflect.TypeOf((*MockIConsentsEvents)(nil).ParseConsentedFromReceipt), receipt)
}

// WatchConsented mocks base method
func (m *MockIConsentsEvents) WatchConsented(opts *bind.WatchOpts, sink chan<- *adapter.ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchConsented", opts, sink, action, userId, appAddr)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchConsented indicates an expected call of WatchConsented
func (mr *MockIConsentsEventsMockRecorder) WatchConsented(opts, sink, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchConsented", reflect.TypeOf((*MockIConsentsEvents)(nil).WatchConsented), opts, sink, action, userId, appAddr)
}

// MockIConsentsFilterer is a mock of IConsentsFilterer interface
type MockIConsentsFilterer struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsFiltererMockRecorder
}

// MockIConsentsFiltererMockRecorder is the mock recorder for MockIConsentsFilterer
type MockIConsentsFiltererMockRecorder struct {
	mock *MockIConsentsFilterer
}

// NewMockIConsentsFilterer creates a new mock instance
func NewMockIConsentsFilterer(ctrl *gomock.Controller) *MockIConsentsFilterer {
	mock := &MockIConsentsFilterer{ctrl: ctrl}
	mock.recorder = &MockIConsentsFiltererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsFilterer) EXPECT() *MockIConsentsFiltererMockRecorder {
	return m.recorder
}

// FilterConsented mocks base method
func (m *MockIConsentsFilterer) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*adapter.ConsentsConsentedIterator, error) {
	ret := m.ctrl.Call(m, "FilterConsented", opts, action, userId, appAddr)
	ret0, _ := ret[0].(*adapter.ConsentsConsentedIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterConsented indicates an expected call of FilterConsented
func (mr *MockIConsentsFiltererMockRecorder) FilterConsented(opts, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterConsented", reflect.TypeOf((*MockIConsentsFilterer)(nil).FilterConsented), opts, action, userId, appAddr)
}

// MockIConsentsParser is a mock of IConsentsParser interface
type MockIConsentsParser struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsParserMockRecorder
}

// MockIConsentsParserMockRecorder is the mock recorder for MockIConsentsParser
type MockIConsentsParserMockRecorder struct {
	mock *MockIConsentsParser
}

// NewMockIConsentsParser creates a new mock instance
func NewMockIConsentsParser(ctrl *gomock.Controller) *MockIConsentsParser {
	mock := &MockIConsentsParser{ctrl: ctrl}
	mock.recorder = &MockIConsentsParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsParser) EXPECT() *MockIConsentsParserMockRecorder {
	return m.recorder
}

// ParseConsented mocks base method
func (m *MockIConsentsParser) ParseConsented(log types0.Log) (*adapter.ConsentsConsented, error) {
	ret := m.ctrl.Call(m, "ParseConsented", log)
	ret0, _ := ret[0].(*adapter.ConsentsConsented)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseConsented indicates an expected call of ParseConsented
func (mr *MockIConsentsParserMockRecorder) ParseConsented(log interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseConsented", reflect.TypeOf((*MockIConsentsParser)(nil).ParseConsented), log)
}

// ParseConsentedFromReceipt mocks base method
func (m *MockIConsentsParser) ParseConsentedFromReceipt(receipt *types0.Receipt) ([]*adapter.ConsentsConsented, error) {
	ret := m.ctrl.Call(m, "ParseConsentedFromReceipt", receipt)
	ret0, _ := ret[0].([]*adapter.ConsentsConsented)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseConsentedFromReceipt indicates an expected call of ParseConsentedFromReceipt
func (mr *MockIConsentsParserMockRecorder) ParseConsentedFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseConsentedFromReceipt", reflect.TypeOf((*MockIConsentsParser)(nil).ParseConsentedFromReceipt), receipt)
}

// MockIConsentsWatcher is a mock of IConsentsWatcher interface
type MockIConsentsWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsWatcherMockRecorder
}

// MockIConsentsWatcherMockRecorder is the mock recorder for MockIConsentsWatcher
type MockIConsentsWatcherMockRecorder struct {
	mock *MockIConsentsWatcher
}

// NewMockIConsentsWatcher creates a new mock instance
func NewMockIConsentsWatcher(ctrl *gomock.Controller) *MockIConsentsWatcher {
	mock := &MockIConsentsWatcher{ctrl: ctrl}
	mock.recorder = &MockIConsentsWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsWatcher) EXPECT() *MockIConsentsWatcherMockRecorder {
	return m.recorder
}

// WatchConsented mocks base method
func (m *MockIConsentsWatcher) WatchConsented(opts *bind.WatchOpts, sink chan<- *adapter.ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchConsented", opts, sink, action, userId, appAddr)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchConsented indicates an expected call of WatchConsented
func (mr *MockIConsentsWatcherMockRecorder) WatchConsented(opts, sink, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchConsented", reflect.TypeOf((*MockIConsentsWatcher)(nil).WatchConsented), opts, sink, action, userId, appAddr)
}

// MockIConsentsContract is a mock of IConsentsContract interface
type MockIConsentsContract struct {
	ctrl     *gomock.Controller
	recorder *MockIConsentsContractMockRecorder
}

// MockIConsentsContractMockRecorder is the mock recorder for MockIConsentsContract
type MockIConsentsContractMockRecorder struct {
	mock *MockIConsentsContract
}

// NewMockIConsentsContract creates a new mock instance
func NewMockIConsentsContract(ctrl *gomock.Controller) *MockIConsentsContract {
	mock := &MockIConsentsContract{ctrl: ctrl}
	mock.recorder = &MockIConsentsContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConsentsContract) EXPECT() *MockIConsentsContractMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockIConsentsContract) Address() common.Address {
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockIConsentsContractMockRecorder) Address() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockIConsentsContract)(nil).Address))
}

// TxHash mocks base method
func (m *MockIConsentsContract) TxHash() common.Hash {
	ret := m.ctrl.Call(m, "TxHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// TxHash indicates an expected call of TxHash
func (mr *MockIConsentsContractMockRecorder) TxHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxHash", reflect.TypeOf((*MockIConsentsContract)(nil).TxHash))
}

// CreatedAt mocks base method
func (m *MockIConsentsContract) CreatedAt() *big.Int {
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CreatedAt indicates an expected call of CreatedAt
func (mr *MockIConsentsContractMockRecorder) CreatedAt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockIConsentsContract)(nil).CreatedAt))
}

// IsAllowed mocks base method
func (m *MockIConsentsContract) IsAllowed(userId types.ID, appName string, action uint8, dataType string) (bool, error) {
	ret := m.ctrl.Call(m, "IsAllowed", userId, appName, action, dataType)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAllowed indicates an expected call of IsAllowed
func (mr *MockIConsentsContractMockRecorder) IsAllowed(userId, appName, action, dataType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAllowed", reflect.TypeOf((*MockIConsentsContract)(nil).IsAllowed), userId, appName, action, dataType)
}

// IsAllowedAt mocks base method
func (m *MockIConsentsContract) IsAllowedAt(userId types.ID, appName string, action uint8, dataType string, blockNumber *big.Int) (bool, error) {
	ret := m.ctrl.Call(m, "IsAllowedAt", userId, appName, action, dataType, blockNumber)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAllowedAt indicates an expected call of IsAllowedAt
func (mr *MockIConsentsContractMockRecorder) IsAllowedAt(userId, appName, action, dataType, blockNumber interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAllowedAt", reflect.TypeOf((*MockIConsentsContract)(nil).IsAllowedAt), userId, appName, action, dataType, blockNumber)
}

// Consent mocks base method
func (m *MockIConsentsContract) Consent(ctx context.Context, opts *blockchain.TransactOpts, appName string, consentData types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Consent", ctx, opts, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consent indicates an expected call of Consent
func (mr *MockIConsentsContractMockRecorder) Consent(ctx, opts, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consent", reflect.TypeOf((*MockIConsentsContract)(nil).Consent), ctx, opts, appName, consentData)
}

// ConsentByController mocks base method
func (m *MockIConsentsContract) ConsentByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ConsentByController", ctx, opts, userId, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsentByController indicates an expected call of ConsentByController
func (mr *MockIConsentsContractMockRecorder) ConsentByController(ctx, opts, userId, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentByController", reflect.TypeOf((*MockIConsentsContract)(nil).ConsentByController), ctx, opts, userId, appName, consentData)
}

// ConsentMany mocks base method
func (m *MockIConsentsContract) ConsentMany(ctx context.Context, opts *blockchain.TransactOpts, appName string, consentData []types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ConsentMany", ctx, opts, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsentMany indicates an expected call of ConsentMany
func (mr *MockIConsentsContractMockRecorder) ConsentMany(ctx, opts, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentMany", reflect.TypeOf((*MockIConsentsContract)(nil).ConsentMany), ctx, opts, appName, consentData)
}

// ConsentManyByController mocks base method
func (m *MockIConsentsContract) ConsentManyByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ConsentManyByController", ctx, opts, userId, appName, consentData)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsentManyByController indicates an expected call of ConsentManyByController
func (mr *MockIConsentsContractMockRecorder) ConsentManyByController(ctx, opts, userId, appName, consentData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentManyByController", reflect.TypeOf((*MockIConsentsContract)(nil).ConsentManyByController), ctx, opts, userId, appName, consentData)
}

// ModifyConsentByController mocks base method
func (m *MockIConsentsContract) ModifyConsentByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData types.ConsentData, passwordSignature []byte) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ModifyConsentByController", ctx, opts, userId, appName, consentData, passwordSignature)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyConsentByController indicates an expected call of ModifyConsentByController
func (mr *MockIConsentsContractMockRecorder) ModifyConsentByController(ctx, opts, userId, appName, consentData, passwordSignature interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyConsentByController", reflect.TypeOf((*MockIConsentsContract)(nil).ModifyConsentByController), ctx, opts, userId, appName, consentData, passwordSignature)
}

// ModifyConsentManyByController mocks base method
func (m *MockIConsentsContract) ModifyConsentManyByController(ctx context.Context, opts *blockchain.TransactOpts, userId types.ID, appName string, consentData []types.ConsentData, passwordSignature []byte) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "ModifyConsentManyByController", ctx, opts, userId, appName, consentData, passwordSignature)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyConsentManyByController indicates an expected call of ModifyConsentManyByController
func (mr *MockIConsentsContractMockRecorder) ModifyConsentManyByController(ctx, opts, userId, appName, consentData, passwordSignature interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyConsentManyByController", reflect.TypeOf((*MockIConsentsContract)(nil).ModifyConsentManyByController), ctx, opts, userId, appName, consentData, passwordSignature)
}

// FilterConsented mocks base method
func (m *MockIConsentsContract) FilterConsented(opts *bind.FilterOpts, action []uint8, userId []types.ID, appAddr []common.Address) (*adapter.ConsentsConsentedIterator, error) {
	ret := m.ctrl.Call(m, "FilterConsented", opts, action, userId, appAddr)
	ret0, _ := ret[0].(*adapter.ConsentsConsentedIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterConsented indicates an expected call of FilterConsented
func (mr *MockIConsentsContractMockRecorder) FilterConsented(opts, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterConsented", reflect.TypeOf((*MockIConsentsContract)(nil).FilterConsented), opts, action, userId, appAddr)
}

// ParseConsented mocks base method
func (m *MockIConsentsContract) ParseConsented(log types0.Log) (*adapter.ConsentsConsented, error) {
	ret := m.ctrl.Call(m, "ParseConsented", log)
	ret0, _ := ret[0].(*adapter.ConsentsConsented)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseConsented indicates an expected call of ParseConsented
func (mr *MockIConsentsContractMockRecorder) ParseConsented(log interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseConsented", reflect.TypeOf((*MockIConsentsContract)(nil).ParseConsented), log)
}

// ParseConsentedFromReceipt mocks base method
func (m *MockIConsentsContract) ParseConsentedFromReceipt(receipt *types0.Receipt) ([]*adapter.ConsentsConsented, error) {
	ret := m.ctrl.Call(m, "ParseConsentedFromReceipt", receipt)
	ret0, _ := ret[0].([]*adapter.ConsentsConsented)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseConsentedFromReceipt indicates an expected call of ParseConsentedFromReceipt
func (mr *MockIConsentsContractMockRecorder) ParseConsentedFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseConsentedFromReceipt", reflect.TypeOf((*MockIConsentsContract)(nil).ParseConsentedFromReceipt), receipt)
}

// WatchConsented mocks base method
func (m *MockIConsentsContract) WatchConsented(opts *bind.WatchOpts, sink chan<- *adapter.ConsentsConsented, action []uint8, userId []types.ID, appAddr []common.Address) (event.Subscription, error) {
	ret := m.ctrl.Call(m, "WatchConsented", opts, sink, action, userId, appAddr)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchConsented indicates an expected call of WatchConsented
func (mr *MockIConsentsContractMockRecorder) WatchConsented(opts, sink, action, userId, appAddr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchConsented", reflect.TypeOf((*MockIConsentsContract)(nil).WatchConsented), opts, sink, action, userId, appAddr)
}
