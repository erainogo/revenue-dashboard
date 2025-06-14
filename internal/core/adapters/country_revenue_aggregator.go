package adapters

import (
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type CountryRevenueAggregator interface {
    Aggregate(key entities.Transaction)
	GetOutput() map[entities.SummaryKey]*entities.CountryLevelRevenue
}
