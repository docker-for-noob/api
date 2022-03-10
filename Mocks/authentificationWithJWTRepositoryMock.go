// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/authentification/authentificationWithJWTPorts.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/docker-generator/api/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthentificationWithJWTService is a mock of AuthentificationWithJWTService interface.
type MockAuthentificationWithJWTService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthentificationWithJWTServiceMockRecorder
}

// MockAuthentificationWithJWTServiceMockRecorder is the mock recorder for MockAuthentificationWithJWTService.
type MockAuthentificationWithJWTServiceMockRecorder struct {
	mock *MockAuthentificationWithJWTService
}

// NewMockAuthentificationWithJWTService creates a new mock instance.
func NewMockAuthentificationWithJWTService(ctrl *gomock.Controller) *MockAuthentificationWithJWTService {
	mock := &MockAuthentificationWithJWTService{ctrl: ctrl}
	mock.recorder = &MockAuthentificationWithJWTServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthentificationWithJWTService) EXPECT() *MockAuthentificationWithJWTServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthentificationWithJWTService) Login(credentials domain.Credentials) (domain.JwtToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", credentials)
	ret0, _ := ret[0].(domain.JwtToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthentificationWithJWTServiceMockRecorder) Login(credentials interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthentificationWithJWTService)(nil).Login), credentials)
}

// Logout mocks base method.
func (m *MockAuthentificationWithJWTService) Logout(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthentificationWithJWTServiceMockRecorder) Logout(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthentificationWithJWTService)(nil).Logout), id)
}