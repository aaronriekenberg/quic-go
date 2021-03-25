// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lucas-clemente/quic-go/internal/congestion (interfaces: SendAlgorithmWithDebugInfos)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/lucas-clemente/quic-go/internal/protocol"
)

// MockSendAlgorithmWithDebugInfos is a mock of SendAlgorithmWithDebugInfos interface.
type MockSendAlgorithmWithDebugInfos struct {
	ctrl     *gomock.Controller
	recorder *MockSendAlgorithmWithDebugInfosMockRecorder
}

// MockSendAlgorithmWithDebugInfosMockRecorder is the mock recorder for MockSendAlgorithmWithDebugInfos.
type MockSendAlgorithmWithDebugInfosMockRecorder struct {
	mock *MockSendAlgorithmWithDebugInfos
}

// NewMockSendAlgorithmWithDebugInfos creates a new mock instance.
func NewMockSendAlgorithmWithDebugInfos(ctrl *gomock.Controller) *MockSendAlgorithmWithDebugInfos {
	mock := &MockSendAlgorithmWithDebugInfos{ctrl: ctrl}
	mock.recorder = &MockSendAlgorithmWithDebugInfosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendAlgorithmWithDebugInfos) EXPECT() *MockSendAlgorithmWithDebugInfosMockRecorder {
	return m.recorder
}

// CanSend mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) CanSend(arg0 protocol.ByteCount) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSend", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSend indicates an expected call of CanSend.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) CanSend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSend", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).CanSend), arg0)
}

// GetCongestionWindow mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) GetCongestionWindow() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCongestionWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetCongestionWindow indicates an expected call of GetCongestionWindow.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) GetCongestionWindow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCongestionWindow", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).GetCongestionWindow))
}

// HasPacingBudget mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) HasPacingBudget() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPacingBudget")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPacingBudget indicates an expected call of HasPacingBudget.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) HasPacingBudget() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPacingBudget", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).HasPacingBudget))
}

// InRecovery mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) InRecovery() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InRecovery")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InRecovery indicates an expected call of InRecovery.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) InRecovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InRecovery", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).InRecovery))
}

// InSlowStart mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) InSlowStart() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InSlowStart")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InSlowStart indicates an expected call of InSlowStart.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) InSlowStart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InSlowStart", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).InSlowStart))
}

// MaybeExitSlowStart mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) MaybeExitSlowStart() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MaybeExitSlowStart")
}

// MaybeExitSlowStart indicates an expected call of MaybeExitSlowStart.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) MaybeExitSlowStart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeExitSlowStart", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).MaybeExitSlowStart))
}

// OnPacketAcked mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnPacketAcked(arg0 protocol.PacketNumber, arg1, arg2 protocol.ByteCount, arg3 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPacketAcked", arg0, arg1, arg2, arg3)
}

// OnPacketAcked indicates an expected call of OnPacketAcked.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnPacketAcked(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketAcked", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnPacketAcked), arg0, arg1, arg2, arg3)
}

// OnPacketLost mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnPacketLost(arg0 protocol.PacketNumber, arg1, arg2 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPacketLost", arg0, arg1, arg2)
}

// OnPacketLost indicates an expected call of OnPacketLost.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnPacketLost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketLost", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnPacketLost), arg0, arg1, arg2)
}

// OnPacketSent mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnPacketSent(arg0 time.Time, arg1 protocol.ByteCount, arg2 protocol.PacketNumber, arg3 protocol.ByteCount, arg4 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPacketSent", arg0, arg1, arg2, arg3, arg4)
}

// OnPacketSent indicates an expected call of OnPacketSent.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnPacketSent(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketSent", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnPacketSent), arg0, arg1, arg2, arg3, arg4)
}

// OnRetransmissionTimeout mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnRetransmissionTimeout(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRetransmissionTimeout", arg0)
}

// OnRetransmissionTimeout indicates an expected call of OnRetransmissionTimeout.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnRetransmissionTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRetransmissionTimeout", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnRetransmissionTimeout), arg0)
}

// SetMaxDatagramSize mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) SetMaxDatagramSize(arg0 protocol.ByteCount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMaxDatagramSize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMaxDatagramSize indicates an expected call of SetMaxDatagramSize.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) SetMaxDatagramSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxDatagramSize", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).SetMaxDatagramSize), arg0)
}

// TimeUntilSend mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) TimeUntilSend(arg0 protocol.ByteCount) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeUntilSend", arg0)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// TimeUntilSend indicates an expected call of TimeUntilSend.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) TimeUntilSend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeUntilSend", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).TimeUntilSend), arg0)
}
