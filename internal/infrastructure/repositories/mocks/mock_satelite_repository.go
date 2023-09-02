// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	models "prueba_meli/internal/domain/models"

	mock "github.com/stretchr/testify/mock"
)

// SateliteRepository is an autogenerated mock type for the SateliteRepository type
type SateliteRepository struct {
	mock.Mock
}

// GetAllSatelites provides a mock function with given fields:
func (_m *SateliteRepository) GetAllSatelites() ([]models.Satellite, error) {
	ret := _m.Called()

	var r0 []models.Satellite
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Satellite, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Satellite); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Satellite)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSateliteByName provides a mock function with given fields: sateliteName
func (_m *SateliteRepository) GetSateliteByName(sateliteName string) (models.Satellite, error) {
	ret := _m.Called(sateliteName)

	var r0 models.Satellite
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Satellite, error)); ok {
		return rf(sateliteName)
	}
	if rf, ok := ret.Get(0).(func(string) models.Satellite); ok {
		r0 = rf(sateliteName)
	} else {
		r0 = ret.Get(0).(models.Satellite)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sateliteName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSatelite provides a mock function with given fields: satelites
func (_m *SateliteRepository) SaveSatelite(satelites models.Satellite) error {
	ret := _m.Called(satelites)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Satellite) error); ok {
		r0 = rf(satelites)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSateliteRepository creates a new instance of SateliteRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSateliteRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SateliteRepository {
	mock := &SateliteRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}