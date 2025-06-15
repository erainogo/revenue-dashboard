package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"

	"github.com/erainogo/revenue-dashboard/cmd/initializations"
	"github.com/erainogo/revenue-dashboard/internal/app/repositories"
	"github.com/erainogo/revenue-dashboard/internal/app/services"
	"github.com/erainogo/revenue-dashboard/internal/config"
	"github.com/erainogo/revenue-dashboard/internal/handlers"
)

func main() {
	logger := initializations.SetUpLogger()

	// context for the application
	ctx, cancel := context.WithCancel(context.Background())

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", *config.Config.HttpPort),
		WriteTimeout: time.Duration(*config.Config.WriteTimeOut) * time.Second,
		ReadTimeout:  time.Duration(*config.Config.ReadTimeOut) * time.Second,
	}

	mongoClient, err := initializations.CreateMongoClient(ctx, logger)

	if err != nil {
		logger.Fatal("failed to connect to mongo: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// background routine to shut down server if signal received
	// this will wait for the ch chan to receive the exit signals from the os.
	go func() {
		sig := <-ch
		logger.Infof("Got %s signal. Cancelling", sig)
		// shut down background routines
		defer cancel()

		shutdownCtx, shutdownRelease := context.WithTimeout(ctx, 5*time.Second)

		defer shutdownRelease()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Errorf("Shutdown error: %s", err)
		}

		defer func() {
			err = mongoClient.Disconnect(ctx)
			if err != nil {
				logger.Error(err)
			}
		}()

		defer func() {
			if err := logger.Sync(); err != nil && !errors.Is(err, os.ErrInvalid) {
				logger.Errorf("Failed to sync logger: %v", err)
			}
		}()

		logger.Info("Server gracefully stopped")
	}()

	if *config.Config.MongoDBMigrate {
		logger.Info("Migrating ... ")

		err = initializations.RunMigration(mongoClient)

		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			logger.Fatal(err)
		}

		return
	}

	// create repositories for db access layer.
	productSummeryRepository := repositories.NewProductSummeryRepository(ctx,
		mongoClient.Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoCountryProductSummaryCollection),
		repositories.WithLoggerP(logger))

	purchaseSummeryRepository := repositories.NewPurchaseSummeryRepository(ctx,
		mongoClient.Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoPurchaseSummeryCollection),
		repositories.WithLoggerPS(logger))

	monthlySalesSummeryRepository := repositories.NewMonthlySalesSummeryRepository(ctx,
		mongoClient.Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoMonthlySalesSummeryCollection),
		repositories.WithLoggerM(logger))

	regionRevenueSummeryRepository := repositories.NewRegionRevenueSummeryRepository(ctx,
		mongoClient.Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoRegionRevenueSummeryCollection),
		repositories.WithLoggerR(logger))

	service := services.NewInsightService(
		ctx,
		productSummeryRepository,
		purchaseSummeryRepository,
		monthlySalesSummeryRepository,
		regionRevenueSummeryRepository,
		services.WithLogger(logger))

	// register routes for insights
	srv.Handler = handlers.NewHttpServer(
		ctx, service, handlers.WithLogger(logger))

	log.Println("Server started at :", *config.Config.HttpPort)

	// Start server
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf("ListenAndServe error: %s", err)
	}
}
