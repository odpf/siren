// Code generated by mockery v2.10.0. DO NOT EDIT.

package rules

import (
	context "context"

	domain "github.com/odpf/siren/domain"
	mock "github.com/stretchr/testify/mock"
)

// RuleRepositoryMock is an autogenerated mock type for the RuleRepositoryMock type
type RuleRepositoryMock struct {
	mock.Mock
}

// Commit provides a mock function with given fields: ctx
func (_m *RuleRepositoryMock) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *RuleRepositoryMock) Get(_a0 context.Context, _a1 string, _a2 string, _a3 string, _a4 string, _a5 uint64) ([]domain.Rule, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 []domain.Rule
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, uint64) []domain.Rule); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields:
func (_m *RuleRepositoryMock) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields: ctx
func (_m *RuleRepositoryMock) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *RuleRepositoryMock) Upsert(_a0 context.Context, _a1 *domain.Rule) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Rule) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithTransaction provides a mock function with given fields: ctx
func (_m *RuleRepositoryMock) WithTransaction(ctx context.Context) context.Context {
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
