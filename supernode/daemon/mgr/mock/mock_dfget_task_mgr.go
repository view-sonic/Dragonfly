// Code generated by MockGen. DO NOT EDIT.
// Source: supernode/daemon/mgr/dfget_task_mgr.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/dragonflyoss/Dragonfly/apis/types"
)

// MockDfgetTaskMgr is a mock of DfgetTaskMgr interface
type MockDfgetTaskMgr struct {
	ctrl     *gomock.Controller
	recorder *MockDfgetTaskMgrMockRecorder
}

// MockDfgetTaskMgrMockRecorder is the mock recorder for MockDfgetTaskMgr
type MockDfgetTaskMgrMockRecorder struct {
	mock *MockDfgetTaskMgr
}

// NewMockDfgetTaskMgr creates a new mock instance
func NewMockDfgetTaskMgr(ctrl *gomock.Controller) *MockDfgetTaskMgr {
	mock := &MockDfgetTaskMgr{ctrl: ctrl}
	mock.recorder = &MockDfgetTaskMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDfgetTaskMgr) EXPECT() *MockDfgetTaskMgrMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockDfgetTaskMgr) Add(ctx context.Context, dfgetTask *types.DfGetTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, dfgetTask)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockDfgetTaskMgrMockRecorder) Add(ctx, dfgetTask interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDfgetTaskMgr)(nil).Add), ctx, dfgetTask)
}

// Get mocks base method
func (m *MockDfgetTaskMgr) Get(ctx context.Context, clientID, taskID string) (*types.DfGetTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, clientID, taskID)
	ret0, _ := ret[0].(*types.DfGetTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDfgetTaskMgrMockRecorder) Get(ctx, clientID, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDfgetTaskMgr)(nil).Get), ctx, clientID, taskID)
}

// List mocks base method
func (m *MockDfgetTaskMgr) List(ctx context.Context, filter map[string]string) ([]*types.DfGetTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filter)
	ret0, _ := ret[0].([]*types.DfGetTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDfgetTaskMgrMockRecorder) List(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDfgetTaskMgr)(nil).List), ctx, filter)
}

// Delete mocks base method
func (m *MockDfgetTaskMgr) Delete(ctx context.Context, clientID, taskID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, clientID, taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDfgetTaskMgrMockRecorder) Delete(ctx, clientID, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDfgetTaskMgr)(nil).Delete), ctx, clientID, taskID)
}

// UpdateStatus mocks base method
func (m *MockDfgetTaskMgr) UpdateStatus(ctx context.Context, clientID, taskID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, clientID, taskID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockDfgetTaskMgrMockRecorder) UpdateStatus(ctx, clientID, taskID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockDfgetTaskMgr)(nil).UpdateStatus), ctx, clientID, taskID, status)
}
