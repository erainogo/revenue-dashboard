package insight

import (
	"context"
	"errors"
	"time"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
	"github.com/stretchr/testify/mock"
)

func (suite *InsightServiceTestSuite) TestGetCountryLevelRevenue_Success() {
	expected := []*entities.CountryLevelRevenue{
		{
			Country:          "USA",
			ProductName:      "Product A",
			TotalRevenue:     1000.0,
			TransactionCount: 10,
			UpdatedAt:        time.Now(),
		},
	}

	suite.productSummeryRepository.
		On("GetCountryLevelRevenueSortedByTotal", mock.Anything, 1, 10).
		Return(expected, nil)

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetCountryLevelRevenue(context.TODO(), 1, 10)
	suite.NoError(err)
	suite.Equal(expected, result)

	suite.productSummeryRepository.AssertExpectations(suite.T())
}

func (suite *InsightServiceTestSuite) TestGetCountryLevelRevenue_Error() {
	suite.productSummeryRepository.
		On("GetCountryLevelRevenueSortedByTotal", mock.Anything, 1, 10).
		Return(nil, errors.New("db error"))

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetCountryLevelRevenue(context.TODO(), 1, 10)
	suite.Nil(result)
	suite.Error(err)
	suite.Contains(err.Error(), "db error")

	suite.productSummeryRepository.AssertExpectations(suite.T())
}

func (suite *InsightServiceTestSuite) TestGetFrequentlyPurchasedProducts_Success() {
	expected := []*entities.ProductPurchaseSummary{
		{
			ProductID:     "P123",
			ProductName:   "Product X",
			PurchaseCount: 25,
			StockQuantity: 200,
			UpdatedAt:     time.Now(),
		},
	}

	suite.purchaseSummeryRepository.
		On("GetFrequentlyPurchasedProducts", mock.Anything).
		Return(expected, nil)

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetFrequentlyPurchasedProducts(context.TODO())
	suite.NoError(err)
	suite.Equal(expected, result)

	suite.purchaseSummeryRepository.AssertExpectations(suite.T())
}

func (suite *InsightServiceTestSuite) TestGetFrequentlyPurchasedProducts_Error() {
	suite.purchaseSummeryRepository.
		On("GetFrequentlyPurchasedProducts", mock.Anything).
		Return(nil, errors.New("fetch error"))

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetFrequentlyPurchasedProducts(context.TODO())
	suite.Nil(result)
	suite.Error(err)
	suite.Contains(err.Error(), "fetch error")

	suite.purchaseSummeryRepository.AssertExpectations(suite.T())
}

func (suite *InsightServiceTestSuite) TestGetRegionRevenue_Success() {
	expected := []*entities.RegionRevenue{
		{
			Region:        "California",
			TotalRevenue:  100000.0,
			TotalQuantity: 5000,
			UpdatedAt:     time.Now(),
		},
	}

	suite.regionRevenueSummeryRepository.
		On("GetRegionRevenue", mock.Anything).
		Return(expected, nil)

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetRegionRevenue(context.TODO())
	suite.NoError(err)
	suite.Equal(expected, result)

	suite.regionRevenueSummeryRepository.AssertExpectations(suite.T())
}

func (suite *InsightServiceTestSuite) TestGetRegionRevenue_Error() {
	suite.regionRevenueSummeryRepository.
		On("GetRegionRevenue", mock.Anything).
		Return(nil, errors.New("region fetch error"))

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetRegionRevenue(context.TODO())
	suite.Nil(result)
	suite.Error(err)
	suite.Contains(err.Error(), "region fetch error")

	suite.regionRevenueSummeryRepository.AssertExpectations(suite.T())
}

func (suite *InsightServiceTestSuite) TestGetMonthlyRevenue_Success() {
	expected := []*entities.MonthlySales{
		{
			Year:          2024,
			Month:         "January",
			TotalQuantity: 1000,
			UpdatedAt:     time.Now(),
		},
	}

	suite.monthlySalesSummeryRepository.
		On("GetMonthlyRevenue", mock.Anything).
		Return(expected, nil)

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetMonthlyRevenue(context.TODO())
	suite.NoError(err)
	suite.Equal(expected, result)

	suite.monthlySalesSummeryRepository.AssertExpectations(suite.T())
}

func (suite *InsightServiceTestSuite) TestGetMonthlyRevenue_Error() {
	suite.monthlySalesSummeryRepository.
		On("GetMonthlyRevenue", mock.Anything).
		Return(nil, errors.New("monthly error"))

	service := NewInsightService(
		context.TODO(),
		suite.productSummeryRepository,
		suite.purchaseSummeryRepository,
		suite.monthlySalesSummeryRepository,
		suite.regionRevenueSummeryRepository,
		WithLogger(suite.logger),
	)

	result, err := service.GetMonthlyRevenue(context.TODO())
	suite.Nil(result)
	suite.Error(err)
	suite.Contains(err.Error(), "monthly error")

	suite.monthlySalesSummeryRepository.AssertExpectations(suite.T())
}
