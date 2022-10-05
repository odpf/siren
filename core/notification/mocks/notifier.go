// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	notification "github.com/odpf/siren/core/notification"
	mock "github.com/stretchr/testify/mock"
)

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

type Notifier_Expecter struct {
	mock *mock.Mock
}

func (_m *Notifier) EXPECT() *Notifier_Expecter {
	return &Notifier_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: ctx, message
func (_m *Notifier) Publish(ctx context.Context, message notification.Message) (bool, error) {
	ret := _m.Called(ctx, message)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, notification.Message) bool); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, notification.Message) error); ok {
		r1 = rf(ctx, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Notifier_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type Notifier_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//  - ctx context.Context
//  - message notification.Message
func (_e *Notifier_Expecter) Publish(ctx interface{}, message interface{}) *Notifier_Publish_Call {
	return &Notifier_Publish_Call{Call: _e.mock.On("Publish", ctx, message)}
}

func (_c *Notifier_Publish_Call) Run(run func(ctx context.Context, message notification.Message)) *Notifier_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(notification.Message))
	})
	return _c
}

func (_c *Notifier_Publish_Call) Return(_a0 bool, _a1 error) *Notifier_Publish_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ValidateConfigMap provides a mock function with given fields: notificationConfigMap
func (_m *Notifier) ValidateConfigMap(notificationConfigMap map[string]interface{}) error {
	ret := _m.Called(notificationConfigMap)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) error); ok {
		r0 = rf(notificationConfigMap)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Notifier_ValidateConfigMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateConfigMap'
type Notifier_ValidateConfigMap_Call struct {
	*mock.Call
}

// ValidateConfigMap is a helper method to define mock.On call
//  - notificationConfigMap map[string]interface{}
func (_e *Notifier_Expecter) ValidateConfigMap(notificationConfigMap interface{}) *Notifier_ValidateConfigMap_Call {
	return &Notifier_ValidateConfigMap_Call{Call: _e.mock.On("ValidateConfigMap", notificationConfigMap)}
}

func (_c *Notifier_ValidateConfigMap_Call) Run(run func(notificationConfigMap map[string]interface{})) *Notifier_ValidateConfigMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]interface{}))
	})
	return _c
}

func (_c *Notifier_ValidateConfigMap_Call) Return(_a0 error) *Notifier_ValidateConfigMap_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewNotifier interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotifier creates a new instance of Notifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotifier(t mockConstructorTestingTNewNotifier) *Notifier {
	mock := &Notifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
