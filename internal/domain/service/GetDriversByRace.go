package domain_service

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

type GetDriversByRaceDependencies interface {
	FetchDriversByRace(raceId int) ([]domain_model.Driver, error)
}

func GetDriversByRace(races []domain_model.Race, dependencies GetDriversByRaceDependencies) (map[int][]domain_model.Driver, error) {

	driversByRace := make(map[int][]domain_model.Driver)

	for _, race := range races {
		result, error := dependencies.FetchDriversByRace(race.Id)
		if error != nil {
			return nil, error
		}

		driversByRace[race.Id] = result
	}

	return driversByRace, nil
}
