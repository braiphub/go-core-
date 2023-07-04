// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/braiphub/ms-orders/pkg/core/cache/redis (interfaces: ClientI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
)

// MockClientI is a mock of ClientI interface.
type MockClientI struct {
	ctrl     *gomock.Controller
	recorder *MockClientIMockRecorder
}

// MockClientIMockRecorder is the mock recorder for MockClientI.
type MockClientIMockRecorder struct {
	mock *MockClientI
}

// NewMockClientI creates a new mock instance.
func NewMockClientI(ctrl *gomock.Controller) *MockClientI {
	mock := &MockClientI{ctrl: ctrl}
	mock.recorder = &MockClientIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientI) EXPECT() *MockClientIMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockClientI) Exists(arg0 context.Context, arg1 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exists", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockClientIMockRecorder) Exists(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockClientI)(nil).Exists), varargs...)
}

// Get mocks base method.
func (m *MockClientI) Get(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockClientIMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClientI)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockClientI) Set(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockClientIMockRecorder) Set(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockClientI)(nil).Set), arg0, arg1, arg2, arg3)
}
