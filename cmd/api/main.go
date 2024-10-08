package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"fmt"

	infrastructure_router "github.com/mmadariaga/go-api/internal/infrastructure/router"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Starting GO API service")

	router := infrastructure_router.RouterFactory()

	port := ":8080"
	log.Print("listening on localhost" + port)
	fmt.Println(`
	▗▄▄▖ ▗▄▖      ▗▄▄▖ ▗▄▖      ▗▄▄▖ ▗▄▖ 
	▐▌   ▐▌ ▐▌    ▐▌   ▐▌ ▐▌    ▐▌   ▐▌ ▐▌
	▐▌▝▜▌▐▌ ▐▌    ▐▌▝▜▌▐▌ ▐▌    ▐▌▝▜▌▐▌ ▐▌
	▝▚▄▞▘▝▚▄▞▘    ▝▚▄▞▘▝▚▄▞▘    ▝▚▄▞▘▝▚▄▞▘
	`)
	serverError := http.ListenAndServe(port, router)
	if serverError != nil {
		log.Error().Msg(serverError.Error())
	}
}
