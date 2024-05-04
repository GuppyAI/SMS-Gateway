// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/messaging/message_handler.go
//
// Generated by this command:
//
//	mockgen -source=./internal/messaging/message_handler.go -package=messaging -destination=./internal/messaging/mock_message_handler.go
//

// Package messaging is a generated GoMock package.
package messaging

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMessageHandler is a mock of MessageHandler interface.
type MockMessageHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMessageHandlerMockRecorder
}

// MockMessageHandlerMockRecorder is the mock recorder for MockMessageHandler.
type MockMessageHandlerMockRecorder struct {
	mock *MockMessageHandler
}

// NewMockMessageHandler creates a new mock instance.
func NewMockMessageHandler(ctrl *gomock.Controller) *MockMessageHandler {
	mock := &MockMessageHandler{ctrl: ctrl}
	mock.recorder = &MockMessageHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageHandler) EXPECT() *MockMessageHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockMessageHandler) Handle(arg0 Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockMessageHandlerMockRecorder) Handle(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockMessageHandler)(nil).Handle), arg0)
}