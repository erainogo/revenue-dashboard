package adapters

import (
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type Aggregator interface {
	Aggregate(tx entities.Transaction)
	GetOutput() any
}
