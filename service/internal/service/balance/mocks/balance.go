// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "user_balance/service/internal/models"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	vo "user_balance/service/internal/vo"
)

// Balance is an autogenerated mock type for the Balance type
type Balance struct {
	mock.Mock
}

// Add provides a mock function with given fields: input
func (_m *Balance) Add(input *models.TransactionFields) (sql.Result, error) {
	ret := _m.Called(input)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(*models.TransactionFields) sql.Result); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.TransactionFields) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: userId
func (_m *Balance) Get(userId *vo.IntID) (*models.BalanceFields, error) {
	ret := _m.Called(userId)

	var r0 *models.BalanceFields
	if rf, ok := ret.Get(0).(func(*vo.IntID) *models.BalanceFields); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BalanceFields)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*vo.IntID) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBalance interface {
	mock.TestingT
	Cleanup(func())
}

// NewBalance creates a new instance of Balance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBalance(t mockConstructorTestingTNewBalance) *Balance {
	mock := &Balance{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
