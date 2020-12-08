// Code generated by MockGen. DO NOT EDIT.
// Source: watchers.go

// Package watchers is a generated GoMock package.
package watchers

import (
	context "context"
	watcher "github.com/choria-io/go-choria/aagent/watchers/watcher"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	sync "sync"
	time "time"
)

// MockWatcher is a mock of Watcher interface
type MockWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherMockRecorder
}

// MockWatcherMockRecorder is the mock recorder for MockWatcher
type MockWatcherMockRecorder struct {
	mock *MockWatcher
}

// NewMockWatcher creates a new mock instance
func NewMockWatcher(ctrl *gomock.Controller) *MockWatcher {
	mock := &MockWatcher{ctrl: ctrl}
	mock.recorder = &MockWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatcher) EXPECT() *MockWatcherMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockWatcher) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockWatcherMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockWatcher)(nil).Name))
}

// Type mocks base method
func (m *MockWatcher) Type() string {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockWatcherMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockWatcher)(nil).Type))
}

// Run mocks base method
func (m *MockWatcher) Run(arg0 context.Context, arg1 *sync.WaitGroup) {
	m.ctrl.Call(m, "Run", arg0, arg1)
}

// Run indicates an expected call of Run
func (mr *MockWatcherMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWatcher)(nil).Run), arg0, arg1)
}

// NotifyStateChance mocks base method
func (m *MockWatcher) NotifyStateChance() {
	m.ctrl.Call(m, "NotifyStateChance")
}

// NotifyStateChance indicates an expected call of NotifyStateChance
func (mr *MockWatcherMockRecorder) NotifyStateChance() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyStateChance", reflect.TypeOf((*MockWatcher)(nil).NotifyStateChance))
}

// CurrentState mocks base method
func (m *MockWatcher) CurrentState() interface{} {
	ret := m.ctrl.Call(m, "CurrentState")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// CurrentState indicates an expected call of CurrentState
func (mr *MockWatcherMockRecorder) CurrentState() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentState", reflect.TypeOf((*MockWatcher)(nil).CurrentState))
}

// AnnounceInterval mocks base method
func (m *MockWatcher) AnnounceInterval() time.Duration {
	ret := m.ctrl.Call(m, "AnnounceInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// AnnounceInterval indicates an expected call of AnnounceInterval
func (mr *MockWatcherMockRecorder) AnnounceInterval() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceInterval", reflect.TypeOf((*MockWatcher)(nil).AnnounceInterval))
}

// Delete mocks base method
func (m *MockWatcher) Delete() {
	m.ctrl.Call(m, "Delete")
}

// Delete indicates an expected call of Delete
func (mr *MockWatcherMockRecorder) Delete() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWatcher)(nil).Delete))
}

// MockMachine is a mock of Machine interface
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockMachine) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockMachineMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMachine)(nil).Name))
}

// State mocks base method
func (m *MockMachine) State() string {
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(string)
	return ret0
}

// State indicates an expected call of State
func (mr *MockMachineMockRecorder) State() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockMachine)(nil).State))
}

// Directory mocks base method
func (m *MockMachine) Directory() string {
	ret := m.ctrl.Call(m, "Directory")
	ret0, _ := ret[0].(string)
	return ret0
}

// Directory indicates an expected call of Directory
func (mr *MockMachineMockRecorder) Directory() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Directory", reflect.TypeOf((*MockMachine)(nil).Directory))
}

// Transition mocks base method
func (m *MockMachine) Transition(t string, args ...interface{}) error {
	varargs := []interface{}{t}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transition indicates an expected call of Transition
func (mr *MockMachineMockRecorder) Transition(t interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{t}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transition", reflect.TypeOf((*MockMachine)(nil).Transition), varargs...)
}

// NotifyWatcherState mocks base method
func (m *MockMachine) NotifyWatcherState(arg0 string, arg1 interface{}) {
	m.ctrl.Call(m, "NotifyWatcherState", arg0, arg1)
}

// NotifyWatcherState indicates an expected call of NotifyWatcherState
func (mr *MockMachineMockRecorder) NotifyWatcherState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWatcherState", reflect.TypeOf((*MockMachine)(nil).NotifyWatcherState), arg0, arg1)
}

// Watchers mocks base method
func (m *MockMachine) Watchers() []*WatcherDef {
	ret := m.ctrl.Call(m, "Watchers")
	ret0, _ := ret[0].([]*WatcherDef)
	return ret0
}

// Watchers indicates an expected call of Watchers
func (mr *MockMachineMockRecorder) Watchers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watchers", reflect.TypeOf((*MockMachine)(nil).Watchers))
}

// Identity mocks base method
func (m *MockMachine) Identity() string {
	ret := m.ctrl.Call(m, "Identity")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identity indicates an expected call of Identity
func (mr *MockMachineMockRecorder) Identity() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockMachine)(nil).Identity))
}

// InstanceID mocks base method
func (m *MockMachine) InstanceID() string {
	ret := m.ctrl.Call(m, "InstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// InstanceID indicates an expected call of InstanceID
func (mr *MockMachineMockRecorder) InstanceID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockMachine)(nil).InstanceID))
}

// Version mocks base method
func (m *MockMachine) Version() string {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockMachineMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockMachine)(nil).Version))
}

// TimeStampSeconds mocks base method
func (m *MockMachine) TimeStampSeconds() int64 {
	ret := m.ctrl.Call(m, "TimeStampSeconds")
	ret0, _ := ret[0].(int64)
	return ret0
}

// TimeStampSeconds indicates an expected call of TimeStampSeconds
func (mr *MockMachineMockRecorder) TimeStampSeconds() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeStampSeconds", reflect.TypeOf((*MockMachine)(nil).TimeStampSeconds))
}

// TextFileDirectory mocks base method
func (m *MockMachine) TextFileDirectory() string {
	ret := m.ctrl.Call(m, "TextFileDirectory")
	ret0, _ := ret[0].(string)
	return ret0
}

// TextFileDirectory indicates an expected call of TextFileDirectory
func (mr *MockMachineMockRecorder) TextFileDirectory() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TextFileDirectory", reflect.TypeOf((*MockMachine)(nil).TextFileDirectory))
}

// OverrideData mocks base method
func (m *MockMachine) OverrideData() ([]byte, error) {
	ret := m.ctrl.Call(m, "OverrideData")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OverrideData indicates an expected call of OverrideData
func (mr *MockMachineMockRecorder) OverrideData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverrideData", reflect.TypeOf((*MockMachine)(nil).OverrideData))
}

// Debugf mocks base method
func (m *MockMachine) Debugf(name, format string, args ...interface{}) {
	varargs := []interface{}{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockMachineMockRecorder) Debugf(name, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockMachine)(nil).Debugf), varargs...)
}

// Infof mocks base method
func (m *MockMachine) Infof(name, format string, args ...interface{}) {
	varargs := []interface{}{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockMachineMockRecorder) Infof(name, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockMachine)(nil).Infof), varargs...)
}

// Errorf mocks base method
func (m *MockMachine) Errorf(name, format string, args ...interface{}) {
	varargs := []interface{}{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockMachineMockRecorder) Errorf(name, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockMachine)(nil).Errorf), varargs...)
}

// MockWatcherConstructor is a mock of WatcherConstructor interface
type MockWatcherConstructor struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherConstructorMockRecorder
}

// MockWatcherConstructorMockRecorder is the mock recorder for MockWatcherConstructor
type MockWatcherConstructorMockRecorder struct {
	mock *MockWatcherConstructor
}

// NewMockWatcherConstructor creates a new mock instance
func NewMockWatcherConstructor(ctrl *gomock.Controller) *MockWatcherConstructor {
	mock := &MockWatcherConstructor{ctrl: ctrl}
	mock.recorder = &MockWatcherConstructorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatcherConstructor) EXPECT() *MockWatcherConstructorMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockWatcherConstructor) New(machine watcher.Machine, name string, states []string, failEvent, successEvent, interval string, ai time.Duration, properties map[string]interface{}) (interface{}, error) {
	ret := m.ctrl.Call(m, "New", machine, name, states, failEvent, successEvent, interval, ai, properties)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockWatcherConstructorMockRecorder) New(machine, name, states, failEvent, successEvent, interval, ai, properties interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockWatcherConstructor)(nil).New), machine, name, states, failEvent, successEvent, interval, ai, properties)
}

// Type mocks base method
func (m *MockWatcherConstructor) Type() string {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockWatcherConstructorMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockWatcherConstructor)(nil).Type))
}

// EventType mocks base method
func (m *MockWatcherConstructor) EventType() string {
	ret := m.ctrl.Call(m, "EventType")
	ret0, _ := ret[0].(string)
	return ret0
}

// EventType indicates an expected call of EventType
func (mr *MockWatcherConstructorMockRecorder) EventType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventType", reflect.TypeOf((*MockWatcherConstructor)(nil).EventType))
}

// UnmarshalNotification mocks base method
func (m *MockWatcherConstructor) UnmarshalNotification(n []byte) (interface{}, error) {
	ret := m.ctrl.Call(m, "UnmarshalNotification", n)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnmarshalNotification indicates an expected call of UnmarshalNotification
func (mr *MockWatcherConstructorMockRecorder) UnmarshalNotification(n interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalNotification", reflect.TypeOf((*MockWatcherConstructor)(nil).UnmarshalNotification), n)
}
