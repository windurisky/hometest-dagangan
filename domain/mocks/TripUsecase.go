// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TripUsecase is an autogenerated mock type for the TripUsecase type
type TripUsecase struct {
	mock.Mock
}

// CalculateFare provides a mock function with given fields: _a0
func (_m *TripUsecase) CalculateFare(_a0 uint64) (uint64, error) {
	ret := _m.Called(_a0)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (uint64, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTripUsecase creates a new instance of TripUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTripUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *TripUsecase {
	mock := &TripUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}