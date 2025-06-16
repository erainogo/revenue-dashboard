package repositories

import (
	"context"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.uber.org/zap"
)

func (suite *RepositoriesTestSuite) TestGetRegionRevenue_Success() {
	mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock)).Run("returns_region_revenue", func(mt *mtest.T) {
		expected := bson.D{
			{"Region", "NA"},
			{"total_revenue", 5000.0},
			{"total_quantity", 100},
		}

		// Required: Add mock response for the Find operation
		first := bson.D{
			{"cursor", bson.D{
				{"firstBatch", bson.A{expected}},
				{"id", int64(0)},
				{"ns", "test.regionrevenue"},
			}},
			{"ok", 1},
		}

		mt.AddMockResponses(first)

		repo := NewRegionRevenueSummeryRepository(
			context.TODO(),
			mt.Coll,
			WithLoggerR(zap.NewNop().Sugar()),
		)

		results, err := repo.GetRegionRevenue(context.TODO())

		suite.asserts.NoError(err)
		suite.asserts.Equal("NA", results[0].Region)
		suite.asserts.Equal(float64(5000), results[0].TotalRevenue)
		suite.asserts.Equal(100, results[0].TotalQuantity)
	})
}

func (suite *RepositoriesTestSuite) TestGetRegionRevenue_Failure() {
	suite.mt.Run("returns error on query", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    1,
			Message: "mocked error",
			Name:    "CommandError",
		}))

		repo := NewRegionRevenueSummeryRepository(context.TODO(), mt.Coll, WithLoggerR(suite.logger))
		result, err := repo.GetRegionRevenue(context.TODO())

		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), result)
	})
}
