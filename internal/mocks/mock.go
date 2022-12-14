// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockStore) GetInfo(find string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", find)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockStoreMockRecorder) GetInfo(find interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockStore)(nil).GetInfo), find)
}

// PostInfo mocks base method.
func (m *MockStore) PostInfo(info string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostInfo", info)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostInfo indicates an expected call of PostInfo.
func (mr *MockStoreMockRecorder) PostInfo(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostInfo", reflect.TypeOf((*MockStore)(nil).PostInfo), info)
}
