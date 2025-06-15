package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type IngestService interface {
	IngestTransactionData(ctx context.Context, tc <-chan entities.Transaction)
	IngestCountrySummery(ctx context.Context) error
	IngestPurchaseSummery(ctx context.Context) error
	IngestMonthlySalesSummery(ctx context.Context) error
	IngestRegionRevenueSummery(ctx context.Context) error
}
