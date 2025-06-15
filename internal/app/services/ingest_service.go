package services

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/constants"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type IngestService struct {
	ctx                            context.Context
	logger                         *zap.SugaredLogger
	transactionRepository          adapters.TransactionRepository
	productSummeryRepository       adapters.ProductSummeryRepository
	purchaseRepository             adapters.PurchaseSummeryRepository
	countryAggregator              adapters.Aggregator
	purchaseAggregator             adapters.Aggregator
	monthlySalesAggregator         adapters.Aggregator
	regionRevenueAggregator        adapters.Aggregator
	monthlySalesSummeryRepository  adapters.MonthlySalesSummeryRepository
	regionRevenueSummeryRepository adapters.RegionRevenueSummeryRepository
}

type IngestServiceOptions func(*IngestService)

func WithLoggerI(logger *zap.SugaredLogger) IngestServiceOptions {
	return func(s *IngestService) {
		s.logger = logger
	}
}

func NewIngestService(
	ctx context.Context,
	transactionRepository adapters.TransactionRepository,
	productSummeryRepository adapters.ProductSummeryRepository,
	purchaseRepository adapters.PurchaseSummeryRepository,
	monthlySalesSummeryRepository adapters.MonthlySalesSummeryRepository,
	regionRevenueSummeryRepository adapters.RegionRevenueSummeryRepository,
	countryAggregator adapters.Aggregator,
	purchaseAggregator adapters.Aggregator,
	monthlySalesAggregator adapters.Aggregator,
	regionRevenueAggregator adapters.Aggregator,
	opts ...IngestServiceOptions,
) adapters.IngestService {
	svc := &IngestService{
		ctx:                           ctx,
		transactionRepository:         transactionRepository,
		productSummeryRepository:      productSummeryRepository,
		purchaseRepository:            purchaseRepository,
		monthlySalesSummeryRepository: monthlySalesSummeryRepository,
		regionRevenueSummeryRepository:regionRevenueSummeryRepository,
		countryAggregator:             countryAggregator,
		purchaseAggregator:            purchaseAggregator,
		monthlySalesAggregator:        monthlySalesAggregator,
		regionRevenueAggregator:       regionRevenueAggregator,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

// IngestTransactionData IngestData worker function for data ingestion to the table.
// listen to the application context and exit early if there is a cancellation.
// prevents memory leaks and ensures the goroutine exits cleanly.
// creates a slice of data to insert to the db as a batch. fixed size 500 - batch size.
// reuse the same slice to avoid reallocating memory to a new slice.
// also update the summery map to later ingest precalculated country product summery data.
func (i *IngestService) IngestTransactionData(ctx context.Context, tc <-chan entities.Transaction) {
	select {
	case <-ctx.Done():
		i.logger.Infow("context done", "error", ctx.Err())

		return
	default:
		// preallocate capacity.
		batch := make([]interface{}, 0, constants.BatchSize)

		for tx := range tc {
			batch = append(batch, tx)

			i.RunAggregators(tx)

			if len(batch) >= constants.BatchSize {
				err := i.transactionRepository.BulkInsert(ctx, batch)
				if err != nil {
					i.logger.Errorf("Insert error: %v", err)
				}

				if err != nil {
					return
				}

				//reset the slice.
				batch = batch[:0]
			}
		}

		if len(batch) > 0 {
			err := i.transactionRepository.BulkInsert(ctx, batch)

			if err != nil {
				i.logger.Errorf("Final insert error: %v", err)
			}
		}
	}
}

func (i *IngestService) RunAggregators(tx entities.Transaction) {
	i.countryAggregator.Aggregate(tx)
	i.purchaseAggregator.Aggregate(tx)
	i.monthlySalesAggregator.Aggregate(tx)
	i.regionRevenueAggregator.Aggregate(tx)
}

func (i *IngestService) IngestCountrySummery(ctx context.Context) error {
	rwo := i.countryAggregator.GetOutput()

	output, ok := rwo.(map[entities.CountrySummaryKey]*entities.CountryLevelRevenue)
	if !ok {
		return fmt.Errorf("unexpected type for output: %T", rwo)
	}

	err := i.productSummeryRepository.BulkInsert(ctx, output)

	if err != nil {
		return err
	}

	return nil
}

func (i *IngestService) IngestPurchaseSummery(ctx context.Context) error {
	rwo := i.purchaseAggregator.GetOutput()

	output, ok := rwo.(map[string]*entities.ProductPurchaseSummary)
	if !ok {
		return fmt.Errorf("unexpected type for output: %T", rwo)
	}

	err := i.purchaseRepository.BulkInsert(ctx, output)

	if err != nil {
		return err
	}

	return nil
}

func (i *IngestService) IngestMonthlySalesSummery(ctx context.Context) error {
	rwo := i.monthlySalesAggregator.GetOutput()

	output, ok := rwo.(map[string]*entities.MonthlySales)
	if !ok {
		return fmt.Errorf("unexpected type for output: %T", rwo)
	}

	err := i.monthlySalesSummeryRepository.BulkInsert(ctx, output)

	if err != nil {
		return err
	}

	return nil
}

func (i *IngestService) IngestRegionRevenueSummery(ctx context.Context) error {
	rwo := i.regionRevenueAggregator.GetOutput()

	output, ok := rwo.(map[string]*entities.RegionRevenue)
	if !ok {
		return fmt.Errorf("unexpected type for output: %T", rwo)
	}

	err := i.regionRevenueSummeryRepository.BulkInsert(ctx, output)

	if err != nil {
		return err
	}

	return nil
}