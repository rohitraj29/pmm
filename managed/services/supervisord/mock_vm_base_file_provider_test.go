// Code generated by mockery v1.0.0. DO NOT EDIT.

package supervisord

import (
	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
)

// mockVmBaseFileProvider is an autogenerated mock type for the vmBaseFileProvider type
type mockVmBaseFileProvider struct {
	mock.Mock
}

// GetBaseFile provides a mock function with given fields:
func (_m *mockVmBaseFileProvider) GetBaseFile() (models.File, error) {
	ret := _m.Called()

	var r0 models.File
	if rf, ok := ret.Get(0).(func() models.File); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
