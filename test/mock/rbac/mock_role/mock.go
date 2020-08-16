// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/data/rbac/role/interface.go

// Package mock_role is a generated GoMock package.
package mock_role

import (
	response "github.com/go-liam/util/response"
	gomock "github.com/golang/mock/gomock"
	role "grape/internal/pkg/data/rbac/role"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockService) Create(item *role.Model) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", item)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServiceMockRecorder) Create(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), item)
}

// FindOne mocks base method
func (m *MockService) FindOne(id int64) (*role.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", id)
	ret0, _ := ret[0].(*role.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockServiceMockRecorder) FindOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockService)(nil).FindOne), id)
}

// FindMulti mocks base method
func (m *MockService) FindMulti(page *response.Pagination, s *response.ListParameter) ([]*role.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMulti", page, s)
	ret0, _ := ret[0].([]*role.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMulti indicates an expected call of FindMulti
func (mr *MockServiceMockRecorder) FindMulti(page, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMulti", reflect.TypeOf((*MockService)(nil).FindMulti), page, s)
}

// Update mocks base method
func (m *MockService) Update(item *role.Model) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", item)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockServiceMockRecorder) Update(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), item)
}

// UpdateStatusByIDs mocks base method
func (m *MockService) UpdateStatusByIDs(status int, ls []int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusByIDs", status, ls)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatusByIDs indicates an expected call of UpdateStatusByIDs
func (mr *MockServiceMockRecorder) UpdateStatusByIDs(status, ls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusByIDs", reflect.TypeOf((*MockService)(nil).UpdateStatusByIDs), status, ls)
}

// CacheMulti mocks base method
func (m *MockService) CacheMulti() ([]*role.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheMulti")
	ret0, _ := ret[0].([]*role.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CacheMulti indicates an expected call of CacheMulti
func (mr *MockServiceMockRecorder) CacheMulti() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheMulti", reflect.TypeOf((*MockService)(nil).CacheMulti))
}
