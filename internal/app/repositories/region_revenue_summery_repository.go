package repositories

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

const (
// Mongodb indexes.
// countryProductTotalCompoundIndex = "country_product_revenue_index"
)

type RegionRevenueSummeryRepository struct {
	ctx        context.Context
	logger     *zap.SugaredLogger
	collection *mongo.Collection
}

type RegionRevenueSummeryRepositoryOptions func(*RegionRevenueSummeryRepository)

func WithLoggerR(logger *zap.SugaredLogger) RegionRevenueSummeryRepositoryOptions {
	return func(s *RegionRevenueSummeryRepository) {
		s.logger = logger
	}
}

func NewRegionRevenueSummeryRepository(
	ctx context.Context,
	col *mongo.Collection,
	opts ...RegionRevenueSummeryRepositoryOptions,
) adapters.RegionRevenueSummeryRepository {
	svc := &RegionRevenueSummeryRepository{
		ctx:        ctx,
		collection: col,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (r RegionRevenueSummeryRepository) BulkInsert(ctx context.Context, summaryMap map[string]*entities.RegionRevenue) error {
	if len(summaryMap) == 0 {
		r.logger.Warn("Empty documents")

		return nil
	}

	var updates []mongo.WriteModel

	for _, summary := range summaryMap {
		filter := bson.M{"Region": summary.Region}

		update := bson.M{
			"$inc": bson.M{
				"total_revenue":  summary.TotalRevenue,
				"total_quantity": summary.TotalQuantity,
			},
			"$set": bson.M{
				"updated_at": summary.UpdatedAt,
			},
		}

		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		updates = append(updates, model)
	}

	bulkOpts := options.BulkWrite().SetOrdered(false)

	res, err := r.collection.BulkWrite(ctx, updates, bulkOpts)
	if err != nil {
		var bulkErr mongo.BulkWriteException

		if errors.As(err, &bulkErr) {
			for _, writeErr := range bulkErr.WriteErrors {
				r.logger.Errorf("Write error: %v", writeErr)
			}
		}

		return err
	}

	r.logger.Infof("Matched: %d, Modified: %d, Upserts: %d",
		res.MatchedCount, res.ModifiedCount, len(res.UpsertedIDs))

	return nil
}

func (r RegionRevenueSummeryRepository) GetRegionRevenue(ctx context.Context) ([]*entities.RegionRevenue, error) {
	//TODO implement me
	panic("implement me")
}
