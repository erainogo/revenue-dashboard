package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)


type PurchaseSummeryRepository interface {
	BulkInsert(ctx context.Context, docs map[string]*entities.ProductPurchaseSummary) error
	GetFrequentlyPurchasedProducts(ctx context.Context) ([]*entities.ProductPurchaseSummary, error)
}

