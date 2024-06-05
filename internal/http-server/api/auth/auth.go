package auth

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func Auth(log *slog.Logger, apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const op = "handler.ApiAuth.New"

			log = log.With(
				slog.String("op", op),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)

			key := r.Header.Get("X-API-KEY")
			if key == "" {
				log.Error("x-api-key is empty")
				http.Error(w, "enter x-api-key", http.StatusUnauthorized)
				return
			}

			if apiKey != key {
				log.Info("key is not valid")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
