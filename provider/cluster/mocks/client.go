// Code generated by mockery 2.12.1. DO NOT EDIT.

package mocks

import (
	akash_networkv2beta1 "github.com/ovrclk/akash/pkg/apis/akash.network/v2beta1"

	context "context"

	io "io"

	mock "github.com/stretchr/testify/mock"

	remotecommand "k8s.io/client-go/tools/remotecommand"

	testing "testing"

	typesv1beta2 "github.com/ovrclk/akash/x/market/types/v1beta2"

	v1beta2 "github.com/ovrclk/akash/provider/cluster/types/v1beta2"

	v2beta1 "github.com/ovrclk/akash/manifest/v2beta1"

	version "k8s.io/apimachinery/pkg/version"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AllHostnames provides a mock function with given fields: _a0
func (_m *Client) AllHostnames(_a0 context.Context) ([]v1beta2.ActiveHostname, error) {
	ret := _m.Called(_a0)

	var r0 []v1beta2.ActiveHostname
	if rf, ok := ret.Get(0).(func(context.Context) []v1beta2.ActiveHostname); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1beta2.ActiveHostname)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectHostnameToDeployment provides a mock function with given fields: ctx, directive
func (_m *Client) ConnectHostnameToDeployment(ctx context.Context, directive v1beta2.ConnectHostnameToDeploymentDirective) error {
	ret := _m.Called(ctx, directive)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1beta2.ConnectHostnameToDeploymentDirective) error); ok {
		r0 = rf(ctx, directive)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeclareHostname provides a mock function with given fields: ctx, lID, host, serviceName, externalPort
func (_m *Client) DeclareHostname(ctx context.Context, lID typesv1beta2.LeaseID, host string, serviceName string, externalPort uint32) error {
	ret := _m.Called(ctx, lID, host, serviceName, externalPort)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID, string, string, uint32) error); ok {
		r0 = rf(ctx, lID, host, serviceName, externalPort)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deploy provides a mock function with given fields: ctx, lID, mgroup
func (_m *Client) Deploy(ctx context.Context, lID typesv1beta2.LeaseID, mgroup *v2beta1.Group) error {
	ret := _m.Called(ctx, lID, mgroup)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID, *v2beta1.Group) error); ok {
		r0 = rf(ctx, lID, mgroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deployments provides a mock function with given fields: _a0
func (_m *Client) Deployments(_a0 context.Context) ([]v1beta2.Deployment, error) {
	ret := _m.Called(_a0)

	var r0 []v1beta2.Deployment
	if rf, ok := ret.Get(0).(func(context.Context) []v1beta2.Deployment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1beta2.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exec provides a mock function with given fields: ctx, lID, service, podIndex, cmd, stdin, stdout, stderr, tty, tsq
func (_m *Client) Exec(ctx context.Context, lID typesv1beta2.LeaseID, service string, podIndex uint, cmd []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, tty bool, tsq remotecommand.TerminalSizeQueue) (v1beta2.ExecResult, error) {
	ret := _m.Called(ctx, lID, service, podIndex, cmd, stdin, stdout, stderr, tty, tsq)

	var r0 v1beta2.ExecResult
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID, string, uint, []string, io.Reader, io.Writer, io.Writer, bool, remotecommand.TerminalSizeQueue) v1beta2.ExecResult); ok {
		r0 = rf(ctx, lID, service, podIndex, cmd, stdin, stdout, stderr, tty, tsq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta2.ExecResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, typesv1beta2.LeaseID, string, uint, []string, io.Reader, io.Writer, io.Writer, bool, remotecommand.TerminalSizeQueue) error); ok {
		r1 = rf(ctx, lID, service, podIndex, cmd, stdin, stdout, stderr, tty, tsq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHostnameDeploymentConnections provides a mock function with given fields: ctx
func (_m *Client) GetHostnameDeploymentConnections(ctx context.Context) ([]v1beta2.LeaseIDHostnameConnection, error) {
	ret := _m.Called(ctx)

	var r0 []v1beta2.LeaseIDHostnameConnection
	if rf, ok := ret.Get(0).(func(context.Context) []v1beta2.LeaseIDHostnameConnection); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1beta2.LeaseIDHostnameConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManifestGroup provides a mock function with given fields: _a0, _a1
func (_m *Client) GetManifestGroup(_a0 context.Context, _a1 typesv1beta2.LeaseID) (bool, akash_networkv2beta1.ManifestGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 akash_networkv2beta1.ManifestGroup
	if rf, ok := ret.Get(1).(func(context.Context, typesv1beta2.LeaseID) akash_networkv2beta1.ManifestGroup); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(akash_networkv2beta1.ManifestGroup)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, typesv1beta2.LeaseID) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Inventory provides a mock function with given fields: _a0
func (_m *Client) Inventory(_a0 context.Context) (v1beta2.Inventory, error) {
	ret := _m.Called(_a0)

	var r0 v1beta2.Inventory
	if rf, ok := ret.Get(0).(func(context.Context) v1beta2.Inventory); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta2.Inventory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubeVersion provides a mock function with given fields:
func (_m *Client) KubeVersion() (*version.Info, error) {
	ret := _m.Called()

	var r0 *version.Info
	if rf, ok := ret.Get(0).(func() *version.Info); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaseEvents provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Client) LeaseEvents(_a0 context.Context, _a1 typesv1beta2.LeaseID, _a2 string, _a3 bool) (v1beta2.EventsWatcher, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 v1beta2.EventsWatcher
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID, string, bool) v1beta2.EventsWatcher); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta2.EventsWatcher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, typesv1beta2.LeaseID, string, bool) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaseLogs provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *Client) LeaseLogs(_a0 context.Context, _a1 typesv1beta2.LeaseID, _a2 string, _a3 bool, _a4 *int64) ([]*v1beta2.ServiceLog, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 []*v1beta2.ServiceLog
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID, string, bool, *int64) []*v1beta2.ServiceLog); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta2.ServiceLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, typesv1beta2.LeaseID, string, bool, *int64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaseStatus provides a mock function with given fields: _a0, _a1
func (_m *Client) LeaseStatus(_a0 context.Context, _a1 typesv1beta2.LeaseID) (*v1beta2.LeaseStatus, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta2.LeaseStatus
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID) *v1beta2.LeaseStatus); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.LeaseStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, typesv1beta2.LeaseID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObserveHostnameState provides a mock function with given fields: ctx
func (_m *Client) ObserveHostnameState(ctx context.Context) (<-chan v1beta2.HostnameResourceEvent, error) {
	ret := _m.Called(ctx)

	var r0 <-chan v1beta2.HostnameResourceEvent
	if rf, ok := ret.Get(0).(func(context.Context) <-chan v1beta2.HostnameResourceEvent); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan v1beta2.HostnameResourceEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurgeDeclaredHostname provides a mock function with given fields: ctx, lID, hostname
func (_m *Client) PurgeDeclaredHostname(ctx context.Context, lID typesv1beta2.LeaseID, hostname string) error {
	ret := _m.Called(ctx, lID, hostname)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID, string) error); ok {
		r0 = rf(ctx, lID, hostname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PurgeDeclaredHostnames provides a mock function with given fields: ctx, lID
func (_m *Client) PurgeDeclaredHostnames(ctx context.Context, lID typesv1beta2.LeaseID) error {
	ret := _m.Called(ctx, lID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID) error); ok {
		r0 = rf(ctx, lID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveHostnameFromDeployment provides a mock function with given fields: ctx, hostname, leaseID, allowMissing
func (_m *Client) RemoveHostnameFromDeployment(ctx context.Context, hostname string, leaseID typesv1beta2.LeaseID, allowMissing bool) error {
	ret := _m.Called(ctx, hostname, leaseID, allowMissing)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, typesv1beta2.LeaseID, bool) error); ok {
		r0 = rf(ctx, hostname, leaseID, allowMissing)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceStatus provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) ServiceStatus(_a0 context.Context, _a1 typesv1beta2.LeaseID, _a2 string) (*v1beta2.ServiceStatus, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *v1beta2.ServiceStatus
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID, string) *v1beta2.ServiceStatus); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.ServiceStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, typesv1beta2.LeaseID, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeardownLease provides a mock function with given fields: _a0, _a1
func (_m *Client) TeardownLease(_a0 context.Context, _a1 typesv1beta2.LeaseID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t testing.TB) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
