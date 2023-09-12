// Code generated by mockery v2.33.0. DO NOT EDIT.

package versioner

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockExecFunctions is an autogenerated mock type for the ExecFunctions type
type MockExecFunctions struct {
	mock.Mock
}

// CommandContext provides a mock function with given fields: ctx, name, arg
func (_m *MockExecFunctions) CommandContext(ctx context.Context, name string, arg ...string) CombinedOutputer {
	_va := make([]interface{}, len(arg))
	for _i := range arg {
		_va[_i] = arg[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 CombinedOutputer
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) CombinedOutputer); ok {
		r0 = rf(ctx, name, arg...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(CombinedOutputer)
		}
	}

	return r0
}

// LookPath provides a mock function with given fields: file
func (_m *MockExecFunctions) LookPath(file string) (string, error) {
	ret := _m.Called(file)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(file)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockExecFunctions creates a new instance of MockExecFunctions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExecFunctions(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MockExecFunctions {
	mock := &MockExecFunctions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
