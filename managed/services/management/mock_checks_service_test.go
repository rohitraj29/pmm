// Code generated by mockery v2.33.0. DO NOT EDIT.

package management

import (
	context "context"

	check "github.com/percona-platform/saas/pkg/check"
	mock "github.com/stretchr/testify/mock"

	services "github.com/percona/pmm/managed/services"
)

// mockChecksService is an autogenerated mock type for the checksService type
type mockChecksService struct {
	mock.Mock
}

// ChangeInterval provides a mock function with given fields: params
func (_m *mockChecksService) ChangeInterval(params map[string]check.Interval) error {
	ret := _m.Called(params)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]check.Interval) error); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableChecks provides a mock function with given fields: checkNames
func (_m *mockChecksService) DisableChecks(checkNames []string) error {
	ret := _m.Called(checkNames)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(checkNames)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableChecks provides a mock function with given fields: checkNames
func (_m *mockChecksService) EnableChecks(checkNames []string) error {
	ret := _m.Called(checkNames)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(checkNames)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdvisors provides a mock function with given fields:
func (_m *mockChecksService) GetAdvisors() ([]check.Advisor, error) {
	ret := _m.Called()

	var r0 []check.Advisor
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]check.Advisor, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []check.Advisor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]check.Advisor)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChecks provides a mock function with given fields:
func (_m *mockChecksService) GetChecks() (map[string]check.Check, error) {
	ret := _m.Called()

	var r0 map[string]check.Check
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]check.Check, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]check.Check); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]check.Check)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChecksResults provides a mock function with given fields: ctx, serviceID
func (_m *mockChecksService) GetChecksResults(ctx context.Context, serviceID string) ([]services.CheckResult, error) {
	ret := _m.Called(ctx, serviceID)

	var r0 []services.CheckResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]services.CheckResult, error)); ok {
		return rf(ctx, serviceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []services.CheckResult); ok {
		r0 = rf(ctx, serviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.CheckResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, serviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDisabledChecks provides a mock function with given fields:
func (_m *mockChecksService) GetDisabledChecks() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecurityCheckResults provides a mock function with given fields:
func (_m *mockChecksService) GetSecurityCheckResults() ([]services.CheckResult, error) {
	ret := _m.Called()

	var r0 []services.CheckResult
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]services.CheckResult, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []services.CheckResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.CheckResult)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartChecks provides a mock function with given fields: checkNames
func (_m *mockChecksService) StartChecks(checkNames []string) error {
	ret := _m.Called(checkNames)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(checkNames)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockChecksService creates a new instance of mockChecksService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockChecksService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockChecksService {
	mock := &mockChecksService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
