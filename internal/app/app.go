package app

import (
	"context"
	"log/slog"

	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/go-chi/chi/v5"
)

type App struct {
	log         *slog.Logger
	DiContainer *DiContainer
	router      *chi.Mux
	HttpServer  *HttpServer
}

func NewApp(log *slog.Logger) *App {
	return &App{
		log: log,
	}
}

func (a *App) Run(ctx context.Context) error {
	a.DiContainer = NewDiContainer(a.log)

	a.router = a.initRouter(ctx, a.DiContainer)
	defer a.DiContainer.dbConn.Close()

	httpServer := NewServer(a.DiContainer.Config(ctx), a.router)

	if err := httpServer.Run(a.log, a.DiContainer.Config(ctx)); err != nil {
		a.log.Error("Failed to start http server: ", sl.Err(err))
		return err
	}

	httpServer.handleSignals(ctx, a.log, a.DiContainer.Config(ctx).HTTPServer.StopTimeout)

	return nil
}
