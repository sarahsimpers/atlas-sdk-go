// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115010/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// AWSClustersDNSApi is an autogenerated mock type for the AWSClustersDNSApi type
type AWSClustersDNSApi struct {
	mock.Mock
}

type AWSClustersDNSApi_Expecter struct {
	mock *mock.Mock
}

func (_m *AWSClustersDNSApi) EXPECT() *AWSClustersDNSApi_Expecter {
	return &AWSClustersDNSApi_Expecter{mock: &_m.Mock}
}

// GetAWSCustomDNS provides a mock function with given fields: ctx, groupId
func (_m *AWSClustersDNSApi) GetAWSCustomDNS(ctx context.Context, groupId string) admin.GetAWSCustomDNSApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetAWSCustomDNS")
	}

	var r0 admin.GetAWSCustomDNSApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetAWSCustomDNSApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.GetAWSCustomDNSApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_GetAWSCustomDNS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAWSCustomDNS'
type AWSClustersDNSApi_GetAWSCustomDNS_Call struct {
	*mock.Call
}

// GetAWSCustomDNS is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *AWSClustersDNSApi_Expecter) GetAWSCustomDNS(ctx interface{}, groupId interface{}) *AWSClustersDNSApi_GetAWSCustomDNS_Call {
	return &AWSClustersDNSApi_GetAWSCustomDNS_Call{Call: _e.mock.On("GetAWSCustomDNS", ctx, groupId)}
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNS_Call) Run(run func(ctx context.Context, groupId string)) *AWSClustersDNSApi_GetAWSCustomDNS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNS_Call) Return(_a0 admin.GetAWSCustomDNSApiRequest) *AWSClustersDNSApi_GetAWSCustomDNS_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNS_Call) RunAndReturn(run func(context.Context, string) admin.GetAWSCustomDNSApiRequest) *AWSClustersDNSApi_GetAWSCustomDNS_Call {
	_c.Call.Return(run)
	return _c
}

// GetAWSCustomDNSExecute provides a mock function with given fields: r
func (_m *AWSClustersDNSApi) GetAWSCustomDNSExecute(r admin.GetAWSCustomDNSApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetAWSCustomDNSExecute")
	}

	var r0 *admin.AWSCustomDNSEnabled
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetAWSCustomDNSApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetAWSCustomDNSApiRequest) *admin.AWSCustomDNSEnabled); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AWSCustomDNSEnabled)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetAWSCustomDNSApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetAWSCustomDNSApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AWSClustersDNSApi_GetAWSCustomDNSExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAWSCustomDNSExecute'
type AWSClustersDNSApi_GetAWSCustomDNSExecute_Call struct {
	*mock.Call
}

// GetAWSCustomDNSExecute is a helper method to define mock.On call
//   - r admin.GetAWSCustomDNSApiRequest
func (_e *AWSClustersDNSApi_Expecter) GetAWSCustomDNSExecute(r interface{}) *AWSClustersDNSApi_GetAWSCustomDNSExecute_Call {
	return &AWSClustersDNSApi_GetAWSCustomDNSExecute_Call{Call: _e.mock.On("GetAWSCustomDNSExecute", r)}
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNSExecute_Call) Run(run func(r admin.GetAWSCustomDNSApiRequest)) *AWSClustersDNSApi_GetAWSCustomDNSExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetAWSCustomDNSApiRequest))
	})
	return _c
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNSExecute_Call) Return(_a0 *admin.AWSCustomDNSEnabled, _a1 *http.Response, _a2 error) *AWSClustersDNSApi_GetAWSCustomDNSExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNSExecute_Call) RunAndReturn(run func(admin.GetAWSCustomDNSApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)) *AWSClustersDNSApi_GetAWSCustomDNSExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAWSCustomDNSWithParams provides a mock function with given fields: ctx, args
func (_m *AWSClustersDNSApi) GetAWSCustomDNSWithParams(ctx context.Context, args *admin.GetAWSCustomDNSApiParams) admin.GetAWSCustomDNSApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetAWSCustomDNSWithParams")
	}

	var r0 admin.GetAWSCustomDNSApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetAWSCustomDNSApiParams) admin.GetAWSCustomDNSApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetAWSCustomDNSApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAWSCustomDNSWithParams'
type AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call struct {
	*mock.Call
}

// GetAWSCustomDNSWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetAWSCustomDNSApiParams
func (_e *AWSClustersDNSApi_Expecter) GetAWSCustomDNSWithParams(ctx interface{}, args interface{}) *AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call {
	return &AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call{Call: _e.mock.On("GetAWSCustomDNSWithParams", ctx, args)}
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call) Run(run func(ctx context.Context, args *admin.GetAWSCustomDNSApiParams)) *AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetAWSCustomDNSApiParams))
	})
	return _c
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call) Return(_a0 admin.GetAWSCustomDNSApiRequest) *AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetAWSCustomDNSApiParams) admin.GetAWSCustomDNSApiRequest) *AWSClustersDNSApi_GetAWSCustomDNSWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleAWSCustomDNS provides a mock function with given fields: ctx, groupId, aWSCustomDNSEnabled
func (_m *AWSClustersDNSApi) ToggleAWSCustomDNS(ctx context.Context, groupId string, aWSCustomDNSEnabled *admin.AWSCustomDNSEnabled) admin.ToggleAWSCustomDNSApiRequest {
	ret := _m.Called(ctx, groupId, aWSCustomDNSEnabled)

	if len(ret) == 0 {
		panic("no return value specified for ToggleAWSCustomDNS")
	}

	var r0 admin.ToggleAWSCustomDNSApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.AWSCustomDNSEnabled) admin.ToggleAWSCustomDNSApiRequest); ok {
		r0 = rf(ctx, groupId, aWSCustomDNSEnabled)
	} else {
		r0 = ret.Get(0).(admin.ToggleAWSCustomDNSApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_ToggleAWSCustomDNS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleAWSCustomDNS'
type AWSClustersDNSApi_ToggleAWSCustomDNS_Call struct {
	*mock.Call
}

// ToggleAWSCustomDNS is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - aWSCustomDNSEnabled *admin.AWSCustomDNSEnabled
func (_e *AWSClustersDNSApi_Expecter) ToggleAWSCustomDNS(ctx interface{}, groupId interface{}, aWSCustomDNSEnabled interface{}) *AWSClustersDNSApi_ToggleAWSCustomDNS_Call {
	return &AWSClustersDNSApi_ToggleAWSCustomDNS_Call{Call: _e.mock.On("ToggleAWSCustomDNS", ctx, groupId, aWSCustomDNSEnabled)}
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNS_Call) Run(run func(ctx context.Context, groupId string, aWSCustomDNSEnabled *admin.AWSCustomDNSEnabled)) *AWSClustersDNSApi_ToggleAWSCustomDNS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.AWSCustomDNSEnabled))
	})
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNS_Call) Return(_a0 admin.ToggleAWSCustomDNSApiRequest) *AWSClustersDNSApi_ToggleAWSCustomDNS_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNS_Call) RunAndReturn(run func(context.Context, string, *admin.AWSCustomDNSEnabled) admin.ToggleAWSCustomDNSApiRequest) *AWSClustersDNSApi_ToggleAWSCustomDNS_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleAWSCustomDNSExecute provides a mock function with given fields: r
func (_m *AWSClustersDNSApi) ToggleAWSCustomDNSExecute(r admin.ToggleAWSCustomDNSApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ToggleAWSCustomDNSExecute")
	}

	var r0 *admin.AWSCustomDNSEnabled
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ToggleAWSCustomDNSApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ToggleAWSCustomDNSApiRequest) *admin.AWSCustomDNSEnabled); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AWSCustomDNSEnabled)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ToggleAWSCustomDNSApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ToggleAWSCustomDNSApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleAWSCustomDNSExecute'
type AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call struct {
	*mock.Call
}

// ToggleAWSCustomDNSExecute is a helper method to define mock.On call
//   - r admin.ToggleAWSCustomDNSApiRequest
func (_e *AWSClustersDNSApi_Expecter) ToggleAWSCustomDNSExecute(r interface{}) *AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call {
	return &AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call{Call: _e.mock.On("ToggleAWSCustomDNSExecute", r)}
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call) Run(run func(r admin.ToggleAWSCustomDNSApiRequest)) *AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ToggleAWSCustomDNSApiRequest))
	})
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call) Return(_a0 *admin.AWSCustomDNSEnabled, _a1 *http.Response, _a2 error) *AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call) RunAndReturn(run func(admin.ToggleAWSCustomDNSApiRequest) (*admin.AWSCustomDNSEnabled, *http.Response, error)) *AWSClustersDNSApi_ToggleAWSCustomDNSExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleAWSCustomDNSWithParams provides a mock function with given fields: ctx, args
func (_m *AWSClustersDNSApi) ToggleAWSCustomDNSWithParams(ctx context.Context, args *admin.ToggleAWSCustomDNSApiParams) admin.ToggleAWSCustomDNSApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ToggleAWSCustomDNSWithParams")
	}

	var r0 admin.ToggleAWSCustomDNSApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ToggleAWSCustomDNSApiParams) admin.ToggleAWSCustomDNSApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ToggleAWSCustomDNSApiRequest)
	}

	return r0
}

// AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleAWSCustomDNSWithParams'
type AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call struct {
	*mock.Call
}

// ToggleAWSCustomDNSWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ToggleAWSCustomDNSApiParams
func (_e *AWSClustersDNSApi_Expecter) ToggleAWSCustomDNSWithParams(ctx interface{}, args interface{}) *AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call {
	return &AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call{Call: _e.mock.On("ToggleAWSCustomDNSWithParams", ctx, args)}
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call) Run(run func(ctx context.Context, args *admin.ToggleAWSCustomDNSApiParams)) *AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ToggleAWSCustomDNSApiParams))
	})
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call) Return(_a0 admin.ToggleAWSCustomDNSApiRequest) *AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call) RunAndReturn(run func(context.Context, *admin.ToggleAWSCustomDNSApiParams) admin.ToggleAWSCustomDNSApiRequest) *AWSClustersDNSApi_ToggleAWSCustomDNSWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewAWSClustersDNSApi creates a new instance of AWSClustersDNSApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAWSClustersDNSApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *AWSClustersDNSApi {
	mock := &AWSClustersDNSApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
