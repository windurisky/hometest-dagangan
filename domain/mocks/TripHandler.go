// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/windurisky/hometest-dagangan/domain"
)

// TripHandler is an autogenerated mock type for the TripHandler type
type TripHandler struct {
	mock.Mock
}

// ParseInput provides a mock function with given fields: _a0
func (_m *TripHandler) ParseInput(_a0 string) (domain.Trip, error) {
	ret := _m.Called(_a0)

	var r0 domain.Trip
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.Trip, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) domain.Trip); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.Trip)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SummarizeTrip provides a mock function with given fields: _a0
func (_m *TripHandler) SummarizeTrip(_a0 []domain.Trip) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.Trip) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTripHandler creates a new instance of TripHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTripHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *TripHandler {
	mock := &TripHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}