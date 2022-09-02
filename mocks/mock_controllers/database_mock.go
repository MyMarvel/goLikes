// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIDatabase is a mock of IDatabase interface.
type MockIDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockIDatabaseMockRecorder
}

// MockIDatabaseMockRecorder is the mock recorder for MockIDatabase.
type MockIDatabaseMockRecorder struct {
	mock *MockIDatabase
}

// NewMockIDatabase creates a new mock instance.
func NewMockIDatabase(ctrl *gomock.Controller) *MockIDatabase {
	mock := &MockIDatabase{ctrl: ctrl}
	mock.recorder = &MockIDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDatabase) EXPECT() *MockIDatabaseMockRecorder {
	return m.recorder
}

// AddLikeToAccount mocks base method.
func (m *MockIDatabase) AddLikeToAccount(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLikeToAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLikeToAccount indicates an expected call of AddLikeToAccount.
func (mr *MockIDatabaseMockRecorder) AddLikeToAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLikeToAccount", reflect.TypeOf((*MockIDatabase)(nil).AddLikeToAccount), arg0, arg1)
}

// IsAccountLikes mocks base method.
func (m *MockIDatabase) IsAccountLikes(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAccountLikes", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAccountLikes indicates an expected call of IsAccountLikes.
func (mr *MockIDatabaseMockRecorder) IsAccountLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAccountLikes", reflect.TypeOf((*MockIDatabase)(nil).IsAccountLikes), arg0, arg1)
}

// SelectAmountLikesForAccount mocks base method.
func (m *MockIDatabase) SelectAmountLikesForAccount(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAmountLikesForAccount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAmountLikesForAccount indicates an expected call of SelectAmountLikesForAccount.
func (mr *MockIDatabaseMockRecorder) SelectAmountLikesForAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAmountLikesForAccount", reflect.TypeOf((*MockIDatabase)(nil).SelectAmountLikesForAccount), arg0)
}
