// Code generated by mockery v2.33.0. DO NOT EDIT.

package management

import mock "github.com/stretchr/testify/mock"

// mockJobsService is an autogenerated mock type for the jobsService type
type mockJobsService struct {
	mock.Mock
}

// StopJob provides a mock function with given fields: jobID
func (_m *mockJobsService) StopJob(jobID string) error {
	ret := _m.Called(jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockJobsService creates a new instance of mockJobsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockJobsService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockJobsService {
	mock := &mockJobsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
