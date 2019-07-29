// Code generated by MockGen. DO NOT EDIT.
// Source: data_type_registry.go

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

// MockIDataTypeRegistryManager is a mock of IDataTypeRegistryManager interface
type MockIDataTypeRegistryManager struct {
	ctrl     *gomock.Controller
	recorder *MockIDataTypeRegistryManagerMockRecorder
}

// MockIDataTypeRegistryManagerMockRecorder is the mock recorder for MockIDataTypeRegistryManager
type MockIDataTypeRegistryManagerMockRecorder struct {
	mock *MockIDataTypeRegistryManager
}

// NewMockIDataTypeRegistryManager creates a new mock instance
func NewMockIDataTypeRegistryManager(ctrl *gomock.Controller) *MockIDataTypeRegistryManager {
	mock := &MockIDataTypeRegistryManager{ctrl: ctrl}
	mock.recorder = &MockIDataTypeRegistryManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDataTypeRegistryManager) EXPECT() *MockIDataTypeRegistryManagerMockRecorder {
	return m.recorder
}

// Exists mocks base method
func (m *MockIDataTypeRegistryManager) Exists(name string) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockIDataTypeRegistryManagerMockRecorder) Exists(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIDataTypeRegistryManager)(nil).Exists), name)
}

// Get mocks base method
func (m *MockIDataTypeRegistryManager) Get(name string) (types.DataType, error) {
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(types.DataType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIDataTypeRegistryManagerMockRecorder) Get(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIDataTypeRegistryManager)(nil).Get), name)
}

// IsOwner mocks base method
func (m *MockIDataTypeRegistryManager) IsOwner(name string, owner common.Address) (bool, error) {
	ret := m.ctrl.Call(m, "IsOwner", name, owner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockIDataTypeRegistryManagerMockRecorder) IsOwner(name, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockIDataTypeRegistryManager)(nil).IsOwner), name, owner)
}

// Register mocks base method
func (m *MockIDataTypeRegistryManager) Register(ctx context.Context, name string, schemaHash common.Hash) error {
	ret := m.ctrl.Call(m, "Register", ctx, name, schemaHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockIDataTypeRegistryManagerMockRecorder) Register(ctx, name, schemaHash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIDataTypeRegistryManager)(nil).Register), ctx, name, schemaHash)
}

// Unregister mocks base method
func (m *MockIDataTypeRegistryManager) Unregister(ctx context.Context, name string) error {
	ret := m.ctrl.Call(m, "Unregister", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister
func (mr *MockIDataTypeRegistryManagerMockRecorder) Unregister(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockIDataTypeRegistryManager)(nil).Unregister), ctx, name)
}

// MockIDataTypeRegistryContract is a mock of IDataTypeRegistryContract interface
type MockIDataTypeRegistryContract struct {
	ctrl     *gomock.Controller
	recorder *MockIDataTypeRegistryContractMockRecorder
}

// MockIDataTypeRegistryContractMockRecorder is the mock recorder for MockIDataTypeRegistryContract
type MockIDataTypeRegistryContractMockRecorder struct {
	mock *MockIDataTypeRegistryContract
}

// NewMockIDataTypeRegistryContract creates a new mock instance
func NewMockIDataTypeRegistryContract(ctrl *gomock.Controller) *MockIDataTypeRegistryContract {
	mock := &MockIDataTypeRegistryContract{ctrl: ctrl}
	mock.recorder = &MockIDataTypeRegistryContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDataTypeRegistryContract) EXPECT() *MockIDataTypeRegistryContractMockRecorder {
	return m.recorder
}

// Exists mocks base method
func (m *MockIDataTypeRegistryContract) Exists(name string) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockIDataTypeRegistryContractMockRecorder) Exists(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIDataTypeRegistryContract)(nil).Exists), name)
}

// Get mocks base method
func (m *MockIDataTypeRegistryContract) Get(name string) (types.DataType, error) {
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(types.DataType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIDataTypeRegistryContractMockRecorder) Get(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIDataTypeRegistryContract)(nil).Get), name)
}

// IsOwner mocks base method
func (m *MockIDataTypeRegistryContract) IsOwner(name string, owner common.Address) (bool, error) {
	ret := m.ctrl.Call(m, "IsOwner", name, owner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockIDataTypeRegistryContractMockRecorder) IsOwner(name, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockIDataTypeRegistryContract)(nil).IsOwner), name, owner)
}

// Register mocks base method
func (m *MockIDataTypeRegistryContract) Register(ctx context.Context, name string, schemaHash common.Hash) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Register", ctx, name, schemaHash)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockIDataTypeRegistryContractMockRecorder) Register(ctx, name, schemaHash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIDataTypeRegistryContract)(nil).Register), ctx, name, schemaHash)
}

// Unregister mocks base method
func (m *MockIDataTypeRegistryContract) Unregister(ctx context.Context, name string) (*types0.Receipt, error) {
	ret := m.ctrl.Call(m, "Unregister", ctx, name)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister
func (mr *MockIDataTypeRegistryContractMockRecorder) Unregister(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockIDataTypeRegistryContract)(nil).Unregister), ctx, name)
}

// ParseRegistrationFromReceipt mocks base method
func (m *MockIDataTypeRegistryContract) ParseRegistrationFromReceipt(receipt *types0.Receipt) (*adapter.DataTypeRegistryRegistration, error) {
	ret := m.ctrl.Call(m, "ParseRegistrationFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.DataTypeRegistryRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseRegistrationFromReceipt indicates an expected call of ParseRegistrationFromReceipt
func (mr *MockIDataTypeRegistryContractMockRecorder) ParseRegistrationFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRegistrationFromReceipt", reflect.TypeOf((*MockIDataTypeRegistryContract)(nil).ParseRegistrationFromReceipt), receipt)
}

// ParseUnregistrationFromReceipt mocks base method
func (m *MockIDataTypeRegistryContract) ParseUnregistrationFromReceipt(receipt *types0.Receipt) (*adapter.DataTypeRegistryUnregistration, error) {
	ret := m.ctrl.Call(m, "ParseUnregistrationFromReceipt", receipt)
	ret0, _ := ret[0].(*adapter.DataTypeRegistryUnregistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseUnregistrationFromReceipt indicates an expected call of ParseUnregistrationFromReceipt
func (mr *MockIDataTypeRegistryContractMockRecorder) ParseUnregistrationFromReceipt(receipt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUnregistrationFromReceipt", reflect.TypeOf((*MockIDataTypeRegistryContract)(nil).ParseUnregistrationFromReceipt), receipt)
}