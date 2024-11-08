// Code generated by mockery v2.46.3. DO NOT EDIT.

package affected_apps

import (
	context "context"

	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MockargoClient is an autogenerated mock type for the argoClient type
type MockargoClient struct {
	mock.Mock
}

type MockargoClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockargoClient) EXPECT() *MockargoClient_Expecter {
	return &MockargoClient_Expecter{mock: &_m.Mock}
}

// GetApplications provides a mock function with given fields: ctx
func (_m *MockargoClient) GetApplications(ctx context.Context) (*v1alpha1.ApplicationList, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetApplications")
	}

	var r0 *v1alpha1.ApplicationList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*v1alpha1.ApplicationList, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *v1alpha1.ApplicationList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ApplicationList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockargoClient_GetApplications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApplications'
type MockargoClient_GetApplications_Call struct {
	*mock.Call
}

// GetApplications is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockargoClient_Expecter) GetApplications(ctx interface{}) *MockargoClient_GetApplications_Call {
	return &MockargoClient_GetApplications_Call{Call: _e.mock.On("GetApplications", ctx)}
}

func (_c *MockargoClient_GetApplications_Call) Run(run func(ctx context.Context)) *MockargoClient_GetApplications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockargoClient_GetApplications_Call) Return(_a0 *v1alpha1.ApplicationList, _a1 error) *MockargoClient_GetApplications_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockargoClient_GetApplications_Call) RunAndReturn(run func(context.Context) (*v1alpha1.ApplicationList, error)) *MockargoClient_GetApplications_Call {
	_c.Call.Return(run)
	return _c
}

// GetApplicationsByAppset provides a mock function with given fields: ctx, appsetName
func (_m *MockargoClient) GetApplicationsByAppset(ctx context.Context, appsetName string) (*v1alpha1.ApplicationList, error) {
	ret := _m.Called(ctx, appsetName)

	if len(ret) == 0 {
		panic("no return value specified for GetApplicationsByAppset")
	}

	var r0 *v1alpha1.ApplicationList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*v1alpha1.ApplicationList, error)); ok {
		return rf(ctx, appsetName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *v1alpha1.ApplicationList); ok {
		r0 = rf(ctx, appsetName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ApplicationList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, appsetName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockargoClient_GetApplicationsByAppset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApplicationsByAppset'
type MockargoClient_GetApplicationsByAppset_Call struct {
	*mock.Call
}

// GetApplicationsByAppset is a helper method to define mock.On call
//   - ctx context.Context
//   - appsetName string
func (_e *MockargoClient_Expecter) GetApplicationsByAppset(ctx interface{}, appsetName interface{}) *MockargoClient_GetApplicationsByAppset_Call {
	return &MockargoClient_GetApplicationsByAppset_Call{Call: _e.mock.On("GetApplicationsByAppset", ctx, appsetName)}
}

func (_c *MockargoClient_GetApplicationsByAppset_Call) Run(run func(ctx context.Context, appsetName string)) *MockargoClient_GetApplicationsByAppset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockargoClient_GetApplicationsByAppset_Call) Return(_a0 *v1alpha1.ApplicationList, _a1 error) *MockargoClient_GetApplicationsByAppset_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockargoClient_GetApplicationsByAppset_Call) RunAndReturn(run func(context.Context, string) (*v1alpha1.ApplicationList, error)) *MockargoClient_GetApplicationsByAppset_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockargoClient creates a new instance of MockargoClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockargoClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockargoClient {
	mock := &MockargoClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
