package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	myMiddlewares "github.com/mmadariaga/go-api/internal/middleware"
)

func setMiddlewares(router *chi.Mux) *chi.Mux {

	router.Use(middleware.Logger)
	router.Use(middleware.StripSlashes)

	router.Use(myMiddlewares.BasicAuthMiddleware)

	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	return router
}
