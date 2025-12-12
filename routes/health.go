package routes

import (
	"log/slog"
	"net/http"
)

func HandleHealth(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := Response{Message: "ok"}
		encode(w, r, http.StatusOK, m)
	})
}
