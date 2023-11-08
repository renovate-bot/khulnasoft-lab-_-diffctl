// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	cloudfront "github.com/aws/aws-sdk-go/service/cloudfront"
	mock "github.com/stretchr/testify/mock"
)

// MockCloudfrontRepository is an autogenerated mock type for the CloudfrontRepository type
type MockCloudfrontRepository struct {
	mock.Mock
}

// ListAllDistributions provides a mock function with given fields:
func (_m *MockCloudfrontRepository) ListAllDistributions() ([]*cloudfront.DistributionSummary, error) {
	ret := _m.Called()

	var r0 []*cloudfront.DistributionSummary
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*cloudfront.DistributionSummary, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*cloudfront.DistributionSummary); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*cloudfront.DistributionSummary)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockCloudfrontRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCloudfrontRepository creates a new instance of MockCloudfrontRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCloudfrontRepository(t mockConstructorTestingTNewMockCloudfrontRepository) *MockCloudfrontRepository {
	mock := &MockCloudfrontRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
