package domain_model

import "time"

type Race struct {
	Id          int       `json:"id"`
	Year        int       `json:"year"`
	StartDate   time.Time `json:"date_start"`
	CountryName string    `json:"country_name"`
	CircuitName string    `json:"circuit_name"`
}
