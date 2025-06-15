package handlers

import (
	"context"
	"net/http"
)

func (h *HttpServer) registerRoutes(ctx context.Context) {
	routes := map[string]http.HandlerFunc{
		"/api/insights/getfrequentlypurchasedproducts": h.GetFrequentlyPurchasedProductsHandler(ctx),
		"/api/insights/getcountrylevelrevenue":         h.GetCountryLevelRevenueHandler(ctx),
		"/api/insights/getmonthlysalessummary":         h.GetMonthlySalesSummeryHandler(ctx),
		"/api/insights/getregionrevenyesummary":        h.GetRegionRevenueSummeryHandler(ctx),
		"/api/health":                                  h.HealthHandler(),
	}

	for path, handler := range routes {
		h.mux.HandleFunc(path, handler)
	}
}
