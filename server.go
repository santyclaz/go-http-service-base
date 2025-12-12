package main

import (
	"log/slog"
	"net/http"

	"go-example/stores"
)

// NewServer creates a new HTTP server with the provided dependencies and middleware
func NewServer(
	logger *slog.Logger,
	config *Config,
	loggingMiddleware Middleware,
	sessionMiddleware Middleware,
	pokemonStore stores.PokemonStore,
) http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		mux,
		logger,
		config,
		sessionMiddleware,
		pokemonStore,
	)

	var handler http.Handler = mux
	// can add top-level middleware here like logging, metrics, etc.
	handler = loggingMiddleware(handler)

	return handler
}
