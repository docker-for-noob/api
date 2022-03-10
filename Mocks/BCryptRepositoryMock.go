// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/user/BCryptPorts.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBCryptRepository is a mock of BCryptRepository interface.
type MockBCryptRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBCryptRepositoryMockRecorder
}

// MockBCryptRepositoryMockRecorder is the mock recorder for MockBCryptRepository.
type MockBCryptRepositoryMockRecorder struct {
	mock *MockBCryptRepository
}

// NewMockBCryptRepository creates a new mock instance.
func NewMockBCryptRepository(ctrl *gomock.Controller) *MockBCryptRepository {
	mock := &MockBCryptRepository{ctrl: ctrl}
	mock.recorder = &MockBCryptRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBCryptRepository) EXPECT() *MockBCryptRepositoryMockRecorder {
	return m.recorder
}

// CheckPasswordHash mocks base method.
func (m *MockBCryptRepository) CheckPasswordHash(password, hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordHash", password, hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckPasswordHash indicates an expected call of CheckPasswordHash.
func (mr *MockBCryptRepositoryMockRecorder) CheckPasswordHash(password, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordHash", reflect.TypeOf((*MockBCryptRepository)(nil).CheckPasswordHash), password, hash)
}

// HashPassword mocks base method.
func (m *MockBCryptRepository) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockBCryptRepositoryMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockBCryptRepository)(nil).HashPassword), password)
}