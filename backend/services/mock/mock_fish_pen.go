// Code generated by MockGen. DO NOT EDIT.
// Source: fish_pen.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
)

// MockIFishPenService is a mock of IFishPenService interface.
type MockIFishPenService struct {
	ctrl     *gomock.Controller
	recorder *MockIFishPenServiceMockRecorder
}

// MockIFishPenServiceMockRecorder is the mock recorder for MockIFishPenService.
type MockIFishPenServiceMockRecorder struct {
	mock *MockIFishPenService
}

// NewMockIFishPenService creates a new mock instance.
func NewMockIFishPenService(ctrl *gomock.Controller) *MockIFishPenService {
	mock := &MockIFishPenService{ctrl: ctrl}
	mock.recorder = &MockIFishPenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFishPenService) EXPECT() *MockIFishPenServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIFishPenService) Create(farmID uint, fishPen *models.FishPen) (*models.FishPen, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", farmID, fishPen)
	ret0, _ := ret[0].(*models.FishPen)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIFishPenServiceMockRecorder) Create(farmID, fishPen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIFishPenService)(nil).Create), farmID, fishPen)
}

// Delete mocks base method.
func (m *MockIFishPenService) Delete(farmID, fishPenID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", farmID, fishPenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIFishPenServiceMockRecorder) Delete(farmID, fishPenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIFishPenService)(nil).Delete), farmID, fishPenID)
}

// GetAllForFarm mocks base method.
func (m *MockIFishPenService) GetAllForFarm(farmID uint) ([]*models.FishPen, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllForFarm", farmID)
	ret0, _ := ret[0].([]*models.FishPen)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllForFarm indicates an expected call of GetAllForFarm.
func (mr *MockIFishPenServiceMockRecorder) GetAllForFarm(farmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllForFarm", reflect.TypeOf((*MockIFishPenService)(nil).GetAllForFarm), farmID)
}

// GetSingle mocks base method.
func (m *MockIFishPenService) GetSingle(farmID, fishPenID uint) (*models.FishPen, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSingle", farmID, fishPenID)
	ret0, _ := ret[0].(*models.FishPen)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSingle indicates an expected call of GetSingle.
func (mr *MockIFishPenServiceMockRecorder) GetSingle(farmID, fishPenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingle", reflect.TypeOf((*MockIFishPenService)(nil).GetSingle), farmID, fishPenID)
}

// Update mocks base method.
func (m *MockIFishPenService) Update(farmID uint, fishPen *models.FishPen) (*models.FishPen, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", farmID, fishPen)
	ret0, _ := ret[0].(*models.FishPen)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIFishPenServiceMockRecorder) Update(farmID, fishPen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIFishPenService)(nil).Update), farmID, fishPen)
}
