package mongo_rp

import (
	"context"
	"errors"
	"fmt"
	"github.com/romeros69/basket/internal/apperrors"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	mongodb "github.com/romeros69/basket/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameRepo struct {
	mngCollection *mongo.Collection
}

func NewGameRepo(mng *mongodb.Mongo, collectionName string) *GameRepo {
	return &GameRepo{
		mngCollection: mng.Db.Collection(collectionName),
	}
}

var _ usecase.GameRp = (*GameRepo)(nil)

func (g *GameRepo) CreateGame(ctx context.Context, game *entity.Game) (string, error) {
	res, err := g.mngCollection.InsertOne(ctx, game)
	if err != nil {
		return "", fmt.Errorf("create game")
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (g *GameRepo) UpdateGame(ctx context.Context, gameID string, game *entity.Game) (*entity.Game, error) {
	objID, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return nil, apperrors.ErrInvalidGameID
	}

	filter := bson.M{
		"_id": objID,
	}

	if err = g.mngCollection.FindOneAndReplace(ctx, filter, game).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrGameNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return game, nil
}

func (g *GameRepo) GetGame(ctx context.Context, gameID string) (*entity.Game, error) {
	objID, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return nil, apperrors.ErrInvalidGameID
	}

	filter := bson.M{
		"_id": objID,
	}

	game := new(entity.Game)
	if err := g.mngCollection.FindOne(ctx, filter).Decode(game); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrGameNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return game, nil
}

func (g *GameRepo) DeleteGame(ctx context.Context, gameID string) error {
	objID, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return apperrors.ErrInvalidGameID
	}

	filter := bson.M{
		"_id": objID,
	}

	if err = g.mngCollection.FindOneAndDelete(ctx, filter).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return apperrors.ErrGameNotFound
		}
		return fmt.Errorf("mongo error: %w", err)
	}

	return nil
}

func (g *GameRepo) GetGameList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Game, error) {
	cursor, err := g.mngCollection.Find(ctx, bson.M{}, options.Find().SetLimit(pageSize).SetSkip((pageNumber-1)*pageSize))
	if err != nil {
		return nil, fmt.Errorf("mongo error: %w", err)
	}
	defer cursor.Close(ctx)

	var games []*entity.Game
	for cursor.Next(ctx) {
		var game *entity.Game
		err := cursor.Decode(&game)
		if err != nil {
			return nil, fmt.Errorf("mongo error: %w", err)
		}
		games = append(games, game)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return games, nil
}
