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

	// Award - use case
	Award interface {
		CreateAward(ctx context.Context, award *entity.Award) (string, error)
		UpdateAward(ctx context.Context, awardID string, award *entity.Award) (*entity.Award, error)
		GetAward(ctx context.Context, awardID string) (*entity.Award, error)
		DeleteAward(ctx context.Context, awardID string) error
		GetAwardList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Award, error)
	}

	// AwardRp - mongodb
	AwardRp interface {
		CreateAward(ctx context.Context, award *entity.Award) (string, error)
		UpdateAward(ctx context.Context, awardID string, award *entity.Award) (*entity.Award, error)
		GetAward(ctx context.Context, awardID string) (*entity.Award, error)
		DeleteAward(ctx context.Context, awardID string) error
		GetAwardList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Award, error)
	}
)
