package handlers

import (
	"bufio"
	"context"
	"encoding/csv"
	"os"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
	"github.com/erainogo/revenue-dashboard/pkg/constants"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

type Cli struct {
	ctx     context.Context
	logger  *zap.SugaredLogger
	service adapters.IngestService
}

type CliOptions func(*Cli)

func WithLoggerC(logger *zap.SugaredLogger) CliOptions {
	return func(s *Cli) {
		s.logger = logger
	}
}

func NewCli(
	ctx context.Context,
	service adapters.IngestService,
	opts ...CliOptions,
) *Cli {
	svc := &Cli{
		ctx:     ctx,
		service: service,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *Cli) Ingest(ctx context.Context, inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		s.logger.Errorf("Failed to open input file")

		os.Exit(1)
	}

	// to reduce read() system calls use bufio
	// It reads large chunks of data into memory and then serves that buffer to the consumer
	r := csv.NewReader(bufio.NewReader(file))
	_, err = r.Read()
	if err != nil {
		s.logger.Errorf("Failed reading csv")

		os.Exit(1)
	}

	// wg for wait all ingest workers to be done
	var wg sync.WaitGroup
	//make channel to send transactions .
	tc := make(chan entities.Transaction, 500)
	// deploy a worker pool for concurrent run ingestion
	for i := 0; i < constants.WorkerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.service.IngestTransactionData(ctx, tc)
		}()
	}

	ln := 0
	sln := 0

	for {
		record, err := r.Read()
		if err != nil {
			break
		}

		ln++

		tx, err := parseRecord(record)
		if err != nil {
			s.logger.Warnf("Skipping line %d: %v", ln, err)

			sln++

			continue
		}
		// send record to the channel
		tc <- tx
	}

	s.logger.Info("Finished sending transactions, closing channel")
	// close channel when done.
	close(tc)

	wg.Wait()

	s.logger.Info("All ingestion workers done")

	// we can close the file after file has been imported.
	// pre aggregated in memory cache (map) will be read to insert insight after that.
	err = file.Close()
	if err != nil {
		s.logger.Error("failed to close file", zap.Error(err))
	}

	// when transaction ingesting to the db,
	// in the background in memory cache is being used to pre-aggregate the results for insights
	// after file data ingestion is done, those maps are being inserted here as upsert bulks.
	// doing this to reduce calculating data when fetching via API
	// also easy to cache for the frond-end in the future
	// if we want these data to be updated real time, ( if file ingestion happens more often we have to update the summaries )
	// we can deploy background thread to update the db.
	err = s.service.IngestCountrySummery(ctx)
	if err != nil {
		s.logger.Warnf("Failed to insert bulk product summery: %v", err)

		return err
	}

	err = s.service.IngestPurchaseSummery(ctx)
	if err != nil {
		s.logger.Warnf("Failed to insert bulk purchase summery: %v", err)

		return err
	}

	s.logger.Infof("total number of %v lines has been processed ", ln)
	s.logger.Infof("total number of  %v lines has been skipped ", sln)

	return nil
}

func parseRecord(record []string) (entities.Transaction, error) {
	transactionDate, err := time.Parse(constants.TimeLayout, record[1])
	if err != nil {
		return entities.Transaction{}, err
	}

	price, _ := strconv.ParseFloat(record[8], 64)
	quantity, _ := strconv.Atoi(record[9])
	totalPrice, _ := strconv.ParseFloat(record[10], 64)
	stockQuantity, _ := strconv.Atoi(record[11])
	addedDate, err := time.Parse(constants.TimeLayout, record[12])

	if err != nil {
		return entities.Transaction{}, err
	}

	return entities.Transaction{
		TransactionID:   record[0],
		TransactionDate: transactionDate,
		UserID:          record[2],
		Country:         record[3],
		Region:          record[4],
		Product: entities.Product{
			ID:            record[5],
			Name:          record[6],
			Category:      record[7],
			StockQuantity: stockQuantity,
		},
		Price:      price,
		Quantity:   quantity,
		TotalPrice: totalPrice,
		AddedDate:  addedDate,
	}, nil
}
