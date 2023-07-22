// Code generated by MockGen. DO NOT EDIT.
// Source: farm.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
)

// MockIFarmRepository is a mock of IFarmRepository interface.
type MockIFarmRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFarmRepositoryMockRecorder
}

// MockIFarmRepositoryMockRecorder is the mock recorder for MockIFarmRepository.
type MockIFarmRepositoryMockRecorder struct {
	mock *MockIFarmRepository
}

// NewMockIFarmRepository creates a new mock instance.
func NewMockIFarmRepository(ctrl *gomock.Controller) *MockIFarmRepository {
	mock := &MockIFarmRepository{ctrl: ctrl}
	mock.recorder = &MockIFarmRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFarmRepository) EXPECT() *MockIFarmRepositoryMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockIFarmRepository) All() ([]*models.Farm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]*models.Farm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockIFarmRepositoryMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIFarmRepository)(nil).All))
}

// ExistsByID mocks base method.
func (m *MockIFarmRepository) ExistsByID(id uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByID", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByID indicates an expected call of ExistsByID.
func (mr *MockIFarmRepositoryMockRecorder) ExistsByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByID", reflect.TypeOf((*MockIFarmRepository)(nil).ExistsByID), id)
}

// GetByID mocks base method.
func (m *MockIFarmRepository) GetByID(id uint) (*models.Farm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*models.Farm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIFarmRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIFarmRepository)(nil).GetByID), id)
}