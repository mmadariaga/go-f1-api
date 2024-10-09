package infrastructure_service

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	helper "github.com/mmadariaga/go-api/internal/infrastructure/service/helper"
)

type Race = domain_model.Race

func FetchRacesByYear(year int) ([]Race, error) {

	targetUrl := "https://api.openf1.org/v1/sessions?year=" + strconv.Itoa(year) + "&session_type=Race"
	body, err := helper.HttpGet(targetUrl, &helper.HttpGetExtraArgs{UseCache: true, Retry: true})
	if err != nil {
		return nil, err
	}

	// JSON to struct
	var sessions []struct {
		Id          int    `json:"session_key"`
		DateStart   string `json:"date_start"`
		CountryName string `json:"country_name"`
		CircuitName string `json:"circuit_short_name"`
	}

	if err := json.Unmarshal(body, &sessions); err != nil {
		errorMsg := "FetchRacesByYear JSON deserialization error: " + err.Error() + " in " + string(body)
		log.Error().Msg(errorMsg)
		return nil, errors.New(errorMsg)
	}

	// Session to domain model
	response := make([]Race, 0, len(sessions))
	for _, value := range sessions {

		date, err := time.Parse("2006-01-02T15:04:05-07:00", value.DateStart)
		if err != nil {
			errorMsg := "Unable to parse datetime: " + value.DateStart
			log.Error().Msg(errorMsg)
			return nil, errors.New(errorMsg)
		}

		race := Race{
			Id:          value.Id,
			Year:        date.Year(),
			StartDate:   date,
			CountryName: value.CountryName,
			CircuitName: value.CircuitName,
		}
		response = append(response, race)
	}

	return response, nil
}
