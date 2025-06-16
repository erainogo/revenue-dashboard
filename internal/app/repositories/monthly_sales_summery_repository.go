package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type MonthlySalesSummeryRepository struct {
	ctx        context.Context
	logger     *zap.SugaredLogger
	collection *mongo.Collection
}

type MonthlySalesSummeryRepositoryOptions func(*MonthlySalesSummeryRepository)

func WithLoggerM(logger *zap.SugaredLogger) MonthlySalesSummeryRepositoryOptions {
	return func(s *MonthlySalesSummeryRepository) {
		s.logger = logger
	}
}

func NewMonthlySalesSummeryRepository(
	ctx context.Context,
	col *mongo.Collection,
	opts ...MonthlySalesSummeryRepositoryOptions,
) adapters.MonthlySalesSummeryRepository {
	svc := &MonthlySalesSummeryRepository{
		ctx:        ctx,
		collection: col,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (m MonthlySalesSummeryRepository) BulkInsert(ctx context.Context, summaryMap map[string]*entities.MonthlySales) error {
	if len(summaryMap) == 0 {
		m.logger.Warn("Empty documents")

		return nil
	}

	var updates []mongo.WriteModel

	for _, summary := range summaryMap {
		filter := bson.M{"year": summary.Year, "month": summary.Month}

		update := bson.M{
			"$inc": bson.M{"total_quantity": summary.TotalQuantity},
			"$set": bson.M{
				"updated_at":     summary.UpdatedAt,
			},
		}

		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		updates = append(updates, model)
	}

	bulkOpts := options.BulkWrite().SetOrdered(false)

	res, err := m.collection.BulkWrite(ctx, updates, bulkOpts)
	if err != nil {
		var bulkErr mongo.BulkWriteException

		if errors.As(err, &bulkErr) {
			for _, writeErr := range bulkErr.WriteErrors {
				m.logger.Errorf("Write error: %v", writeErr)
			}
		}

		return err
	}

	m.logger.Infof("Matched: %d, Modified: %d, Upserts: %d",
		res.MatchedCount, res.ModifiedCount, len(res.UpsertedIDs))

	return nil
}

func (m MonthlySalesSummeryRepository) GetMonthlyRevenue(ctx context.Context) ([]*entities.MonthlySales, error) {
	opts := options.Find()
	opts.SetSort(bson.M{
		"total_quantity": -1,
	})
	opts.SetLimit(30)

	cursor, err := m.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		m.logger.Errorf("Failed to query monthly sales: %v", err)
		return nil, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			m.logger.Errorf("Failed to close cursor: %v", err)
		}
	}()

	var results []*entities.MonthlySales

	for cursor.Next(ctx) {
		var summary entities.MonthlySales

		if err := cursor.Decode(&summary); err != nil {
			return nil, err
		}

		results = append(results, &summary)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}