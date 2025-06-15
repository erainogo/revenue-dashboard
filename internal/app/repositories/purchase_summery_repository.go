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
// countryProductTotalCompoundIndex = "country_product_revenue_index"
)

type PurchaseSummeryRepository struct {
	ctx        context.Context
	logger     *zap.SugaredLogger
	collection *mongo.Collection
}

type PurchaseSummeryRepositoryOptions func(*PurchaseSummeryRepository)

func WithLoggerPS(logger *zap.SugaredLogger) PurchaseSummeryRepositoryOptions {
	return func(s *PurchaseSummeryRepository) {
		s.logger = logger
	}
}

func NewPurchaseSummeryRepository(
	ctx context.Context,
	col *mongo.Collection,
	opts ...PurchaseSummeryRepositoryOptions,
) adapters.PurchaseSummeryRepository {
	svc := &PurchaseSummeryRepository{
		ctx:        ctx,
		collection: col,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (p *PurchaseSummeryRepository) BulkInsert(
	ctx context.Context,
	summaryMap map[string]*entities.ProductPurchaseSummary,
) error {
	if len(summaryMap) == 0 {
		p.logger.Warn("Empty documents")

		return nil
	}

	var updates []mongo.WriteModel

	for _, summary := range summaryMap {
		filter := bson.M{"product_id": summary.ProductID}

		update := bson.M{
			"$inc": bson.M{"purchase_count": summary.PurchaseCount},
			"$set": bson.M{
				"product_name":   summary.ProductName,
				"stock_quantity": summary.StockQuantity,
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

func (p *PurchaseSummeryRepository) GetFrequentlyPurchasedProducts(ctx context.Context,
) ([]*entities.ProductPurchaseSummary, error) {
	opts := options.Find()
	opts.SetSort(bson.M{"purchase_count": -1})
	opts.SetLimit(20)

	cursor, err := p.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		p.logger.Errorf("Failed to query frequently purchased products: %v", err)
		return nil, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			p.logger.Errorf("Failed to close cursor: %v", err)
		}
	}()

	var results []*entities.ProductPurchaseSummary

	for cursor.Next(ctx) {
		var summary entities.ProductPurchaseSummary

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
