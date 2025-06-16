package insight

import (
	"context"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type InsightService struct {
	ctx                            context.Context
	logger                         *zap.SugaredLogger
	productSummeryRepository       adapters.ProductSummeryRepository
	purchaseSummeryRepository      adapters.PurchaseSummeryRepository
	monthlySalesSummeryRepository  adapters.MonthlySalesSummeryRepository
	regionRevenueSummeryRepository adapters.RegionRevenueSummeryRepository
}

type InsightServiceOptions func(*InsightService)

func WithLogger(logger *zap.SugaredLogger) InsightServiceOptions {
	return func(s *InsightService) {
		s.logger = logger
	}
}

func NewInsightService(
	ctx context.Context,
	productSummeryRepository adapters.ProductSummeryRepository,
	purchaseSummeryRepository adapters.PurchaseSummeryRepository,
	monthlySalesSummeryRepository adapters.MonthlySalesSummeryRepository,
	regionRevenueSummeryRepository adapters.RegionRevenueSummeryRepository,
	opts ...InsightServiceOptions,
) adapters.InsightService {
	svc := &InsightService{
		ctx:                       ctx,
		productSummeryRepository:  productSummeryRepository,
		purchaseSummeryRepository: purchaseSummeryRepository,
		monthlySalesSummeryRepository: monthlySalesSummeryRepository,
		regionRevenueSummeryRepository:regionRevenueSummeryRepository,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (i *InsightService) GetCountryLevelRevenue(
	ctx context.Context,
	page int,
	limit int) ([]*entities.CountryLevelRevenue, error) {
	aggregator, err := i.productSummeryRepository.GetCountryLevelRevenueSortedByTotal(
		ctx, page, limit)

	if err != nil {
		return nil, err
	}

	return aggregator, nil
}

func (i *InsightService) GetFrequentlyPurchasedProducts(ctx context.Context,
) ([]*entities.ProductPurchaseSummary, error) {
	aggregator, err := i.purchaseSummeryRepository.GetFrequentlyPurchasedProducts(ctx)

	if err != nil {
		return nil, err
	}

	return aggregator, nil
}

func (i *InsightService) GetRegionRevenue(ctx context.Context) ([]*entities.RegionRevenue, error) {
	aggregator, err := i.regionRevenueSummeryRepository.GetRegionRevenue(ctx)

	if err != nil {
		return nil, err
	}

	return aggregator, nil
}

func (i *InsightService) GetMonthlyRevenue(ctx context.Context) ([]*entities.MonthlySales, error) {
	aggregator, err := i.monthlySalesSummeryRepository.GetMonthlyRevenue(ctx)

	if err != nil {
		return nil, err
	}

	return aggregator, nil
}
