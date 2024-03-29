// Code generated by mockery v2.25.1. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "microservice/shared/dto"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: c, id
func (_m *UserUsecase) Delete(c context.Context, id string) error {
	ret := _m.Called(c, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: c, pagination
func (_m *UserUsecase) Fetch(c context.Context, pagination dto.Pagination) (dto.Pagination, error) {
	ret := _m.Called(c, pagination)

	var r0 dto.Pagination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.Pagination) (dto.Pagination, error)); ok {
		return rf(c, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.Pagination) dto.Pagination); ok {
		r0 = rf(c, pagination)
	} else {
		r0 = ret.Get(0).(dto.Pagination)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.Pagination) error); ok {
		r1 = rf(c, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, id
func (_m *UserUsecase) GetByID(c context.Context, id string) (dto.UserResponse, error) {
	ret := _m.Called(c, id)

	var r0 dto.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (dto.UserResponse, error)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) dto.UserResponse); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(dto.UserResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: c, ureq
func (_m *UserUsecase) Login(c context.Context, ureq dto.UserRequestLogin) (dto.UserResponseToken, error) {
	ret := _m.Called(c, ureq)

	var r0 dto.UserResponseToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestLogin) (dto.UserResponseToken, error)); ok {
		return rf(c, ureq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestLogin) dto.UserResponseToken); ok {
		r0 = rf(c, ureq)
	} else {
		r0 = ret.Get(0).(dto.UserResponseToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UserRequestLogin) error); ok {
		r1 = rf(c, ureq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: c, metadata
func (_m *UserUsecase) Logout(c context.Context, metadata *dto.AccessDetails) error {
	ret := _m.Called(c, metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.AccessDetails) error); ok {
		r0 = rf(c, metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Refresh provides a mock function with given fields: c, ureq
func (_m *UserUsecase) Refresh(c context.Context, ureq dto.UserRequestRefresh) (dto.UserResponseToken, error) {
	ret := _m.Called(c, ureq)

	var r0 dto.UserResponseToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestRefresh) (dto.UserResponseToken, error)); ok {
		return rf(c, ureq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestRefresh) dto.UserResponseToken); ok {
		r0 = rf(c, ureq)
	} else {
		r0 = ret.Get(0).(dto.UserResponseToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UserRequestRefresh) error); ok {
		r1 = rf(c, ureq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: c, metadata, ureq
func (_m *UserUsecase) ResetPassword(c context.Context, metadata *dto.AccessDetails, ureq dto.UserRequestPasswordUpdate) error {
	ret := _m.Called(c, metadata, ureq)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.AccessDetails, dto.UserRequestPasswordUpdate) error); ok {
		r0 = rf(c, metadata, ureq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: c, ureq
func (_m *UserUsecase) Store(c context.Context, ureq dto.UserRequestCreate) (dto.UserResponse, error) {
	ret := _m.Called(c, ureq)

	var r0 dto.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestCreate) (dto.UserResponse, error)); ok {
		return rf(c, ureq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestCreate) dto.UserResponse); ok {
		r0 = rf(c, ureq)
	} else {
		r0 = ret.Get(0).(dto.UserResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UserRequestCreate) error); ok {
		r1 = rf(c, ureq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, ureq
func (_m *UserUsecase) Update(c context.Context, ureq dto.UserRequestUpdate) (dto.UserResponse, error) {
	ret := _m.Called(c, ureq)

	var r0 dto.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestUpdate) (dto.UserResponse, error)); ok {
		return rf(c, ureq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserRequestUpdate) dto.UserResponse); ok {
		r0 = rf(c, ureq)
	} else {
		r0 = ret.Get(0).(dto.UserResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UserRequestUpdate) error); ok {
		r1 = rf(c, ureq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
