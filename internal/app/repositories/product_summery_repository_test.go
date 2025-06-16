package repositories

import (
	"context"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func (suite *RepositoriesTestSuite) TestGetCountryLevelRevenueSortedByTotal_Success() {
	mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock)).Run("returns_country_level_revenue", func(mt *mtest.T) {
		expected := bson.D{
			{"country", "USA"},
			{"product_name", "ProductX"},
			{"total_revenue", 10000.0},
			{"transaction_count", 50},
			{"updated_at", time.Now()},
		}

		first := bson.D{
			{"cursor", bson.D{
				{"firstBatch", bson.A{expected}},
				{"id", int64(0)}, // Must be int64
				{"ns", "test.productsummary"},
			}},
			{"ok", 1},
		}

		mt.AddMockResponses(first)

		repo := NewProductSummeryRepository(
			context.TODO(),
			mt.Coll,
			WithLoggerP(suite.logger),
		)

		results, err := repo.GetCountryLevelRevenueSortedByTotal(context.TODO(), 0, 50)

		suite.asserts.NoError(err)
		suite.asserts.Len(results, 1)

		actual := results[0]
		suite.asserts.Equal("USA", actual.Country)
		suite.asserts.Equal("ProductX", actual.ProductName)
		suite.asserts.Equal(float64(10000.0), actual.TotalRevenue)
		suite.asserts.Equal(int32(50), actual.TransactionCount)
	})
}

func (suite *RepositoriesTestSuite) TestGetCountryLevelRevenueSortedByTotal_Failure() {
	suite.mt.Run("returns error on find query", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    123,
			Message: "mocked find error",
			Name:    "FindError",
		}))

		repo := NewProductSummeryRepository(
			context.TODO(),
			mt.Coll,
			WithLoggerP(suite.logger),
		)

		result, err := repo.GetCountryLevelRevenueSortedByTotal(context.TODO(), 0, 10)

		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), result)
	})
}
