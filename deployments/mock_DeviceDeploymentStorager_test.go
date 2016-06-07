// Copyright 2016 Mender Software AS
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

package deployments

import "github.com/stretchr/testify/mock"

// MockDeviceDeploymentStorager is an autogenerated mock type for the DeviceDeploymentStorager type
type MockDeviceDeploymentStorager struct {
	mock.Mock
}

// ExistAssignedImageWithIDAndStatuses provides a mock function with given fields: id, statuses
func (_m *MockDeviceDeploymentStorager) ExistAssignedImageWithIDAndStatuses(id string, statuses ...string) (bool, error) {
	ret := _m.Called(id, statuses)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...string) bool); ok {
		r0 = rf(id, statuses...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(id, statuses...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOldestDeploymentForDeviceIDWithStatuses provides a mock function with given fields: deviceID, statuses
func (_m *MockDeviceDeploymentStorager) FindOldestDeploymentForDeviceIDWithStatuses(deviceID string, statuses ...string) (*DeviceDeployment, error) {
	ret := _m.Called(deviceID, statuses)

	var r0 *DeviceDeployment
	if rf, ok := ret.Get(0).(func(string, ...string) *DeviceDeployment); ok {
		r0 = rf(deviceID, statuses...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(deviceID, statuses...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertMany provides a mock function with given fields: deployment
func (_m *MockDeviceDeploymentStorager) InsertMany(deployment ...*DeviceDeployment) error {
	ret := _m.Called(deployment)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*DeviceDeployment) error); ok {
		r0 = rf(deployment...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
