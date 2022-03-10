// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/security/JWTPorts.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/docker-generator/api/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockJWTRepository is a mock of JWTRepository interface.
type MockJWTRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJWTRepositoryMockRecorder
}

// MockJWTRepositoryMockRecorder is the mock recorder for MockJWTRepository.
type MockJWTRepositoryMockRecorder struct {
	mock *MockJWTRepository
}

// NewMockJWTRepository creates a new mock instance.
func NewMockJWTRepository(ctrl *gomock.Controller) *MockJWTRepository {
	mock := &MockJWTRepository{ctrl: ctrl}
	mock.recorder = &MockJWTRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTRepository) EXPECT() *MockJWTRepositoryMockRecorder {
	return m.recorder
}

// CreateJWTTokenString mocks base method.
func (m *MockJWTRepository) CreateJWTTokenString(user domain.User) (domain.JwtToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJWTTokenString", user)
	ret0, _ := ret[0].(domain.JwtToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJWTTokenString indicates an expected call of CreateJWTTokenString.
func (mr *MockJWTRepositoryMockRecorder) CreateJWTTokenString(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJWTTokenString", reflect.TypeOf((*MockJWTRepository)(nil).CreateJWTTokenString), user)
}