// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gosom/scrapemate (interfaces: HttpFetcher)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	scrapemate "github.com/gosom/scrapemate"
)

// MockHttpFetcher is a mock of HttpFetcher interface.
type MockHttpFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockHttpFetcherMockRecorder
}

// MockHttpFetcherMockRecorder is the mock recorder for MockHttpFetcher.
type MockHttpFetcherMockRecorder struct {
	mock *MockHttpFetcher
}

// NewMockHttpFetcher creates a new mock instance.
func NewMockHttpFetcher(ctrl *gomock.Controller) *MockHttpFetcher {
	mock := &MockHttpFetcher{ctrl: ctrl}
	mock.recorder = &MockHttpFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpFetcher) EXPECT() *MockHttpFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockHttpFetcher) Fetch(arg0 context.Context, arg1 scrapemate.IJob) scrapemate.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1)
	ret0, _ := ret[0].(scrapemate.Response)
	return ret0
}

// Fetch indicates an expected call of Fetch.
func (mr *MockHttpFetcherMockRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockHttpFetcher)(nil).Fetch), arg0, arg1)
}
