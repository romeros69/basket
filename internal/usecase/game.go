package usecase

import (
	"context"
	"github.com/romeros69/basket/internal/entity"
)

type GameUC struct {
	gameRp GameRp
}

func NewGameUC(gameRp GameRp) *GameUC {
	return &GameUC{
		gameRp: gameRp,
	}
}

var _ Game = (*GameUC)(nil)

func (g *GameUC) CreateGame(ctx context.Context, game *entity.Game) (string, error) {
	return g.gameRp.CreateGame(ctx, game)
}

func (g *GameUC) UpdateGame(ctx context.Context, gameID string, game *entity.Game) (*entity.Game, error) {
	return g.gameRp.UpdateGame(ctx, gameID, game)
}

func (g *GameUC) GetGame(ctx context.Context, gameID string) (*entity.Game, error) {
	return g.gameRp.GetGame(ctx, gameID)
}

func (g *GameUC) DeleteGame(ctx context.Context, gameID string) error {
	return g.gameRp.DeleteGame(ctx, gameID)
}

func (g *GameUC) GetGameList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Game, error) {
	return g.gameRp.GetGameList(ctx, pageSize, pageNumber)
}
