// Code generated by MockGen. DO NOT EDIT.
// Source: model/actions.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	timer "github.com/charmbracelet/bubbles/timer"
	bubbletea "github.com/charmbracelet/bubbletea"
	gomock "github.com/golang/mock/gomock"
)

// MockModelHandler is a mock of ModelHandler interface.
type MockModelHandler struct {
	ctrl     *gomock.Controller
	recorder *MockModelHandlerMockRecorder
}

// MockModelHandlerMockRecorder is the mock recorder for MockModelHandler.
type MockModelHandlerMockRecorder struct {
	mock *MockModelHandler
}

// NewMockModelHandler creates a new mock instance.
func NewMockModelHandler(ctrl *gomock.Controller) *MockModelHandler {
	mock := &MockModelHandler{ctrl: ctrl}
	mock.recorder = &MockModelHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelHandler) EXPECT() *MockModelHandlerMockRecorder {
	return m.recorder
}

// HandleConfirm mocks base method.
func (m *MockModelHandler) HandleConfirm(arg0 bubbletea.Model, arg1 bubbletea.Cmd) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleConfirm", arg0, arg1)
}

// HandleConfirm indicates an expected call of HandleConfirm.
func (mr *MockModelHandlerMockRecorder) HandleConfirm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleConfirm", reflect.TypeOf((*MockModelHandler)(nil).HandleConfirm), arg0, arg1)
}

// HandleContinue mocks base method.
func (m *MockModelHandler) HandleContinue(arg0 bubbletea.Model, arg1 bubbletea.Cmd) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleContinue", arg0, arg1)
}

// HandleContinue indicates an expected call of HandleContinue.
func (mr *MockModelHandlerMockRecorder) HandleContinue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleContinue", reflect.TypeOf((*MockModelHandler)(nil).HandleContinue), arg0, arg1)
}

// HandleDown mocks base method.
func (m *MockModelHandler) HandleDown(arg0 bubbletea.Model, arg1 bubbletea.Cmd) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleDown", arg0, arg1)
}

// HandleDown indicates an expected call of HandleDown.
func (mr *MockModelHandlerMockRecorder) HandleDown(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDown", reflect.TypeOf((*MockModelHandler)(nil).HandleDown), arg0, arg1)
}

// HandleEnter mocks base method.
func (m *MockModelHandler) HandleEnter(arg0 bubbletea.Model, arg1 bubbletea.Cmd) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleEnter", arg0, arg1)
}

// HandleEnter indicates an expected call of HandleEnter.
func (mr *MockModelHandlerMockRecorder) HandleEnter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEnter", reflect.TypeOf((*MockModelHandler)(nil).HandleEnter), arg0, arg1)
}

// HandleQuit mocks base method.
func (m *MockModelHandler) HandleQuit(arg0 bubbletea.Model, arg1 bubbletea.Cmd) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleQuit", arg0, arg1)
}

// HandleQuit indicates an expected call of HandleQuit.
func (mr *MockModelHandlerMockRecorder) HandleQuit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleQuit", reflect.TypeOf((*MockModelHandler)(nil).HandleQuit), arg0, arg1)
}

// HandleStartStop mocks base method.
func (m *MockModelHandler) HandleStartStop(arg0 bubbletea.Model, arg1 bubbletea.Cmd) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleStartStop", arg0, arg1)
}

// HandleStartStop indicates an expected call of HandleStartStop.
func (mr *MockModelHandlerMockRecorder) HandleStartStop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStartStop", reflect.TypeOf((*MockModelHandler)(nil).HandleStartStop), arg0, arg1)
}

// HandleTimerStartStopMsg mocks base method.
func (m *MockModelHandler) HandleTimerStartStopMsg(arg0 timer.StartStopMsg) (bubbletea.Model, bubbletea.Cmd) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTimerStartStopMsg", arg0)
	ret0, _ := ret[0].(bubbletea.Model)
	ret1, _ := ret[1].(bubbletea.Cmd)
	return ret0, ret1
}

// HandleTimerStartStopMsg indicates an expected call of HandleTimerStartStopMsg.
func (mr *MockModelHandlerMockRecorder) HandleTimerStartStopMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTimerStartStopMsg", reflect.TypeOf((*MockModelHandler)(nil).HandleTimerStartStopMsg), arg0)
}

// HandleTimerTickMsg mocks base method.
func (m *MockModelHandler) HandleTimerTickMsg(arg0 timer.TickMsg) (bubbletea.Model, bubbletea.Cmd) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTimerTickMsg", arg0)
	ret0, _ := ret[0].(bubbletea.Model)
	ret1, _ := ret[1].(bubbletea.Cmd)
	return ret0, ret1
}

// HandleTimerTickMsg indicates an expected call of HandleTimerTickMsg.
func (mr *MockModelHandlerMockRecorder) HandleTimerTickMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTimerTickMsg", reflect.TypeOf((*MockModelHandler)(nil).HandleTimerTickMsg), arg0)
}

// HandleTimerTimeout mocks base method.
func (m *MockModelHandler) HandleTimerTimeout() (bubbletea.Model, bubbletea.Cmd) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTimerTimeout")
	ret0, _ := ret[0].(bubbletea.Model)
	ret1, _ := ret[1].(bubbletea.Cmd)
	return ret0, ret1
}

// HandleTimerTimeout indicates an expected call of HandleTimerTimeout.
func (mr *MockModelHandlerMockRecorder) HandleTimerTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTimerTimeout", reflect.TypeOf((*MockModelHandler)(nil).HandleTimerTimeout))
}

// HandleUp mocks base method.
func (m *MockModelHandler) HandleUp(arg0 bubbletea.Model, arg1 bubbletea.Cmd) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleUp", arg0, arg1)
}

// HandleUp indicates an expected call of HandleUp.
func (mr *MockModelHandlerMockRecorder) HandleUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUp", reflect.TypeOf((*MockModelHandler)(nil).HandleUp), arg0, arg1)
}

// HandleUpdate mocks base method.
func (m *MockModelHandler) HandleUpdate(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpdate", msg)
	ret0, _ := ret[0].(bubbletea.Model)
	ret1, _ := ret[1].(bubbletea.Cmd)
	return ret0, ret1
}

// HandleUpdate indicates an expected call of HandleUpdate.
func (mr *MockModelHandlerMockRecorder) HandleUpdate(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockModelHandler)(nil).HandleUpdate), msg)
}
