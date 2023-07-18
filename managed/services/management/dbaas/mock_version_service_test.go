// Code generated by mockery v2.32.0. DO NOT EDIT.

package dbaas

import (
	context "context"

	version "github.com/hashicorp/go-version"
	mock "github.com/stretchr/testify/mock"
)

// mockVersionService is an autogenerated mock type for the versionService type
type mockVersionService struct {
	mock.Mock
}

// GetNextDatabaseImage provides a mock function with given fields: ctx, operatorType, operatorVersion, installedDBVersion
func (_m *mockVersionService) GetNextDatabaseImage(ctx context.Context, operatorType string, operatorVersion string, installedDBVersion string) (string, error) {
	ret := _m.Called(ctx, operatorType, operatorVersion, installedDBVersion)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (string, error)); ok {
		return rf(ctx, operatorType, operatorVersion, installedDBVersion)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) string); ok {
		r0 = rf(ctx, operatorType, operatorVersion, installedDBVersion)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, operatorType, operatorVersion, installedDBVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVersionServiceURL provides a mock function with given fields:
func (_m *mockVersionService) GetVersionServiceURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsDatabaseVersionSupportedByOperator provides a mock function with given fields: ctx, operatorType, operatorVersion, databaseVersion
func (_m *mockVersionService) IsDatabaseVersionSupportedByOperator(ctx context.Context, operatorType string, operatorVersion string, databaseVersion string) (bool, error) {
	ret := _m.Called(ctx, operatorType, operatorVersion, databaseVersion)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (bool, error)); ok {
		return rf(ctx, operatorType, operatorVersion, databaseVersion)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) bool); ok {
		r0 = rf(ctx, operatorType, operatorVersion, databaseVersion)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, operatorType, operatorVersion, databaseVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestOperatorVersion provides a mock function with given fields: ctx, pmmVersion
func (_m *mockVersionService) LatestOperatorVersion(ctx context.Context, pmmVersion string) (*version.Version, *version.Version, error) {
	ret := _m.Called(ctx, pmmVersion)

	var r0 *version.Version
	var r1 *version.Version
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*version.Version, *version.Version, error)); ok {
		return rf(ctx, pmmVersion)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *version.Version); ok {
		r0 = rf(ctx, pmmVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *version.Version); ok {
		r1 = rf(ctx, pmmVersion)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*version.Version)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, pmmVersion)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Matrix provides a mock function with given fields: ctx, params
func (_m *mockVersionService) Matrix(ctx context.Context, params componentsParams) (*VersionServiceResponse, error) {
	ret := _m.Called(ctx, params)

	var r0 *VersionServiceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, componentsParams) (*VersionServiceResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, componentsParams) *VersionServiceResponse); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VersionServiceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, componentsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NextOperatorVersion provides a mock function with given fields: ctx, operatorType, installedVersion
func (_m *mockVersionService) NextOperatorVersion(ctx context.Context, operatorType string, installedVersion string) (*version.Version, error) {
	ret := _m.Called(ctx, operatorType, installedVersion)

	var r0 *version.Version
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*version.Version, error)); ok {
		return rf(ctx, operatorType, installedVersion)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *version.Version); ok {
		r0 = rf(ctx, operatorType, installedVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, operatorType, installedVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportedOperatorVersionsList provides a mock function with given fields: ctx, pmmVersion
func (_m *mockVersionService) SupportedOperatorVersionsList(ctx context.Context, pmmVersion string) (map[string][]string, error) {
	ret := _m.Called(ctx, pmmVersion)

	var r0 map[string][]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (map[string][]string, error)); ok {
		return rf(ctx, pmmVersion)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string][]string); ok {
		r0 = rf(ctx, pmmVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pmmVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockVersionService creates a new instance of mockVersionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockVersionService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockVersionService {
	mock := &mockVersionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
