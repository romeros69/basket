package usecase

import (
	"context"

	"github.com/romeros69/basket/internal/entity"
)

type (
	// Player - use case
	Player interface {
		CreatePlayer(ctx context.Context, player *entity.Player) (string, error)
		UpdatePlayer(ctx context.Context, playerID string, player *entity.Player) (*entity.Player, error)
		GetPlayer(ctx context.Context, playerID string) (*entity.Player, error)
		DeletePlayer(ctx context.Context, playerID string) error
		GetPlayerList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Player, error)
	}

	// PlayerRp - mongodb
	PlayerRp interface {
		CreatePlayer(ctx context.Context, player *entity.Player) (string, error)
		UpdatePlayer(ctx context.Context, playerID string, player *entity.Player) (*entity.Player, error)
		GetPlayer(ctx context.Context, playerID string) (*entity.Player, error)
		DeletePlayer(ctx context.Context, playerID string) error
		GetPlayerList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Player, error)
	}
)
