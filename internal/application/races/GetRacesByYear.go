package application_races

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	domain_service "github.com/mmadariaga/go-api/internal/domain/service"
)

type Response struct {
	domain_model.Race
	Podium [3]domain_model.Podium `json:"podium"`
}

type GetRacesByYearDependencies interface {
	domain_service.GetPodiumByRaceDependencies
	FetchRacesByYear(year int) ([]domain_model.Race, error)
}

func GetRacesByYear(
	year int,
	dependencies GetRacesByYearDependencies,
) ([]Response, error) {

	races, error := dependencies.FetchRacesByYear(year)
	if error != nil {
		return nil, error
	}

	podiumsByRace, error := domain_service.GetPodiumByRace(
		races,
		dependencies,
	)
	if error != nil {
		return nil, error
	}

	response := make([]Response, 0, len(races))
	for _, race := range races {

		podium := podiumsByRace[race.Id]

		response = append(
			response,
			Response{
				Race: domain_model.Race{
					Id:          race.Id,
					Year:        race.Year,
					StartDate:   race.StartDate,
					CountryName: race.CountryName,
					CircuitName: race.CircuitName,
				},
				Podium: podium,
			},
		)
	}

	return response, nil
}
