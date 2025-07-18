package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/erainogo/revenue-dashboard/pkg/constants"
	"github.com/erainogo/revenue-dashboard/pkg/entities"
)

func (h *HttpServer) GetCountryLevelRevenueHandler(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		pageStr := query.Get("page")
		limitStr := query.Get("limit")

		page := constants.FirstPage
		limit := constants.PaginationCount

		if pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}

		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
				limit = l
			}
		}

		rg, err := h.service.GetCountryLevelRevenue(ctx, page, limit)

		if err != nil {
			h.logger.Errorw("failed to get revenue", "error", err)

			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)

			return
		}

		response := entities.CountryLevelRevenueResponse{
			Page:  page,
			Limit: limit,
			Data:  rg,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			h.logger.Errorw("failed to encode result", "error", err)

			http.Error(w, "Failed to encode response", http.StatusInternalServerError)

			return
		}
	}
}

func (h *HttpServer) GetFrequentlyPurchasedProductsHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rg, err := h.service.GetFrequentlyPurchasedProducts(ctx)

		if err != nil {
			h.logger.Errorw("failed to get top frequently purchased products", "error", err)

			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)

			return
		}

		response := entities.FrequentlyPurchasedProductsResponse{
			Data: rg,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			h.logger.Errorw("failed to encode result", "error", err)

			http.Error(w, "Failed to encode response", http.StatusInternalServerError)

			return
		}
	}
}

func (h *HttpServer) GetMonthlySalesSummeryHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rg, err := h.service.GetMonthlyRevenue(ctx)

		if err != nil {
			h.logger.Errorw("failed to get monthly sales summery", "error", err)

			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)

			return
		}

		response := entities.MonthlySalesSummaryResponse{
			Data: rg,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			h.logger.Errorw("failed to encode result", "error", err)

			http.Error(w, "Failed to encode response", http.StatusInternalServerError)

			return
		}
	}
}

func (h *HttpServer) GetRegionRevenueSummeryHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rg, err := h.service.GetRegionRevenue(ctx)

		if err != nil {
			h.logger.Errorw("failed to get region revenue summery", "error", err)

			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)

			return
		}

		response := entities.RegionRevenueSummaryResponse{
			Data: rg,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			h.logger.Errorw("failed to encode result", "error", err)

			http.Error(w, "Failed to encode response", http.StatusInternalServerError)

			return
		}
	}
}

func (h *HttpServer) HealthHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(`{"status":"ok"}`))
		if err != nil {
			return
		}
	}
}
