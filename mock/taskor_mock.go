// Code generated by MockGen. DO NOT EDIT.
// Source: taskor.go

// Package mock_taskor is a generated GoMock package.
package mock_taskor

import (
	gomock "github.com/golang/mock/gomock"
	task "github.com/scaleway/taskor/task"
	reflect "reflect"
)

// MockTaskManager is a mock of TaskManager interface
type MockTaskManager struct {
	ctrl     *gomock.Controller
	recorder *MockTaskManagerMockRecorder
}

// MockTaskManagerMockRecorder is the mock recorder for MockTaskManager
type MockTaskManagerMockRecorder struct {
	mock *MockTaskManager
}

// NewMockTaskManager creates a new mock instance
func NewMockTaskManager(ctrl *gomock.Controller) *MockTaskManager {
	mock := &MockTaskManager{ctrl: ctrl}
	mock.recorder = &MockTaskManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskManager) EXPECT() *MockTaskManagerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockTaskManager) Send(task *task.Task) error {
	ret := m.ctrl.Call(m, "Send", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockTaskManagerMockRecorder) Send(task interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTaskManager)(nil).Send), task)
}

// Handle mocks base method
func (m *MockTaskManager) Handle(Definition *task.Definition) error {
	ret := m.ctrl.Call(m, "Handle", Definition)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle
func (mr *MockTaskManagerMockRecorder) Handle(Definition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockTaskManager)(nil).Handle), Definition)
}

// GetHandled mocks base method
func (m *MockTaskManager) GetHandled() []*task.Definition {
	ret := m.ctrl.Call(m, "GetHandled")
	ret0, _ := ret[0].([]*task.Definition)
	return ret0
}

// GetHandled indicates an expected call of GetHandled
func (mr *MockTaskManagerMockRecorder) GetHandled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandled", reflect.TypeOf((*MockTaskManager)(nil).GetHandled))
}

// RunWorker mocks base method
func (m *MockTaskManager) RunWorker() error {
	ret := m.ctrl.Call(m, "RunWorker")
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWorker indicates an expected call of RunWorker
func (mr *MockTaskManagerMockRecorder) RunWorker() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWorker", reflect.TypeOf((*MockTaskManager)(nil).RunWorker))
}

// StopWorker mocks base method
func (m *MockTaskManager) StopWorker() {
	m.ctrl.Call(m, "StopWorker")
}

// StopWorker indicates an expected call of StopWorker
func (mr *MockTaskManagerMockRecorder) StopWorker() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopWorker", reflect.TypeOf((*MockTaskManager)(nil).StopWorker))
}
