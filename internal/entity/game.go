package entity

type Game struct {
	FirstTeam  string `json:"first_team,omitempty" default:"LA Lakers"`
	SecondTeam string `json:"second_team,omitempty" default:"Chicago Bulls"`
	Date       string `json:"date,omitempty" default:"12.03.24"`
	Type       string `json:"type,omitempty" default:"final"`
	League     string `json:"league,omitempty" default:"NBA"`
}
