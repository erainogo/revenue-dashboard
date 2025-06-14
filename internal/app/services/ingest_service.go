package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/constants"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type IngestService struct {
	ctx                      context.Context
	logger                   *zap.SugaredLogger
	transactionRepository    adapters.TransactionRepository
	productSummeryRepository adapters.ProductSummeryRepository
	countryAggregator        adapters.CountryRevenueAggregator
}

type IngestServiceOptions func(*IngestService)

func WithLoggerI(logger *zap.SugaredLogger) IngestServiceOptions {
	return func(s *IngestService) {
		s.logger = logger
	}
}

func NewIngestService(
	ctx context.Context,
	transactionRepository adapters.TransactionRepository,
	productSummeryRepository adapters.ProductSummeryRepository,
	countryAggregator adapters.CountryRevenueAggregator,
	opts ...IngestServiceOptions,
) adapters.IngestService {
	svc := &IngestService{
		ctx:                      ctx,
		transactionRepository:    transactionRepository,
		productSummeryRepository: productSummeryRepository,
		countryAggregator:        countryAggregator,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

// IngestTransactionData IngestData worker function for data ingestion to the table.
// listen to the application context and exit early if there is a cancellation.
// prevents memory leaks and ensures the goroutine exits cleanly.
// creates a slice of data to insert to the db as a batch. fixed size 500 - batch size.
// reuse the same slice to avoid reallocating memory to a new slice.
// also update the summery map to later ingest precalculated country product summery data.
func (i IngestService) IngestTransactionData(ctx context.Context, tc <-chan entities.Transaction) {
	select {
	case <-ctx.Done():
		i.logger.Infow("context done", "error", ctx.Err())

		return
	default:
		// preallocate capacity.
		batch := make([]interface{}, 0, constants.BatchSize)

		for tx := range tc {
			batch = append(batch, tx)

			i.countryAggregator.Aggregate(tx)

			if len(batch) >= constants.BatchSize {
				err := i.transactionRepository.BulkInsert(ctx, batch)
				if err != nil {
					i.logger.Errorf("Insert error: %v", err)
				}

				if err != nil {
					return
				}

				//reset the slice.
				batch = batch[:0]
			}
		}

		if len(batch) > 0 {
			err := i.transactionRepository.BulkInsert(ctx, batch)

			if err != nil {
				i.logger.Errorf("Final insert error: %v", err)
			}
		}
	}
}

func (i IngestService) IngestProductSummery(ctx context.Context) error {
	err := i.productSummeryRepository.BulkInsert(ctx, i.countryAggregator.GetOutput())

	if err != nil {
		return err
	}

	return nil
}