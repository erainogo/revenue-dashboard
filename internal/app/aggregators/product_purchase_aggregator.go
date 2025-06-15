package aggregators

import (
	"context"
	"sync"
	"time"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
	"go.uber.org/zap"
)

type ProductPurchaseAggregator struct {
	ctx                      context.Context
	logger                   *zap.SugaredLogger
	SummeryMap               map[string]*entities.ProductPurchaseSummary // in memory cache to calculate purchase summery
	mutex                    sync.Mutex
}

type ProductPurchaseAggregatorOptions func(*ProductPurchaseAggregator)

func WithLoggerP(logger *zap.SugaredLogger) ProductPurchaseAggregatorOptions {
	return func(s *ProductPurchaseAggregator) {
		s.logger = logger
	}
}

func NewProductPurchaseAggregator(
	ctx context.Context,
	opts ...ProductPurchaseAggregatorOptions,
) adapters.Aggregator {
	svc := &ProductPurchaseAggregator{
		ctx:                      ctx,
		SummeryMap:               make(map[string]*entities.ProductPurchaseSummary),
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *ProductPurchaseAggregator) Aggregate(tx entities.Transaction) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := tx.Product.ID

	if _, exists := s.SummeryMap[key]; !exists {
		s.SummeryMap[key] = &entities.ProductPurchaseSummary{
			ProductID:     tx.Product.ID,
			ProductName:   tx.Product.Name,
			PurchaseCount: 0,
			StockQuantity: tx.Product.StockQuantity,
			UpdatedAt:     time.Now(),
		}
	}

	summary := s.SummeryMap[key]
	summary.PurchaseCount += tx.Quantity
	summary.StockQuantity = tx.Product.StockQuantity
	summary.UpdatedAt = time.Now()
}

func (s *ProductPurchaseAggregator) GetOutput() any {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.SummeryMap
}
