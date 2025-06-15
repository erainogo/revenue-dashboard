package aggregators

import (
	"context"
	"sync"
	"time"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
	"go.uber.org/zap"
)

type CountryRevenueAggregator struct {
	ctx                      context.Context
	logger                   *zap.SugaredLogger
	SummeryMap               map[entities.CountrySummaryKey]*entities.CountryLevelRevenue // in memory cache to calculate revenue by country summery
	mutex                    sync.Mutex
}

type CountryRevenueAggregatorOptions func(*CountryRevenueAggregator)

func WithLogger(logger *zap.SugaredLogger) CountryRevenueAggregatorOptions {
	return func(s *CountryRevenueAggregator) {
		s.logger = logger
	}
}

func NewCountryRevenueAggregator(
	ctx context.Context,
	opts ...CountryRevenueAggregatorOptions,
) adapters.Aggregator {
	svc := &CountryRevenueAggregator{
		ctx:                      ctx,
		SummeryMap:               make(map[entities.CountrySummaryKey]*entities.CountryLevelRevenue),
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *CountryRevenueAggregator) Aggregate(tx entities.Transaction) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := entities.CountrySummaryKey{Country: tx.Country, ProductName: tx.Product.Name}

	if _, exists := s.SummeryMap[key]; !exists {
		s.SummeryMap[key] = &entities.CountryLevelRevenue{
			Country:          tx.Country,
			ProductName:      tx.Product.Name,
			TotalRevenue:     0,
			TransactionCount: 0,
			UpdatedAt:        time.Now(),
		}
	}

	summary := s.SummeryMap[key]
	summary.TotalRevenue += tx.TotalPrice
	summary.TransactionCount++
}

func (s *CountryRevenueAggregator) GetOutput() any {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.SummeryMap
}
