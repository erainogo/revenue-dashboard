// Code generated by mockery v2.53.3. DO NOT EDIT.

package adapters

import (
	context "context"

	entities "github.com/erainogo/revenue-dashboard/pkg/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockInsightService is an autogenerated mock type for the InsightService type
type MockInsightService struct {
	mock.Mock
}

type MockInsightService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInsightService) EXPECT() *MockInsightService_Expecter {
	return &MockInsightService_Expecter{mock: &_m.Mock}
}

// GetCountryLevelRevenue provides a mock function with given fields: ctx, page, limit
func (_m *MockInsightService) GetCountryLevelRevenue(ctx context.Context, page int, limit int) ([]*entities.CountryLevelRevenue, error) {
	ret := _m.Called(ctx, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetCountryLevelRevenue")
	}

	var r0 []*entities.CountryLevelRevenue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]*entities.CountryLevelRevenue, error)); ok {
		return rf(ctx, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*entities.CountryLevelRevenue); ok {
		r0 = rf(ctx, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.CountryLevelRevenue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInsightService_GetCountryLevelRevenue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCountryLevelRevenue'
type MockInsightService_GetCountryLevelRevenue_Call struct {
	*mock.Call
}

// GetCountryLevelRevenue is a helper method to define mock.On call
//   - ctx context.Context
//   - page int
//   - limit int
func (_e *MockInsightService_Expecter) GetCountryLevelRevenue(ctx interface{}, page interface{}, limit interface{}) *MockInsightService_GetCountryLevelRevenue_Call {
	return &MockInsightService_GetCountryLevelRevenue_Call{Call: _e.mock.On("GetCountryLevelRevenue", ctx, page, limit)}
}

func (_c *MockInsightService_GetCountryLevelRevenue_Call) Run(run func(ctx context.Context, page int, limit int)) *MockInsightService_GetCountryLevelRevenue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockInsightService_GetCountryLevelRevenue_Call) Return(_a0 []*entities.CountryLevelRevenue, _a1 error) *MockInsightService_GetCountryLevelRevenue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInsightService_GetCountryLevelRevenue_Call) RunAndReturn(run func(context.Context, int, int) ([]*entities.CountryLevelRevenue, error)) *MockInsightService_GetCountryLevelRevenue_Call {
	_c.Call.Return(run)
	return _c
}

// GetFrequentlyPurchasedProducts provides a mock function with given fields: ctx
func (_m *MockInsightService) GetFrequentlyPurchasedProducts(ctx context.Context) ([]*entities.ProductPurchaseSummary, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetFrequentlyPurchasedProducts")
	}

	var r0 []*entities.ProductPurchaseSummary
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entities.ProductPurchaseSummary, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entities.ProductPurchaseSummary); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ProductPurchaseSummary)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInsightService_GetFrequentlyPurchasedProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFrequentlyPurchasedProducts'
type MockInsightService_GetFrequentlyPurchasedProducts_Call struct {
	*mock.Call
}

// GetFrequentlyPurchasedProducts is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInsightService_Expecter) GetFrequentlyPurchasedProducts(ctx interface{}) *MockInsightService_GetFrequentlyPurchasedProducts_Call {
	return &MockInsightService_GetFrequentlyPurchasedProducts_Call{Call: _e.mock.On("GetFrequentlyPurchasedProducts", ctx)}
}

func (_c *MockInsightService_GetFrequentlyPurchasedProducts_Call) Run(run func(ctx context.Context)) *MockInsightService_GetFrequentlyPurchasedProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInsightService_GetFrequentlyPurchasedProducts_Call) Return(_a0 []*entities.ProductPurchaseSummary, _a1 error) *MockInsightService_GetFrequentlyPurchasedProducts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInsightService_GetFrequentlyPurchasedProducts_Call) RunAndReturn(run func(context.Context) ([]*entities.ProductPurchaseSummary, error)) *MockInsightService_GetFrequentlyPurchasedProducts_Call {
	_c.Call.Return(run)
	return _c
}

// GetMonthlyRevenue provides a mock function with given fields: ctx
func (_m *MockInsightService) GetMonthlyRevenue(ctx context.Context) ([]*entities.MonthlySales, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetMonthlyRevenue")
	}

	var r0 []*entities.MonthlySales
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entities.MonthlySales, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entities.MonthlySales); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.MonthlySales)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInsightService_GetMonthlyRevenue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMonthlyRevenue'
type MockInsightService_GetMonthlyRevenue_Call struct {
	*mock.Call
}

// GetMonthlyRevenue is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInsightService_Expecter) GetMonthlyRevenue(ctx interface{}) *MockInsightService_GetMonthlyRevenue_Call {
	return &MockInsightService_GetMonthlyRevenue_Call{Call: _e.mock.On("GetMonthlyRevenue", ctx)}
}

func (_c *MockInsightService_GetMonthlyRevenue_Call) Run(run func(ctx context.Context)) *MockInsightService_GetMonthlyRevenue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInsightService_GetMonthlyRevenue_Call) Return(_a0 []*entities.MonthlySales, _a1 error) *MockInsightService_GetMonthlyRevenue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInsightService_GetMonthlyRevenue_Call) RunAndReturn(run func(context.Context) ([]*entities.MonthlySales, error)) *MockInsightService_GetMonthlyRevenue_Call {
	_c.Call.Return(run)
	return _c
}

// GetRegionRevenue provides a mock function with given fields: ctx
func (_m *MockInsightService) GetRegionRevenue(ctx context.Context) ([]*entities.RegionRevenue, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRegionRevenue")
	}

	var r0 []*entities.RegionRevenue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entities.RegionRevenue, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entities.RegionRevenue); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.RegionRevenue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInsightService_GetRegionRevenue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRegionRevenue'
type MockInsightService_GetRegionRevenue_Call struct {
	*mock.Call
}

// GetRegionRevenue is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInsightService_Expecter) GetRegionRevenue(ctx interface{}) *MockInsightService_GetRegionRevenue_Call {
	return &MockInsightService_GetRegionRevenue_Call{Call: _e.mock.On("GetRegionRevenue", ctx)}
}

func (_c *MockInsightService_GetRegionRevenue_Call) Run(run func(ctx context.Context)) *MockInsightService_GetRegionRevenue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInsightService_GetRegionRevenue_Call) Return(_a0 []*entities.RegionRevenue, _a1 error) *MockInsightService_GetRegionRevenue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInsightService_GetRegionRevenue_Call) RunAndReturn(run func(context.Context) ([]*entities.RegionRevenue, error)) *MockInsightService_GetRegionRevenue_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInsightService creates a new instance of MockInsightService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInsightService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInsightService {
	mock := &MockInsightService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
