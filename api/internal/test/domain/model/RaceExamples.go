package test_domain_model

import (
	"time"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

const startDateFormat = "2006-01-02T15:04:05-07:00"

var races = make(map[string]domain_model.Race)

func InitRaces() {

	if len(races) > 0 {
		return
	}

	races["Sakhir"] = domain_model.Race{
		Id:   9472,
		Year: 2024,
		StartDate: func() time.Time {
			startTime, _ := time.Parse(startDateFormat, "2024-03-02T15:00:00Z")
			return startTime
		}(),
		CountryName: "Bahrain",
		CircuitName: "Sakhir",
	}

	races["Melbourne"] = domain_model.Race{
		Id:   9488,
		Year: 2024,
		StartDate: func() time.Time {
			startTime, _ := time.Parse(startDateFormat, "2024-03-24T04:00:00Z")
			return startTime
		}(),
		CountryName: "Australia",
		CircuitName: "Melbourne",
	}

	races["Suzuka"] = domain_model.Race{
		Id:   9496,
		Year: 2024,
		StartDate: func() time.Time {
			startTime, _ := time.Parse(startDateFormat, "2024-04-07T05:00:00Z")
			return startTime
		}(),
		CountryName: "Japan",
		CircuitName: "Suzuka",
	}

	races["Imola"] = domain_model.Race{
		Id:   9515,
		Year: 2024,
		StartDate: func() time.Time {
			startTime, _ := time.Parse(startDateFormat, "22024-05-19T13:00:00Z")
			return startTime
		}(),
		CountryName: "Italy",
		CircuitName: "Imola",
	}
}

func GetAllRaceExamples() map[string]domain_model.Race {
	InitRaces()

	return races
}

func Suzuka() []domain_model.Race {
	InitRaces()

	var response = make([]domain_model.Race, 0, 1)
	response = append(response, races["Suzuka"])

	return response
}
