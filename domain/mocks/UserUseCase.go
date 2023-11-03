// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "DiskusiTugas/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// BlockStudent provides a mock function with given fields: id
func (_m *UserUseCase) BlockStudent(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Destroy provides a mock function with given fields: id
func (_m *UserUseCase) Destroy(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FecthStudent provides a mock function with given fields: page, pageSize
func (_m *UserUseCase) FecthStudent(page int, pageSize int) ([]domain.User, int, error) {
	ret := _m.Called(page, pageSize)

	var r0 []domain.User
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]domain.User, int, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []domain.User); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(page, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Fetch provides a mock function with given fields:
func (_m *UserUseCase) Fetch() ([]domain.User, error) {
	ret := _m.Called()

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchWithPagination provides a mock function with given fields: page, pageSize
func (_m *UserUseCase) FetchWithPagination(page int, pageSize int) ([]domain.User, int, error) {
	ret := _m.Called(page, pageSize)

	var r0 []domain.User
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]domain.User, int, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []domain.User); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(page, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserUseCase) GetByEmail(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *UserUseCase) GetByID(id int) (domain.User, error) {
	ret := _m.Called(id)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) domain.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: user
func (_m *UserUseCase) Store(user *domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: user
func (_m *UserUseCase) Update(user *domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}