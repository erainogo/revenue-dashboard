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

func (i *InsightService) IngestData(ctx context.Context, tc <-chan entities.Transaction) {
	batch := make([]interface{}, 0, constants.BatchSize)

	for tx := range tc {
		batch = append(batch, tx)

		if len(batch) >= constants.BatchSize {
			err := i.repository.BulkInsert(ctx, batch)
			if err != nil {
				i.logger.Errorf("Insert error: %v", err)
			}

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
