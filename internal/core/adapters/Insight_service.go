package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type InsightService interface {
	 IngestData(ctx context.Context,tc <-chan entities.Transaction)
	 GetCountryLevelRevenue(ctx context.Context,  page int, limit int) ([]*entities.CountryLevelRevenue, error)
	 InsertBulkProductSummery(ctx context.Context) error
}
