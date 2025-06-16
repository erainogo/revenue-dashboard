# 📊 Revenue Dashboard

> High-performance dashboard to analyze and visualize key business metrics from a provided dataset.

## 🏗️ System Design

<img width="1358" alt="Screenshot 2025-06-16 at 07 10 34" src="https://github.com/user-attachments/assets/132e27d3-3857-4922-83de-e07e685c931a" />

## 🏛️ Architecture

### 🔧 Components
- **🖥️ HTTP Server**: Handles API requests and serves dashboard endpoints
- **🔍 Insight Service**: Fetches revenue analytics queries
- **📥 Ingestion Service/Aggregators/Repositories**: Batch processes CSV data and aggregates metrics
- **🗄️ MongoDB**: Document database with optimized collections for analytics

## 📚 Database Collections

| Collection | Description |
|------------|-------------|
| `transactions` | Raw transaction data |
| `region_revenue_summary` | Aggregated regional metrics |
| `country_product_summary` | Country-wise product performance |
| `monthly_sales_summary` | Monthly sales aggregations |

## ⚙️ Setup

### 📦 Installation
```bash
git clone <repository-url>
cd revenue-dashboard
```

### 🔧 Configuration
You can modify the configuration in the `config.go` located in `/app/internal/config`

## 🚀 API Endpoints

### Start API Server
Navigate to `cmd/server`
```bash
go run main.go
```
> 💡 You can also build the binary and run this. This will start the web server at port 8090

### 🔗 Available Endpoints
| Endpoint | Description |
|----------|-------------|
| `GET /getfrequentlypurchasedproducts` | Most popular products |
| `GET /getcountrylevelrevenue` | Country-wise product performance |
| `GET /getmonthlysalessummary` | Monthly sales aggregation |
| `GET /getregionrevenyesummary` | Regional revenue breakdown |

## 📊 Data Ingestion

Navigate to `cmd/ingest`

Run:
```bash
go run main.go <path_to_input.csv>
```

> 💡 You can also build the binary and run this. This will read the file from the disk and aggregates the data to mongo

The system processes CSV files through a worker pool of go routines and pre-aggregates data into summary collections for optimized read performance.

## 🧪 Run Test Cases

```bash
make test
```

### For Better View
```bash
go test -coverprofile=test_coverage/coverage.out ./internal/app/repositories && go tool cover -html=test_coverage/coverage.out -o test_coverage/repositories_coverage.html
```

## 🔍 Run Lint

```bash
make lint
```

