// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	testing "testing"

	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

// StorageV1Interface is an autogenerated mock type for the StorageV1Interface type
type StorageV1Interface struct {
	mock.Mock
}

// CSIDrivers provides a mock function with given fields:
func (_m *StorageV1Interface) CSIDrivers() v1.CSIDriverInterface {
	ret := _m.Called()

	var r0 v1.CSIDriverInterface
	if rf, ok := ret.Get(0).(func() v1.CSIDriverInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.CSIDriverInterface)
		}
	}

	return r0
}

// CSINodes provides a mock function with given fields:
func (_m *StorageV1Interface) CSINodes() v1.CSINodeInterface {
	ret := _m.Called()

	var r0 v1.CSINodeInterface
	if rf, ok := ret.Get(0).(func() v1.CSINodeInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.CSINodeInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *StorageV1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// StorageClasses provides a mock function with given fields:
func (_m *StorageV1Interface) StorageClasses() v1.StorageClassInterface {
	ret := _m.Called()

	var r0 v1.StorageClassInterface
	if rf, ok := ret.Get(0).(func() v1.StorageClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.StorageClassInterface)
		}
	}

	return r0
}

// VolumeAttachments provides a mock function with given fields:
func (_m *StorageV1Interface) VolumeAttachments() v1.VolumeAttachmentInterface {
	ret := _m.Called()

	var r0 v1.VolumeAttachmentInterface
	if rf, ok := ret.Get(0).(func() v1.VolumeAttachmentInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.VolumeAttachmentInterface)
		}
	}

	return r0
}

// NewStorageV1Interface creates a new instance of StorageV1Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorageV1Interface(t testing.TB) *StorageV1Interface {
	mock := &StorageV1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
