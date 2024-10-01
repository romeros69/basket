package usecase

import (
	"context"

	"github.com/romeros69/basket/internal/entity"
)

type StatPlayerUC struct {
	statPlayerRp StatPlayerRp
}

func NewStatPlayerUC(statPlayerRp StatPlayerRp) *StatPlayerUC {
	return &StatPlayerUC{
		statPlayerRp: statPlayerRp,
	}
}

var _ StatPlayer = (*StatPlayerUC)(nil)


func (sp *StatPlayerUC) InsertPlayerStat(ctx context.Context, stat entity.PlayerStat) error {
	return sp.statPlayerRp.InsertPlayerStat(ctx, stat)
}

func (sp *StatPlayerUC) GetPlayerStatsByIDAndMatch(ctx context.Context, playerID, matchID string) ([]entity.PlayerStat, error) {
	return sp.statPlayerRp.GetPlayerStatsByIDAndMatch(ctx, playerID, matchID)
}

func (sp *StatPlayerUC) GetPlayersWithAvgGoalsGreaterThanByMatch(ctx context.Context, minAvgGoals float64, matchID string) ([]entity.PlayerStat, error) {
	return sp.statPlayerRp.GetPlayersWithAvgGoalsGreaterThanByMatch(ctx, minAvgGoals, matchID)
}

func (sp *StatPlayerUC) GetPlayersWithTotalAvgStatsGreaterThanByMatch(ctx context.Context, minTotalAvg float64, matchID string) ([]entity.PlayerStat, error) {
	return sp.statPlayerRp.GetPlayersWithTotalAvgStatsGreaterThanByMatch(ctx, minTotalAvg, matchID)
}
