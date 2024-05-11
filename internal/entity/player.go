package entity

type Player struct {
	Name        string `json:"name,omitempty" default:"Jimmi"`
	Surname     string `json:"surname,omitempty" default:"Butler"`
	MiddleName  string `json:"middle_name,omitempty"`
	Age         int    `json:"age,omitempty" default:"34"`
	Height      int    `json:"height,omitempty" default:"201"`
	Weight      int    `json:"weight,omitempty" default:"104"`
	Team        string `json:"team,omitempty" default:"Miami Heat"`
	Role        string `json:"role,omitempty" default:"heavy forward"`
	Citizenship string `json:"citizenship,omitempty" default:"USA"`
}
