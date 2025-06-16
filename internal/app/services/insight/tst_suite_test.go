package insight

import (
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/suite"

	"github.com/erainogo/revenue-dashboard/mocks/adapters"
)

type InsightServiceTestSuite struct {
	suite.Suite
	logger                         *zap.SugaredLogger
	productSummeryRepository       *adapters.MockProductSummeryRepository
	purchaseSummeryRepository      *adapters.MockPurchaseSummeryRepository
	monthlySalesSummeryRepository  *adapters.MockMonthlySalesSummeryRepository
	regionRevenueSummeryRepository *adapters.MockRegionRevenueSummeryRepository
}

func (suite *InsightServiceTestSuite) SetupTest() {
	suite.productSummeryRepository    =   adapters.NewMockProductSummeryRepository(suite.T())
	suite.purchaseSummeryRepository    =  adapters.NewMockPurchaseSummeryRepository(suite.T())
	suite.monthlySalesSummeryRepository  = adapters.NewMockMonthlySalesSummeryRepository(suite.T())
	suite.regionRevenueSummeryRepository = adapters.NewMockRegionRevenueSummeryRepository(suite.T())
	suite.logger = zap.NewNop().Sugar()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(InsightServiceTestSuite))
}
