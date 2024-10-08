package domain_model

type Podium struct {
	Driver
	Position int `json:"position"`
}
