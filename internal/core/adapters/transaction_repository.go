package adapters

import (
	"context"
)

type TransactionRepository interface {
	BulkInsert(ctx context.Context, docs []interface{}) error
}
