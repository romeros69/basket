package entity

type PlayerStat struct {
	PlayerID      string  `json:"playerId,omitempty"`
	MatchID       string  `json:"matchId,omitempty"`
	Goals         int     `json:"goals,omitempty"`
	Assists       int     `json:"assists,omitempty"`
	Interceptions int     `json:"interceptions,omitempty"`
	Rebounds      int     `json:"rebounds,omitempty"`
	AVGGoals      float64 `json:"avgGoals,omitempty"`
	TotalAVGStats float64 `json:"totalAvgStats,omitempty"`
}
