package main

import (
	"log/slog"
	"net/http"
)

func newSessionMiddleware(logger *slog.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("session middleware")
			h.ServeHTTP(w, r)
		})
	}
}

func newLoggingMiddleware(logger *slog.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("logging middleware")
			h.ServeHTTP(w, r)
		})
	}
}
