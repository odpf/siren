// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	alert "github.com/odpf/siren/core/alert"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AlertService is an autogenerated mock type for the AlertService type
type AlertService struct {
	mock.Mock
}

type AlertService_Expecter struct {
	mock *mock.Mock
}

func (_m *AlertService) EXPECT() *AlertService_Expecter {
	return &AlertService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *AlertService) Create(_a0 context.Context, _a1 []*alert.Alert) ([]alert.Alert, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []alert.Alert
	if rf, ok := ret.Get(0).(func(context.Context, []*alert.Alert) []alert.Alert); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]alert.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*alert.Alert) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AlertService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []*alert.Alert
func (_e *AlertService_Expecter) Create(_a0 interface{}, _a1 interface{}) *AlertService_Create_Call {
	return &AlertService_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *AlertService_Create_Call) Run(run func(_a0 context.Context, _a1 []*alert.Alert)) *AlertService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*alert.Alert))
	})
	return _c
}

func (_c *AlertService_Create_Call) Return(_a0 []alert.Alert, _a1 error) *AlertService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *AlertService) List(_a0 context.Context, _a1 alert.Filter) ([]alert.Alert, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []alert.Alert
	if rf, ok := ret.Get(0).(func(context.Context, alert.Filter) []alert.Alert); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]alert.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, alert.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type AlertService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 alert.Filter
func (_e *AlertService_Expecter) List(_a0 interface{}, _a1 interface{}) *AlertService_List_Call {
	return &AlertService_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *AlertService_List_Call) Run(run func(_a0 context.Context, _a1 alert.Filter)) *AlertService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(alert.Filter))
	})
	return _c
}

func (_c *AlertService_List_Call) Return(_a0 []alert.Alert, _a1 error) *AlertService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewAlertService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAlertService creates a new instance of AlertService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAlertService(t mockConstructorTestingTNewAlertService) *AlertService {
	mock := &AlertService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
