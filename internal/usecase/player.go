package usecase

import (
	"context"
	"github.com/romeros69/basket/internal/entity"
)

type PlayerUC struct {
	playerRp PlayerRp
}

func NewPlayerUC(playerRp PlayerRp) *PlayerUC {
	return &PlayerUC{
		playerRp: playerRp,
	}
}

var _ Player = (*PlayerUC)(nil)

func (p *PlayerUC) CreatePlayer(ctx context.Context, player *entity.Player) (string, error) {
	return p.playerRp.CreatePlayer(ctx, player)
}

func (p *PlayerUC) UpdatePlayer(ctx context.Context, player *entity.Player) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerUC) GetPlayer(ctx context.Context, playerID string) (*entity.Player, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerUC) DeletePlayer(ctx context.Context, playerID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerUC) GetPlayerList(ctx context.Context) ([]*entity.Player, error) {
	//TODO implement me
	panic("implement me")
}
