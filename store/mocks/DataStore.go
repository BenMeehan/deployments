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

	model "github.com/mendersoftware/deployments/model"
	mock "github.com/stretchr/testify/mock"

	store "github.com/mendersoftware/deployments/store"

	time "time"
)

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// AbortDeviceDeployments provides a mock function with given fields: ctx, deploymentID
func (_m *DataStore) AbortDeviceDeployments(ctx context.Context, deploymentID string) error {
	ret := _m.Called(ctx, deploymentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, deploymentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AggregateDeviceDeploymentByStatus provides a mock function with given fields: ctx, id
func (_m *DataStore) AggregateDeviceDeploymentByStatus(ctx context.Context, id string) (model.Stats, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Stats
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Stats); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssignArtifact provides a mock function with given fields: ctx, deviceID, deploymentID, artifact
func (_m *DataStore) AssignArtifact(ctx context.Context, deviceID string, deploymentID string, artifact *model.Image) error {
	ret := _m.Called(ctx, deviceID, deploymentID, artifact)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *model.Image) error); ok {
		r0 = rf(ctx, deviceID, deploymentID, artifact)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecommissionDeviceDeployments provides a mock function with given fields: ctx, deviceId
func (_m *DataStore) DecommissionDeviceDeployments(ctx context.Context, deviceId string) error {
	ret := _m.Called(ctx, deviceId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, deviceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeployment provides a mock function with given fields: ctx, id
func (_m *DataStore) DeleteDeployment(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceDeploymentsHistory provides a mock function with given fields: ctx, deviceId
func (_m *DataStore) DeleteDeviceDeploymentsHistory(ctx context.Context, deviceId string) error {
	ret := _m.Called(ctx, deviceId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, deviceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteImage provides a mock function with given fields: ctx, id
func (_m *DataStore) DeleteImage(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceCountByDeployment provides a mock function with given fields: ctx, id
func (_m *DataStore) DeviceCountByDeployment(ctx context.Context, id string) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistAssignedImageWithIDAndStatuses provides a mock function with given fields: ctx, id, statuses
func (_m *DataStore) ExistAssignedImageWithIDAndStatuses(ctx context.Context, id string, statuses ...model.DeviceDeploymentStatus) (bool, error) {
	_va := make([]interface{}, len(statuses))
	for _i := range statuses {
		_va[_i] = statuses[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, ...model.DeviceDeploymentStatus) bool); ok {
		r0 = rf(ctx, id, statuses...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...model.DeviceDeploymentStatus) error); ok {
		r1 = rf(ctx, id, statuses...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistByArtifactId provides a mock function with given fields: ctx, id
func (_m *DataStore) ExistByArtifactId(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistUnfinishedByArtifactId provides a mock function with given fields: ctx, id
func (_m *DataStore) ExistUnfinishedByArtifactId(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistUnfinishedByArtifactName provides a mock function with given fields: ctx, artifactName
func (_m *DataStore) ExistUnfinishedByArtifactName(ctx context.Context, artifactName string) (bool, error) {
	ret := _m.Called(ctx, artifactName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, artifactName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, artifactName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields: ctx, id
func (_m *DataStore) Exists(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, query
func (_m *DataStore) Find(ctx context.Context, query model.Query) ([]*model.Deployment, int64, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, model.Query) []*model.Deployment); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Deployment)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, model.Query) int64); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.Query) error); ok {
		r2 = rf(ctx, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindDeploymentByID provides a mock function with given fields: ctx, id
func (_m *DataStore) FindDeploymentByID(ctx context.Context, id string) (*model.Deployment, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Deployment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDeploymentStatsByIDs provides a mock function with given fields: ctx, ids
func (_m *DataStore) FindDeploymentStatsByIDs(ctx context.Context, ids ...string) ([]*model.DeploymentStats, error) {
	_va := make([]interface{}, len(ids))
	for _i := range ids {
		_va[_i] = ids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*model.DeploymentStats
	if rf, ok := ret.Get(0).(func(context.Context, ...string) []*model.DeploymentStats); ok {
		r0 = rf(ctx, ids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DeploymentStats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, ids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindImageByID provides a mock function with given fields: ctx, id
func (_m *DataStore) FindImageByID(ctx context.Context, id string) (*model.Image, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Image
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Image); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLatestInactiveDeviceDeployment provides a mock function with given fields: ctx, deviceID
func (_m *DataStore) FindLatestInactiveDeviceDeployment(ctx context.Context, deviceID string) (*model.DeviceDeployment, error) {
	ret := _m.Called(ctx, deviceID)

	var r0 *model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.DeviceDeployment); ok {
		r0 = rf(ctx, deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNewerActiveDeployments provides a mock function with given fields: ctx, createdAfter, skip, limit
func (_m *DataStore) FindNewerActiveDeployments(ctx context.Context, createdAfter *time.Time, skip int, limit int) ([]*model.Deployment, error) {
	ret := _m.Called(ctx, createdAfter, skip, limit)

	var r0 []*model.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *time.Time, int, int) []*model.Deployment); ok {
		r0 = rf(ctx, createdAfter, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *time.Time, int, int) error); ok {
		r1 = rf(ctx, createdAfter, skip, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOldestActiveDeviceDeployment provides a mock function with given fields: ctx, deviceID
func (_m *DataStore) FindOldestActiveDeviceDeployment(ctx context.Context, deviceID string) (*model.DeviceDeployment, error) {
	ret := _m.Called(ctx, deviceID)

	var r0 *model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.DeviceDeployment); ok {
		r0 = rf(ctx, deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUnfinishedByID provides a mock function with given fields: ctx, id
func (_m *DataStore) FindUnfinishedByID(ctx context.Context, id string) (*model.Deployment, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Deployment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceDeployment provides a mock function with given fields: ctx, deploymentID, deviceID, includeDeleted
func (_m *DataStore) GetDeviceDeployment(ctx context.Context, deploymentID string, deviceID string, includeDeleted bool) (*model.DeviceDeployment, error) {
	ret := _m.Called(ctx, deploymentID, deviceID, includeDeleted)

	var r0 *model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) *model.DeviceDeployment); ok {
		r0 = rf(ctx, deploymentID, deviceID, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool) error); ok {
		r1 = rf(ctx, deploymentID, deviceID, includeDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceDeploymentLog provides a mock function with given fields: ctx, deviceID, deploymentID
func (_m *DataStore) GetDeviceDeploymentLog(ctx context.Context, deviceID string, deploymentID string) (*model.DeploymentLog, error) {
	ret := _m.Called(ctx, deviceID, deploymentID)

	var r0 *model.DeploymentLog
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.DeploymentLog); ok {
		r0 = rf(ctx, deviceID, deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeploymentLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, deviceID, deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceDeployments provides a mock function with given fields: ctx, skip, limit, deviceID, active, includeDeleted
func (_m *DataStore) GetDeviceDeployments(ctx context.Context, skip int, limit int, deviceID string, active *bool, includeDeleted bool) ([]model.DeviceDeployment, error) {
	ret := _m.Called(ctx, skip, limit, deviceID, active, includeDeleted)

	var r0 []model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string, *bool, bool) []model.DeviceDeployment); ok {
		r0 = rf(ctx, skip, limit, deviceID, active, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, string, *bool, bool) error); ok {
		r1 = rf(ctx, skip, limit, deviceID, active, includeDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceDeploymentsForDevice provides a mock function with given fields: ctx, query
func (_m *DataStore) GetDeviceDeploymentsForDevice(ctx context.Context, query store.ListQueryDeviceDeployments) ([]model.DeviceDeployment, int, error) {
	ret := _m.Called(ctx, query)

	var r0 []model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, store.ListQueryDeviceDeployments) []model.DeviceDeployment); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceDeployment)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, store.ListQueryDeviceDeployments) int); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, store.ListQueryDeviceDeployments) error); ok {
		r2 = rf(ctx, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDeviceStatusesForDeployment provides a mock function with given fields: ctx, deploymentID
func (_m *DataStore) GetDeviceStatusesForDeployment(ctx context.Context, deploymentID string) ([]model.DeviceDeployment, error) {
	ret := _m.Called(ctx, deploymentID)

	var r0 []model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.DeviceDeployment); ok {
		r0 = rf(ctx, deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicesListForDeployment provides a mock function with given fields: ctx, query
func (_m *DataStore) GetDevicesListForDeployment(ctx context.Context, query store.ListQuery) ([]model.DeviceDeployment, int, error) {
	ret := _m.Called(ctx, query)

	var r0 []model.DeviceDeployment
	if rf, ok := ret.Get(0).(func(context.Context, store.ListQuery) []model.DeviceDeployment); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceDeployment)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, store.ListQuery) int); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, store.ListQuery) error); ok {
		r2 = rf(ctx, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLimit provides a mock function with given fields: ctx, name
func (_m *DataStore) GetLimit(ctx context.Context, name string) (*model.Limit, error) {
	ret := _m.Called(ctx, name)

	var r0 *model.Limit
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Limit); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Limit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleases provides a mock function with given fields: ctx, filt
func (_m *DataStore) GetReleases(ctx context.Context, filt *model.ReleaseOrImageFilter) ([]model.Release, int, error) {
	ret := _m.Called(ctx, filt)

	var r0 []model.Release
	if rf, ok := ret.Get(0).(func(context.Context, *model.ReleaseOrImageFilter) []model.Release); ok {
		r0 = rf(ctx, filt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Release)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, *model.ReleaseOrImageFilter) int); ok {
		r1 = rf(ctx, filt)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *model.ReleaseOrImageFilter) error); ok {
		r2 = rf(ctx, filt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetStorageSettings provides a mock function with given fields: ctx
func (_m *DataStore) GetStorageSettings(ctx context.Context) (*model.StorageSettings, error) {
	ret := _m.Called(ctx)

	var r0 *model.StorageSettings
	if rf, ok := ret.Get(0).(func(context.Context) *model.StorageSettings); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StorageSettings)
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

// GetTenantDbs provides a mock function with given fields:
func (_m *DataStore) GetTenantDbs() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// HasDeploymentForDevice provides a mock function with given fields: ctx, deploymentID, deviceID
func (_m *DataStore) HasDeploymentForDevice(ctx context.Context, deploymentID string, deviceID string) (bool, error) {
	ret := _m.Called(ctx, deploymentID, deviceID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, deploymentID, deviceID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, deploymentID, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageByIdsAndDeviceType provides a mock function with given fields: ctx, ids, deviceType
func (_m *DataStore) ImageByIdsAndDeviceType(ctx context.Context, ids []string, deviceType string) (*model.Image, error) {
	ret := _m.Called(ctx, ids, deviceType)

	var r0 *model.Image
	if rf, ok := ret.Get(0).(func(context.Context, []string, string) *model.Image); ok {
		r0 = rf(ctx, ids, deviceType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, string) error); ok {
		r1 = rf(ctx, ids, deviceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageByNameAndDeviceType provides a mock function with given fields: ctx, name, deviceType
func (_m *DataStore) ImageByNameAndDeviceType(ctx context.Context, name string, deviceType string) (*model.Image, error) {
	ret := _m.Called(ctx, name, deviceType)

	var r0 *model.Image
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Image); ok {
		r0 = rf(ctx, name, deviceType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, deviceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImagesByName provides a mock function with given fields: ctx, artifactName
func (_m *DataStore) ImagesByName(ctx context.Context, artifactName string) ([]*model.Image, error) {
	ret := _m.Called(ctx, artifactName)

	var r0 []*model.Image
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.Image); ok {
		r0 = rf(ctx, artifactName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, artifactName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementDeploymentDeviceCount provides a mock function with given fields: ctx, deploymentID, increment
func (_m *DataStore) IncrementDeploymentDeviceCount(ctx context.Context, deploymentID string, increment int) error {
	ret := _m.Called(ctx, deploymentID, increment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, deploymentID, increment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IncrementDeploymentTotalSize provides a mock function with given fields: ctx, deploymentID, increment
func (_m *DataStore) IncrementDeploymentTotalSize(ctx context.Context, deploymentID string, increment int64) error {
	ret := _m.Called(ctx, deploymentID, increment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, deploymentID, increment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertDeployment provides a mock function with given fields: ctx, deployment
func (_m *DataStore) InsertDeployment(ctx context.Context, deployment *model.Deployment) error {
	ret := _m.Called(ctx, deployment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Deployment) error); ok {
		r0 = rf(ctx, deployment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertDeviceDeployment provides a mock function with given fields: ctx, deviceDeployment, incrementDeviceCount
func (_m *DataStore) InsertDeviceDeployment(ctx context.Context, deviceDeployment *model.DeviceDeployment, incrementDeviceCount bool) error {
	ret := _m.Called(ctx, deviceDeployment, incrementDeviceCount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.DeviceDeployment, bool) error); ok {
		r0 = rf(ctx, deviceDeployment, incrementDeviceCount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertImage provides a mock function with given fields: ctx, image
func (_m *DataStore) InsertImage(ctx context.Context, image *model.Image) error {
	ret := _m.Called(ctx, image)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Image) error); ok {
		r0 = rf(ctx, image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMany provides a mock function with given fields: ctx, deployment
func (_m *DataStore) InsertMany(ctx context.Context, deployment ...*model.DeviceDeployment) error {
	_va := make([]interface{}, len(deployment))
	for _i := range deployment {
		_va[_i] = deployment[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*model.DeviceDeployment) error); ok {
		r0 = rf(ctx, deployment...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertUploadIntent provides a mock function with given fields: ctx, link
func (_m *DataStore) InsertUploadIntent(ctx context.Context, link *model.UploadLink) error {
	ret := _m.Called(ctx, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UploadLink) error); ok {
		r0 = rf(ctx, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsArtifactUnique provides a mock function with given fields: ctx, artifactName, deviceTypesCompatible
func (_m *DataStore) IsArtifactUnique(ctx context.Context, artifactName string, deviceTypesCompatible []string) (bool, error) {
	ret := _m.Called(ctx, artifactName, deviceTypesCompatible)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) bool); ok {
		r0 = rf(ctx, artifactName, deviceTypesCompatible)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, artifactName, deviceTypesCompatible)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImages provides a mock function with given fields: ctx, filt
func (_m *DataStore) ListImages(ctx context.Context, filt *model.ReleaseOrImageFilter) ([]*model.Image, int, error) {
	ret := _m.Called(ctx, filt)

	var r0 []*model.Image
	if rf, ok := ret.Get(0).(func(context.Context, *model.ReleaseOrImageFilter) []*model.Image); ok {
		r0 = rf(ctx, filt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Image)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, *model.ReleaseOrImageFilter) int); ok {
		r1 = rf(ctx, filt)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *model.ReleaseOrImageFilter) error); ok {
		r2 = rf(ctx, filt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Ping provides a mock function with given fields: ctx
func (_m *DataStore) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionTenant provides a mock function with given fields: ctx, tenantId
func (_m *DataStore) ProvisionTenant(ctx context.Context, tenantId string) error {
	ret := _m.Called(ctx, tenantId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tenantId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDeviceDeploymentLog provides a mock function with given fields: ctx, log
func (_m *DataStore) SaveDeviceDeploymentLog(ctx context.Context, log model.DeploymentLog) error {
	ret := _m.Called(ctx, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeploymentLog) error); ok {
		r0 = rf(ctx, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDeviceDeploymentRequest provides a mock function with given fields: ctx, ID, request
func (_m *DataStore) SaveDeviceDeploymentRequest(ctx context.Context, ID string, request *model.DeploymentNextRequest) error {
	ret := _m.Called(ctx, ID, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.DeploymentNextRequest) error); ok {
		r0 = rf(ctx, ID, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeploymentDeviceCount provides a mock function with given fields: ctx, deploymentID, count
func (_m *DataStore) SetDeploymentDeviceCount(ctx context.Context, deploymentID string, count int) error {
	ret := _m.Called(ctx, deploymentID, count)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, deploymentID, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeploymentStatus provides a mock function with given fields: ctx, id, status, now
func (_m *DataStore) SetDeploymentStatus(ctx context.Context, id string, status model.DeploymentStatus, now time.Time) error {
	ret := _m.Called(ctx, id, status, now)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.DeploymentStatus, time.Time) error); ok {
		r0 = rf(ctx, id, status, now)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetStorageSettings provides a mock function with given fields: ctx, storageSettings
func (_m *DataStore) SetStorageSettings(ctx context.Context, storageSettings *model.StorageSettings) error {
	ret := _m.Called(ctx, storageSettings)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.StorageSettings) error); ok {
		r0 = rf(ctx, storageSettings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, image
func (_m *DataStore) Update(ctx context.Context, image *model.Image) (bool, error) {
	ret := _m.Called(ctx, image)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *model.Image) bool); ok {
		r0 = rf(ctx, image)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Image) error); ok {
		r1 = rf(ctx, image)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeploymentsWithArtifactName provides a mock function with given fields: ctx, artifactName, artifactIDs
func (_m *DataStore) UpdateDeploymentsWithArtifactName(ctx context.Context, artifactName string, artifactIDs []string) error {
	ret := _m.Called(ctx, artifactName, artifactIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, artifactName, artifactIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceDeploymentLogAvailability provides a mock function with given fields: ctx, deviceID, deploymentID, log
func (_m *DataStore) UpdateDeviceDeploymentLogAvailability(ctx context.Context, deviceID string, deploymentID string, log bool) error {
	ret := _m.Called(ctx, deviceID, deploymentID, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, deviceID, deploymentID, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceDeploymentStatus provides a mock function with given fields: ctx, deviceID, deploymentID, state
func (_m *DataStore) UpdateDeviceDeploymentStatus(ctx context.Context, deviceID string, deploymentID string, state model.DeviceDeploymentState) (model.DeviceDeploymentStatus, error) {
	ret := _m.Called(ctx, deviceID, deploymentID, state)

	var r0 model.DeviceDeploymentStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.DeviceDeploymentState) model.DeviceDeploymentStatus); ok {
		r0 = rf(ctx, deviceID, deploymentID, state)
	} else {
		r0 = ret.Get(0).(model.DeviceDeploymentStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, model.DeviceDeploymentState) error); ok {
		r1 = rf(ctx, deviceID, deploymentID, state)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStats provides a mock function with given fields: ctx, id, stats
func (_m *DataStore) UpdateStats(ctx context.Context, id string, stats model.Stats) error {
	ret := _m.Called(ctx, id, stats)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Stats) error); ok {
		r0 = rf(ctx, id, stats)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatsInc provides a mock function with given fields: ctx, id, stateFrom, stateTo
func (_m *DataStore) UpdateStatsInc(ctx context.Context, id string, stateFrom model.DeviceDeploymentStatus, stateTo model.DeviceDeploymentStatus) error {
	ret := _m.Called(ctx, id, stateFrom, stateTo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.DeviceDeploymentStatus, model.DeviceDeploymentStatus) error); ok {
		r0 = rf(ctx, id, stateFrom, stateTo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUploadIntentStatus provides a mock function with given fields: ctx, id, from, to
func (_m *DataStore) UpdateUploadIntentStatus(ctx context.Context, id string, from model.LinkStatus, to model.LinkStatus) error {
	ret := _m.Called(ctx, id, from, to)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.LinkStatus, model.LinkStatus) error); ok {
		r0 = rf(ctx, id, from, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDataStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataStore creates a new instance of DataStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataStore(t mockConstructorTestingTNewDataStore) *DataStore {
	mock := &DataStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
