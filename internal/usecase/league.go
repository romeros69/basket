package usecase

import (
	"context"
	"github.com/romeros69/basket/internal/entity"
)

type LeagueUC struct {
	leagueRp LeagueRp
}

func NewLeagueUC(leagueRp LeagueRp) *LeagueUC {
	return &LeagueUC{
		leagueRp: leagueRp,
	}
}

var _ League = (*LeagueUC)(nil)

func (l *LeagueUC) CreateLeague(ctx context.Context, league *entity.League) (string, error) {
	return l.leagueRp.CreateLeague(ctx, league)
}

func (l *LeagueUC) UpdateLeague(ctx context.Context, leagueID string, league *entity.League) (*entity.League, error) {
	return l.leagueRp.UpdateLeague(ctx, leagueID, league)
}

func (l *LeagueUC) GetLeague(ctx context.Context, leagueID string) (*entity.League, error) {
	return l.leagueRp.GetLeague(ctx, leagueID)
}

func (l *LeagueUC) DeleteLeague(ctx context.Context, leagueID string) error {
	return l.leagueRp.DeleteLeague(ctx, leagueID)
}

func (l *LeagueUC) GetLeagueList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.League, error) {
	return l.leagueRp.GetLeagueList(ctx, pageSize, pageNumber)
}
