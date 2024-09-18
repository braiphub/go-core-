// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/braiphub/go-core/cache (interfaces: Cacherer)

// Package cache is a generated GoMock package.
package cache

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockCacherer is a mock of Cacherer interface.
type MockCacherer struct {
	ctrl     *gomock.Controller
	recorder *MockCachererMockRecorder
}

// MockCachererMockRecorder is the mock recorder for MockCacherer.
type MockCachererMockRecorder struct {
	mock *MockCacherer
}

// NewMockCacherer creates a new mock instance.
func NewMockCacherer(ctrl *gomock.Controller) *MockCacherer {
	mock := &MockCacherer{ctrl: ctrl}
	mock.recorder = &MockCachererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacherer) EXPECT() *MockCachererMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCacherer) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCachererMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacherer)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockCacherer) Get(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCachererMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacherer)(nil).Get), arg0, arg1)
}

// GetInt mocks base method.
func (m *MockCacherer) GetInt(arg0 context.Context, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt indicates an expected call of GetInt.
func (mr *MockCachererMockRecorder) GetInt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockCacherer)(nil).GetInt), arg0, arg1)
}

// GetString mocks base method.
func (m *MockCacherer) GetString(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetString indicates an expected call of GetString.
func (mr *MockCachererMockRecorder) GetString(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockCacherer)(nil).GetString), arg0, arg1)
}

// GetUint mocks base method.
func (m *MockCacherer) GetUint(arg0 context.Context, arg1 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUint indicates an expected call of GetUint.
func (mr *MockCachererMockRecorder) GetUint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint", reflect.TypeOf((*MockCacherer)(nil).GetUint), arg0, arg1)
}

// Set mocks base method.
func (m *MockCacherer) Set(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCachererMockRecorder) Set(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacherer)(nil).Set), arg0, arg1, arg2, arg3)
}

// TestConnection mocks base method.
func (m *MockCacherer) TestConnection(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestConnection", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TestConnection indicates an expected call of TestConnection.
func (mr *MockCachererMockRecorder) TestConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestConnection", reflect.TypeOf((*MockCacherer)(nil).TestConnection), arg0)
}
