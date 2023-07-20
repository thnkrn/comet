// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, dbName
func (_m *UserUsecase) Count(ctx context.Context, dbName string) (string, error) {
	ret := _m.Called(ctx, dbName)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, dbName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, dbName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dbName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, dbName, key, value
func (_m *UserUsecase) Create(ctx context.Context, dbName string, key string, value []byte) ([]byte, error) {
	ret := _m.Called(ctx, dbName, key, value)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte) ([]byte, error)); ok {
		return rf(ctx, dbName, key, value)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte) []byte); ok {
		r0 = rf(ctx, dbName, key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, []byte) error); ok {
		r1 = rf(ctx, dbName, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, dbName, key
func (_m *UserUsecase) Delete(ctx context.Context, dbName string, key string) error {
	ret := _m.Called(ctx, dbName, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, dbName, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, dbName, key
func (_m *UserUsecase) Get(ctx context.Context, dbName string, key string) ([]byte, error) {
	ret := _m.Called(ctx, dbName, key)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]byte, error)); ok {
		return rf(ctx, dbName, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []byte); ok {
		r0 = rf(ctx, dbName, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, dbName, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MultiGet provides a mock function with given fields: ctx, dbName, keys
func (_m *UserUsecase) MultiGet(ctx context.Context, dbName string, keys []string) ([]string, error) {
	ret := _m.Called(ctx, dbName, keys)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]string, error)); ok {
		return rf(ctx, dbName, keys)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []string); ok {
		r0 = rf(ctx, dbName, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, dbName, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
