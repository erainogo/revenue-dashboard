package adapters

import (
	"context"

	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type IngestService interface {
	IngestTransactionData(ctx context.Context,tc <-chan entities.Transaction)
	IngestProductSummery(ctx context.Context) error
}