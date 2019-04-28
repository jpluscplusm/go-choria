// Code generated by MockGen. DO NOT EDIT.
// Source: machine.go

// Package machine is a generated GoMock package.
package machine

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	sync "sync"
)

// MockWatcherManager is a mock of WatcherManager interface
type MockWatcherManager struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherManagerMockRecorder
}

// MockWatcherManagerMockRecorder is the mock recorder for MockWatcherManager
type MockWatcherManagerMockRecorder struct {
	mock *MockWatcherManager
}

// NewMockWatcherManager creates a new mock instance
func NewMockWatcherManager(ctrl *gomock.Controller) *MockWatcherManager {
	mock := &MockWatcherManager{ctrl: ctrl}
	mock.recorder = &MockWatcherManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatcherManager) EXPECT() *MockWatcherManagerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockWatcherManager) Run(arg0 context.Context, arg1 *sync.WaitGroup) error {
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockWatcherManagerMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWatcherManager)(nil).Run), arg0, arg1)
}

// NotifyStateChance mocks base method
func (m *MockWatcherManager) NotifyStateChance() {
	m.ctrl.Call(m, "NotifyStateChance")
}

// NotifyStateChance indicates an expected call of NotifyStateChance
func (mr *MockWatcherManagerMockRecorder) NotifyStateChance() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyStateChance", reflect.TypeOf((*MockWatcherManager)(nil).NotifyStateChance))
}

// SetMachine mocks base method
func (m *MockWatcherManager) SetMachine(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SetMachine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMachine indicates an expected call of SetMachine
func (mr *MockWatcherManagerMockRecorder) SetMachine(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMachine", reflect.TypeOf((*MockWatcherManager)(nil).SetMachine), arg0)
}
