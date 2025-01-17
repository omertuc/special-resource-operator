// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package cluster is a generated GoMock package.
package cluster

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockCluster is a mock of Cluster interface.
type MockCluster struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMockRecorder
}

// MockClusterMockRecorder is the mock recorder for MockCluster.
type MockClusterMockRecorder struct {
	mock *MockCluster
}

// NewMockCluster creates a new mock instance.
func NewMockCluster(ctrl *gomock.Controller) *MockCluster {
	mock := &MockCluster{ctrl: ctrl}
	mock.recorder = &MockClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCluster) EXPECT() *MockClusterMockRecorder {
	return m.recorder
}

// GetDTKImages mocks base method.
func (m *MockCluster) GetDTKImages(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDTKImages", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDTKImages indicates an expected call of GetDTKImages.
func (mr *MockClusterMockRecorder) GetDTKImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDTKImages", reflect.TypeOf((*MockCluster)(nil).GetDTKImages), arg0)
}

// IsClusterInUpgrade mocks base method.
func (m *MockCluster) IsClusterInUpgrade(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClusterInUpgrade", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsClusterInUpgrade indicates an expected call of IsClusterInUpgrade.
func (mr *MockClusterMockRecorder) IsClusterInUpgrade(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClusterInUpgrade", reflect.TypeOf((*MockCluster)(nil).IsClusterInUpgrade), arg0)
}

// NextUpgradeVersion mocks base method.
func (m *MockCluster) NextUpgradeVersion(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextUpgradeVersion", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextUpgradeVersion indicates an expected call of NextUpgradeVersion.
func (mr *MockClusterMockRecorder) NextUpgradeVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextUpgradeVersion", reflect.TypeOf((*MockCluster)(nil).NextUpgradeVersion), arg0)
}

// OSImageURL mocks base method.
func (m *MockCluster) OSImageURL(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OSImageURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OSImageURL indicates an expected call of OSImageURL.
func (mr *MockClusterMockRecorder) OSImageURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OSImageURL", reflect.TypeOf((*MockCluster)(nil).OSImageURL), arg0)
}

// OperatingSystem mocks base method.
func (m *MockCluster) OperatingSystem(arg0 *v1.NodeList) (string, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperatingSystem", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// OperatingSystem indicates an expected call of OperatingSystem.
func (mr *MockClusterMockRecorder) OperatingSystem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperatingSystem", reflect.TypeOf((*MockCluster)(nil).OperatingSystem), arg0)
}

// Version mocks base method.
func (m *MockCluster) Version(arg0 context.Context) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Version indicates an expected call of Version.
func (mr *MockClusterMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockCluster)(nil).Version), arg0)
}
