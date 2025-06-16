package repositories

import (
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"go.uber.org/zap"
)

type RepositoriesTestSuite struct {
	suite.Suite
	asserts                        *assert.Assertions
	logger                         *zap.SugaredLogger
	mt                             *mtest.T
}

func (suite *RepositoriesTestSuite) SetupTest() {
	suite.asserts = assert.New(suite.T())
	suite.mt = mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock))
	suite.logger = zap.NewNop().Sugar()
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(RepositoriesTestSuite))
}

