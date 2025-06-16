package repositories

import (
	"context"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"time"
)

func (suite *RepositoriesTestSuite) TestGetFrequentlyPurchasedProducts_Success() {
	suite.mt.Run("returns purchased products", func(mt *mtest.T) {
		expected := &entities.ProductPurchaseSummary{
			ProductID:     "P001",
			ProductName:   "Wireless Mouse",
			PurchaseCount: 50,
			StockQuantity: 100,
			UpdatedAt:     time.Now(),
		}

		doc := bson.D{
			{"product_id", expected.ProductID},
			{"product_name", expected.ProductName},
			{"purchase_count", expected.PurchaseCount},
			{"stock_quantity", expected.StockQuantity},
			{"updated_at", expected.UpdatedAt},
		}

		first := bson.D{
			{"cursor", bson.D{
				{"firstBatch", bson.A{doc}},
				{"id", int64(0)},
				{"ns", "test.product_purchase"},
			}},
			{"ok", 1},
		}

		mt.AddMockResponses(first)

		repo := NewPurchaseSummeryRepository(context.TODO(), mt.Coll, WithLoggerPS(suite.logger))
		result, err := repo.GetFrequentlyPurchasedProducts(context.TODO())

		assert.NoError(suite.T(), err)
		assert.Len(suite.T(), result, 1)
		assert.Equal(suite.T(), expected.ProductID, result[0].ProductID)
	})
}

func (suite *RepositoriesTestSuite) TestGetFrequentlyPurchasedProducts_Failure() {
	suite.mt.Run("returns error on find", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    2,
			Message: "mocked error",
		}))

		repo := NewPurchaseSummeryRepository(context.TODO(), mt.Coll, WithLoggerPS(suite.logger))
		result, err := repo.GetFrequentlyPurchasedProducts(context.TODO())

		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), result)
	})
}