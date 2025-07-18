// Code generated by mockery v2.53.3. DO NOT EDIT.

package adapters

import (
	context "context"

	entities "github.com/erainogo/revenue-dashboard/pkg/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockRegionRevenueSummeryRepository is an autogenerated mock type for the RegionRevenueSummeryRepository type
type MockRegionRevenueSummeryRepository struct {
	mock.Mock
}

type MockRegionRevenueSummeryRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRegionRevenueSummeryRepository) EXPECT() *MockRegionRevenueSummeryRepository_Expecter {
	return &MockRegionRevenueSummeryRepository_Expecter{mock: &_m.Mock}
}

// BulkInsert provides a mock function with given fields: ctx, docs
func (_m *MockRegionRevenueSummeryRepository) BulkInsert(ctx context.Context, docs map[string]*entities.RegionRevenue) error {
	ret := _m.Called(ctx, docs)

	if len(ret) == 0 {
		panic("no return value specified for BulkInsert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]*entities.RegionRevenue) error); ok {
		r0 = rf(ctx, docs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRegionRevenueSummeryRepository_BulkInsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkInsert'
type MockRegionRevenueSummeryRepository_BulkInsert_Call struct {
	*mock.Call
}

// BulkInsert is a helper method to define mock.On call
//   - ctx context.Context
//   - docs map[string]*entities.RegionRevenue
func (_e *MockRegionRevenueSummeryRepository_Expecter) BulkInsert(ctx interface{}, docs interface{}) *MockRegionRevenueSummeryRepository_BulkInsert_Call {
	return &MockRegionRevenueSummeryRepository_BulkInsert_Call{Call: _e.mock.On("BulkInsert", ctx, docs)}
}

func (_c *MockRegionRevenueSummeryRepository_BulkInsert_Call) Run(run func(ctx context.Context, docs map[string]*entities.RegionRevenue)) *MockRegionRevenueSummeryRepository_BulkInsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]*entities.RegionRevenue))
	})
	return _c
}

func (_c *MockRegionRevenueSummeryRepository_BulkInsert_Call) Return(_a0 error) *MockRegionRevenueSummeryRepository_BulkInsert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegionRevenueSummeryRepository_BulkInsert_Call) RunAndReturn(run func(context.Context, map[string]*entities.RegionRevenue) error) *MockRegionRevenueSummeryRepository_BulkInsert_Call {
	_c.Call.Return(run)
	return _c
}

// GetRegionRevenue provides a mock function with given fields: ctx
func (_m *MockRegionRevenueSummeryRepository) GetRegionRevenue(ctx context.Context) ([]*entities.RegionRevenue, error) {
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

// MockRegionRevenueSummeryRepository_GetRegionRevenue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRegionRevenue'
type MockRegionRevenueSummeryRepository_GetRegionRevenue_Call struct {
	*mock.Call
}

// GetRegionRevenue is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRegionRevenueSummeryRepository_Expecter) GetRegionRevenue(ctx interface{}) *MockRegionRevenueSummeryRepository_GetRegionRevenue_Call {
	return &MockRegionRevenueSummeryRepository_GetRegionRevenue_Call{Call: _e.mock.On("GetRegionRevenue", ctx)}
}

func (_c *MockRegionRevenueSummeryRepository_GetRegionRevenue_Call) Run(run func(ctx context.Context)) *MockRegionRevenueSummeryRepository_GetRegionRevenue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRegionRevenueSummeryRepository_GetRegionRevenue_Call) Return(_a0 []*entities.RegionRevenue, _a1 error) *MockRegionRevenueSummeryRepository_GetRegionRevenue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegionRevenueSummeryRepository_GetRegionRevenue_Call) RunAndReturn(run func(context.Context) ([]*entities.RegionRevenue, error)) *MockRegionRevenueSummeryRepository_GetRegionRevenue_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRegionRevenueSummeryRepository creates a new instance of MockRegionRevenueSummeryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRegionRevenueSummeryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRegionRevenueSummeryRepository {
	mock := &MockRegionRevenueSummeryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
