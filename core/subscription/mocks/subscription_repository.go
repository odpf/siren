// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	subscription "github.com/odpf/siren/core/subscription"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// SubscriptionRepository is an autogenerated mock type for the Repository type
type SubscriptionRepository struct {
	mock.Mock
}

type SubscriptionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *SubscriptionRepository) EXPECT() *SubscriptionRepository_Expecter {
	return &SubscriptionRepository_Expecter{mock: &_m.Mock}
}

// Commit provides a mock function with given fields: ctx
func (_m *SubscriptionRepository) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionRepository_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type SubscriptionRepository_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
//  - ctx context.Context
func (_e *SubscriptionRepository_Expecter) Commit(ctx interface{}) *SubscriptionRepository_Commit_Call {
	return &SubscriptionRepository_Commit_Call{Call: _e.mock.On("Commit", ctx)}
}

func (_c *SubscriptionRepository_Commit_Call) Run(run func(ctx context.Context)) *SubscriptionRepository_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SubscriptionRepository_Commit_Call) Return(_a0 error) *SubscriptionRepository_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionRepository) Create(_a0 context.Context, _a1 *subscription.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type SubscriptionRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *subscription.Subscription
func (_e *SubscriptionRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *SubscriptionRepository_Create_Call {
	return &SubscriptionRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *SubscriptionRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *subscription.Subscription)) *SubscriptionRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*subscription.Subscription))
	})
	return _c
}

func (_c *SubscriptionRepository_Create_Call) Return(_a0 error) *SubscriptionRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionRepository) Delete(_a0 context.Context, _a1 uint64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type SubscriptionRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 uint64
func (_e *SubscriptionRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *SubscriptionRepository_Delete_Call {
	return &SubscriptionRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *SubscriptionRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 uint64)) *SubscriptionRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *SubscriptionRepository_Delete_Call) Return(_a0 error) *SubscriptionRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionRepository) Get(_a0 context.Context, _a1 uint64) (*subscription.Subscription, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *subscription.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *subscription.Subscription); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*subscription.Subscription)
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

// SubscriptionRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type SubscriptionRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 uint64
func (_e *SubscriptionRepository_Expecter) Get(_a0 interface{}, _a1 interface{}) *SubscriptionRepository_Get_Call {
	return &SubscriptionRepository_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *SubscriptionRepository_Get_Call) Run(run func(_a0 context.Context, _a1 uint64)) *SubscriptionRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *SubscriptionRepository_Get_Call) Return(_a0 *subscription.Subscription, _a1 error) *SubscriptionRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: _a0
func (_m *SubscriptionRepository) List(_a0 context.Context) ([]*subscription.Subscription, error) {
	ret := _m.Called(_a0)

	var r0 []*subscription.Subscription
	if rf, ok := ret.Get(0).(func(context.Context) []*subscription.Subscription); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*subscription.Subscription)
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

// SubscriptionRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type SubscriptionRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//  - _a0 context.Context
func (_e *SubscriptionRepository_Expecter) List(_a0 interface{}) *SubscriptionRepository_List_Call {
	return &SubscriptionRepository_List_Call{Call: _e.mock.On("List", _a0)}
}

func (_c *SubscriptionRepository_List_Call) Run(run func(_a0 context.Context)) *SubscriptionRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SubscriptionRepository_List_Call) Return(_a0 []*subscription.Subscription, _a1 error) *SubscriptionRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Rollback provides a mock function with given fields: ctx
func (_m *SubscriptionRepository) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionRepository_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type SubscriptionRepository_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//  - ctx context.Context
func (_e *SubscriptionRepository_Expecter) Rollback(ctx interface{}) *SubscriptionRepository_Rollback_Call {
	return &SubscriptionRepository_Rollback_Call{Call: _e.mock.On("Rollback", ctx)}
}

func (_c *SubscriptionRepository_Rollback_Call) Run(run func(ctx context.Context)) *SubscriptionRepository_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SubscriptionRepository_Rollback_Call) Return(_a0 error) *SubscriptionRepository_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *SubscriptionRepository) Update(_a0 context.Context, _a1 *subscription.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type SubscriptionRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *subscription.Subscription
func (_e *SubscriptionRepository_Expecter) Update(_a0 interface{}, _a1 interface{}) *SubscriptionRepository_Update_Call {
	return &SubscriptionRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *SubscriptionRepository_Update_Call) Run(run func(_a0 context.Context, _a1 *subscription.Subscription)) *SubscriptionRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*subscription.Subscription))
	})
	return _c
}

func (_c *SubscriptionRepository_Update_Call) Return(_a0 error) *SubscriptionRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

// WithTransaction provides a mock function with given fields: ctx
func (_m *SubscriptionRepository) WithTransaction(ctx context.Context) context.Context {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// SubscriptionRepository_WithTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTransaction'
type SubscriptionRepository_WithTransaction_Call struct {
	*mock.Call
}

// WithTransaction is a helper method to define mock.On call
//  - ctx context.Context
func (_e *SubscriptionRepository_Expecter) WithTransaction(ctx interface{}) *SubscriptionRepository_WithTransaction_Call {
	return &SubscriptionRepository_WithTransaction_Call{Call: _e.mock.On("WithTransaction", ctx)}
}

func (_c *SubscriptionRepository_WithTransaction_Call) Run(run func(ctx context.Context)) *SubscriptionRepository_WithTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SubscriptionRepository_WithTransaction_Call) Return(_a0 context.Context) *SubscriptionRepository_WithTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewSubscriptionRepository creates a new instance of SubscriptionRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubscriptionRepository(t testing.TB) *SubscriptionRepository {
	mock := &SubscriptionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
