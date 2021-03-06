// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/imageReferencePort.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/docker-generator/api/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockImageReferenceRepository is a mock of ImageReferenceRepository interface.
type MockImageReferenceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockImageReferenceRepositoryMockRecorder
}

// MockImageReferenceRepositoryMockRecorder is the mock recorder for MockImageReferenceRepository.
type MockImageReferenceRepositoryMockRecorder struct {
	mock *MockImageReferenceRepository
}

// NewMockImageReferenceRepository creates a new mock instance.
func NewMockImageReferenceRepository(ctrl *gomock.Controller) *MockImageReferenceRepository {
	mock := &MockImageReferenceRepository{ctrl: ctrl}
	mock.recorder = &MockImageReferenceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageReferenceRepository) EXPECT() *MockImageReferenceRepositoryMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockImageReferenceRepository) Read(imageName string) (domain.ImageReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", imageName)
	ret0, _ := ret[0].(domain.ImageReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockImageReferenceRepositoryMockRecorder) Read(imageName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockImageReferenceRepository)(nil).Read), imageName)
}

// MockImageReferenceService is a mock of ImageReferenceService interface.
type MockImageReferenceService struct {
	ctrl     *gomock.Controller
	recorder *MockImageReferenceServiceMockRecorder
}

// MockImageReferenceServiceMockRecorder is the mock recorder for MockImageReferenceService.
type MockImageReferenceServiceMockRecorder struct {
	mock *MockImageReferenceService
}

// NewMockImageReferenceService creates a new mock instance.
func NewMockImageReferenceService(ctrl *gomock.Controller) *MockImageReferenceService {
	mock := &MockImageReferenceService{ctrl: ctrl}
	mock.recorder = &MockImageReferenceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageReferenceService) EXPECT() *MockImageReferenceServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockImageReferenceService) Get(imageName string) (domain.ImageReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", imageName)
	ret0, _ := ret[0].(domain.ImageReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockImageReferenceServiceMockRecorder) Get(imageName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockImageReferenceService)(nil).Get), imageName)
}
