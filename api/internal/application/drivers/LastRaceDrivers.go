package application_drivers

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

type GetLastRaceDriversDependencies interface {
	FetchLastRace() (domain_model.Race, error)
	FetchDriversByRace(raceId int) ([]domain_model.Driver, error)
}

type LastRaceDrivers struct {
	dependencies GetLastRaceDriversDependencies
}

func NewLastRaceDrivers(dependencies GetLastRaceDriversDependencies) *LastRaceDrivers {
	return &LastRaceDrivers{
		dependencies: dependencies,
	}
}

func (r *LastRaceDrivers) Get() ([]domain_model.Driver, error) {

	race, error := r.dependencies.FetchLastRace()
	if error != nil {
		return nil, error
	}

	return r.dependencies.FetchDriversByRace(race.Id)
}
