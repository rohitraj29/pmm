// Code generated by mockery v2.33.0. DO NOT EDIT.

package client

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	agentpb "github.com/percona/pmm/api/agentpb"
)

// mockServiceInfoBroker is an autogenerated mock type for the serviceInfoBroker type
type mockServiceInfoBroker struct {
	mock.Mock
}

// GetInfoFromService provides a mock function with given fields: ctx, req, id
func (_m *mockServiceInfoBroker) GetInfoFromService(ctx context.Context, req *agentpb.ServiceInfoRequest, id uint32) *agentpb.ServiceInfoResponse {
	ret := _m.Called(ctx, req, id)

	var r0 *agentpb.ServiceInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *agentpb.ServiceInfoRequest, uint32) *agentpb.ServiceInfoResponse); ok {
		r0 = rf(ctx, req, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*agentpb.ServiceInfoResponse)
		}
	}

	return r0
}

// newMockServiceInfoBroker creates a new instance of mockServiceInfoBroker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockServiceInfoBroker(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockServiceInfoBroker {
	mock := &mockServiceInfoBroker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
