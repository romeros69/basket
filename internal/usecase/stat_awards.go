package usecase

import (
	"context"

	"github.com/romeros69/basket/internal/entity"
)

type StatAwardsUC struct {
	statAwardsRp StatAwardsRp
}

func NewStatAwardsUC(statAwardsRp StatAwardsRp) *StatAwardsUC {
	return &StatAwardsUC{
		statAwardsRp: statAwardsRp,
	}
}

var _ StatAwards = (*StatAwardsUC)(nil)

func (sa *StatAwardsUC) CreateRecord(ctx context.Context, rewardStat entity.RewardStat) error {
	return sa.statAwardsRp.CreateRecord(ctx, rewardStat)
}
func (sa *StatAwardsUC) ViewPlayersAndRewardsInTournament(ctx context.Context, tournamentId string) ([]entity.RewardStat, error) {
	return sa.statAwardsRp.ViewPlayersAndRewardsInTournament(ctx, tournamentId)
}
func (sa *StatAwardsUC) ViewPlayersAndRewardsInMatch(ctx context.Context, matchId string) ([]entity.RewardStat, error) {
	return sa.statAwardsRp.ViewPlayersAndRewardsInMatch(ctx, matchId)
}
func (sa *StatAwardsUC) ViewRewardsForPlayer(ctx context.Context, playerId string) ([]entity.RewardStat, error) {
	return sa.statAwardsRp.ViewRewardsForPlayer(ctx, playerId)
}
func (sa *StatAwardsUC) ViewWhoGotSpecificReward(ctx context.Context, rewardId string) ([]entity.RewardStat, error) {
	return sa.statAwardsRp.ViewWhoGotSpecificReward(ctx, rewardId)
}
