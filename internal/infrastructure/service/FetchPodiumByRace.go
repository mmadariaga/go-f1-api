package infrastructure_service

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	helper "github.com/mmadariaga/go-api/internal/infrastructure/service/helper"
)

type Podium = domain_model.Podium

type Position struct {
	DriverNumber int `json:"driver_number"`
	Position     int `json:"position"`
}

func FetchPodiumByRace(raceId int, drivers []domain_model.Driver) ([3]Podium, error) {

	targetUrl := "https://api.openf1.org/v1/position?position<=3&session_key=" + strconv.Itoa(raceId)
	body, err := helper.HttpGet(targetUrl, true)
	if err != nil {
		return [3]Podium{}, err
	}

	// JSON to struct
	var positions []Position

	if err := json.Unmarshal(body, &positions); err != nil {
		errorMsg := "FetchPodiumByRace JSON deserialization error: " + err.Error() + " in " + string(body)
		log.Error().Msg(errorMsg)
		return [3]Podium{}, errors.New(errorMsg)
	}

	latestPositions := getLatestPositions(positions)

	return positionsToPodium(latestPositions, drivers), nil
}

func getLatestPositions(values []Position) [3]Position {

	response := [3]Position{}
	for i := len(values) - 1; i >= 0; i-- {

		position := values[i].Position - 1
		if response[position] != (Position{}) {
			continue
		}

		response[position] = values[i]
	}

	return response
}

func positionsToPodium(values [3]Position, drivers []domain_model.Driver) [3]Podium {

	response := [3]Podium{}
	for i := len(values) - 1; i >= 0; i-- {

		position := values[i].Position - 1
		if response[position] != (Podium{}) {
			continue
		}

		driver, found := findDriverByDriverNum(
			values[i].DriverNumber,
			drivers,
		)

		if !found {
			continue
		}

		response[position] = Podium{
			Position: values[i].DriverNumber,
			Driver:   driver,
		}
	}

	return response
}

func findDriverByDriverNum(driverNumber int, drivers []domain_model.Driver) (domain_model.Driver, bool) {

	for _, value := range drivers {
		if value.Number == driverNumber {
			return value, true
		}
	}

	return domain_model.Driver{}, false
}
