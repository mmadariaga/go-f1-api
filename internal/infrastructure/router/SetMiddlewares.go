package infrastructure_router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	infrastructure_middleware "github.com/mmadariaga/go-api/internal/infrastructure/middleware"
)

func setMiddlewares(router *chi.Mux) *chi.Mux {

	// third party middlewares
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(middleware.SupressNotFound(router))

	// my middlewares
	router.Use(infrastructure_middleware.BasicAuthMiddleware)

	return router
}
