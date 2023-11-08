// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	iam "github.com/aws/aws-sdk-go/service/iam"
	mock "github.com/stretchr/testify/mock"
)

// MockIAMRepository is an autogenerated mock type for the IAMRepository type
type MockIAMRepository struct {
	mock.Mock
}

// ListAllAccessKeys provides a mock function with given fields: _a0
func (_m *MockIAMRepository) ListAllAccessKeys(_a0 []*iam.User) ([]*iam.AccessKeyMetadata, error) {
	ret := _m.Called(_a0)

	var r0 []*iam.AccessKeyMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func([]*iam.User) ([]*iam.AccessKeyMetadata, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]*iam.User) []*iam.AccessKeyMetadata); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*iam.AccessKeyMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func([]*iam.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllGroupPolicies provides a mock function with given fields: _a0
func (_m *MockIAMRepository) ListAllGroupPolicies(_a0 []*iam.Group) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func([]*iam.Group) ([]string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]*iam.Group) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]*iam.Group) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllGroupPolicyAttachments provides a mock function with given fields: _a0
func (_m *MockIAMRepository) ListAllGroupPolicyAttachments(_a0 []*iam.Group) ([]*AttachedGroupPolicy, error) {
	ret := _m.Called(_a0)

	var r0 []*AttachedGroupPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func([]*iam.Group) ([]*AttachedGroupPolicy, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]*iam.Group) []*AttachedGroupPolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*AttachedGroupPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func([]*iam.Group) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllGroups provides a mock function with given fields:
func (_m *MockIAMRepository) ListAllGroups() ([]*iam.Group, error) {
	ret := _m.Called()

	var r0 []*iam.Group
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*iam.Group, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*iam.Group); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*iam.Group)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllPolicies provides a mock function with given fields:
func (_m *MockIAMRepository) ListAllPolicies() ([]*iam.Policy, error) {
	ret := _m.Called()

	var r0 []*iam.Policy
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*iam.Policy, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*iam.Policy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*iam.Policy)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllRolePolicies provides a mock function with given fields: _a0
func (_m *MockIAMRepository) ListAllRolePolicies(_a0 []*iam.Role) ([]RolePolicy, error) {
	ret := _m.Called(_a0)

	var r0 []RolePolicy
	var r1 error
	if rf, ok := ret.Get(0).(func([]*iam.Role) ([]RolePolicy, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]*iam.Role) []RolePolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RolePolicy)
		}
	}

	if rf, ok := ret.Get(1).(func([]*iam.Role) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllRolePolicyAttachments provides a mock function with given fields: _a0
func (_m *MockIAMRepository) ListAllRolePolicyAttachments(_a0 []*iam.Role) ([]*AttachedRolePolicy, error) {
	ret := _m.Called(_a0)

	var r0 []*AttachedRolePolicy
	var r1 error
	if rf, ok := ret.Get(0).(func([]*iam.Role) ([]*AttachedRolePolicy, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]*iam.Role) []*AttachedRolePolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*AttachedRolePolicy)
		}
	}

	if rf, ok := ret.Get(1).(func([]*iam.Role) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllRoles provides a mock function with given fields:
func (_m *MockIAMRepository) ListAllRoles() ([]*iam.Role, error) {
	ret := _m.Called()

	var r0 []*iam.Role
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*iam.Role, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*iam.Role); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*iam.Role)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllUserPolicies provides a mock function with given fields: _a0
func (_m *MockIAMRepository) ListAllUserPolicies(_a0 []*iam.User) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func([]*iam.User) ([]string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]*iam.User) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]*iam.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllUserPolicyAttachments provides a mock function with given fields: _a0
func (_m *MockIAMRepository) ListAllUserPolicyAttachments(_a0 []*iam.User) ([]*AttachedUserPolicy, error) {
	ret := _m.Called(_a0)

	var r0 []*AttachedUserPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func([]*iam.User) ([]*AttachedUserPolicy, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]*iam.User) []*AttachedUserPolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*AttachedUserPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func([]*iam.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllUsers provides a mock function with given fields:
func (_m *MockIAMRepository) ListAllUsers() ([]*iam.User, error) {
	ret := _m.Called()

	var r0 []*iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*iam.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*iam.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIAMRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIAMRepository creates a new instance of MockIAMRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIAMRepository(t mockConstructorTestingTNewMockIAMRepository) *MockIAMRepository {
	mock := &MockIAMRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
