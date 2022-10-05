// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// HTTPCaller is an autogenerated mock type for the HTTPCaller type
type HTTPCaller struct {
	mock.Mock
}

type HTTPCaller_Expecter struct {
	mock *mock.Mock
}

func (_m *HTTPCaller) EXPECT() *HTTPCaller_Expecter {
	return &HTTPCaller_Expecter{mock: &_m.Mock}
}

// Notify provides a mock function with given fields: ctx, apiURL, body
func (_m *HTTPCaller) Notify(ctx context.Context, apiURL string, body []byte) error {
	ret := _m.Called(ctx, apiURL, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, apiURL, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HTTPCaller_Notify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Notify'
type HTTPCaller_Notify_Call struct {
	*mock.Call
}

// Notify is a helper method to define mock.On call
//  - ctx context.Context
//  - apiURL string
//  - body []byte
func (_e *HTTPCaller_Expecter) Notify(ctx interface{}, apiURL interface{}, body interface{}) *HTTPCaller_Notify_Call {
	return &HTTPCaller_Notify_Call{Call: _e.mock.On("Notify", ctx, apiURL, body)}
}

func (_c *HTTPCaller_Notify_Call) Run(run func(ctx context.Context, apiURL string, body []byte)) *HTTPCaller_Notify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *HTTPCaller_Notify_Call) Return(_a0 error) *HTTPCaller_Notify_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewHTTPCaller interface {
	mock.TestingT
	Cleanup(func())
}

// NewHTTPCaller creates a new instance of HTTPCaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHTTPCaller(t mockConstructorTestingTNewHTTPCaller) *HTTPCaller {
	mock := &HTTPCaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
