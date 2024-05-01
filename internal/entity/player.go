package entity

type Player struct {
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	Age         int    `json:"age,omitempty"`
	Height      int    `json:"height,omitempty"`
	Weight      int    `json:"weight,omitempty"`
	Team        string `json:"team,omitempty"`
	Role        string `json:"role,omitempty"`
	Citizenship string `json:"citizenship,omitempty"`
}
