package app

import (
	"context"

	"github.com/ShevelevEvgeniy/app/config/routes"
	apiAuth "github.com/ShevelevEvgeniy/app/internal/http-server/api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) initRouter(ctx context.Context, DiContainer *DiContainer) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route(routes.ApiV1Group, func(router chi.Router) {
		router.Use(apiAuth.Auth(a.log, DiContainer.Config(ctx).Auth.ApiKey))

		router.Route(routes.LandPlotsGroup, func(router chi.Router) {
			router.Get(routes.GetCoordinates, DiContainer.LandPlotsHandler(ctx).GetCoordinates(ctx))
		})

		router.Route(routes.KptGroup, func(router chi.Router) {
			router.Post(routes.SaveKpt, DiContainer.SaveKptHandler(ctx).SaveKpt(ctx))
			router.Get(routes.GetDownloadLinkForKpt, DiContainer.GetDownloadLinkForKptHandler(ctx).GetDownloadLinkKpt(ctx))
		})
	})

	return router
}
