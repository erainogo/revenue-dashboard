package config

import (
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

const (
	ReadTimeOut    int = 30
	WriteTimeOut   int = 30
	bootupWaitTime int = 5
)

type Configuration struct {
	Prefix                               *string //nolint:goimports,gofmt
	HttpPort                             *string
	BootUpWaitTime                       *int
	WriteTimeOut                         *int
	ReadTimeOut                          *int
	MongoDBEndpoint                      *string
	MongoDBDatabase                      *string
	FEURL                                *string
	MongoTransactionCollectionName       *string
	MongoDBMigrate                       *bool
	MongoCountryProductSummaryCollection *string
}

var (
	Config *Configuration

	prefix = flag.String(
		"prefix",
		"local",
		"Prefix")

	httpPort = flag.String(
		"http-port",
		"8090",
		"the port to serve on")

	bootupWaittime = flag.Int(
		"bootup-wait-time",
		bootupWaitTime,
		"Number of seconds to wait.")

	readTimeOut = flag.Int(
		"http-read-timeout",
		ReadTimeOut,
		"read time out")

	writeTimeOut = flag.Int(
		"http-write-timeout",
		WriteTimeOut,
		"write time out")

	mongoDBEndpoint = flag.String(
		"mongodb-endpoint",
		"mongodb://127.0.0.1:27017",
		"Mongodb Endpoint.")

	mongoDBDatabase = flag.String(
		"mongodb-database",
		"local-transactions-db",
		"Mongodb Database.")

	mongoDBMigrate = flag.Bool(
		"mongodb-migrate",
		false,
		"Mongodb migrate.")

	mongoTransactionCollection = flag.String(
		"transactions-collection",
		"transactions",
		"transactions Mongodb Collection")

	mongoCountryProductSummaryCollection = flag.String(
		"country-product-summary-collection",
		"country_product_summary",
		"country product summary Mongodb Collection")

	feUrl = flag.String(
		"fe-url",
		"http://localhost:5173", // change your front end url
		"fe url",
	)
)

func updateStringEnvVariable(defValue *string, key string) *string {
	val := os.Getenv(key)

	if val == "" {
		return defValue
	}

	return &val
}

func updateIntEnvVariable(defValue *int, key string) *int {
	sVal := os.Getenv(key)
	if sVal == "" {
		return defValue
	}

	iVal, err := strconv.Atoi(sVal)
	if err != nil {
		return defValue
	}

	return &iVal
}

func init() {
	flag.Parse()

	prefix = updateStringEnvVariable(prefix, "BRAND_NAME")
	httpPort = updateStringEnvVariable(httpPort, "HTTP_PORT")
	writeTimeOut = updateIntEnvVariable(writeTimeOut, "WRITE_TIMEOUT")
	readTimeOut = updateIntEnvVariable(readTimeOut, "READ_TIMEOUT")
	feUrl = updateStringEnvVariable(feUrl, "FEURL")

	mongoDBEndpoint = updateStringEnvVariable(mongoDBEndpoint, "MONGODB_ENDPOINT")
	mongoDBDatabase = updateStringEnvVariable(mongoDBDatabase, "MONGODB_DATABASE")

	Config = &Configuration{
		Prefix:                         prefix, //nolint:goimports,gofmt
		HttpPort:                       httpPort,
		WriteTimeOut:                   writeTimeOut,
		ReadTimeOut:                    readTimeOut,
		BootUpWaitTime:                 bootupWaittime,
		FEURL:                          feUrl,
		MongoTransactionCollectionName: mongoTransactionCollection,
		MongoCountryProductSummaryCollection: mongoCountryProductSummaryCollection,
		MongoDBMigrate:                 mongoDBMigrate,
		MongoDBEndpoint:                mongoDBEndpoint,
		MongoDBDatabase:                mongoDBDatabase,
	}
}
