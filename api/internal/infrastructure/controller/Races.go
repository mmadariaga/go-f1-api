package infrastructure_controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/rs/zerolog/log"

	application_races "github.com/mmadariaga/go-api/internal/application/races"
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	infrastructure_service "github.com/mmadariaga/go-api/internal/infrastructure/service"
)

func Races(w http.ResponseWriter, r *http.Request) {

	racesByYear := application_races.NewRacesByYear(
		&infraDependencies{},
	)

	races, error := racesByYear.Get(
		getYearFromUrl(r),
	)
	if error != nil {
		log.Panic().Msg(error.Error())
	}

	racesJson, error := json.Marshal(races)
	if error != nil {
		log.Panic().Msg(error.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(racesJson)
}

func getYearFromUrl(r *http.Request) int {

	yearStr := chi.URLParam(r, "year")
	if yearStr == "" {
		yearStr = "2024"
	}

	year, error := strconv.Atoi(yearStr)
	if error != nil {
		log.Panic().Msg(error.Error())
	}

	return year
}

type infraDependencies struct{}

func (d *infraDependencies) FetchRacesByYear(year int) ([]domain_model.Race, error) {
	return infrastructure_service.FetchRacesByYear(year)
}
func (d *infraDependencies) FetchPodiumByRace(raceId int, drivers []domain_model.Driver) ([3]domain_model.Podium, error) {
	return infrastructure_service.FetchPodiumByRace(raceId, drivers)
}
func (d *infraDependencies) FetchDriversByRace(raceId int) ([]domain_model.Driver, error) {
	return infrastructure_service.FetchDriversByRace(raceId)
}
