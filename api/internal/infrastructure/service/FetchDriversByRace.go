package infrastructure_service

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	helper "github.com/mmadariaga/go-api/internal/infrastructure/service/helper"
)

type Driver = domain_model.Driver

func FetchDriversByRace(raceId int) ([]Driver, error) {

	targetUrl := "https://api.openf1.org/v1/drivers?session_key=" + strconv.Itoa(raceId)
	body, err := helper.HttpGet(targetUrl, &helper.HttpGetExtraArgs{UseCache: true, Retry: true})
	if err != nil {
		return nil, err
	}

	// JSON to struct
	var drivers []struct {
		Number   int    `json:"driver_number"`
		FullName string `json:"full_name"`
		Country  string `json:"country_code"`
		Avatar   string `json:"headshot_url"`
		TeamName string `json:"team_name"`
	}

	if err := json.Unmarshal(body, &drivers); err != nil {
		errorMsg := "FetchDriversByRace JSON deserialization error: " + err.Error() + " in " + string(body)
		log.Error().Msg(errorMsg)
		return nil, errors.New(errorMsg)
	}

	// Session to domain model
	response := make([]Driver, 0, len(drivers))
	for _, value := range drivers {

		driver := Driver{
			Number:   value.Number,
			FullName: value.FullName,
			Country:  value.Country,
			Avatar:   value.Avatar,
			TeamName: value.TeamName,
		}
		response = append(response, driver)
	}

	return response, nil
}
