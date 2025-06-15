package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type MonthlySalesSummeryRepository interface {
	BulkInsert(ctx context.Context, docs map[string]*entities.MonthlySales) error
	GetMonthlyRevenue(ctx context.Context) ([]*entities.MonthlySales, error)
}