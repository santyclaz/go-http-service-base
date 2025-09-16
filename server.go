package main

import (
	"log/slog"
	"net/http"
)

// NewServer creates a new HTTP server with the provided dependencies and middleware
func NewServer(
	logger *slog.Logger,
	config *Config,
) http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		mux,
		logger,
		config,
	)

	// top-level middleware
	var handler http.Handler = mux
	handler = newLoggingMiddleware(logger)(handler)

	return handler
}
