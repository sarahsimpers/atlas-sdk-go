// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240530003/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ThirdPartyIntegrationsApi is an autogenerated mock type for the ThirdPartyIntegrationsApi type
type ThirdPartyIntegrationsApi struct {
	mock.Mock
}

type ThirdPartyIntegrationsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *ThirdPartyIntegrationsApi) EXPECT() *ThirdPartyIntegrationsApi_Expecter {
	return &ThirdPartyIntegrationsApi_Expecter{mock: &_m.Mock}
}

// CreateThirdPartyIntegration provides a mock function with given fields: ctx, integrationType, groupId, thirdPartyIntegration
func (_m *ThirdPartyIntegrationsApi) CreateThirdPartyIntegration(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *admin.ThirdPartyIntegration) admin.CreateThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, integrationType, groupId, thirdPartyIntegration)

	if len(ret) == 0 {
		panic("no return value specified for CreateThirdPartyIntegration")
	}

	var r0 admin.CreateThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.ThirdPartyIntegration) admin.CreateThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, integrationType, groupId, thirdPartyIntegration)
	} else {
		r0 = ret.Get(0).(admin.CreateThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateThirdPartyIntegration'
type ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call struct {
	*mock.Call
}

// CreateThirdPartyIntegration is a helper method to define mock.On call
//   - ctx context.Context
//   - integrationType string
//   - groupId string
//   - thirdPartyIntegration *admin.ThirdPartyIntegration
func (_e *ThirdPartyIntegrationsApi_Expecter) CreateThirdPartyIntegration(ctx interface{}, integrationType interface{}, groupId interface{}, thirdPartyIntegration interface{}) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call {
	return &ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call{Call: _e.mock.On("CreateThirdPartyIntegration", ctx, integrationType, groupId, thirdPartyIntegration)}
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call) Run(run func(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *admin.ThirdPartyIntegration)) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.ThirdPartyIntegration))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call) Return(_a0 admin.CreateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call) RunAndReturn(run func(context.Context, string, string, *admin.ThirdPartyIntegration) admin.CreateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegration_Call {
	_c.Call.Return(run)
	return _c
}

// CreateThirdPartyIntegrationExecute provides a mock function with given fields: r
func (_m *ThirdPartyIntegrationsApi) CreateThirdPartyIntegrationExecute(r admin.CreateThirdPartyIntegrationApiRequest) (*admin.PaginatedIntegration, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateThirdPartyIntegrationExecute")
	}

	var r0 *admin.PaginatedIntegration
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateThirdPartyIntegrationApiRequest) (*admin.PaginatedIntegration, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateThirdPartyIntegrationApiRequest) *admin.PaginatedIntegration); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedIntegration)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateThirdPartyIntegrationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateThirdPartyIntegrationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateThirdPartyIntegrationExecute'
type ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call struct {
	*mock.Call
}

// CreateThirdPartyIntegrationExecute is a helper method to define mock.On call
//   - r admin.CreateThirdPartyIntegrationApiRequest
func (_e *ThirdPartyIntegrationsApi_Expecter) CreateThirdPartyIntegrationExecute(r interface{}) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call {
	return &ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call{Call: _e.mock.On("CreateThirdPartyIntegrationExecute", r)}
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call) Run(run func(r admin.CreateThirdPartyIntegrationApiRequest)) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateThirdPartyIntegrationApiRequest))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call) Return(_a0 *admin.PaginatedIntegration, _a1 *http.Response, _a2 error) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call) RunAndReturn(run func(admin.CreateThirdPartyIntegrationApiRequest) (*admin.PaginatedIntegration, *http.Response, error)) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateThirdPartyIntegrationWithParams provides a mock function with given fields: ctx, args
func (_m *ThirdPartyIntegrationsApi) CreateThirdPartyIntegrationWithParams(ctx context.Context, args *admin.CreateThirdPartyIntegrationApiParams) admin.CreateThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateThirdPartyIntegrationWithParams")
	}

	var r0 admin.CreateThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateThirdPartyIntegrationApiParams) admin.CreateThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateThirdPartyIntegrationWithParams'
type ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call struct {
	*mock.Call
}

// CreateThirdPartyIntegrationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateThirdPartyIntegrationApiParams
func (_e *ThirdPartyIntegrationsApi_Expecter) CreateThirdPartyIntegrationWithParams(ctx interface{}, args interface{}) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call {
	return &ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call{Call: _e.mock.On("CreateThirdPartyIntegrationWithParams", ctx, args)}
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateThirdPartyIntegrationApiParams)) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateThirdPartyIntegrationApiParams))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call) Return(_a0 admin.CreateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateThirdPartyIntegrationApiParams) admin.CreateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_CreateThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteThirdPartyIntegration provides a mock function with given fields: ctx, integrationType, groupId
func (_m *ThirdPartyIntegrationsApi) DeleteThirdPartyIntegration(ctx context.Context, integrationType string, groupId string) admin.DeleteThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, integrationType, groupId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteThirdPartyIntegration")
	}

	var r0 admin.DeleteThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.DeleteThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, integrationType, groupId)
	} else {
		r0 = ret.Get(0).(admin.DeleteThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteThirdPartyIntegration'
type ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call struct {
	*mock.Call
}

// DeleteThirdPartyIntegration is a helper method to define mock.On call
//   - ctx context.Context
//   - integrationType string
//   - groupId string
func (_e *ThirdPartyIntegrationsApi_Expecter) DeleteThirdPartyIntegration(ctx interface{}, integrationType interface{}, groupId interface{}) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call {
	return &ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call{Call: _e.mock.On("DeleteThirdPartyIntegration", ctx, integrationType, groupId)}
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call) Run(run func(ctx context.Context, integrationType string, groupId string)) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call) Return(_a0 admin.DeleteThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call) RunAndReturn(run func(context.Context, string, string) admin.DeleteThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegration_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteThirdPartyIntegrationExecute provides a mock function with given fields: r
func (_m *ThirdPartyIntegrationsApi) DeleteThirdPartyIntegrationExecute(r admin.DeleteThirdPartyIntegrationApiRequest) (interface{}, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteThirdPartyIntegrationExecute")
	}

	var r0 interface{}
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DeleteThirdPartyIntegrationApiRequest) (interface{}, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteThirdPartyIntegrationApiRequest) interface{}); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteThirdPartyIntegrationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DeleteThirdPartyIntegrationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteThirdPartyIntegrationExecute'
type ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call struct {
	*mock.Call
}

// DeleteThirdPartyIntegrationExecute is a helper method to define mock.On call
//   - r admin.DeleteThirdPartyIntegrationApiRequest
func (_e *ThirdPartyIntegrationsApi_Expecter) DeleteThirdPartyIntegrationExecute(r interface{}) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call {
	return &ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call{Call: _e.mock.On("DeleteThirdPartyIntegrationExecute", r)}
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call) Run(run func(r admin.DeleteThirdPartyIntegrationApiRequest)) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteThirdPartyIntegrationApiRequest))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call) Return(_a0 interface{}, _a1 *http.Response, _a2 error) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call) RunAndReturn(run func(admin.DeleteThirdPartyIntegrationApiRequest) (interface{}, *http.Response, error)) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteThirdPartyIntegrationWithParams provides a mock function with given fields: ctx, args
func (_m *ThirdPartyIntegrationsApi) DeleteThirdPartyIntegrationWithParams(ctx context.Context, args *admin.DeleteThirdPartyIntegrationApiParams) admin.DeleteThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteThirdPartyIntegrationWithParams")
	}

	var r0 admin.DeleteThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteThirdPartyIntegrationApiParams) admin.DeleteThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteThirdPartyIntegrationWithParams'
type ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call struct {
	*mock.Call
}

// DeleteThirdPartyIntegrationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteThirdPartyIntegrationApiParams
func (_e *ThirdPartyIntegrationsApi_Expecter) DeleteThirdPartyIntegrationWithParams(ctx interface{}, args interface{}) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call {
	return &ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call{Call: _e.mock.On("DeleteThirdPartyIntegrationWithParams", ctx, args)}
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteThirdPartyIntegrationApiParams)) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteThirdPartyIntegrationApiParams))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call) Return(_a0 admin.DeleteThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteThirdPartyIntegrationApiParams) admin.DeleteThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_DeleteThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetThirdPartyIntegration provides a mock function with given fields: ctx, groupId, integrationType
func (_m *ThirdPartyIntegrationsApi) GetThirdPartyIntegration(ctx context.Context, groupId string, integrationType string) admin.GetThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, groupId, integrationType)

	if len(ret) == 0 {
		panic("no return value specified for GetThirdPartyIntegration")
	}

	var r0 admin.GetThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, groupId, integrationType)
	} else {
		r0 = ret.Get(0).(admin.GetThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetThirdPartyIntegration'
type ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call struct {
	*mock.Call
}

// GetThirdPartyIntegration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - integrationType string
func (_e *ThirdPartyIntegrationsApi_Expecter) GetThirdPartyIntegration(ctx interface{}, groupId interface{}, integrationType interface{}) *ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call {
	return &ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call{Call: _e.mock.On("GetThirdPartyIntegration", ctx, groupId, integrationType)}
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call) Run(run func(ctx context.Context, groupId string, integrationType string)) *ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call) Return(_a0 admin.GetThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call) RunAndReturn(run func(context.Context, string, string) admin.GetThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_GetThirdPartyIntegration_Call {
	_c.Call.Return(run)
	return _c
}

// GetThirdPartyIntegrationExecute provides a mock function with given fields: r
func (_m *ThirdPartyIntegrationsApi) GetThirdPartyIntegrationExecute(r admin.GetThirdPartyIntegrationApiRequest) (*admin.ThirdPartyIntegration, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetThirdPartyIntegrationExecute")
	}

	var r0 *admin.ThirdPartyIntegration
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetThirdPartyIntegrationApiRequest) (*admin.ThirdPartyIntegration, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetThirdPartyIntegrationApiRequest) *admin.ThirdPartyIntegration); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ThirdPartyIntegration)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetThirdPartyIntegrationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetThirdPartyIntegrationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetThirdPartyIntegrationExecute'
type ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call struct {
	*mock.Call
}

// GetThirdPartyIntegrationExecute is a helper method to define mock.On call
//   - r admin.GetThirdPartyIntegrationApiRequest
func (_e *ThirdPartyIntegrationsApi_Expecter) GetThirdPartyIntegrationExecute(r interface{}) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call {
	return &ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call{Call: _e.mock.On("GetThirdPartyIntegrationExecute", r)}
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call) Run(run func(r admin.GetThirdPartyIntegrationApiRequest)) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetThirdPartyIntegrationApiRequest))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call) Return(_a0 *admin.ThirdPartyIntegration, _a1 *http.Response, _a2 error) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call) RunAndReturn(run func(admin.GetThirdPartyIntegrationApiRequest) (*admin.ThirdPartyIntegration, *http.Response, error)) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetThirdPartyIntegrationWithParams provides a mock function with given fields: ctx, args
func (_m *ThirdPartyIntegrationsApi) GetThirdPartyIntegrationWithParams(ctx context.Context, args *admin.GetThirdPartyIntegrationApiParams) admin.GetThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetThirdPartyIntegrationWithParams")
	}

	var r0 admin.GetThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetThirdPartyIntegrationApiParams) admin.GetThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetThirdPartyIntegrationWithParams'
type ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call struct {
	*mock.Call
}

// GetThirdPartyIntegrationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetThirdPartyIntegrationApiParams
func (_e *ThirdPartyIntegrationsApi_Expecter) GetThirdPartyIntegrationWithParams(ctx interface{}, args interface{}) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call {
	return &ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call{Call: _e.mock.On("GetThirdPartyIntegrationWithParams", ctx, args)}
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call) Run(run func(ctx context.Context, args *admin.GetThirdPartyIntegrationApiParams)) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetThirdPartyIntegrationApiParams))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call) Return(_a0 admin.GetThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetThirdPartyIntegrationApiParams) admin.GetThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_GetThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListThirdPartyIntegrations provides a mock function with given fields: ctx, groupId
func (_m *ThirdPartyIntegrationsApi) ListThirdPartyIntegrations(ctx context.Context, groupId string) admin.ListThirdPartyIntegrationsApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListThirdPartyIntegrations")
	}

	var r0 admin.ListThirdPartyIntegrationsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListThirdPartyIntegrationsApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListThirdPartyIntegrationsApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListThirdPartyIntegrations'
type ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call struct {
	*mock.Call
}

// ListThirdPartyIntegrations is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *ThirdPartyIntegrationsApi_Expecter) ListThirdPartyIntegrations(ctx interface{}, groupId interface{}) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call {
	return &ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call{Call: _e.mock.On("ListThirdPartyIntegrations", ctx, groupId)}
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call) Run(run func(ctx context.Context, groupId string)) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call) Return(_a0 admin.ListThirdPartyIntegrationsApiRequest) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call) RunAndReturn(run func(context.Context, string) admin.ListThirdPartyIntegrationsApiRequest) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrations_Call {
	_c.Call.Return(run)
	return _c
}

// ListThirdPartyIntegrationsExecute provides a mock function with given fields: r
func (_m *ThirdPartyIntegrationsApi) ListThirdPartyIntegrationsExecute(r admin.ListThirdPartyIntegrationsApiRequest) (*admin.PaginatedIntegration, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListThirdPartyIntegrationsExecute")
	}

	var r0 *admin.PaginatedIntegration
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListThirdPartyIntegrationsApiRequest) (*admin.PaginatedIntegration, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListThirdPartyIntegrationsApiRequest) *admin.PaginatedIntegration); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedIntegration)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListThirdPartyIntegrationsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListThirdPartyIntegrationsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListThirdPartyIntegrationsExecute'
type ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call struct {
	*mock.Call
}

// ListThirdPartyIntegrationsExecute is a helper method to define mock.On call
//   - r admin.ListThirdPartyIntegrationsApiRequest
func (_e *ThirdPartyIntegrationsApi_Expecter) ListThirdPartyIntegrationsExecute(r interface{}) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call {
	return &ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call{Call: _e.mock.On("ListThirdPartyIntegrationsExecute", r)}
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call) Run(run func(r admin.ListThirdPartyIntegrationsApiRequest)) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListThirdPartyIntegrationsApiRequest))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call) Return(_a0 *admin.PaginatedIntegration, _a1 *http.Response, _a2 error) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call) RunAndReturn(run func(admin.ListThirdPartyIntegrationsApiRequest) (*admin.PaginatedIntegration, *http.Response, error)) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListThirdPartyIntegrationsWithParams provides a mock function with given fields: ctx, args
func (_m *ThirdPartyIntegrationsApi) ListThirdPartyIntegrationsWithParams(ctx context.Context, args *admin.ListThirdPartyIntegrationsApiParams) admin.ListThirdPartyIntegrationsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListThirdPartyIntegrationsWithParams")
	}

	var r0 admin.ListThirdPartyIntegrationsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListThirdPartyIntegrationsApiParams) admin.ListThirdPartyIntegrationsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListThirdPartyIntegrationsApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListThirdPartyIntegrationsWithParams'
type ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call struct {
	*mock.Call
}

// ListThirdPartyIntegrationsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListThirdPartyIntegrationsApiParams
func (_e *ThirdPartyIntegrationsApi_Expecter) ListThirdPartyIntegrationsWithParams(ctx interface{}, args interface{}) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call {
	return &ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call{Call: _e.mock.On("ListThirdPartyIntegrationsWithParams", ctx, args)}
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListThirdPartyIntegrationsApiParams)) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListThirdPartyIntegrationsApiParams))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call) Return(_a0 admin.ListThirdPartyIntegrationsApiRequest) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListThirdPartyIntegrationsApiParams) admin.ListThirdPartyIntegrationsApiRequest) *ThirdPartyIntegrationsApi_ListThirdPartyIntegrationsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateThirdPartyIntegration provides a mock function with given fields: ctx, integrationType, groupId, thirdPartyIntegration
func (_m *ThirdPartyIntegrationsApi) UpdateThirdPartyIntegration(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *admin.ThirdPartyIntegration) admin.UpdateThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, integrationType, groupId, thirdPartyIntegration)

	if len(ret) == 0 {
		panic("no return value specified for UpdateThirdPartyIntegration")
	}

	var r0 admin.UpdateThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.ThirdPartyIntegration) admin.UpdateThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, integrationType, groupId, thirdPartyIntegration)
	} else {
		r0 = ret.Get(0).(admin.UpdateThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateThirdPartyIntegration'
type ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call struct {
	*mock.Call
}

// UpdateThirdPartyIntegration is a helper method to define mock.On call
//   - ctx context.Context
//   - integrationType string
//   - groupId string
//   - thirdPartyIntegration *admin.ThirdPartyIntegration
func (_e *ThirdPartyIntegrationsApi_Expecter) UpdateThirdPartyIntegration(ctx interface{}, integrationType interface{}, groupId interface{}, thirdPartyIntegration interface{}) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call {
	return &ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call{Call: _e.mock.On("UpdateThirdPartyIntegration", ctx, integrationType, groupId, thirdPartyIntegration)}
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call) Run(run func(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *admin.ThirdPartyIntegration)) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.ThirdPartyIntegration))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call) Return(_a0 admin.UpdateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call) RunAndReturn(run func(context.Context, string, string, *admin.ThirdPartyIntegration) admin.UpdateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegration_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateThirdPartyIntegrationExecute provides a mock function with given fields: r
func (_m *ThirdPartyIntegrationsApi) UpdateThirdPartyIntegrationExecute(r admin.UpdateThirdPartyIntegrationApiRequest) (*admin.PaginatedIntegration, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateThirdPartyIntegrationExecute")
	}

	var r0 *admin.PaginatedIntegration
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.UpdateThirdPartyIntegrationApiRequest) (*admin.PaginatedIntegration, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.UpdateThirdPartyIntegrationApiRequest) *admin.PaginatedIntegration); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedIntegration)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.UpdateThirdPartyIntegrationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.UpdateThirdPartyIntegrationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateThirdPartyIntegrationExecute'
type ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call struct {
	*mock.Call
}

// UpdateThirdPartyIntegrationExecute is a helper method to define mock.On call
//   - r admin.UpdateThirdPartyIntegrationApiRequest
func (_e *ThirdPartyIntegrationsApi_Expecter) UpdateThirdPartyIntegrationExecute(r interface{}) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call {
	return &ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call{Call: _e.mock.On("UpdateThirdPartyIntegrationExecute", r)}
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call) Run(run func(r admin.UpdateThirdPartyIntegrationApiRequest)) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.UpdateThirdPartyIntegrationApiRequest))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call) Return(_a0 *admin.PaginatedIntegration, _a1 *http.Response, _a2 error) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call) RunAndReturn(run func(admin.UpdateThirdPartyIntegrationApiRequest) (*admin.PaginatedIntegration, *http.Response, error)) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateThirdPartyIntegrationWithParams provides a mock function with given fields: ctx, args
func (_m *ThirdPartyIntegrationsApi) UpdateThirdPartyIntegrationWithParams(ctx context.Context, args *admin.UpdateThirdPartyIntegrationApiParams) admin.UpdateThirdPartyIntegrationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for UpdateThirdPartyIntegrationWithParams")
	}

	var r0 admin.UpdateThirdPartyIntegrationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.UpdateThirdPartyIntegrationApiParams) admin.UpdateThirdPartyIntegrationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.UpdateThirdPartyIntegrationApiRequest)
	}

	return r0
}

// ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateThirdPartyIntegrationWithParams'
type ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call struct {
	*mock.Call
}

// UpdateThirdPartyIntegrationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.UpdateThirdPartyIntegrationApiParams
func (_e *ThirdPartyIntegrationsApi_Expecter) UpdateThirdPartyIntegrationWithParams(ctx interface{}, args interface{}) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call {
	return &ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call{Call: _e.mock.On("UpdateThirdPartyIntegrationWithParams", ctx, args)}
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call) Run(run func(ctx context.Context, args *admin.UpdateThirdPartyIntegrationApiParams)) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.UpdateThirdPartyIntegrationApiParams))
	})
	return _c
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call) Return(_a0 admin.UpdateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call) RunAndReturn(run func(context.Context, *admin.UpdateThirdPartyIntegrationApiParams) admin.UpdateThirdPartyIntegrationApiRequest) *ThirdPartyIntegrationsApi_UpdateThirdPartyIntegrationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewThirdPartyIntegrationsApi creates a new instance of ThirdPartyIntegrationsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewThirdPartyIntegrationsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *ThirdPartyIntegrationsApi {
	mock := &ThirdPartyIntegrationsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
