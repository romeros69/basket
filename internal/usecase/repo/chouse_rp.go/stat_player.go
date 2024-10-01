package chouse_rp

import (
	"context"
	"fmt"

	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/chouse"
)

type ChouseRepo struct {
	cHouseDB *chouse.Chouse
}

func NewChouseRepo(chouse *chouse.Chouse) *ChouseRepo {
	return &ChouseRepo{
		cHouseDB: chouse,
	}
}

var _ usecase.StatPlayerRp = (*ChouseRepo)(nil)

// Функция для вставки данных в таблицу player_stats
func (c *ChouseRepo) InsertPlayerStat(ctx context.Context, stat entity.PlayerStat) error {
	insertQuery := `
		INSERT INTO player_stats (player_id, match_id, goals, assists, interceptions, rebounds)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := c.cHouseDB.DB.ExecContext(ctx, insertQuery, stat.PlayerID, stat.MatchID, stat.Goals, stat.Assists, stat.Interceptions, stat.Rebounds)
	if err != nil {
		return fmt.Errorf("ошибка при вставке данных: %w", err)
	}
	return nil
}

// Поиск статистики игрока по его идентификатору (player_id) и матчу (match_id)
func (c *ChouseRepo) GetPlayerStatsByIDAndMatch(ctx context.Context, playerID, matchID string) ([]entity.PlayerStat, error) {
	query := `
		SELECT player_id, match_id, goals, assists, interceptions, rebounds
		FROM player_stats
		WHERE player_id = ? AND match_id = ?
	`
	rows, err := c.cHouseDB.DB.QueryContext(ctx, query, playerID, matchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []entity.PlayerStat
	for rows.Next() {
		var stat entity.PlayerStat
		err := rows.Scan(&stat.PlayerID, &stat.MatchID, &stat.Goals, &stat.Assists, &stat.Interceptions, &stat.Rebounds)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

// Поиск игроков с средним количеством голов больше определённого значения по id матча
func (c *ChouseRepo) GetPlayersWithAvgGoalsGreaterThanByMatch(ctx context.Context, minAvgGoals float64, matchID string) ([]entity.PlayerStat, error) {
	query := `
		SELECT player_id, AVG(goals) AS avg_goals
		FROM player_stats
		WHERE match_id = ?
		GROUP BY player_id
		HAVING avg_goals > ?
	`
	rows, err := c.cHouseDB.DB.QueryContext(ctx, query, matchID, minAvgGoals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []entity.PlayerStat
	for rows.Next() {
		var stat entity.PlayerStat
		err := rows.Scan(&stat.PlayerID, &stat.AVGGoals)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

// Поиск игроков, у которых сумма средних значений голов, перехватов, подборов и передач больше определённого значения по id матча
func (c *ChouseRepo) GetPlayersWithTotalAvgStatsGreaterThanByMatch(ctx context.Context, minTotalAvg float64, matchID string) ([]entity.PlayerStat, error) {
	query := `
		SELECT player_id,
		       AVG(goals) + AVG(assists) + AVG(interceptions) + AVG(rebounds) AS total_avg_stats
		FROM player_stats
		WHERE match_id = ?
		GROUP BY player_id
		HAVING total_avg_stats > ?
	`
	rows, err := c.cHouseDB.DB.QueryContext(ctx, query, matchID, minTotalAvg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []entity.PlayerStat
	for rows.Next() {
		var stat entity.PlayerStat
		err := rows.Scan(&stat.PlayerID, &stat.TotalAVGStats)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}
