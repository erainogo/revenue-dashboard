package aggregators

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type RegionRevenueAggregator struct {
	ctx                      context.Context
	logger                   *zap.SugaredLogger
	SummeryMap               map[string]*entities.RegionRevenue // in memory cache to calculate purchase summery
	mutex                    sync.Mutex
}

type RegionRevenueAggregatorOptions func(*RegionRevenueAggregator)

func WithLoggerR(logger *zap.SugaredLogger) RegionRevenueAggregatorOptions {
	return func(s *RegionRevenueAggregator) {
		s.logger = logger
	}
}

func NewRegionRevenueAggregator(
	ctx context.Context,
	opts ...RegionRevenueAggregatorOptions,
) adapters.Aggregator {
	svc := &RegionRevenueAggregator{
		ctx:                      ctx,
		SummeryMap:               make(map[string]*entities.RegionRevenue),
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *RegionRevenueAggregator) Aggregate(tx entities.Transaction) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := tx.Region

	if _, exists := s.SummeryMap[key]; !exists {
		s.SummeryMap[key] = &entities.RegionRevenue{
		    Region:        key,
			TotalQuantity: 0,
			TotalRevenue:  0,
			UpdatedAt:     time.Now(),
		}
	}

	summary := s.SummeryMap[key]
	summary.TotalQuantity += tx.Quantity
	summary.TotalRevenue += tx.TotalPrice
	summary.UpdatedAt = time.Now()
}

func (s *RegionRevenueAggregator) GetOutput() any {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.SummeryMap
}
