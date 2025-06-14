package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/erainogo/revenue-dashboard/cmd/init"
	"github.com/erainogo/revenue-dashboard/internal/app/repositories"
	"github.com/erainogo/revenue-dashboard/internal/app/services"
	"github.com/erainogo/revenue-dashboard/internal/config"
	"github.com/erainogo/revenue-dashboard/pkg/constants"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

func main() {
	if len(os.Args) < constants.ARGS {
		fmt.Println("please provide input file: <input.csv>")

		os.Exit(1)
	}

	logger := init.SetUpLogger()

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	mongoClient, err := init.CreateMongoClient(ctx, logger)

	if err != nil {
		logger.Fatal("failed to connect to mongo: %v", err)
	}

	// background routine to shut down server if signal received
	// this will wait for the ch chan to receive the exit signals from the os.
	// if received cancel the context.
	go func() {
		sig := <-ch
		logger.Infof("Got %s signal. Cancelling", sig)

		cancel()

		defer func() {
			err = mongoClient.Disconnect(ctx)
			if err != nil {
				logger.Error(err)
			}
		}()

		logger.Info("Server gracefully stopped")
	}()

	repository := repositories.NewTransactionRepository(ctx,
		mongoClient.Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoTransactionCollectionName),
		repositories.WithLogger(logger))

	service := services.NewInsightService(
		ctx, repository, services.WithLogger(logger))

	inputPath := os.Args[1]

	logger.Info("Started ingesting report")

	file, err := os.Open(inputPath)
	if err != nil {
		logger.Errorf("Failed to open input file")

		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Error("failed to close file", zap.Error(err))
		}
	}(file)

	// to reduce read() system calls use bufio
	// It reads large chunks of data into memory and then serves that buffer to the consumer
	r := csv.NewReader(bufio.NewReader(file))
	_, err = r.Read()
	if err != nil {
		logger.Errorf("Failed reading csv")

		os.Exit(1)
	}

	// wg for wait all ingest workers to be done
	var wg sync.WaitGroup
	//make channel to send transactions .
	tc := make(chan entities.Transaction, 500)
    // close channel when done.
	defer close(tc)
	// deploy a worker pool for concurrent run ingestion
	for i := 0; i < constants.WorkerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			service.IngestData(ctx, tc)
		}()
	}

	ln := 0

	for {
		record, err := r.Read()
		if err != nil {
			break
		}

		ln++

		tx, err := parseRecord(record)
		if err != nil {
			logger.Warnf("Skipping line %d: %v", ln, err)

			continue
		}
		// send record to the channel
		tc <- tx
	}

	wg.Wait()

	logger.Infof("total number of %v lines has been processed", ln)

	logger.Info("Ingestion completed successfully.")

	os.Exit(1)
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
