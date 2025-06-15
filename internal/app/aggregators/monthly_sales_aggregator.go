package aggregators

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
	"go.uber.org/zap"
)

type MonthlySalesAggregator struct {
	ctx                      context.Context
	logger                   *zap.SugaredLogger
	SummeryMap               map[string]*entities.MonthlySales // in memory cache to calculate purchase summery
	mutex                    sync.Mutex
}

type MonthlySalesAggregatorOptions func(*MonthlySalesAggregator)

func WithLoggerM(logger *zap.SugaredLogger) MonthlySalesAggregatorOptions {
	return func(s *MonthlySalesAggregator) {
		s.logger = logger
	}
}

func NewMonthlySalesAggregator(
	ctx context.Context,
	opts ...MonthlySalesAggregatorOptions,
) adapters.Aggregator {
	svc := &MonthlySalesAggregator{
		ctx:                      ctx,
		SummeryMap:               make(map[string]*entities.MonthlySales),
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *MonthlySalesAggregator) Aggregate(tx entities.Transaction) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	year := tx.TransactionDate.Year()
	month := tx.TransactionDate.Month()

	key := fmt.Sprintf("%d-%02d", year, int(month))

	if _, exists := s.SummeryMap[key]; !exists {
		s.SummeryMap[key] = &entities.MonthlySales{
			Year: year,
			Month: month.String(),
			TotalQuantity: 0,
		}
	}

	summary := s.SummeryMap[key]
	summary.TotalQuantity += tx.Quantity
	summary.UpdatedAt = time.Now()
}

func (s *MonthlySalesAggregator) GetOutput() any {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.SummeryMap
}
