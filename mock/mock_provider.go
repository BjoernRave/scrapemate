// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gosom/scrapemate (interfaces: JobProvider)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	scrapemate "github.com/gosom/scrapemate"
)

// MockJobProvider is a mock of JobProvider interface.
type MockJobProvider struct {
	ctrl     *gomock.Controller
	recorder *MockJobProviderMockRecorder
}

// MockJobProviderMockRecorder is the mock recorder for MockJobProvider.
type MockJobProviderMockRecorder struct {
	mock *MockJobProvider
}

// NewMockJobProvider creates a new mock instance.
func NewMockJobProvider(ctrl *gomock.Controller) *MockJobProvider {
	mock := &MockJobProvider{ctrl: ctrl}
	mock.recorder = &MockJobProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobProvider) EXPECT() *MockJobProviderMockRecorder {
	return m.recorder
}

// Jobs mocks base method.
func (m *MockJobProvider) Jobs(arg0 context.Context) (<-chan scrapemate.IJob, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jobs", arg0)
	ret0, _ := ret[0].(<-chan scrapemate.IJob)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// Jobs indicates an expected call of Jobs.
func (mr *MockJobProviderMockRecorder) Jobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockJobProvider)(nil).Jobs), arg0)
}

// Push mocks base method.
func (m *MockJobProvider) Push(arg0 context.Context, arg1 scrapemate.IJob) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockJobProviderMockRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockJobProvider)(nil).Push), arg0, arg1)
}
