// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/odpf/siren/domain"
	mock "github.com/stretchr/testify/mock"
)

// AlertmanagerService is an autogenerated mock type for the AlertmanagerService type
type AlertmanagerService struct {
	mock.Mock
}

// Get provides a mock function with given fields: teamName
func (_m *AlertmanagerService) Get(teamName string) (domain.AlertCredential, error) {
	ret := _m.Called(teamName)

	var r0 domain.AlertCredential
	if rf, ok := ret.Get(0).(func(string) domain.AlertCredential); ok {
		r0 = rf(teamName)
	} else {
		r0 = ret.Get(0).(domain.AlertCredential)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(teamName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields:
func (_m *AlertmanagerService) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: credential
func (_m *AlertmanagerService) Upsert(credential domain.AlertCredential) error {
	ret := _m.Called(credential)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.AlertCredential) error); ok {
		r0 = rf(credential)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
