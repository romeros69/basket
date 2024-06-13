package entity

type League struct {
	Name   string `json:"name,omitempty" default:"NBA"`
	Season string `json:"season,omitempty" default:"2023/2024"`
}
