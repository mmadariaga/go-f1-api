package domain_service

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

type GetPodiumByRaceDependencies interface {
	GetDriversByRaceDependencies
	FetchPodiumByRace(raceId int, drivers []domain_model.Driver) ([3]domain_model.Podium, error)
}

func GetPodiumByRace(
	races []domain_model.Race,
	dependencies GetPodiumByRaceDependencies,
) (map[int][3]domain_model.Podium, error) {

	podiumsByRace := make(map[int][3]domain_model.Podium)

	drivers, error := GetDriversByRace(races, dependencies)
	if error != nil {
		return nil, error
	}

	for _, race := range races {
		result, error := dependencies.FetchPodiumByRace(race.Id, drivers[race.Id])
		if error != nil {
			return nil, error
		}

		podiumsByRace[race.Id] = result
	}

	return podiumsByRace, nil
}
