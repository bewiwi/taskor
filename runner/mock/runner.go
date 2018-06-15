// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scaleway/taskor/runner (interfaces: Runner)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	task "github.com/scaleway/taskor/task"
	reflect "reflect"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockRunner) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRunnerMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRunner)(nil).Init))
}

// RunWorkerTaskAck mocks base method
func (m *MockRunner) RunWorkerTaskAck(arg0 <-chan task.Task) {
	m.ctrl.Call(m, "RunWorkerTaskAck", arg0)
}

// RunWorkerTaskAck indicates an expected call of RunWorkerTaskAck
func (mr *MockRunnerMockRecorder) RunWorkerTaskAck(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWorkerTaskAck", reflect.TypeOf((*MockRunner)(nil).RunWorkerTaskAck), arg0)
}

// RunWorkerTaskProvider mocks base method
func (m *MockRunner) RunWorkerTaskProvider(arg0 chan task.Task, arg1 <-chan bool) error {
	ret := m.ctrl.Call(m, "RunWorkerTaskProvider", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWorkerTaskProvider indicates an expected call of RunWorkerTaskProvider
func (mr *MockRunnerMockRecorder) RunWorkerTaskProvider(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWorkerTaskProvider", reflect.TypeOf((*MockRunner)(nil).RunWorkerTaskProvider), arg0, arg1)
}

// Send mocks base method
func (m *MockRunner) Send(arg0 *task.Task) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockRunnerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockRunner)(nil).Send), arg0)
}

// Stop mocks base method
func (m *MockRunner) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockRunnerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRunner)(nil).Stop))
}
