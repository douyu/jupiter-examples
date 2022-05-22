// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RedisInterface is an autogenerated mock type for the RedisInterface type
type RedisInterface struct {
	mock.Mock
}

// GetNodeId provides a mock function with given fields:
func (_m *RedisInterface) GetNodeId() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRedisInterface creates a new instance of RedisInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRedisInterface(t testing.TB) *RedisInterface {
	mock := &RedisInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
