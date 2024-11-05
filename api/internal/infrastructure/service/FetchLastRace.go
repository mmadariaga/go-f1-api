package infrastructure_service

import (
	"errors"
	"time"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

const no_race_found_error = "NO RACE FOUND"

func FetchLastRace() (domain_model.Race, error) {

	currentYear := time.Now().Year()
	race, err := fetchLastRace(currentYear)

	if err != nil {

		if err.Error() != no_race_found_error {
			return domain_model.Race{}, err
		}

		race, err = fetchLastRace(currentYear - 1)
	}

	if err != nil {
		return domain_model.Race{}, err
	}

	return race, nil
}

func fetchLastRace(year int) (domain_model.Race, error) {

	races, err := FetchRacesByYear(year)

	if err != nil {
		return domain_model.Race{}, err
	}

	if len(races) == 0 {
		err := errors.New(no_race_found_error)
		return domain_model.Race{}, err
	}

	lastRace := races[len(races)-1]

	return lastRace, nil
}
