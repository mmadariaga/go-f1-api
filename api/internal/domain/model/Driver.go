package domain_model

type Driver struct {
	Number   int    `json:"driver_number"`
	FullName string `json:"full_name"`
	Country  string `json:"country_code"`
	Avatar   string `json:"avatar"`
	TeamName string `json:"team_name"`
}
