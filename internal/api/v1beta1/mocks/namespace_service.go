// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	namespace "github.com/odpf/siren/core/namespace"
	mock "github.com/stretchr/testify/mock"
)

// NamespaceService is an autogenerated mock type for the NamespaceService type
type NamespaceService struct {
	mock.Mock
}

type NamespaceService_Expecter struct {
	mock *mock.Mock
}

func (_m *NamespaceService) EXPECT() *NamespaceService_Expecter {
	return &NamespaceService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *NamespaceService) Create(_a0 context.Context, _a1 *namespace.Namespace) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *namespace.Namespace) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type NamespaceService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *namespace.Namespace
func (_e *NamespaceService_Expecter) Create(_a0 interface{}, _a1 interface{}) *NamespaceService_Create_Call {
	return &NamespaceService_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *NamespaceService_Create_Call) Run(run func(_a0 context.Context, _a1 *namespace.Namespace)) *NamespaceService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*namespace.Namespace))
	})
	return _c
}

func (_c *NamespaceService_Create_Call) Return(_a0 error) *NamespaceService_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *NamespaceService) Delete(_a0 context.Context, _a1 uint64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type NamespaceService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *NamespaceService_Expecter) Delete(_a0 interface{}, _a1 interface{}) *NamespaceService_Delete_Call {
	return &NamespaceService_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *NamespaceService_Delete_Call) Run(run func(_a0 context.Context, _a1 uint64)) *NamespaceService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *NamespaceService_Delete_Call) Return(_a0 error) *NamespaceService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *NamespaceService) Get(_a0 context.Context, _a1 uint64) (*namespace.Namespace, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *namespace.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *namespace.Namespace); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*namespace.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type NamespaceService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *NamespaceService_Expecter) Get(_a0 interface{}, _a1 interface{}) *NamespaceService_Get_Call {
	return &NamespaceService_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *NamespaceService_Get_Call) Run(run func(_a0 context.Context, _a1 uint64)) *NamespaceService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *NamespaceService_Get_Call) Return(_a0 *namespace.Namespace, _a1 error) *NamespaceService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: _a0
func (_m *NamespaceService) List(_a0 context.Context) ([]namespace.Namespace, error) {
	ret := _m.Called(_a0)

	var r0 []namespace.Namespace
	if rf, ok := ret.Get(0).(func(context.Context) []namespace.Namespace); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]namespace.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type NamespaceService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *NamespaceService_Expecter) List(_a0 interface{}) *NamespaceService_List_Call {
	return &NamespaceService_List_Call{Call: _e.mock.On("List", _a0)}
}

func (_c *NamespaceService_List_Call) Run(run func(_a0 context.Context)) *NamespaceService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *NamespaceService_List_Call) Return(_a0 []namespace.Namespace, _a1 error) *NamespaceService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *NamespaceService) Update(_a0 context.Context, _a1 *namespace.Namespace) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *namespace.Namespace) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type NamespaceService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *namespace.Namespace
func (_e *NamespaceService_Expecter) Update(_a0 interface{}, _a1 interface{}) *NamespaceService_Update_Call {
	return &NamespaceService_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *NamespaceService_Update_Call) Run(run func(_a0 context.Context, _a1 *namespace.Namespace)) *NamespaceService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*namespace.Namespace))
	})
	return _c
}

func (_c *NamespaceService_Update_Call) Return(_a0 error) *NamespaceService_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewNamespaceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewNamespaceService creates a new instance of NamespaceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNamespaceService(t mockConstructorTestingTNewNamespaceService) *NamespaceService {
	mock := &NamespaceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
