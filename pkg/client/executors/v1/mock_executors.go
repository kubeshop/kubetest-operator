// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube-operator/pkg/client/executors/v1 (interfaces: Interface)

// Package executors is a generated GoMock package.
package executors

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/kubeshop/testkube-operator/api/executor/v1"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockInterface) Create(arg0 *v1.Executor) (*v1.Executor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.Executor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockInterface) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), arg0)
}

// DeleteByLabels mocks base method.
func (m *MockInterface) DeleteByLabels(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByLabels", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByLabels indicates an expected call of DeleteByLabels.
func (mr *MockInterfaceMockRecorder) DeleteByLabels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByLabels", reflect.TypeOf((*MockInterface)(nil).DeleteByLabels), arg0)
}

// Get mocks base method.
func (m *MockInterface) Get(arg0 string) (*v1.Executor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v1.Executor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), arg0)
}

// GetByType mocks base method.
func (m *MockInterface) GetByType(arg0 string) (*v1.Executor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByType", arg0)
	ret0, _ := ret[0].(*v1.Executor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByType indicates an expected call of GetByType.
func (mr *MockInterfaceMockRecorder) GetByType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByType", reflect.TypeOf((*MockInterface)(nil).GetByType), arg0)
}

// List mocks base method.
func (m *MockInterface) List(arg0 string) (*v1.ExecutorList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.ExecutorList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockInterface) Update(arg0 *v1.Executor) (*v1.Executor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.Executor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInterface)(nil).Update), arg0)
}
