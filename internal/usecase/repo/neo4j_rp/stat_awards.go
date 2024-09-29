package neo4j_rp

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	neo4jdb "github.com/romeros69/basket/pkg/neo4j"
)

type StatAwardsRepo struct {
	neoDB *neo4jdb.Neo4j
}

func NewStatAwardsRepo(neoDB *neo4jdb.Neo4j) *StatAwardsRepo {
	return &StatAwardsRepo{
		neoDB: neoDB,
	}
}

var _ usecase.StatAwardsRp = (*StatAwardsRepo)(nil)

// CreateRecord - Функция для создания записи (награждение игрока в рамках матча и турнира)
func (sa *StatAwardsRepo) CreateRecord(ctx context.Context, rewardStat entity.RewardStat) error {
	_, err := sa.neoDB.DB.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MERGE (r:Reward {id: $rewardId})
			MERGE (p:Player {id: $playerId})
			MERGE (m:Match {id: $matchId})
			MERGE (t:Tournament {id: $tournamentId})
			MERGE (r)-[:AWARDED_TO]->(p)
			MERGE (r)-[:AWARDED_FOR_MATCH]->(m)
			MERGE (m)-[:PART_OF_TOURNAMENT]->(t)
			RETURN r, p, m, t`
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"rewardId":     rewardStat.Reward,
			"playerId":     rewardStat.Player,
			"matchId":      rewardStat.Match,
			"tournamentId": rewardStat.Tournament,
		})
		return nil, err
	})

	return err
}

// ViewPlayersAndRewardsInTournament - Функция для просмотра какие игроки получили награды в рамках турнира
func (sa *StatAwardsRepo) ViewPlayersAndRewardsInTournament(ctx context.Context, tournamentId string) ([]entity.RewardStat, error) {
	var rewards []entity.RewardStat

	_, err := sa.neoDB.DB.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (t:Tournament {id: $tournamentId})<-[:PART_OF_TOURNAMENT]-(m:Match)<-[:AWARDED_FOR_MATCH]-(r:Reward)-[:AWARDED_TO]->(p:Player)
			RETURN p.id AS player, r.id AS reward, m.id AS match, t.id AS tournament`
		result, err := tx.Run(ctx, query, map[string]interface{}{
			"tournamentId": tournamentId,
		})
		if err != nil {
			return nil, err
		}

		for result.Next(ctx) {
			record := result.Record()

			player, _ := record.Get("player")
			reward, _ := record.Get("reward")
			match, _ := record.Get("match")
			tournament, _ := record.Get("tournament")

			rewards = append(rewards, entity.RewardStat{
				Player:     player.(string),
				Reward:     reward.(string),
				Match:      match.(string),
				Tournament: tournament.(string),
			})
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return rewards, nil
}

// ViewPlayersAndRewardsInMatch - Функция для просмотра какие игроки получили награды в рамках матча
func (sa *StatAwardsRepo) ViewPlayersAndRewardsInMatch(ctx context.Context, matchId string) ([]entity.RewardStat, error) {
	var rewards []entity.RewardStat

	_, err := sa.neoDB.DB.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (m:Match {id: $matchId})<-[:AWARDED_FOR_MATCH]-(r:Reward)-[:AWARDED_TO]->(p:Player)
			RETURN p.id AS player, r.id AS reward, m.id AS match`
		result, err := tx.Run(ctx, query, map[string]interface{}{
			"matchId": matchId,
		})
		if err != nil {
			return nil, err
		}

		for result.Next(ctx) {
			record := result.Record()

			player, _ := record.Get("player")
			reward, _ := record.Get("reward")
			match, _ := record.Get("match")

			rewards = append(rewards, entity.RewardStat{
				Player: player.(string),
				Reward: reward.(string),
				Match:  match.(string),
			})
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return rewards, nil
}

// ViewRewardsForPlayer - Функция для просмотра наград игрока и информации о матче и турнире, где была получена награда
func (sa *StatAwardsRepo) ViewRewardsForPlayer(ctx context.Context, playerId string) ([]entity.RewardStat, error) {
	var rewards []entity.RewardStat

	_, err := sa.neoDB.DB.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (p:Player {id: $playerId})<-[:AWARDED_TO]-(r:Reward)-[:AWARDED_FOR_MATCH]->(m:Match)-[:PART_OF_TOURNAMENT]->(t:Tournament)
			RETURN r.id AS reward, m.id AS match, t.id AS tournament`
		result, err := tx.Run(ctx, query, map[string]interface{}{
			"playerId": playerId,
		})
		if err != nil {
			return nil, err
		}

		for result.Next(ctx) {
			record := result.Record()

			reward, _ := record.Get("reward")
			match, _ := record.Get("match")
			tournament, _ := record.Get("tournament")

			rewards = append(rewards, entity.RewardStat{
				Reward:     reward.(string),
				Match:      match.(string),
				Tournament: tournament.(string),
				Player:     playerId,
			})
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return rewards, nil
}

// ViewWhoGotSpecificReward - Функция для просмотра кто получил конкретную награду (с информацией о матче и турнире)
func (sa *StatAwardsRepo) ViewWhoGotSpecificReward(ctx context.Context, rewardId string) ([]entity.RewardStat, error) {
	var rewards []entity.RewardStat

	_, err := sa.neoDB.DB.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (r:Reward {id: $rewardId})-[:AWARDED_TO]->(p:Player), 
			      (r)-[:AWARDED_FOR_MATCH]->(m:Match)-[:PART_OF_TOURNAMENT]->(t:Tournament)
			RETURN p.id AS player, m.id AS match, t.id AS tournament`
		result, err := tx.Run(ctx, query, map[string]interface{}{
			"rewardId": rewardId,
		})
		if err != nil {
			return nil, err
		}

		for result.Next(ctx) {
			record := result.Record()

			player, _ := record.Get("player")
			match, _ := record.Get("match")
			tournament, _ := record.Get("tournament")

			rewards = append(rewards, entity.RewardStat{
				Player:     player.(string),
				Match:      match.(string),
				Tournament: tournament.(string),
				Reward:     rewardId,
			})
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return rewards, nil
}