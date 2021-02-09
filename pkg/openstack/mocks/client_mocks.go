// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/machine-controller-manager-provider-openstack/pkg/openstack (interfaces: Factory,Compute,Network)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	openstack "github.com/gardener/machine-controller-manager-provider-openstack/pkg/openstack"
	gomock "github.com/golang/mock/gomock"
	servers "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	ports "github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	subnets "github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Compute mocks base method.
func (m *MockFactory) Compute(arg0 ...openstack.Option) (openstack.Compute, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Compute", varargs...)
	ret0, _ := ret[0].(openstack.Compute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compute indicates an expected call of Compute.
func (mr *MockFactoryMockRecorder) Compute(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compute", reflect.TypeOf((*MockFactory)(nil).Compute), arg0...)
}

// Network mocks base method.
func (m *MockFactory) Network(arg0 ...openstack.Option) (openstack.Network, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Network", varargs...)
	ret0, _ := ret[0].(openstack.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Network indicates an expected call of Network.
func (mr *MockFactoryMockRecorder) Network(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockFactory)(nil).Network), arg0...)
}

// MockCompute is a mock of Compute interface.
type MockCompute struct {
	ctrl     *gomock.Controller
	recorder *MockComputeMockRecorder
}

// MockComputeMockRecorder is the mock recorder for MockCompute.
type MockComputeMockRecorder struct {
	mock *MockCompute
}

// NewMockCompute creates a new mock instance.
func NewMockCompute(ctrl *gomock.Controller) *MockCompute {
	mock := &MockCompute{ctrl: ctrl}
	mock.recorder = &MockComputeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompute) EXPECT() *MockComputeMockRecorder {
	return m.recorder
}

// BootFromVolume mocks base method.
func (m *MockCompute) BootFromVolume(arg0 servers.CreateOptsBuilder) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootFromVolume", arg0)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BootFromVolume indicates an expected call of BootFromVolume.
func (mr *MockComputeMockRecorder) BootFromVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootFromVolume", reflect.TypeOf((*MockCompute)(nil).BootFromVolume), arg0)
}

// CreateServer mocks base method.
func (m *MockCompute) CreateServer(arg0 servers.CreateOptsBuilder) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", arg0)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer.
func (mr *MockComputeMockRecorder) CreateServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockCompute)(nil).CreateServer), arg0)
}

// DeleteServer mocks base method.
func (m *MockCompute) DeleteServer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServer indicates an expected call of DeleteServer.
func (mr *MockComputeMockRecorder) DeleteServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockCompute)(nil).DeleteServer), arg0)
}

// FlavorIDFromName mocks base method.
func (m *MockCompute) FlavorIDFromName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlavorIDFromName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FlavorIDFromName indicates an expected call of FlavorIDFromName.
func (mr *MockComputeMockRecorder) FlavorIDFromName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlavorIDFromName", reflect.TypeOf((*MockCompute)(nil).FlavorIDFromName), arg0)
}

// GetServer mocks base method.
func (m *MockCompute) GetServer(arg0 string) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServer", arg0)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServer indicates an expected call of GetServer.
func (mr *MockComputeMockRecorder) GetServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockCompute)(nil).GetServer), arg0)
}

// ImageIDFromName mocks base method.
func (m *MockCompute) ImageIDFromName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageIDFromName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageIDFromName indicates an expected call of ImageIDFromName.
func (mr *MockComputeMockRecorder) ImageIDFromName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageIDFromName", reflect.TypeOf((*MockCompute)(nil).ImageIDFromName), arg0)
}

// ListServers mocks base method.
func (m *MockCompute) ListServers(arg0 servers.ListOptsBuilder) ([]servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServers", arg0)
	ret0, _ := ret[0].([]servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServers indicates an expected call of ListServers.
func (mr *MockComputeMockRecorder) ListServers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockCompute)(nil).ListServers), arg0)
}

// MockNetwork is a mock of Network interface.
type MockNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkMockRecorder
}

// MockNetworkMockRecorder is the mock recorder for MockNetwork.
type MockNetworkMockRecorder struct {
	mock *MockNetwork
}

// NewMockNetwork creates a new mock instance.
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// CreatePort mocks base method.
func (m *MockNetwork) CreatePort(arg0 ports.CreateOptsBuilder) (*ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePort", arg0)
	ret0, _ := ret[0].(*ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePort indicates an expected call of CreatePort.
func (mr *MockNetworkMockRecorder) CreatePort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePort", reflect.TypeOf((*MockNetwork)(nil).CreatePort), arg0)
}

// DeletePort mocks base method.
func (m *MockNetwork) DeletePort(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePort", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePort indicates an expected call of DeletePort.
func (mr *MockNetworkMockRecorder) DeletePort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePort", reflect.TypeOf((*MockNetwork)(nil).DeletePort), arg0)
}

// GetSubnet mocks base method.
func (m *MockNetwork) GetSubnet(arg0 string) (*subnets.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", arg0)
	ret0, _ := ret[0].(*subnets.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockNetworkMockRecorder) GetSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockNetwork)(nil).GetSubnet), arg0)
}

// GroupIDFromName mocks base method.
func (m *MockNetwork) GroupIDFromName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupIDFromName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupIDFromName indicates an expected call of GroupIDFromName.
func (mr *MockNetworkMockRecorder) GroupIDFromName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupIDFromName", reflect.TypeOf((*MockNetwork)(nil).GroupIDFromName), arg0)
}

// ListPorts mocks base method.
func (m *MockNetwork) ListPorts(arg0 ports.ListOptsBuilder) ([]ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPorts", arg0)
	ret0, _ := ret[0].([]ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPorts indicates an expected call of ListPorts.
func (mr *MockNetworkMockRecorder) ListPorts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPorts", reflect.TypeOf((*MockNetwork)(nil).ListPorts), arg0)
}

// NetworkIDFromName mocks base method.
func (m *MockNetwork) NetworkIDFromName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkIDFromName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkIDFromName indicates an expected call of NetworkIDFromName.
func (mr *MockNetworkMockRecorder) NetworkIDFromName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkIDFromName", reflect.TypeOf((*MockNetwork)(nil).NetworkIDFromName), arg0)
}

// PortIDFromName mocks base method.
func (m *MockNetwork) PortIDFromName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PortIDFromName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PortIDFromName indicates an expected call of PortIDFromName.
func (mr *MockNetworkMockRecorder) PortIDFromName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PortIDFromName", reflect.TypeOf((*MockNetwork)(nil).PortIDFromName), arg0)
}

// UpdatePort mocks base method.
func (m *MockNetwork) UpdatePort(arg0 string, arg1 ports.UpdateOptsBuilder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePort", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePort indicates an expected call of UpdatePort.
func (mr *MockNetworkMockRecorder) UpdatePort(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePort", reflect.TypeOf((*MockNetwork)(nil).UpdatePort), arg0, arg1)
}