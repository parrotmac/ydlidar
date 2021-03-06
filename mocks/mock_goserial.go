// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cesanta/go-serial/serial (interfaces: Serial)

// Package mock_serial is a generated GoMock package.
package mock_serial

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockSerial is a mock of Serial interface
type MockSerial struct {
	ctrl     *gomock.Controller
	recorder *MockSerialMockRecorder
}

// MockSerialMockRecorder is the mock recorder for MockSerial
type MockSerialMockRecorder struct {
	mock *MockSerial
}

// NewMockSerial creates a new mock instance
func NewMockSerial(ctrl *gomock.Controller) *MockSerial {
	mock := &MockSerial{ctrl: ctrl}
	mock.recorder = &MockSerialMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSerial) EXPECT() *MockSerialMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSerial) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSerialMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSerial)(nil).Close))
}

// Flush mocks base method
func (m *MockSerial) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockSerialMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockSerial)(nil).Flush))
}

// Read mocks base method
func (m *MockSerial) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockSerialMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSerial)(nil).Read), arg0)
}

// SetBaudRate mocks base method
func (m *MockSerial) SetBaudRate(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBaudRate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBaudRate indicates an expected call of SetBaudRate
func (mr *MockSerialMockRecorder) SetBaudRate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBaudRate", reflect.TypeOf((*MockSerial)(nil).SetBaudRate), arg0)
}

// SetBreak mocks base method
func (m *MockSerial) SetBreak(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBreak", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBreak indicates an expected call of SetBreak
func (mr *MockSerialMockRecorder) SetBreak(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBreak", reflect.TypeOf((*MockSerial)(nil).SetBreak), arg0)
}

// SetDTR mocks base method
func (m *MockSerial) SetDTR(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDTR", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDTR indicates an expected call of SetDTR
func (mr *MockSerialMockRecorder) SetDTR(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDTR", reflect.TypeOf((*MockSerial)(nil).SetDTR), arg0)
}

// SetRTS mocks base method
func (m *MockSerial) SetRTS(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRTS", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRTS indicates an expected call of SetRTS
func (mr *MockSerialMockRecorder) SetRTS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRTS", reflect.TypeOf((*MockSerial)(nil).SetRTS), arg0)
}

// SetRTSDTR mocks base method
func (m *MockSerial) SetRTSDTR(arg0, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRTSDTR", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRTSDTR indicates an expected call of SetRTSDTR
func (mr *MockSerialMockRecorder) SetRTSDTR(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRTSDTR", reflect.TypeOf((*MockSerial)(nil).SetRTSDTR), arg0, arg1)
}

// SetReadTimeout mocks base method
func (m *MockSerial) SetReadTimeout(arg0 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadTimeout", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadTimeout indicates an expected call of SetReadTimeout
func (mr *MockSerialMockRecorder) SetReadTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadTimeout", reflect.TypeOf((*MockSerial)(nil).SetReadTimeout), arg0)
}

// Write mocks base method
func (m *MockSerial) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockSerialMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSerial)(nil).Write), arg0)
}
