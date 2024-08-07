// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/servicebus/receiver.go
//
// Generated by this command:
//
//	mockgen -source=./internal/servicebus/receiver.go -package=servicebus -destination=./internal/servicebus/mock_receiver.go
//

// Package servicebus is a generated GoMock package.
package servicebus

import (
	context "context"
	reflect "reflect"

	azservicebus "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	gomock "go.uber.org/mock/gomock"
)

// MockReceiver is a mock of Receiver interface.
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverMockRecorder
}

// MockReceiverMockRecorder is the mock recorder for MockReceiver.
type MockReceiverMockRecorder struct {
	mock *MockReceiver
}

// NewMockReceiver creates a new mock instance.
func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiver) EXPECT() *MockReceiverMockRecorder {
	return m.recorder
}

// AbandonMessage mocks base method.
func (m *MockReceiver) AbandonMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.AbandonMessageOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbandonMessage", ctx, message, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbandonMessage indicates an expected call of AbandonMessage.
func (mr *MockReceiverMockRecorder) AbandonMessage(ctx, message, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbandonMessage", reflect.TypeOf((*MockReceiver)(nil).AbandonMessage), ctx, message, options)
}

// Close mocks base method.
func (m *MockReceiver) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReceiverMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReceiver)(nil).Close), ctx)
}

// CompleteMessage mocks base method.
func (m *MockReceiver) CompleteMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.CompleteMessageOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteMessage", ctx, message, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteMessage indicates an expected call of CompleteMessage.
func (mr *MockReceiverMockRecorder) CompleteMessage(ctx, message, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteMessage", reflect.TypeOf((*MockReceiver)(nil).CompleteMessage), ctx, message, options)
}

// DeadLetterMessage mocks base method.
func (m *MockReceiver) DeadLetterMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.DeadLetterOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeadLetterMessage", ctx, message, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeadLetterMessage indicates an expected call of DeadLetterMessage.
func (mr *MockReceiverMockRecorder) DeadLetterMessage(ctx, message, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeadLetterMessage", reflect.TypeOf((*MockReceiver)(nil).DeadLetterMessage), ctx, message, options)
}

// DeferMessage mocks base method.
func (m *MockReceiver) DeferMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.DeferMessageOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeferMessage", ctx, message, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeferMessage indicates an expected call of DeferMessage.
func (mr *MockReceiverMockRecorder) DeferMessage(ctx, message, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeferMessage", reflect.TypeOf((*MockReceiver)(nil).DeferMessage), ctx, message, options)
}

// PeekMessages mocks base method.
func (m *MockReceiver) PeekMessages(ctx context.Context, maxMessageCount int, options *azservicebus.PeekMessagesOptions) ([]*azservicebus.ReceivedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekMessages", ctx, maxMessageCount, options)
	ret0, _ := ret[0].([]*azservicebus.ReceivedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeekMessages indicates an expected call of PeekMessages.
func (mr *MockReceiverMockRecorder) PeekMessages(ctx, maxMessageCount, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekMessages", reflect.TypeOf((*MockReceiver)(nil).PeekMessages), ctx, maxMessageCount, options)
}

// ReceiveDeferredMessages mocks base method.
func (m *MockReceiver) ReceiveDeferredMessages(ctx context.Context, sequenceNumbers []int64, options *azservicebus.ReceiveDeferredMessagesOptions) ([]*azservicebus.ReceivedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveDeferredMessages", ctx, sequenceNumbers, options)
	ret0, _ := ret[0].([]*azservicebus.ReceivedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveDeferredMessages indicates an expected call of ReceiveDeferredMessages.
func (mr *MockReceiverMockRecorder) ReceiveDeferredMessages(ctx, sequenceNumbers, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveDeferredMessages", reflect.TypeOf((*MockReceiver)(nil).ReceiveDeferredMessages), ctx, sequenceNumbers, options)
}

// ReceiveMessages mocks base method.
func (m *MockReceiver) ReceiveMessages(ctx context.Context, maxMessages int, options *azservicebus.ReceiveMessagesOptions) ([]*azservicebus.ReceivedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessages", ctx, maxMessages, options)
	ret0, _ := ret[0].([]*azservicebus.ReceivedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessages indicates an expected call of ReceiveMessages.
func (mr *MockReceiverMockRecorder) ReceiveMessages(ctx, maxMessages, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessages", reflect.TypeOf((*MockReceiver)(nil).ReceiveMessages), ctx, maxMessages, options)
}

// RenewMessageLock mocks base method.
func (m *MockReceiver) RenewMessageLock(ctx context.Context, msg *azservicebus.ReceivedMessage, options *azservicebus.RenewMessageLockOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewMessageLock", ctx, msg, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenewMessageLock indicates an expected call of RenewMessageLock.
func (mr *MockReceiverMockRecorder) RenewMessageLock(ctx, msg, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewMessageLock", reflect.TypeOf((*MockReceiver)(nil).RenewMessageLock), ctx, msg, options)
}
