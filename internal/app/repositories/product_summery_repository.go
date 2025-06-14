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

const (
	// Mongodb indexes.
	countryProductTotalCompoundIndex = "country_product_revenue_index"
)

type ProductSummeryRepository struct {
	ctx        context.Context
	logger     *zap.SugaredLogger
	collection *mongo.Collection
}

type ProductSummeryRepositoryOptions func(*ProductSummeryRepository)

func WithLoggerP(logger *zap.SugaredLogger) ProductSummeryRepositoryOptions {
	return func(s *ProductSummeryRepository) {
		s.logger = logger
	}
}

func NewProductSummeryRepository(
	ctx context.Context,
	col *mongo.Collection,
	opts ...ProductSummeryRepositoryOptions,
) adapters.ProductSummeryRepository {
	svc := &ProductSummeryRepository{
		ctx:        ctx,
		collection: col,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (p *ProductSummeryRepository) BulkInsert(
	ctx context.Context,
	summaryMap map[entities.SummaryKey]*entities.CountryLevelRevenue,
) error {
	if len(summaryMap) == 0 {
		p.logger.Warn("Empty documents")

		return nil
	}

	var updates []mongo.WriteModel

	for key, summary := range summaryMap {
		filter := bson.M{
			"country":      key.Country,
			"product_name": key.ProductName,
		}

		update := bson.M{
			"$inc": bson.M{
				"total_revenue":     summary.TotalRevenue,
				"transaction_count": summary.TransactionCount,
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

	res, err := p.collection.BulkWrite(ctx, updates, bulkOpts)
	if err != nil {
		var bulkErr mongo.BulkWriteException
		if errors.As(err, &bulkErr) {
			for _, writeErr := range bulkErr.WriteErrors {
				p.logger.Errorf("Write error: %v", writeErr)
			}
		}
		return err
	}

	p.logger.Infof("Matched: %d, Modified: %d, Upserts: %d",
		res.MatchedCount, res.ModifiedCount, len(res.UpsertedIDs))

	return nil
}

func (p *ProductSummeryRepository) GetCountryLevelRevenueSortedByTotal(
	ctx context.Context,
	offset int,
	limit int,
) ([]*entities.CountryLevelRevenue, error) {
	opts := options.Find()
	opts.SetHint(countryProductTotalCompoundIndex)
	opts.SetSort(bson.D{{Key: "total_revenue", Value: -1}})

	//opts.SetSkip(int64(offset))
	//opts.SetLimit(int64(limit))

	cursor, err := p.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		p.logger.Errorf("Failed to query country summary: %v", err)
		return nil, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			p.logger.Errorf("Failed to close cursor: %v", err)
		}
	}()

	var results []*entities.CountryLevelRevenue

	for cursor.Next(ctx) {
		var summary entities.CountryLevelRevenue

		if err := cursor.Decode(&summary); err != nil {
			return nil, err
		}

		p.logger.Infof("Found country summary: %+v", summary)

		results = append(results, &summary)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
