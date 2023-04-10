// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gosom/scrapemate (interfaces: Cacher)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	scrapemate "github.com/gosom/scrapemate"
)

// MockCacher is a mock of Cacher interface.
type MockCacher struct {
	ctrl     *gomock.Controller
	recorder *MockCacherMockRecorder
}

// MockCacherMockRecorder is the mock recorder for MockCacher.
type MockCacherMockRecorder struct {
	mock *MockCacher
}

// NewMockCacher creates a new mock instance.
func NewMockCacher(ctrl *gomock.Controller) *MockCacher {
	mock := &MockCacher{ctrl: ctrl}
	mock.recorder = &MockCacherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacher) EXPECT() *MockCacherMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCacher) Get(arg0 context.Context, arg1 string) (scrapemate.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(scrapemate.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacherMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacher)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockCacher) Set(arg0 context.Context, arg1 string, arg2 scrapemate.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacherMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacher)(nil).Set), arg0, arg1, arg2)
}