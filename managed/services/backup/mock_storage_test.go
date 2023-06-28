// Code generated by mockery v2.30.16. DO NOT EDIT.

package backup

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	minio "github.com/percona/pmm/managed/services/minio"
)

// MockStorage is an autogenerated mock type for the Storage type
type MockStorage struct {
	mock.Mock
}

// FileStat provides a mock function with given fields: ctx, endpoint, accessKey, secretKey, bucketName, name
func (_m *MockStorage) FileStat(ctx context.Context, endpoint string, accessKey string, secretKey string, bucketName string, name string) (minio.FileInfo, error) {
	ret := _m.Called(ctx, endpoint, accessKey, secretKey, bucketName, name)

	var r0 minio.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) (minio.FileInfo, error)); ok {
		return rf(ctx, endpoint, accessKey, secretKey, bucketName, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) minio.FileInfo); ok {
		r0 = rf(ctx, endpoint, accessKey, secretKey, bucketName, name)
	} else {
		r0 = ret.Get(0).(minio.FileInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string) error); ok {
		r1 = rf(ctx, endpoint, accessKey, secretKey, bucketName, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, endpoint, accessKey, secretKey, bucketName, prefix, suffix
func (_m *MockStorage) List(ctx context.Context, endpoint string, accessKey string, secretKey string, bucketName string, prefix string, suffix string) ([]minio.FileInfo, error) {
	ret := _m.Called(ctx, endpoint, accessKey, secretKey, bucketName, prefix, suffix)

	var r0 []minio.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string) ([]minio.FileInfo, error)); ok {
		return rf(ctx, endpoint, accessKey, secretKey, bucketName, prefix, suffix)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string) []minio.FileInfo); ok {
		r0 = rf(ctx, endpoint, accessKey, secretKey, bucketName, prefix, suffix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]minio.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string, string) error); ok {
		r1 = rf(ctx, endpoint, accessKey, secretKey, bucketName, prefix, suffix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, endpoint, accessKey, secretKey, bucketName, objectName
func (_m *MockStorage) Remove(ctx context.Context, endpoint string, accessKey string, secretKey string, bucketName string, objectName string) error {
	ret := _m.Called(ctx, endpoint, accessKey, secretKey, bucketName, objectName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) error); ok {
		r0 = rf(ctx, endpoint, accessKey, secretKey, bucketName, objectName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveRecursive provides a mock function with given fields: ctx, endpoint, accessKey, secretKey, bucketName, prefix
func (_m *MockStorage) RemoveRecursive(ctx context.Context, endpoint string, accessKey string, secretKey string, bucketName string, prefix string) error {
	ret := _m.Called(ctx, endpoint, accessKey, secretKey, bucketName, prefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) error); ok {
		r0 = rf(ctx, endpoint, accessKey, secretKey, bucketName, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockStorage creates a new instance of MockStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorage(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MockStorage {
	mock := &MockStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
