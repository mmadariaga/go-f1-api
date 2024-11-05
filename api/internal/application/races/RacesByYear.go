package application_races

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	domain_service "github.com/mmadariaga/go-api/internal/domain/service"
)

type RacesByYearResponse struct {
	domain_model.Race
	Podium [3]domain_model.Podium `json:"podium"`
}

type GetRacesByYearDependencies interface {
	domain_service.GetPodiumByRaceDependencies
	FetchRacesByYear(year int) ([]domain_model.Race, error)
}

type RacesByYear struct {
	dependencies GetRacesByYearDependencies
}

func NewRacesByYear(dependencies GetRacesByYearDependencies) *RacesByYear {
	return &RacesByYear{
		dependencies: dependencies,
	}
}

func (r *RacesByYear) Get(
	year int,
) ([]RacesByYearResponse, error) {

	races, error := r.dependencies.FetchRacesByYear(year)
	if error != nil {
		return nil, error
	}

	podiumByRace := domain_service.NewPodiumByRace(r.dependencies)

	podiumsByRace, error := podiumByRace.Get(races)
	if error != nil {
		return nil, error
	}

	response := make([]RacesByYearResponse, 0, len(races))
	for _, race := range races {

		podium := podiumsByRace[race.Id]

		response = append(
			response,
			RacesByYearResponse{
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
