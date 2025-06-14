package services

import (
	"context"
	"github.com/erainogo/revenue-dashboard/pkg/constants"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
)

type InsightService struct {
	ctx        context.Context
	logger     *zap.SugaredLogger
	repository adapters.TransactionRepository
}

type InsightServiceOptions func(*InsightService)

func WithLogger(logger *zap.SugaredLogger) InsightServiceOptions {
	return func(s *InsightService) {
		s.logger = logger
	}
}

func NewInsightService(
	ctx context.Context,
	repository adapters.TransactionRepository,
	opts ...InsightServiceOptions,
) adapters.InsightService {
	svc := &InsightService{
		ctx:        ctx,
		repository: repository,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

// IngestData worker function for data ingestion to the table.
// listen to the application context and exit early if there is a cancellation.
// prevents memory leaks and ensures the goroutine exits cleanly.
// creates a slice of data to insert to the db as a batch. fixed size 500 - batch size.
// reuse the same slice to avoid reallocating memory to a new slice.
func (i *InsightService) IngestData(ctx context.Context, tc <-chan entities.Transaction) {
	select {
	case <-ctx.Done():
        i.logger.Infow("context done", "error", ctx.Err())

		return
	default:
		// preallocate capacity.
		batch := make([]interface{}, 0, constants.BatchSize)

		for tx := range tc {
			batch = append(batch, tx)

			if len(batch) >= constants.BatchSize {
				err := i.repository.BulkInsert(ctx, batch)
				if err != nil {
					i.logger.Errorf("Insert error: %v", err)
				}

				//reset the slice.
				batch = batch[:0]
			}
		}

		if len(batch) > 0 {
			err := i.repository.BulkInsert(ctx, batch)

			if err != nil {
				i.logger.Errorf("Final insert error: %v", err)
			}
		}
	}
}
