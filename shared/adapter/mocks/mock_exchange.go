// Code generated by MockGen. DO NOT EDIT.
// Source: exchange.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	adapter "github.com/airbloc/airbloc-go/shared/adapter"
	types "github.com/airbloc/airbloc-go/shared/types"
	common "github.com/ethereum/go-ethereum/common"
	types0 "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockIExchangeManager is a mock of IExchangeManager interface
type MockIExchangeManager struct {
	ctrl     *gomock.Controller
	recorder *MockIExchangeManagerMockRecorder
}

// MockIExchangeManagerMockRecorder is the mock recorder for MockIExchangeManager
type MockIExchangeManagerMockRecorder struct {
	mock *MockIExchangeManager
}

// NewMockIExchangeManager creates a new mock instance
func NewMockIExchangeManager(ctrl *gomock.Controller) *MockIExchangeManager {
	mock := &MockIExchangeManager{ctrl: ctrl}
	mock.recorder = &MockIExchangeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIExchangeManager) EXPECT() *MockIExchangeManagerMockRecorder {
	return m.recorder
}

// GetOffer mocks base method
func (m *MockIExchangeManager) GetOffer(offerId types.ID) (types.Offer, error) {
	ret := m.ctrl.Call(m, "GetOffer", offerId)
	ret0, _ := ret[0].(types.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffer indicates an expected call of GetOffer
func (mr *MockIExchangeManagerMockRecorder) GetOffer(offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffer", reflect.TypeOf((*MockIExchangeManager)(nil).GetOffer), offerId)
}

// GetOfferMembers mocks base method
func (m *MockIExchangeManager) GetOfferMembers(offerId types.ID) (common.Address, common.Address, error) {
	ret := m.ctrl.Call(m, "GetOfferMembers", offerId)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(common.Address)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOfferMembers indicates an expected call of GetOfferMembers
func (mr *MockIExchangeManagerMockRecorder) GetOfferMembers(offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOfferMembers", reflect.TypeOf((*MockIExchangeManager)(nil).GetOfferMembers), offerId)
}

// OfferExists mocks base method
func (m *MockIExchangeManager) OfferExists(offerId types.ID) (bool, error) {
	ret := m.ctrl.Call(m, "OfferExists", offerId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OfferExists indicates an expected call of OfferExists
func (mr *MockIExchangeManagerMockRecorder) OfferExists(offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferExists", reflect.TypeOf((*MockIExchangeManager)(nil).OfferExists), offerId)
}

// AddDataIds mocks base method
func (m *MockIExchangeManager) AddDataIds(ctx context.Context, offerId types.ID, dataIds []types.DataId) error {
	ret := m.ctrl.Call(m, "AddDataIds", ctx, offerId, dataIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDataIds indicates an expected call of AddDataIds
func (mr *MockIExchangeManagerMockRecorder) AddDataIds(ctx, offerId, dataIds interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDataIds", reflect.TypeOf((*MockIExchangeManager)(nil).AddDataIds), ctx, offerId, dataIds)
}

// Cancel mocks base method
func (m *MockIExchangeManager) Cancel(ctx context.Context, offerId types.ID) error {
	ret := m.ctrl.Call(m, "Cancel", ctx, offerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockIExchangeManagerMockRecorder) Cancel(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockIExchangeManager)(nil).Cancel), ctx, offerId)
}

// Order mocks base method
func (m *MockIExchangeManager) Order(ctx context.Context, offerId types.ID) error {
	ret := m.ctrl.Call(m, "Order", ctx, offerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Order indicates an expected call of Order
func (mr *MockIExchangeManagerMockRecorder) Order(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockIExchangeManager)(nil).Order), ctx, offerId)
}

// Prepare mocks base method
func (m *MockIExchangeManager) Prepare(ctx context.Context, provider string, consumer, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (types.ID, error) {
	ret := m.ctrl.Call(m, "Prepare", ctx, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
	ret0, _ := ret[0].(types.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare
func (mr *MockIExchangeManagerMockRecorder) Prepare(ctx, provider, consumer, escrow, escrowSign, escrowArgs, dataIds interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockIExchangeManager)(nil).Prepare), ctx, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Reject mocks base method
func (m *MockIExchangeManager) Reject(ctx context.Context, offerId types.ID) error {
	ret := m.ctrl.Call(m, "Reject", ctx, offerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject
func (mr *MockIExchangeManagerMockRecorder) Reject(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockIExchangeManager)(nil).Reject), ctx, offerId)
}

// Settle mocks base method
func (m *MockIExchangeManager) Settle(ctx context.Context, offerId types.ID) error {
	ret := m.ctrl.Call(m, "Settle", ctx, offerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Settle indicates an expected call of Settle
func (mr *MockIExchangeManagerMockRecorder) Settle(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settle", reflect.TypeOf((*MockIExchangeManager)(nil).Settle), ctx, offerId)
}

// MockIExchangeContract is a mock of IExchangeContract interface
type MockIExchangeContract struct {
	ctrl     *gomock.Controller
	recorder *MockIExchangeContractMockRecorder
}

// MockIExchangeContractMockRecorder is the mock recorder for MockIExchangeContract
type MockIExchangeContractMockRecorder struct {
	mock *MockIExchangeContract
}

// NewMockIExchangeContract creates a new mock instance
func NewMockIExchangeContract(ctrl *gomock.Controller) *MockIExchangeContract {
	mock := &MockIExchangeContract{ctrl: ctrl}
	mock.recorder = &MockIExchangeContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIExchangeContract) EXPECT() *MockIExchangeContractMockRecorder {
	return m.recorder
}

// GetOffer mocks base method
func (m *MockIExchangeContract) GetOffer(offerId types.ID) (types.Offer, error) {
	ret := m.ctrl.Call(m, "GetOffer", offerId)
	ret0, _ := ret[0].(types.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffer indicates an expected call of GetOffer
func (mr *MockIExchangeContractMockRecorder) GetOffer(offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffer", reflect.TypeOf((*MockIExchangeContract)(nil).GetOffer), offerId)
}

// GetOfferMembers mocks base method
func (m *MockIExchangeContract) GetOfferMembers(offerId types.ID) (common.Address, common.Address, error) {
	ret := m.ctrl.Call(m, "GetOfferMembers", offerId)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(common.Address)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOfferMembers indicates an expected call of GetOfferMembers
func (mr *MockIExchangeContractMockRecorder) GetOfferMembers(offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOfferMembers", reflect.TypeOf((*MockIExchangeContract)(nil).GetOfferMembers), offerId)
}

// OfferExists mocks base method
func (m *MockIExchangeContract) OfferExists(offerId types.ID) (bool, error) {
	ret := m.ctrl.Call(m, "OfferExists", offerId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OfferExists indicates an expected call of OfferExists
func (mr *MockIExchangeContractMockRecorder) OfferExists(offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferExists", reflect.TypeOf((*MockIExchangeContract)(nil).OfferExists), offerId)
}

// AddDataIds mocks base method
func (m *MockIExchangeContract) AddDataIds(ctx context.Context, offerId types.ID, dataIds []types.DataId) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "AddDataIds", ctx, offerId, dataIds)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDataIds indicates an expected call of AddDataIds
func (mr *MockIExchangeContractMockRecorder) AddDataIds(ctx, offerId, dataIds interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDataIds", reflect.TypeOf((*MockIExchangeContract)(nil).AddDataIds), ctx, offerId, dataIds)
}

// Cancel mocks base method
func (m *MockIExchangeContract) Cancel(ctx context.Context, offerId types.ID) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Cancel", ctx, offerId)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel
func (mr *MockIExchangeContractMockRecorder) Cancel(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockIExchangeContract)(nil).Cancel), ctx, offerId)
}

// Order mocks base method
func (m *MockIExchangeContract) Order(ctx context.Context, offerId types.ID) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Order", ctx, offerId)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Order indicates an expected call of Order
func (mr *MockIExchangeContractMockRecorder) Order(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockIExchangeContract)(nil).Order), ctx, offerId)
}

// Prepare mocks base method
func (m *MockIExchangeContract) Prepare(ctx context.Context, provider string, consumer, escrow common.Address, escrowSign [4]byte, escrowArgs []byte, dataIds []types.DataId) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Prepare", ctx, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare
func (mr *MockIExchangeContractMockRecorder) Prepare(ctx, provider, consumer, escrow, escrowSign, escrowArgs, dataIds interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockIExchangeContract)(nil).Prepare), ctx, provider, consumer, escrow, escrowSign, escrowArgs, dataIds)
}

// Reject mocks base method
func (m *MockIExchangeContract) Reject(ctx context.Context, offerId types.ID) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Reject", ctx, offerId)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reject indicates an expected call of Reject
func (mr *MockIExchangeContractMockRecorder) Reject(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockIExchangeContract)(nil).Reject), ctx, offerId)
}

// Settle mocks base method
func (m *MockIExchangeContract) Settle(ctx context.Context, offerId types.ID) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Settle", ctx, offerId)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Settle indicates an expected call of Settle
func (mr *MockIExchangeContractMockRecorder) Settle(ctx, offerId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settle", reflect.TypeOf((*MockIExchangeContract)(nil).Settle), ctx, offerId)
}

// ParseEscrowExecutionFailedFromReceipt mocks base method
func (m *MockIExchangeContract) ParseEscrowExecutionFailedFromReceipt(receipt *types0.Receipt) (*adapter.ExchangeEscrowExecutionFailed, error) {
	ret := m.ctrl.Call(m, "ParseEscrowExecutionFailedFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.ExchangeEscrowExecutionFailed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseEscrowExecutionFailedFromReceipt indicates an expected call of ParseEscrowExecutionFailedFromReceipt
func (mr *MockIExchangeContractMockRecorder) ParseEscrowExecutionFailedFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseEscrowExecutionFailedFromReceipt", reflect.TypeOf((*MockIExchangeContract)(nil).ParseEscrowExecutionFailedFromReceipt), receipt)
}

// ParseOfferCanceledFromReceipt mocks base method
func (m *MockIExchangeContract) ParseOfferCanceledFromReceipt(receipt *types0.Receipt) (*adapter.ExchangeOfferCanceled, error) {
	ret := m.ctrl.Call(m, "ParseOfferCanceledFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.ExchangeOfferCanceled)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseOfferCanceledFromReceipt indicates an expected call of ParseOfferCanceledFromReceipt
func (mr *MockIExchangeContractMockRecorder) ParseOfferCanceledFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseOfferCanceledFromReceipt", reflect.TypeOf((*MockIExchangeContract)(nil).ParseOfferCanceledFromReceipt), receipt)
}

// ParseOfferPreparedFromReceipt mocks base method
func (m *MockIExchangeContract) ParseOfferPreparedFromReceipt(receipt *types0.Receipt) (*adapter.ExchangeOfferPrepared, error) {
	ret := m.ctrl.Call(m, "ParseOfferPreparedFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.ExchangeOfferPrepared)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseOfferPreparedFromReceipt indicates an expected call of ParseOfferPreparedFromReceipt
func (mr *MockIExchangeContractMockRecorder) ParseOfferPreparedFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseOfferPreparedFromReceipt", reflect.TypeOf((*MockIExchangeContract)(nil).ParseOfferPreparedFromReceipt), receipt)
}

// ParseOfferPresentedFromReceipt mocks base method
func (m *MockIExchangeContract) ParseOfferPresentedFromReceipt(receipt *types0.Receipt) (*adapter.ExchangeOfferPresented, error) {
	ret := m.ctrl.Call(m, "ParseOfferPresentedFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.ExchangeOfferPresented)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseOfferPresentedFromReceipt indicates an expected call of ParseOfferPresentedFromReceipt
func (mr *MockIExchangeContractMockRecorder) ParseOfferPresentedFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseOfferPresentedFromReceipt", reflect.TypeOf((*MockIExchangeContract)(nil).ParseOfferPresentedFromReceipt), receipt)
}

// ParseOfferReceiptFromReceipt mocks base method
func (m *MockIExchangeContract) ParseOfferReceiptFromReceipt(receipt *types0.Receipt) (*adapter.ExchangeOfferReceipt, error) {
	ret := m.ctrl.Call(m, "ParseOfferReceiptFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.ExchangeOfferReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseOfferReceiptFromReceipt indicates an expected call of ParseOfferReceiptFromReceipt
func (mr *MockIExchangeContractMockRecorder) ParseOfferReceiptFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseOfferReceiptFromReceipt", reflect.TypeOf((*MockIExchangeContract)(nil).ParseOfferReceiptFromReceipt), receipt)
}

// ParseOfferRejectedFromReceipt mocks base method
func (m *MockIExchangeContract) ParseOfferRejectedFromReceipt(receipt *types0.Receipt) (*adapter.ExchangeOfferRejected, error) {
	ret := m.ctrl.Call(m, "ParseOfferRejectedFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.ExchangeOfferRejected)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseOfferRejectedFromReceipt indicates an expected call of ParseOfferRejectedFromReceipt
func (mr *MockIExchangeContractMockRecorder) ParseOfferRejectedFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseOfferRejectedFromReceipt", reflect.TypeOf((*MockIExchangeContract)(nil).ParseOfferRejectedFromReceipt), receipt)
}

// ParseOfferSettledFromReceipt mocks base method
func (m *MockIExchangeContract) ParseOfferSettledFromReceipt(receipt *types0.Receipt) (*adapter.ExchangeOfferSettled, error) {
	ret := m.ctrl.Call(m, "ParseOfferSettledFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.ExchangeOfferSettled)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseOfferSettledFromReceipt indicates an expected call of ParseOfferSettledFromReceipt
func (mr *MockIExchangeContractMockRecorder) ParseOfferSettledFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseOfferSettledFromReceipt", reflect.TypeOf((*MockIExchangeContract)(nil).ParseOfferSettledFromReceipt), receipt)
}