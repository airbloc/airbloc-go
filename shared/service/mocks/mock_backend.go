// Code generated by MockGen. DO NOT EDIT.
// Source: backend.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	blockchain "github.com/airbloc/airbloc-go/shared/blockchain"
	localdb "github.com/airbloc/airbloc-go/shared/database/localdb"
	metadb "github.com/airbloc/airbloc-go/shared/database/metadb"
	key "github.com/airbloc/airbloc-go/shared/key"
	p2p "github.com/airbloc/airbloc-go/shared/p2p"
	service "github.com/airbloc/airbloc-go/shared/service"
	gomock "github.com/golang/mock/gomock"
)

// MockBackend is a mock of Backend interface
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// Kms mocks base method
func (m *MockBackend) Kms() key.Manager {
	ret := m.ctrl.Call(m, "Kms")
	ret0, _ := ret[0].(key.Manager)
	return ret0
}

// Kms indicates an expected call of Kms
func (mr *MockBackendMockRecorder) Kms() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kms", reflect.TypeOf((*MockBackend)(nil).Kms))
}

// Client mocks base method
func (m *MockBackend) Client() *blockchain.Client {
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*blockchain.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockBackendMockRecorder) Client() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockBackend)(nil).Client))
}

// MetaDatabase mocks base method
func (m *MockBackend) MetaDatabase() metadb.Database {
	ret := m.ctrl.Call(m, "MetaDatabase")
	ret0, _ := ret[0].(metadb.Database)
	return ret0
}

// MetaDatabase indicates an expected call of MetaDatabase
func (mr *MockBackendMockRecorder) MetaDatabase() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetaDatabase", reflect.TypeOf((*MockBackend)(nil).MetaDatabase))
}

// LocalDatabase mocks base method
func (m *MockBackend) LocalDatabase() localdb.Database {
	ret := m.ctrl.Call(m, "LocalDatabase")
	ret0, _ := ret[0].(localdb.Database)
	return ret0
}

// LocalDatabase indicates an expected call of LocalDatabase
func (mr *MockBackendMockRecorder) LocalDatabase() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalDatabase", reflect.TypeOf((*MockBackend)(nil).LocalDatabase))
}

// Config mocks base method
func (m *MockBackend) Config() *service.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*service.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockBackendMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockBackend)(nil).Config))
}

// P2P mocks base method
func (m *MockBackend) P2P() p2p.Server {
	ret := m.ctrl.Call(m, "P2P")
	ret0, _ := ret[0].(p2p.Server)
	return ret0
}

// P2P indicates an expected call of P2P
func (mr *MockBackendMockRecorder) P2P() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "P2P", reflect.TypeOf((*MockBackend)(nil).P2P))
}

// Start mocks base method
func (m *MockBackend) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBackendMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBackend)(nil).Start))
}

// Stop mocks base method
func (m *MockBackend) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockBackendMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBackend)(nil).Stop))
}

// GetService mocks base method
func (m *MockBackend) GetService(arg0 string) service.Service {
	ret := m.ctrl.Call(m, "GetService", arg0)
	ret0, _ := ret[0].(service.Service)
	return ret0
}

// GetService indicates an expected call of GetService
func (mr *MockBackendMockRecorder) GetService(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockBackend)(nil).GetService), arg0)
}

// AttachService mocks base method
func (m *MockBackend) AttachService(arg0 string, arg1 service.Service) {
	m.ctrl.Call(m, "AttachService", arg0, arg1)
}

// AttachService indicates an expected call of AttachService
func (mr *MockBackendMockRecorder) AttachService(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachService", reflect.TypeOf((*MockBackend)(nil).AttachService), arg0, arg1)
}

// DetachService mocks base method
func (m *MockBackend) DetachService(arg0 string) {
	m.ctrl.Call(m, "DetachService", arg0)
}

// DetachService indicates an expected call of DetachService
func (mr *MockBackendMockRecorder) DetachService(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachService", reflect.TypeOf((*MockBackend)(nil).DetachService), arg0)
}
