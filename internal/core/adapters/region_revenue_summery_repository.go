package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type RegionRevenueSummeryRepository interface {
	BulkInsert(ctx context.Context, docs map[string]*entities.RegionRevenue) error
	GetRegionRevenue(ctx context.Context) ([]*entities.RegionRevenue, error)
}