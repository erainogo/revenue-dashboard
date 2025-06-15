package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type ProductSummeryRepository interface {
	BulkInsert(ctx context.Context, docs map[entities.CountrySummaryKey]*entities.CountryLevelRevenue) error
	GetCountryLevelRevenueSortedByTotal(
		ctx context.Context, offset int, imit int) ([]*entities.CountryLevelRevenue, error)
}
