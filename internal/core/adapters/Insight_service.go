package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type InsightService interface {
	GetCountryLevelRevenue(ctx context.Context, page int, limit int) ([]*entities.CountryLevelRevenue, error)
	GetFrequentlyPurchasedProducts(ctx context.Context) ([]*entities.ProductPurchaseSummary, error)
	GetRegionRevenue(ctx context.Context) ([]*entities.RegionRevenue, error)
	GetMonthlyRevenue(ctx context.Context) ([]*entities.MonthlySales, error)
}
