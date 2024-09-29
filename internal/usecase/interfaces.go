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

	// Game - use case
	Game interface {
		CreateGame(ctx context.Context, game *entity.Game) (string, error)
		UpdateGame(ctx context.Context, gameID string, game *entity.Game) (*entity.Game, error)
		GetGame(ctx context.Context, gameID string) (*entity.Game, error)
		DeleteGame(ctx context.Context, gameID string) error
		GetGameList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Game, error)
	}

	// GameRp - mongodb
	GameRp interface {
		CreateGame(ctx context.Context, game *entity.Game) (string, error)
		UpdateGame(ctx context.Context, gameID string, game *entity.Game) (*entity.Game, error)
		GetGame(ctx context.Context, gameID string) (*entity.Game, error)
		DeleteGame(ctx context.Context, gameID string) error
		GetGameList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Game, error)
	}

	// League - use case
	League interface {
		CreateLeague(ctx context.Context, league *entity.League) (string, error)
		UpdateLeague(ctx context.Context, leagueID string, league *entity.League) (*entity.League, error)
		GetLeague(ctx context.Context, leagueID string) (*entity.League, error)
		DeleteLeague(ctx context.Context, leagueID string) error
		GetLeagueList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.League, error)
	}

	// LeagueRp - mongodb
	LeagueRp interface {
		CreateLeague(ctx context.Context, league *entity.League) (string, error)
		UpdateLeague(ctx context.Context, leagueID string, league *entity.League) (*entity.League, error)
		GetLeague(ctx context.Context, leagueID string) (*entity.League, error)
		DeleteLeague(ctx context.Context, leagueID string) error
		GetLeagueList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.League, error)
	}

	// StatAwards - use case
	StatAwards interface {
		CreateRecord(context.Context, entity.RewardStat) error
		ViewPlayersAndRewardsInTournament(context.Context, string) ([]entity.RewardStat, error)
		ViewPlayersAndRewardsInMatch(context.Context, string) ([]entity.RewardStat, error)
		ViewRewardsForPlayer(context.Context, string) ([]entity.RewardStat, error)
		ViewWhoGotSpecificReward(context.Context, string) ([]entity.RewardStat, error)
	}

	// StatAwardsRp - neo4j
	StatAwardsRp interface {
		CreateRecord(context.Context, entity.RewardStat) error
		ViewPlayersAndRewardsInTournament(context.Context, string) ([]entity.RewardStat, error)
		ViewPlayersAndRewardsInMatch(context.Context, string) ([]entity.RewardStat, error)
		ViewRewardsForPlayer(context.Context, string) ([]entity.RewardStat, error)
		ViewWhoGotSpecificReward(context.Context, string) ([]entity.RewardStat, error)
	}
)
