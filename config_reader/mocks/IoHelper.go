// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IoHelper is an autogenerated mock type for the IoHelper type
type IoHelper struct {
	mock.Mock
}

// ReadFile provides a mock function with given fields: _a0
func (_m *IoHelper) ReadFile(_a0 string) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIoHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewIoHelper creates a new instance of IoHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIoHelper(t mockConstructorTestingTNewIoHelper) *IoHelper {
	mock := &IoHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}