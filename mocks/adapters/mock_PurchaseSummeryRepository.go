// Code generated by mockery v2.53.3. DO NOT EDIT.

package adapters

import (
	context "context"

	entities "github.com/erainogo/revenue-dashboard/pkg/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockPurchaseSummeryRepository is an autogenerated mock type for the PurchaseSummeryRepository type
type MockPurchaseSummeryRepository struct {
	mock.Mock
}

type MockPurchaseSummeryRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPurchaseSummeryRepository) EXPECT() *MockPurchaseSummeryRepository_Expecter {
	return &MockPurchaseSummeryRepository_Expecter{mock: &_m.Mock}
}

// BulkInsert provides a mock function with given fields: ctx, docs
func (_m *MockPurchaseSummeryRepository) BulkInsert(ctx context.Context, docs map[string]*entities.ProductPurchaseSummary) error {
	ret := _m.Called(ctx, docs)

	if len(ret) == 0 {
		panic("no return value specified for BulkInsert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]*entities.ProductPurchaseSummary) error); ok {
		r0 = rf(ctx, docs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPurchaseSummeryRepository_BulkInsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkInsert'
type MockPurchaseSummeryRepository_BulkInsert_Call struct {
	*mock.Call
}

// BulkInsert is a helper method to define mock.On call
//   - ctx context.Context
//   - docs map[string]*entities.ProductPurchaseSummary
func (_e *MockPurchaseSummeryRepository_Expecter) BulkInsert(ctx interface{}, docs interface{}) *MockPurchaseSummeryRepository_BulkInsert_Call {
	return &MockPurchaseSummeryRepository_BulkInsert_Call{Call: _e.mock.On("BulkInsert", ctx, docs)}
}

func (_c *MockPurchaseSummeryRepository_BulkInsert_Call) Run(run func(ctx context.Context, docs map[string]*entities.ProductPurchaseSummary)) *MockPurchaseSummeryRepository_BulkInsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]*entities.ProductPurchaseSummary))
	})
	return _c
}

func (_c *MockPurchaseSummeryRepository_BulkInsert_Call) Return(_a0 error) *MockPurchaseSummeryRepository_BulkInsert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPurchaseSummeryRepository_BulkInsert_Call) RunAndReturn(run func(context.Context, map[string]*entities.ProductPurchaseSummary) error) *MockPurchaseSummeryRepository_BulkInsert_Call {
	_c.Call.Return(run)
	return _c
}

// GetFrequentlyPurchasedProducts provides a mock function with given fields: ctx
func (_m *MockPurchaseSummeryRepository) GetFrequentlyPurchasedProducts(ctx context.Context) ([]*entities.ProductPurchaseSummary, error) {
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

// MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFrequentlyPurchasedProducts'
type MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call struct {
	*mock.Call
}

// GetFrequentlyPurchasedProducts is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPurchaseSummeryRepository_Expecter) GetFrequentlyPurchasedProducts(ctx interface{}) *MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call {
	return &MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call{Call: _e.mock.On("GetFrequentlyPurchasedProducts", ctx)}
}

func (_c *MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call) Run(run func(ctx context.Context)) *MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call) Return(_a0 []*entities.ProductPurchaseSummary, _a1 error) *MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call) RunAndReturn(run func(context.Context) ([]*entities.ProductPurchaseSummary, error)) *MockPurchaseSummeryRepository_GetFrequentlyPurchasedProducts_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPurchaseSummeryRepository creates a new instance of MockPurchaseSummeryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPurchaseSummeryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPurchaseSummeryRepository {
	mock := &MockPurchaseSummeryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
