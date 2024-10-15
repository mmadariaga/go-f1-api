package infrastructure_router

import (
	"github.com/go-chi/chi/v5"
	infrastructure_controller "github.com/mmadariaga/go-api/internal/infrastructure/controller"
)

func setRoutes(router *chi.Mux) *chi.Mux {

	router.Get("/ping", infrastructure_controller.Ping)
	router.Get("/races", infrastructure_controller.Races)
	router.Get("/races/{year}", infrastructure_controller.Races)

	return router
}
