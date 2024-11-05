package infrastructure_controller

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	application_drivers "github.com/mmadariaga/go-api/internal/application/drivers"
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	infrastructure_service "github.com/mmadariaga/go-api/internal/infrastructure/service"
)

func Drivers(w http.ResponseWriter, r *http.Request) {

	lastRaceDrivers := application_drivers.NewLastRaceDrivers(
		&driversInfraDependencies{},
	)

	drivers, error := lastRaceDrivers.Get()
	if error != nil {
		log.Panic().Msg(error.Error())
	}

	driversJson, error := json.Marshal(drivers)
	if error != nil {
		log.Panic().Msg(error.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(driversJson)
}

type driversInfraDependencies struct{}

func (d *driversInfraDependencies) FetchLastRace() (domain_model.Race, error) {
	return infrastructure_service.FetchLastRace()
}
func (d *driversInfraDependencies) FetchDriversByRace(raceId int) ([]domain_model.Driver, error) {
	return infrastructure_service.FetchDriversByRace(raceId)
}
