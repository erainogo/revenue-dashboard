# Revenue-dashboard

high-performance dashboard to analyze and visualize key business metrics from a provided dataset.

# System-design

<img width="1358" alt="Screenshot 2025-06-16 at 07 10 34" src="https://github.com/user-attachments/assets/132e27d3-3857-4922-83de-e07e685c931a" />

## Architecture

### Components
- **HTTP Server**: Handles API requests and serves dashboard endpoints
- **Insight Service**: Fetches revenue analytics queries
- **Ingestion Service/ Aggregators/ Repositoris**: Batch processes CSV data and aggregates metrics
- **MongoDB**: Document database with optimized collections for analytics

### API Endpoints
Start API server

Navigate to cmd/server

go run main.go

- You can also build the binary and run this

Navigate to cmd/server

- `GET /getfrequentlypurchasedproducts` - Most popular products
- `GET /getcountrylevelrevenue` - Country-wise product performance
- `GET /getmonthlysalessummary` - Monthly sales aggregation
- `GET /getregionrevenyesummary` - Regional revenue breakdown

## Database Collections
1. `transactions` - Raw transaction data
2. `region_revenue_summary` - Aggregated regional metrics
3. `country_product_summary` - Country-wise product performance
4. `monthly_sales_summary` - Monthly sales aggregations

## Setup

### Installation
```bash
git clone <repository-url>
cd revenue-dashboard
```

### Configuration
you can modify the configuration in the config.go located in /app/internal/config

## Data Ingestion
Navigate to cmd/ingest

Run - ```bash go run main.go <path_to_input.csv>  ```

- You can also build the binary and run this

The system processes CSV files through a worker pool of go routines and pre-aggregates data into summary collections for optimized read performance.

### Run test cases

```bash make test ``` 

For better view

``` bash go test -coverprofile=test_coverage/coverage.out ./internal/app/repositories && go tool cover -html=test_coverage/coverage.out -o test_coverage/repositories_coverage.html ```


