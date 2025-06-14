package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type InsightService struct {
	ctx                      context.Context
	logger                   *zap.SugaredLogger
	productSummeryRepository adapters.ProductSummeryRepository
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
	opts ...InsightServiceOptions,
) adapters.InsightService {
	svc := &InsightService{
		ctx:                      ctx,
		productSummeryRepository: productSummeryRepository,
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
