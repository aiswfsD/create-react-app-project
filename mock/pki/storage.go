// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-go/v2/pki (interfaces: Storage)

// Package pki is a generated GoMock package.
package pki

import (
	models "github.com/baetyl/baetyl-go/v2/pki/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// CountCertByParentId mocks base method
func (m *MockStorage) CountCertByParentId(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountCertByParentId", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountCertByParentId indicates an expected call of CountCertByParentId
func (mr *MockStorageMockRecorder) CountCertByParentId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountCertByParentId", reflect.TypeOf((*MockStorage)(nil).CountCertByParentId), arg0)
}

// CreateCert mocks base method
func (m *MockStorage) CreateCert(arg0 models.Cert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCert indicates an expected call of CreateCert
func (mr *MockStorageMockRecorder) CreateCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCert", reflect.TypeOf((*MockStorage)(nil).CreateCert), arg0)
}

// DeleteCert mocks base method
func (m *MockStorage) DeleteCert(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCert indicates an expected call of DeleteCert
func (mr *MockStorageMockRecorder) DeleteCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCert", reflect.TypeOf((*MockStorage)(nil).DeleteCert), arg0)
}

// GetSubCert mocks base method
func (m *MockStorage) GetCert(arg0 string) (*models.Cert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubCert", arg0)
	ret0, _ := ret[0].(*models.Cert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubCert indicates an expected call of GetSubCert
func (mr *MockStorageMockRecorder) GetCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubCert", reflect.TypeOf((*MockStorage)(nil).GetCert), arg0)
}

// UpdateCert mocks base method
func (m *MockStorage) UpdateCert(arg0 models.Cert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCert indicates an expected call of UpdateCert
func (mr *MockStorageMockRecorder) UpdateCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCert", reflect.TypeOf((*MockStorage)(nil).UpdateCert), arg0)
}
