// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240530002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// InvoicesApi is an autogenerated mock type for the InvoicesApi type
type InvoicesApi struct {
	mock.Mock
}

type InvoicesApi_Expecter struct {
	mock *mock.Mock
}

func (_m *InvoicesApi) EXPECT() *InvoicesApi_Expecter {
	return &InvoicesApi_Expecter{mock: &_m.Mock}
}

// CreateCostExplorerQueryProcess provides a mock function with given fields: ctx, orgId, costExplorerFilterRequestBody
func (_m *InvoicesApi) CreateCostExplorerQueryProcess(ctx context.Context, orgId string, costExplorerFilterRequestBody *admin.CostExplorerFilterRequestBody) admin.CreateCostExplorerQueryProcessApiRequest {
	ret := _m.Called(ctx, orgId, costExplorerFilterRequestBody)

	if len(ret) == 0 {
		panic("no return value specified for CreateCostExplorerQueryProcess")
	}

	var r0 admin.CreateCostExplorerQueryProcessApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.CostExplorerFilterRequestBody) admin.CreateCostExplorerQueryProcessApiRequest); ok {
		r0 = rf(ctx, orgId, costExplorerFilterRequestBody)
	} else {
		r0 = ret.Get(0).(admin.CreateCostExplorerQueryProcessApiRequest)
	}

	return r0
}

// InvoicesApi_CreateCostExplorerQueryProcess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCostExplorerQueryProcess'
type InvoicesApi_CreateCostExplorerQueryProcess_Call struct {
	*mock.Call
}

// CreateCostExplorerQueryProcess is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
//   - costExplorerFilterRequestBody *admin.CostExplorerFilterRequestBody
func (_e *InvoicesApi_Expecter) CreateCostExplorerQueryProcess(ctx interface{}, orgId interface{}, costExplorerFilterRequestBody interface{}) *InvoicesApi_CreateCostExplorerQueryProcess_Call {
	return &InvoicesApi_CreateCostExplorerQueryProcess_Call{Call: _e.mock.On("CreateCostExplorerQueryProcess", ctx, orgId, costExplorerFilterRequestBody)}
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess_Call) Run(run func(ctx context.Context, orgId string, costExplorerFilterRequestBody *admin.CostExplorerFilterRequestBody)) *InvoicesApi_CreateCostExplorerQueryProcess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.CostExplorerFilterRequestBody))
	})
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess_Call) Return(_a0 admin.CreateCostExplorerQueryProcessApiRequest) *InvoicesApi_CreateCostExplorerQueryProcess_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess_Call) RunAndReturn(run func(context.Context, string, *admin.CostExplorerFilterRequestBody) admin.CreateCostExplorerQueryProcessApiRequest) *InvoicesApi_CreateCostExplorerQueryProcess_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCostExplorerQueryProcess1 provides a mock function with given fields: ctx, orgId, token
func (_m *InvoicesApi) CreateCostExplorerQueryProcess1(ctx context.Context, orgId string, token string) admin.CreateCostExplorerQueryProcess1ApiRequest {
	ret := _m.Called(ctx, orgId, token)

	if len(ret) == 0 {
		panic("no return value specified for CreateCostExplorerQueryProcess1")
	}

	var r0 admin.CreateCostExplorerQueryProcess1ApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.CreateCostExplorerQueryProcess1ApiRequest); ok {
		r0 = rf(ctx, orgId, token)
	} else {
		r0 = ret.Get(0).(admin.CreateCostExplorerQueryProcess1ApiRequest)
	}

	return r0
}

// InvoicesApi_CreateCostExplorerQueryProcess1_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCostExplorerQueryProcess1'
type InvoicesApi_CreateCostExplorerQueryProcess1_Call struct {
	*mock.Call
}

// CreateCostExplorerQueryProcess1 is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
//   - token string
func (_e *InvoicesApi_Expecter) CreateCostExplorerQueryProcess1(ctx interface{}, orgId interface{}, token interface{}) *InvoicesApi_CreateCostExplorerQueryProcess1_Call {
	return &InvoicesApi_CreateCostExplorerQueryProcess1_Call{Call: _e.mock.On("CreateCostExplorerQueryProcess1", ctx, orgId, token)}
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1_Call) Run(run func(ctx context.Context, orgId string, token string)) *InvoicesApi_CreateCostExplorerQueryProcess1_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1_Call) Return(_a0 admin.CreateCostExplorerQueryProcess1ApiRequest) *InvoicesApi_CreateCostExplorerQueryProcess1_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1_Call) RunAndReturn(run func(context.Context, string, string) admin.CreateCostExplorerQueryProcess1ApiRequest) *InvoicesApi_CreateCostExplorerQueryProcess1_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCostExplorerQueryProcess1Execute provides a mock function with given fields: r
func (_m *InvoicesApi) CreateCostExplorerQueryProcess1Execute(r admin.CreateCostExplorerQueryProcess1ApiRequest) (string, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateCostExplorerQueryProcess1Execute")
	}

	var r0 string
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateCostExplorerQueryProcess1ApiRequest) (string, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateCostExplorerQueryProcess1ApiRequest) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(admin.CreateCostExplorerQueryProcess1ApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateCostExplorerQueryProcess1ApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCostExplorerQueryProcess1Execute'
type InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call struct {
	*mock.Call
}

// CreateCostExplorerQueryProcess1Execute is a helper method to define mock.On call
//   - r admin.CreateCostExplorerQueryProcess1ApiRequest
func (_e *InvoicesApi_Expecter) CreateCostExplorerQueryProcess1Execute(r interface{}) *InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call {
	return &InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call{Call: _e.mock.On("CreateCostExplorerQueryProcess1Execute", r)}
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call) Run(run func(r admin.CreateCostExplorerQueryProcess1ApiRequest)) *InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateCostExplorerQueryProcess1ApiRequest))
	})
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call) Return(_a0 string, _a1 *http.Response, _a2 error) *InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call) RunAndReturn(run func(admin.CreateCostExplorerQueryProcess1ApiRequest) (string, *http.Response, error)) *InvoicesApi_CreateCostExplorerQueryProcess1Execute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCostExplorerQueryProcess1WithParams provides a mock function with given fields: ctx, args
func (_m *InvoicesApi) CreateCostExplorerQueryProcess1WithParams(ctx context.Context, args *admin.CreateCostExplorerQueryProcess1ApiParams) admin.CreateCostExplorerQueryProcess1ApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateCostExplorerQueryProcess1WithParams")
	}

	var r0 admin.CreateCostExplorerQueryProcess1ApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateCostExplorerQueryProcess1ApiParams) admin.CreateCostExplorerQueryProcess1ApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateCostExplorerQueryProcess1ApiRequest)
	}

	return r0
}

// InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCostExplorerQueryProcess1WithParams'
type InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call struct {
	*mock.Call
}

// CreateCostExplorerQueryProcess1WithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateCostExplorerQueryProcess1ApiParams
func (_e *InvoicesApi_Expecter) CreateCostExplorerQueryProcess1WithParams(ctx interface{}, args interface{}) *InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call {
	return &InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call{Call: _e.mock.On("CreateCostExplorerQueryProcess1WithParams", ctx, args)}
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call) Run(run func(ctx context.Context, args *admin.CreateCostExplorerQueryProcess1ApiParams)) *InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateCostExplorerQueryProcess1ApiParams))
	})
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call) Return(_a0 admin.CreateCostExplorerQueryProcess1ApiRequest) *InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateCostExplorerQueryProcess1ApiParams) admin.CreateCostExplorerQueryProcess1ApiRequest) *InvoicesApi_CreateCostExplorerQueryProcess1WithParams_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCostExplorerQueryProcessExecute provides a mock function with given fields: r
func (_m *InvoicesApi) CreateCostExplorerQueryProcessExecute(r admin.CreateCostExplorerQueryProcessApiRequest) (*admin.CostExplorerFilterResponse, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateCostExplorerQueryProcessExecute")
	}

	var r0 *admin.CostExplorerFilterResponse
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateCostExplorerQueryProcessApiRequest) (*admin.CostExplorerFilterResponse, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateCostExplorerQueryProcessApiRequest) *admin.CostExplorerFilterResponse); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CostExplorerFilterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateCostExplorerQueryProcessApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateCostExplorerQueryProcessApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvoicesApi_CreateCostExplorerQueryProcessExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCostExplorerQueryProcessExecute'
type InvoicesApi_CreateCostExplorerQueryProcessExecute_Call struct {
	*mock.Call
}

// CreateCostExplorerQueryProcessExecute is a helper method to define mock.On call
//   - r admin.CreateCostExplorerQueryProcessApiRequest
func (_e *InvoicesApi_Expecter) CreateCostExplorerQueryProcessExecute(r interface{}) *InvoicesApi_CreateCostExplorerQueryProcessExecute_Call {
	return &InvoicesApi_CreateCostExplorerQueryProcessExecute_Call{Call: _e.mock.On("CreateCostExplorerQueryProcessExecute", r)}
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcessExecute_Call) Run(run func(r admin.CreateCostExplorerQueryProcessApiRequest)) *InvoicesApi_CreateCostExplorerQueryProcessExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateCostExplorerQueryProcessApiRequest))
	})
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcessExecute_Call) Return(_a0 *admin.CostExplorerFilterResponse, _a1 *http.Response, _a2 error) *InvoicesApi_CreateCostExplorerQueryProcessExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcessExecute_Call) RunAndReturn(run func(admin.CreateCostExplorerQueryProcessApiRequest) (*admin.CostExplorerFilterResponse, *http.Response, error)) *InvoicesApi_CreateCostExplorerQueryProcessExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCostExplorerQueryProcessWithParams provides a mock function with given fields: ctx, args
func (_m *InvoicesApi) CreateCostExplorerQueryProcessWithParams(ctx context.Context, args *admin.CreateCostExplorerQueryProcessApiParams) admin.CreateCostExplorerQueryProcessApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateCostExplorerQueryProcessWithParams")
	}

	var r0 admin.CreateCostExplorerQueryProcessApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateCostExplorerQueryProcessApiParams) admin.CreateCostExplorerQueryProcessApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateCostExplorerQueryProcessApiRequest)
	}

	return r0
}

// InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCostExplorerQueryProcessWithParams'
type InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call struct {
	*mock.Call
}

// CreateCostExplorerQueryProcessWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateCostExplorerQueryProcessApiParams
func (_e *InvoicesApi_Expecter) CreateCostExplorerQueryProcessWithParams(ctx interface{}, args interface{}) *InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call {
	return &InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call{Call: _e.mock.On("CreateCostExplorerQueryProcessWithParams", ctx, args)}
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateCostExplorerQueryProcessApiParams)) *InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateCostExplorerQueryProcessApiParams))
	})
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call) Return(_a0 admin.CreateCostExplorerQueryProcessApiRequest) *InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateCostExplorerQueryProcessApiParams) admin.CreateCostExplorerQueryProcessApiRequest) *InvoicesApi_CreateCostExplorerQueryProcessWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadInvoiceCSV provides a mock function with given fields: ctx, orgId, invoiceId
func (_m *InvoicesApi) DownloadInvoiceCSV(ctx context.Context, orgId string, invoiceId string) admin.DownloadInvoiceCSVApiRequest {
	ret := _m.Called(ctx, orgId, invoiceId)

	if len(ret) == 0 {
		panic("no return value specified for DownloadInvoiceCSV")
	}

	var r0 admin.DownloadInvoiceCSVApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.DownloadInvoiceCSVApiRequest); ok {
		r0 = rf(ctx, orgId, invoiceId)
	} else {
		r0 = ret.Get(0).(admin.DownloadInvoiceCSVApiRequest)
	}

	return r0
}

// InvoicesApi_DownloadInvoiceCSV_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadInvoiceCSV'
type InvoicesApi_DownloadInvoiceCSV_Call struct {
	*mock.Call
}

// DownloadInvoiceCSV is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
//   - invoiceId string
func (_e *InvoicesApi_Expecter) DownloadInvoiceCSV(ctx interface{}, orgId interface{}, invoiceId interface{}) *InvoicesApi_DownloadInvoiceCSV_Call {
	return &InvoicesApi_DownloadInvoiceCSV_Call{Call: _e.mock.On("DownloadInvoiceCSV", ctx, orgId, invoiceId)}
}

func (_c *InvoicesApi_DownloadInvoiceCSV_Call) Run(run func(ctx context.Context, orgId string, invoiceId string)) *InvoicesApi_DownloadInvoiceCSV_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *InvoicesApi_DownloadInvoiceCSV_Call) Return(_a0 admin.DownloadInvoiceCSVApiRequest) *InvoicesApi_DownloadInvoiceCSV_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_DownloadInvoiceCSV_Call) RunAndReturn(run func(context.Context, string, string) admin.DownloadInvoiceCSVApiRequest) *InvoicesApi_DownloadInvoiceCSV_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadInvoiceCSVExecute provides a mock function with given fields: r
func (_m *InvoicesApi) DownloadInvoiceCSVExecute(r admin.DownloadInvoiceCSVApiRequest) (string, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DownloadInvoiceCSVExecute")
	}

	var r0 string
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DownloadInvoiceCSVApiRequest) (string, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DownloadInvoiceCSVApiRequest) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(admin.DownloadInvoiceCSVApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DownloadInvoiceCSVApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvoicesApi_DownloadInvoiceCSVExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadInvoiceCSVExecute'
type InvoicesApi_DownloadInvoiceCSVExecute_Call struct {
	*mock.Call
}

// DownloadInvoiceCSVExecute is a helper method to define mock.On call
//   - r admin.DownloadInvoiceCSVApiRequest
func (_e *InvoicesApi_Expecter) DownloadInvoiceCSVExecute(r interface{}) *InvoicesApi_DownloadInvoiceCSVExecute_Call {
	return &InvoicesApi_DownloadInvoiceCSVExecute_Call{Call: _e.mock.On("DownloadInvoiceCSVExecute", r)}
}

func (_c *InvoicesApi_DownloadInvoiceCSVExecute_Call) Run(run func(r admin.DownloadInvoiceCSVApiRequest)) *InvoicesApi_DownloadInvoiceCSVExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DownloadInvoiceCSVApiRequest))
	})
	return _c
}

func (_c *InvoicesApi_DownloadInvoiceCSVExecute_Call) Return(_a0 string, _a1 *http.Response, _a2 error) *InvoicesApi_DownloadInvoiceCSVExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *InvoicesApi_DownloadInvoiceCSVExecute_Call) RunAndReturn(run func(admin.DownloadInvoiceCSVApiRequest) (string, *http.Response, error)) *InvoicesApi_DownloadInvoiceCSVExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadInvoiceCSVWithParams provides a mock function with given fields: ctx, args
func (_m *InvoicesApi) DownloadInvoiceCSVWithParams(ctx context.Context, args *admin.DownloadInvoiceCSVApiParams) admin.DownloadInvoiceCSVApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DownloadInvoiceCSVWithParams")
	}

	var r0 admin.DownloadInvoiceCSVApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DownloadInvoiceCSVApiParams) admin.DownloadInvoiceCSVApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DownloadInvoiceCSVApiRequest)
	}

	return r0
}

// InvoicesApi_DownloadInvoiceCSVWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadInvoiceCSVWithParams'
type InvoicesApi_DownloadInvoiceCSVWithParams_Call struct {
	*mock.Call
}

// DownloadInvoiceCSVWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DownloadInvoiceCSVApiParams
func (_e *InvoicesApi_Expecter) DownloadInvoiceCSVWithParams(ctx interface{}, args interface{}) *InvoicesApi_DownloadInvoiceCSVWithParams_Call {
	return &InvoicesApi_DownloadInvoiceCSVWithParams_Call{Call: _e.mock.On("DownloadInvoiceCSVWithParams", ctx, args)}
}

func (_c *InvoicesApi_DownloadInvoiceCSVWithParams_Call) Run(run func(ctx context.Context, args *admin.DownloadInvoiceCSVApiParams)) *InvoicesApi_DownloadInvoiceCSVWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DownloadInvoiceCSVApiParams))
	})
	return _c
}

func (_c *InvoicesApi_DownloadInvoiceCSVWithParams_Call) Return(_a0 admin.DownloadInvoiceCSVApiRequest) *InvoicesApi_DownloadInvoiceCSVWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_DownloadInvoiceCSVWithParams_Call) RunAndReturn(run func(context.Context, *admin.DownloadInvoiceCSVApiParams) admin.DownloadInvoiceCSVApiRequest) *InvoicesApi_DownloadInvoiceCSVWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetInvoice provides a mock function with given fields: ctx, orgId, invoiceId
func (_m *InvoicesApi) GetInvoice(ctx context.Context, orgId string, invoiceId string) admin.GetInvoiceApiRequest {
	ret := _m.Called(ctx, orgId, invoiceId)

	if len(ret) == 0 {
		panic("no return value specified for GetInvoice")
	}

	var r0 admin.GetInvoiceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetInvoiceApiRequest); ok {
		r0 = rf(ctx, orgId, invoiceId)
	} else {
		r0 = ret.Get(0).(admin.GetInvoiceApiRequest)
	}

	return r0
}

// InvoicesApi_GetInvoice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInvoice'
type InvoicesApi_GetInvoice_Call struct {
	*mock.Call
}

// GetInvoice is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
//   - invoiceId string
func (_e *InvoicesApi_Expecter) GetInvoice(ctx interface{}, orgId interface{}, invoiceId interface{}) *InvoicesApi_GetInvoice_Call {
	return &InvoicesApi_GetInvoice_Call{Call: _e.mock.On("GetInvoice", ctx, orgId, invoiceId)}
}

func (_c *InvoicesApi_GetInvoice_Call) Run(run func(ctx context.Context, orgId string, invoiceId string)) *InvoicesApi_GetInvoice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *InvoicesApi_GetInvoice_Call) Return(_a0 admin.GetInvoiceApiRequest) *InvoicesApi_GetInvoice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_GetInvoice_Call) RunAndReturn(run func(context.Context, string, string) admin.GetInvoiceApiRequest) *InvoicesApi_GetInvoice_Call {
	_c.Call.Return(run)
	return _c
}

// GetInvoiceExecute provides a mock function with given fields: r
func (_m *InvoicesApi) GetInvoiceExecute(r admin.GetInvoiceApiRequest) (string, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetInvoiceExecute")
	}

	var r0 string
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetInvoiceApiRequest) (string, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetInvoiceApiRequest) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(admin.GetInvoiceApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetInvoiceApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvoicesApi_GetInvoiceExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInvoiceExecute'
type InvoicesApi_GetInvoiceExecute_Call struct {
	*mock.Call
}

// GetInvoiceExecute is a helper method to define mock.On call
//   - r admin.GetInvoiceApiRequest
func (_e *InvoicesApi_Expecter) GetInvoiceExecute(r interface{}) *InvoicesApi_GetInvoiceExecute_Call {
	return &InvoicesApi_GetInvoiceExecute_Call{Call: _e.mock.On("GetInvoiceExecute", r)}
}

func (_c *InvoicesApi_GetInvoiceExecute_Call) Run(run func(r admin.GetInvoiceApiRequest)) *InvoicesApi_GetInvoiceExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetInvoiceApiRequest))
	})
	return _c
}

func (_c *InvoicesApi_GetInvoiceExecute_Call) Return(_a0 string, _a1 *http.Response, _a2 error) *InvoicesApi_GetInvoiceExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *InvoicesApi_GetInvoiceExecute_Call) RunAndReturn(run func(admin.GetInvoiceApiRequest) (string, *http.Response, error)) *InvoicesApi_GetInvoiceExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetInvoiceWithParams provides a mock function with given fields: ctx, args
func (_m *InvoicesApi) GetInvoiceWithParams(ctx context.Context, args *admin.GetInvoiceApiParams) admin.GetInvoiceApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetInvoiceWithParams")
	}

	var r0 admin.GetInvoiceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetInvoiceApiParams) admin.GetInvoiceApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetInvoiceApiRequest)
	}

	return r0
}

// InvoicesApi_GetInvoiceWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInvoiceWithParams'
type InvoicesApi_GetInvoiceWithParams_Call struct {
	*mock.Call
}

// GetInvoiceWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetInvoiceApiParams
func (_e *InvoicesApi_Expecter) GetInvoiceWithParams(ctx interface{}, args interface{}) *InvoicesApi_GetInvoiceWithParams_Call {
	return &InvoicesApi_GetInvoiceWithParams_Call{Call: _e.mock.On("GetInvoiceWithParams", ctx, args)}
}

func (_c *InvoicesApi_GetInvoiceWithParams_Call) Run(run func(ctx context.Context, args *admin.GetInvoiceApiParams)) *InvoicesApi_GetInvoiceWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetInvoiceApiParams))
	})
	return _c
}

func (_c *InvoicesApi_GetInvoiceWithParams_Call) Return(_a0 admin.GetInvoiceApiRequest) *InvoicesApi_GetInvoiceWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_GetInvoiceWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetInvoiceApiParams) admin.GetInvoiceApiRequest) *InvoicesApi_GetInvoiceWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListInvoices provides a mock function with given fields: ctx, orgId
func (_m *InvoicesApi) ListInvoices(ctx context.Context, orgId string) admin.ListInvoicesApiRequest {
	ret := _m.Called(ctx, orgId)

	if len(ret) == 0 {
		panic("no return value specified for ListInvoices")
	}

	var r0 admin.ListInvoicesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListInvoicesApiRequest); ok {
		r0 = rf(ctx, orgId)
	} else {
		r0 = ret.Get(0).(admin.ListInvoicesApiRequest)
	}

	return r0
}

// InvoicesApi_ListInvoices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListInvoices'
type InvoicesApi_ListInvoices_Call struct {
	*mock.Call
}

// ListInvoices is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
func (_e *InvoicesApi_Expecter) ListInvoices(ctx interface{}, orgId interface{}) *InvoicesApi_ListInvoices_Call {
	return &InvoicesApi_ListInvoices_Call{Call: _e.mock.On("ListInvoices", ctx, orgId)}
}

func (_c *InvoicesApi_ListInvoices_Call) Run(run func(ctx context.Context, orgId string)) *InvoicesApi_ListInvoices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *InvoicesApi_ListInvoices_Call) Return(_a0 admin.ListInvoicesApiRequest) *InvoicesApi_ListInvoices_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_ListInvoices_Call) RunAndReturn(run func(context.Context, string) admin.ListInvoicesApiRequest) *InvoicesApi_ListInvoices_Call {
	_c.Call.Return(run)
	return _c
}

// ListInvoicesExecute provides a mock function with given fields: r
func (_m *InvoicesApi) ListInvoicesExecute(r admin.ListInvoicesApiRequest) (*admin.PaginatedApiInvoiceMetadata, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListInvoicesExecute")
	}

	var r0 *admin.PaginatedApiInvoiceMetadata
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListInvoicesApiRequest) (*admin.PaginatedApiInvoiceMetadata, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListInvoicesApiRequest) *admin.PaginatedApiInvoiceMetadata); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedApiInvoiceMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListInvoicesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListInvoicesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvoicesApi_ListInvoicesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListInvoicesExecute'
type InvoicesApi_ListInvoicesExecute_Call struct {
	*mock.Call
}

// ListInvoicesExecute is a helper method to define mock.On call
//   - r admin.ListInvoicesApiRequest
func (_e *InvoicesApi_Expecter) ListInvoicesExecute(r interface{}) *InvoicesApi_ListInvoicesExecute_Call {
	return &InvoicesApi_ListInvoicesExecute_Call{Call: _e.mock.On("ListInvoicesExecute", r)}
}

func (_c *InvoicesApi_ListInvoicesExecute_Call) Run(run func(r admin.ListInvoicesApiRequest)) *InvoicesApi_ListInvoicesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListInvoicesApiRequest))
	})
	return _c
}

func (_c *InvoicesApi_ListInvoicesExecute_Call) Return(_a0 *admin.PaginatedApiInvoiceMetadata, _a1 *http.Response, _a2 error) *InvoicesApi_ListInvoicesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *InvoicesApi_ListInvoicesExecute_Call) RunAndReturn(run func(admin.ListInvoicesApiRequest) (*admin.PaginatedApiInvoiceMetadata, *http.Response, error)) *InvoicesApi_ListInvoicesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListInvoicesWithParams provides a mock function with given fields: ctx, args
func (_m *InvoicesApi) ListInvoicesWithParams(ctx context.Context, args *admin.ListInvoicesApiParams) admin.ListInvoicesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListInvoicesWithParams")
	}

	var r0 admin.ListInvoicesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListInvoicesApiParams) admin.ListInvoicesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListInvoicesApiRequest)
	}

	return r0
}

// InvoicesApi_ListInvoicesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListInvoicesWithParams'
type InvoicesApi_ListInvoicesWithParams_Call struct {
	*mock.Call
}

// ListInvoicesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListInvoicesApiParams
func (_e *InvoicesApi_Expecter) ListInvoicesWithParams(ctx interface{}, args interface{}) *InvoicesApi_ListInvoicesWithParams_Call {
	return &InvoicesApi_ListInvoicesWithParams_Call{Call: _e.mock.On("ListInvoicesWithParams", ctx, args)}
}

func (_c *InvoicesApi_ListInvoicesWithParams_Call) Run(run func(ctx context.Context, args *admin.ListInvoicesApiParams)) *InvoicesApi_ListInvoicesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListInvoicesApiParams))
	})
	return _c
}

func (_c *InvoicesApi_ListInvoicesWithParams_Call) Return(_a0 admin.ListInvoicesApiRequest) *InvoicesApi_ListInvoicesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_ListInvoicesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListInvoicesApiParams) admin.ListInvoicesApiRequest) *InvoicesApi_ListInvoicesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListPendingInvoices provides a mock function with given fields: ctx, orgId
func (_m *InvoicesApi) ListPendingInvoices(ctx context.Context, orgId string) admin.ListPendingInvoicesApiRequest {
	ret := _m.Called(ctx, orgId)

	if len(ret) == 0 {
		panic("no return value specified for ListPendingInvoices")
	}

	var r0 admin.ListPendingInvoicesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListPendingInvoicesApiRequest); ok {
		r0 = rf(ctx, orgId)
	} else {
		r0 = ret.Get(0).(admin.ListPendingInvoicesApiRequest)
	}

	return r0
}

// InvoicesApi_ListPendingInvoices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPendingInvoices'
type InvoicesApi_ListPendingInvoices_Call struct {
	*mock.Call
}

// ListPendingInvoices is a helper method to define mock.On call
//   - ctx context.Context
//   - orgId string
func (_e *InvoicesApi_Expecter) ListPendingInvoices(ctx interface{}, orgId interface{}) *InvoicesApi_ListPendingInvoices_Call {
	return &InvoicesApi_ListPendingInvoices_Call{Call: _e.mock.On("ListPendingInvoices", ctx, orgId)}
}

func (_c *InvoicesApi_ListPendingInvoices_Call) Run(run func(ctx context.Context, orgId string)) *InvoicesApi_ListPendingInvoices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *InvoicesApi_ListPendingInvoices_Call) Return(_a0 admin.ListPendingInvoicesApiRequest) *InvoicesApi_ListPendingInvoices_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_ListPendingInvoices_Call) RunAndReturn(run func(context.Context, string) admin.ListPendingInvoicesApiRequest) *InvoicesApi_ListPendingInvoices_Call {
	_c.Call.Return(run)
	return _c
}

// ListPendingInvoicesExecute provides a mock function with given fields: r
func (_m *InvoicesApi) ListPendingInvoicesExecute(r admin.ListPendingInvoicesApiRequest) (*admin.PaginatedApiInvoice, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListPendingInvoicesExecute")
	}

	var r0 *admin.PaginatedApiInvoice
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListPendingInvoicesApiRequest) (*admin.PaginatedApiInvoice, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListPendingInvoicesApiRequest) *admin.PaginatedApiInvoice); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedApiInvoice)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListPendingInvoicesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListPendingInvoicesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvoicesApi_ListPendingInvoicesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPendingInvoicesExecute'
type InvoicesApi_ListPendingInvoicesExecute_Call struct {
	*mock.Call
}

// ListPendingInvoicesExecute is a helper method to define mock.On call
//   - r admin.ListPendingInvoicesApiRequest
func (_e *InvoicesApi_Expecter) ListPendingInvoicesExecute(r interface{}) *InvoicesApi_ListPendingInvoicesExecute_Call {
	return &InvoicesApi_ListPendingInvoicesExecute_Call{Call: _e.mock.On("ListPendingInvoicesExecute", r)}
}

func (_c *InvoicesApi_ListPendingInvoicesExecute_Call) Run(run func(r admin.ListPendingInvoicesApiRequest)) *InvoicesApi_ListPendingInvoicesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListPendingInvoicesApiRequest))
	})
	return _c
}

func (_c *InvoicesApi_ListPendingInvoicesExecute_Call) Return(_a0 *admin.PaginatedApiInvoice, _a1 *http.Response, _a2 error) *InvoicesApi_ListPendingInvoicesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *InvoicesApi_ListPendingInvoicesExecute_Call) RunAndReturn(run func(admin.ListPendingInvoicesApiRequest) (*admin.PaginatedApiInvoice, *http.Response, error)) *InvoicesApi_ListPendingInvoicesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListPendingInvoicesWithParams provides a mock function with given fields: ctx, args
func (_m *InvoicesApi) ListPendingInvoicesWithParams(ctx context.Context, args *admin.ListPendingInvoicesApiParams) admin.ListPendingInvoicesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListPendingInvoicesWithParams")
	}

	var r0 admin.ListPendingInvoicesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListPendingInvoicesApiParams) admin.ListPendingInvoicesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListPendingInvoicesApiRequest)
	}

	return r0
}

// InvoicesApi_ListPendingInvoicesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPendingInvoicesWithParams'
type InvoicesApi_ListPendingInvoicesWithParams_Call struct {
	*mock.Call
}

// ListPendingInvoicesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListPendingInvoicesApiParams
func (_e *InvoicesApi_Expecter) ListPendingInvoicesWithParams(ctx interface{}, args interface{}) *InvoicesApi_ListPendingInvoicesWithParams_Call {
	return &InvoicesApi_ListPendingInvoicesWithParams_Call{Call: _e.mock.On("ListPendingInvoicesWithParams", ctx, args)}
}

func (_c *InvoicesApi_ListPendingInvoicesWithParams_Call) Run(run func(ctx context.Context, args *admin.ListPendingInvoicesApiParams)) *InvoicesApi_ListPendingInvoicesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListPendingInvoicesApiParams))
	})
	return _c
}

func (_c *InvoicesApi_ListPendingInvoicesWithParams_Call) Return(_a0 admin.ListPendingInvoicesApiRequest) *InvoicesApi_ListPendingInvoicesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InvoicesApi_ListPendingInvoicesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListPendingInvoicesApiParams) admin.ListPendingInvoicesApiRequest) *InvoicesApi_ListPendingInvoicesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewInvoicesApi creates a new instance of InvoicesApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInvoicesApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *InvoicesApi {
	mock := &InvoicesApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
