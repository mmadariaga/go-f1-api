package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/mmadariaga/go-api/internal/controller"
)

func setRoutes(router *chi.Mux) *chi.Mux {

	router.Get("/ping", controller.Ping)

	return router
}
