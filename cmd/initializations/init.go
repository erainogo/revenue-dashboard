package initializations

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SetUpLogger() *zap.SugaredLogger {
	appName := fmt.Sprintf("%s-revenue-dashboard", *config.Config.Prefix)

	zapLogger, _ := zap.NewProduction()

	return zapLogger.With(zap.String("app", appName)).Sugar()
}

func RunMigration(client *mongo.Client) error {
	driver, err := mongodb.WithInstance(client, &mongodb.Config{
		DatabaseName: *config.Config.MongoDBDatabase,
	})
	if err != nil {
		return err
	}

	exec, _ := os.Executable()

	migPath := filepath.Join(filepath.Dir(exec), "db", "migrations")

	mig, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:%s", migPath),
		"",
		driver,
	)

	if err != nil {
		return err
	}

	return mig.Up()
}

func CreateMongoClient(ctx context.Context, logger *zap.SugaredLogger) (*mongo.Client, error) {
	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI(*config.Config.MongoDBEndpoint),
	)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logger.Warn("failed to ping mongo: %v", err)

		errDisconnect := client.Disconnect(ctx)
		if errDisconnect != nil {
			logger.Errorf("error in disconnecting mongo %v", errDisconnect)
		}

		return nil, err
	}

	return client, nil
}