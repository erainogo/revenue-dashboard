package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type InsightService interface {
	 GetCountryLevelRevenue(ctx context.Context,  page int, limit int) ([]*entities.CountryLevelRevenue, error)
}
