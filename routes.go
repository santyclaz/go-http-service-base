package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	config *Config,
) {
	sessionMiddleware := newSessionMiddleware(logger)

	mux.Handle("GET /hello", sessionMiddleware(handleGetHello(logger)))
	mux.Handle("POST /hello", sessionMiddleware(handlePostHello(logger)))
	mux.Handle("/health", handleHealth(logger))
	mux.Handle("/", http.NotFoundHandler())
}

type Response struct {
	Message string `json:"message"`
}

func handleGetHello(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := Response{Message: "Hello, World!"}
		encode(w, r, http.StatusOK, m)
	})
}

func handlePostHello(logger *slog.Logger) http.Handler {

	type request struct {
		Name string `json:"name"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoded, problems, err := decode[request](r, nil)
		if err != nil {
			logger.Info(fmt.Sprintf("%v", problems))

			m := Response{Message: err.Error()}
			encode(w, r, http.StatusBadRequest, m)
			return
		}

		name := "World"
		if decoded.Name != "" {
			name = decoded.Name
		}
		m := Response{Message: fmt.Sprintf("Hello, %s!", name)}
		encode(w, r, http.StatusOK, m)
	})
}

func handleHealth(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := Response{Message: "ok"}
		encode(w, r, http.StatusOK, m)
	})
}
