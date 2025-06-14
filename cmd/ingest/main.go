package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/erainogo/revenue-dashboard/cmd/initializations"
	"github.com/erainogo/revenue-dashboard/internal/app/aggregators"
	"github.com/erainogo/revenue-dashboard/internal/app/repositories"
	"github.com/erainogo/revenue-dashboard/internal/app/services"
	"github.com/erainogo/revenue-dashboard/internal/config"
	"github.com/erainogo/revenue-dashboard/internal/handlers"
	"github.com/erainogo/revenue-dashboard/pkg/constants"
)

func main() {
	if len(os.Args) < constants.ARGS {
		fmt.Println("please provide input file: <input.csv>")

		os.Exit(1)
	}

	logger := initializations.SetUpLogger()

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	mongoClient, err := initializations.CreateMongoClient(ctx, logger)

	if err != nil {
		logger.Fatal("failed to connect to mongo: %v", err)
	}

	defer func() {
		err = mongoClient.Disconnect(ctx)
		if err != nil {
			logger.Error(err)
		}
	}()

	// background routine to shut down server if signal received
	// this will wait for the ch chan to receive the exit signals from the os.
	// if received cancel the context.
	go func() {
		sig := <-ch
		logger.Infof("Got %s signal. Cancelling", sig)

		cancel()

		err = mongoClient.Disconnect(ctx)
		if err != nil {
			logger.Error(err)
		}

		logger.Info("Server gracefully stopped")
	}()

	transactionRepository := repositories.NewTransactionRepository(ctx,
		mongoClient.Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoTransactionCollectionName),
		repositories.WithLogger(logger))

	productSummeryRepository := repositories.NewProductSummeryRepository(ctx,
		mongoClient.Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoCountryProductSummaryCollection),
		repositories.WithLoggerP(logger))

	countryAggregator := aggregators.NewCountryRevenueAggregator(
		ctx, productSummeryRepository, aggregators.WithLogger(logger))

	service := services.NewIngestService(
		ctx,
		transactionRepository,
		productSummeryRepository,
		countryAggregator,
		services.WithLoggerI(logger))

	handler := handlers.NewCli(ctx, service, handlers.WithLoggerC(logger))

	inputPath := os.Args[1]

	logger.Info("Started ingesting report")

	err = handler.Ingest(ctx, inputPath)
	if err != nil {
		return
	}

	logger.Info("Ingestion completed successfully.")
}
