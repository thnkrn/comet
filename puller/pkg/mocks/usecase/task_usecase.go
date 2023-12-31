// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// Download provides a mock function with given fields: ctx, gcsObject
func (_m *TaskUsecase) Download(ctx context.Context, gcsObject string) (bool, error) {
	ret := _m.Called(ctx, gcsObject)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, gcsObject)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, gcsObject)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, gcsObject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestSeries provides a mock function with given fields: objects, ingestOnlyLatestDirectory
func (_m *TaskUsecase) GetLatestSeries(objects []string, ingestOnlyLatestDirectory bool) ([]string, error) {
	ret := _m.Called(objects, ingestOnlyLatestDirectory)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func([]string, bool) ([]string, error)); ok {
		return rf(objects, ingestOnlyLatestDirectory)
	}
	if rf, ok := ret.Get(0).(func([]string, bool) []string); ok {
		r0 = rf(objects, ingestOnlyLatestDirectory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]string, bool) error); ok {
		r1 = rf(objects, ingestOnlyLatestDirectory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTargetStorageObjects provides a mock function with given fields: ctx, dbName, source, authorization, ingestOnlyLatestDirectory, ignoreLastIngest
func (_m *TaskUsecase) GetTargetStorageObjects(ctx context.Context, dbName string, source string, authorization string, ingestOnlyLatestDirectory bool, ignoreLastIngest bool) ([]string, error) {
	ret := _m.Called(ctx, dbName, source, authorization, ingestOnlyLatestDirectory, ignoreLastIngest)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, bool, bool) ([]string, error)); ok {
		return rf(ctx, dbName, source, authorization, ingestOnlyLatestDirectory, ignoreLastIngest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, bool, bool) []string); ok {
		r0 = rf(ctx, dbName, source, authorization, ingestOnlyLatestDirectory, ignoreLastIngest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, bool, bool) error); ok {
		r1 = rf(ctx, dbName, source, authorization, ingestOnlyLatestDirectory, ignoreLastIngest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ingest provides a mock function with given fields: ctx, authorization, dbName, directory
func (_m *TaskUsecase) Ingest(ctx context.Context, authorization string, dbName string, directory string) (bool, error) {
	ret := _m.Called(ctx, authorization, dbName, directory)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (bool, error)); ok {
		return rf(ctx, authorization, dbName, directory)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) bool); ok {
		r0 = rf(ctx, authorization, dbName, directory)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, authorization, dbName, directory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Perform provides a mock function with given fields: ctx, db, ignoreLastIngest
func (_m *TaskUsecase) Perform(ctx context.Context, db string, ignoreLastIngest bool) error {
	ret := _m.Called(ctx, db, ignoreLastIngest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) error); ok {
		r0 = rf(ctx, db, ignoreLastIngest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveTemp provides a mock function with given fields: directory
func (_m *TaskUsecase) RemoveTemp(directory string) error {
	ret := _m.Called(directory)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(directory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTaskUsecase creates a new instance of TaskUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskUsecase {
	mock := &TaskUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
