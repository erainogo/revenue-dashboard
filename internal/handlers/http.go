package handlers

import (
	"context"
	
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/erainogo/revenue-dashboard/internal/core/adapters"
)

type HttpServer struct {
	logger  *zap.SugaredLogger
	ctx     context.Context
	mux     *http.ServeMux
	service adapters.InsightService
}

type HttpServerOptions func(*HttpServer)

func WithLogger(logger *zap.SugaredLogger) HttpServerOptions {
	return func(s *HttpServer) {
		s.logger = logger
	}
}

func NewHttpServer(
	ctx context.Context,
	service adapters.InsightService,
	opts ...HttpServerOptions,
) *HttpServer {
	server := &HttpServer{
		ctx: ctx,
		service: service,
		mux: http.NewServeMux(),
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (h *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.serve(w, r)
}

func (h *HttpServer) serve(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	h.mux.ServeHTTP(w, r)

	duration := time.Since(startTime)

	h.logger.Infof("Completed request: method=%s, url=%s, duration=%v",
		r.Method, r.URL.String(), duration)
}
