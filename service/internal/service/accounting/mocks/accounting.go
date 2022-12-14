// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "user_balance/service/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Accounting is an autogenerated mock type for the Accounting type
type Accounting struct {
	mock.Mock
}

// List provides a mock function with given fields: input
func (_m *Accounting) List(input *models.AccountingListIn) (*string, error) {
	ret := _m.Called(input)

	var r0 *string
	if rf, ok := ret.Get(0).(func(*models.AccountingListIn) *string); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.AccountingListIn) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAccounting interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccounting creates a new instance of Accounting. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccounting(t mockConstructorTestingTNewAccounting) *Accounting {
	mock := &Accounting{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
