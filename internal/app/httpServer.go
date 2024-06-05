package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ShevelevEvgeniy/app/config"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	server *http.Server
}

func NewServer(cfg *config.Config, router chi.Router) *HttpServer {
	return &HttpServer{
		server: &http.Server{
			Addr:         ":" + cfg.HTTPServer.Port,
			Handler:      router,
			ReadTimeout:  cfg.HTTPServer.Timeout,
			WriteTimeout: cfg.HTTPServer.Timeout,
			IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		},
	}
}

func (hs *HttpServer) Run(log *slog.Logger, cfg *config.Config) error {
	log.Info("start server", slog.String("port", cfg.HTTPServer.Port))

	go func() {
		if err := hs.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error occurred on server shutting down: ", err)
			return
		}
	}()

	log.Info("server started")

	return nil
}

func (hs *HttpServer) handleSignals(ctx context.Context, log *slog.Logger, stopTimeout time.Duration) {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan

	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(ctx, stopTimeout)
	defer cancel()

	if err := hs.server.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
		return
	}

	log.Info("server stopped")
}
