// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/v2/service (interfaces: FunctionService)

// Package service is a generated GoMock package.
package service

import (
	models "github.com/baetyl/baetyl-cloud/v2/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFunctionService is a mock of FunctionService interface
type MockFunctionService struct {
	ctrl     *gomock.Controller
	recorder *MockFunctionServiceMockRecorder
}

// MockFunctionServiceMockRecorder is the mock recorder for MockFunctionService
type MockFunctionServiceMockRecorder struct {
	mock *MockFunctionService
}

// NewMockFunctionService creates a new mock instance
func NewMockFunctionService(ctrl *gomock.Controller) *MockFunctionService {
	mock := &MockFunctionService{ctrl: ctrl}
	mock.recorder = &MockFunctionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFunctionService) EXPECT() *MockFunctionServiceMockRecorder {
	return m.recorder
}

// GetFunction mocks base method
func (m *MockFunctionService) GetFunction(arg0, arg1, arg2, arg3 string) (*models.Function, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFunction", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*models.Function)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunction indicates an expected call of GetFunction
func (mr *MockFunctionServiceMockRecorder) GetFunction(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunction", reflect.TypeOf((*MockFunctionService)(nil).GetFunction), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockFunctionService) List(arg0, arg1 string) ([]models.Function, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]models.Function)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockFunctionServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFunctionService)(nil).List), arg0, arg1)
}

// ListFunctionVersions mocks base method
func (m *MockFunctionService) ListFunctionVersions(arg0, arg1, arg2 string) ([]models.Function, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFunctionVersions", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Function)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctionVersions indicates an expected call of ListFunctionVersions
func (mr *MockFunctionServiceMockRecorder) ListFunctionVersions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctionVersions", reflect.TypeOf((*MockFunctionService)(nil).ListFunctionVersions), arg0, arg1, arg2)
}

// ListSources mocks base method
func (m *MockFunctionService) ListSources() []models.FunctionSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSources")
	ret0, _ := ret[0].([]models.FunctionSource)
	return ret0
}

// ListSources indicates an expected call of ListSources
func (mr *MockFunctionServiceMockRecorder) ListSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSources", reflect.TypeOf((*MockFunctionService)(nil).ListSources))
}
