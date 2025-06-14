package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type InsightService interface {
	 IngestTransactionData(ctx context.Context,tc <-chan entities.Transaction)
	 GetCountryLevelRevenue(ctx context.Context,  page int, limit int) ([]*entities.CountryLevelRevenue, error)
	 IngestProductSummery(ctx context.Context) error
}
