package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
)

type TransactionRepository struct {
	ctx        context.Context
	logger     *zap.SugaredLogger
	collection *mongo.Collection
}

type TransactionRepositoryOptions func(*TransactionRepository)

func WithLogger(logger *zap.SugaredLogger) TransactionRepositoryOptions {
	return func(s *TransactionRepository) {
		s.logger = logger
	}
}

func NewTransactionRepository(
	ctx context.Context,
	col *mongo.Collection,
	opts ...TransactionRepositoryOptions,
) adapters.TransactionRepository {
	svc := &TransactionRepository{
		ctx: ctx,
		collection: col,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (r *TransactionRepository) BulkInsert(ctx context.Context, docs []interface{}) error {
	if len(docs) == 0 {
		r.logger.Warn("Empty documents")

		return nil
	}

	res, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		var bulkErr mongo.BulkWriteException

		if errors.As(err, &bulkErr) {
			for _, writeErr := range bulkErr.WriteErrors {
				r.logger.Errorf("Write error: %v\n", writeErr)
			}
		}

		return err
	}

	r.logger.Infof("Inserted %d documents\n", len(res.InsertedIDs))

	return nil
}
