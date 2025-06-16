package repositories

import (
	"context"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func (suite *RepositoriesTestSuite) TestGetMonthlyRevenue_Success() {
	mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock)).Run("returns_monthly_revenue", func(mt *mtest.T) {
		expected := bson.D{
			{"year", 2021},
			{"month", "February"},
			{"total_quantity", 100},
		}

		// Required: Add mock response for the Find operation
		first := bson.D{
			{"cursor", bson.D{
				{"firstBatch", bson.A{expected}},
				{"id", int64(0)},
				{"ns", "test.monthlyrevenue"},
			}},
			{"ok", 1},
		}

		mt.AddMockResponses(first)

		repo := NewMonthlySalesSummeryRepository(
			context.TODO(),
			mt.Coll,
			WithLoggerM(suite.logger),
		)

		results, err := repo.GetMonthlyRevenue(context.TODO())

		suite.asserts.NoError(err)
		suite.asserts.Equal(2021, results[0].Year)
		suite.asserts.Equal("February", results[0].Month)
		suite.asserts.Equal(100, results[0].TotalQuantity)
	})
}

func (suite *RepositoriesTestSuite) TestGetMonthlyRevenue_Failure() {
	suite.mt.Run("returns error on query", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    1,
			Message: "mocked error",
			Name:    "CommandError",
		}))

		repo := NewMonthlySalesSummeryRepository(context.TODO(), mt.Coll, WithLoggerM(suite.logger))
		result, err := repo.GetMonthlyRevenue(context.TODO())

		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), result)
	})
}
