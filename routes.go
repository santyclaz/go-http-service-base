package main

import (
	"log/slog"
	"net/http"

	"go-example/routes"
	"go-example/stores"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	config *Config,
	sessionMiddleware Middleware,
	pokemonStore stores.PokemonStore,
) {
	mux.Handle("POST /pokemon", sessionMiddleware(routes.HandlePostPokemon(logger, pokemonStore)))
	mux.Handle("GET /pokemon", sessionMiddleware(routes.HandleGetPokemon(logger, pokemonStore)))

	mux.Handle("/health", routes.HandleHealth(logger))

	mux.Handle("/", http.NotFoundHandler())
}
