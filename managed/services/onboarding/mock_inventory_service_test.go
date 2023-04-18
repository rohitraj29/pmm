// Code generated by mockery v1.0.0. DO NOT EDIT.

package onboarding

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	inventorypb "github.com/percona/pmm/api/inventorypb"
	models "github.com/percona/pmm/managed/models"
)

// mockInventoryService is an autogenerated mock type for the inventoryService type
type mockInventoryService struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx, filters
func (_m *mockInventoryService) List(ctx context.Context, filters models.ServiceFilters) ([]inventorypb.Service, error) {
	ret := _m.Called(ctx, filters)

	var r0 []inventorypb.Service
	if rf, ok := ret.Get(0).(func(context.Context, models.ServiceFilters) []inventorypb.Service); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]inventorypb.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ServiceFilters) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
