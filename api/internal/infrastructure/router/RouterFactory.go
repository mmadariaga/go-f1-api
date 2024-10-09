package infrastructure_router

import (
	"github.com/go-chi/chi/v5"
)

func RouterFactory() *chi.Mux {

	router := chi.NewRouter()
	setMiddlewares(router)
	setRoutes(router)

	return router
}
