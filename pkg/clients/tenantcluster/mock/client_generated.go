// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	v1 "k8s.io/api/core/v1"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// PatchMachine mocks base method
func (m *MockClient) PatchMachine(machine, originMachineCopy *v1beta1.Machine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchMachine", machine, originMachineCopy)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchMachine indicates an expected call of PatchMachine
func (mr *MockClientMockRecorder) PatchMachine(machine, originMachineCopy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchMachine", reflect.TypeOf((*MockClient)(nil).PatchMachine), machine, originMachineCopy)
}

// StatusPatchMachine mocks base method
func (m *MockClient) StatusPatchMachine(machine, originMachineCopy *v1beta1.Machine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusPatchMachine", machine, originMachineCopy)
	ret0, _ := ret[0].(error)
	return ret0
}

// StatusPatchMachine indicates an expected call of StatusPatchMachine
func (mr *MockClientMockRecorder) StatusPatchMachine(machine, originMachineCopy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusPatchMachine", reflect.TypeOf((*MockClient)(nil).StatusPatchMachine), machine, originMachineCopy)
}

// GetSecret mocks base method
func (m *MockClient) GetSecret(secretName, namespace string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", secretName, namespace)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockClientMockRecorder) GetSecret(secretName, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockClient)(nil).GetSecret), secretName, namespace)
}

// GetNamespace mocks base method
func (m *MockClient) GetNamespace() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockClientMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockClient)(nil).GetNamespace))
}

// GetInfraID mocks base method
func (m *MockClient) GetInfraID() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfraID indicates an expected call of GetInfraID
func (mr *MockClientMockRecorder) GetInfraID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraID", reflect.TypeOf((*MockClient)(nil).GetInfraID))
}
