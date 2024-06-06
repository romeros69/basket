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

func (p *PlayerUC) UpdatePlayer(ctx context.Context, playerID string, player *entity.Player) (*entity.Player, error) {
	return p.playerRp.UpdatePlayer(ctx, playerID, player)
}

func (p *PlayerUC) GetPlayer(ctx context.Context, playerID string) (*entity.Player, error) {
	return p.playerRp.GetPlayer(ctx, playerID)
}

func (p *PlayerUC) DeletePlayer(ctx context.Context, playerID string) error {
	return p.playerRp.DeletePlayer(ctx, playerID)
}

func (p *PlayerUC) GetPlayerList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Player, error) {
	return p.playerRp.GetPlayerList(ctx, pageSize, pageNumber)
}
