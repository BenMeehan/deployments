// Copyright 2023 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mendersoftware/deployments/model"

	storage "github.com/mendersoftware/deployments/storage"

	time "time"
)

// ObjectStorage is an autogenerated mock type for the ObjectStorage type
type ObjectStorage struct {
	mock.Mock
}

// DeleteObject provides a mock function with given fields: ctx, path
func (_m *ObjectStorage) DeleteObject(ctx context.Context, path string) error {
	ret := _m.Called(ctx, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRequest provides a mock function with given fields: ctx, path, duration
func (_m *ObjectStorage) DeleteRequest(ctx context.Context, path string, duration time.Duration) (*model.Link, error) {
	ret := _m.Called(ctx, path, duration)

	var r0 *model.Link
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) *model.Link); ok {
		r0 = rf(ctx, path, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration) error); ok {
		r1 = rf(ctx, path, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObject provides a mock function with given fields: ctx, path
func (_m *ObjectStorage) GetObject(ctx context.Context, path string) (io.ReadCloser, error) {
	ret := _m.Called(ctx, path)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string) io.ReadCloser); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRequest provides a mock function with given fields: ctx, path, filename, duration
func (_m *ObjectStorage) GetRequest(ctx context.Context, path string, filename string, duration time.Duration) (*model.Link, error) {
	ret := _m.Called(ctx, path, filename, duration)

	var r0 *model.Link
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Duration) *model.Link); ok {
		r0 = rf(ctx, path, filename, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, time.Duration) error); ok {
		r1 = rf(ctx, path, filename, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *ObjectStorage) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutObject provides a mock function with given fields: ctx, path, src
func (_m *ObjectStorage) PutObject(ctx context.Context, path string, src io.Reader) error {
	ret := _m.Called(ctx, path, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader) error); ok {
		r0 = rf(ctx, path, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRequest provides a mock function with given fields: ctx, path, duration
func (_m *ObjectStorage) PutRequest(ctx context.Context, path string, duration time.Duration) (*model.Link, error) {
	ret := _m.Called(ctx, path, duration)

	var r0 *model.Link
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) *model.Link); ok {
		r0 = rf(ctx, path, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration) error); ok {
		r1 = rf(ctx, path, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatObject provides a mock function with given fields: ctx, path
func (_m *ObjectStorage) StatObject(ctx context.Context, path string) (*storage.ObjectInfo, error) {
	ret := _m.Called(ctx, path)

	var r0 *storage.ObjectInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) *storage.ObjectInfo); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ObjectInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewObjectStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewObjectStorage creates a new instance of ObjectStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewObjectStorage(t mockConstructorTestingTNewObjectStorage) *ObjectStorage {
	mock := &ObjectStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
