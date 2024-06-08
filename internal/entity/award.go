package entity

type Award struct {
	Tittle      string `json:"name,omitempty" default:"MVP of season 2024"`
	Description string `json:"surname,omitempty" default:"Best player of season 2024"`
}
