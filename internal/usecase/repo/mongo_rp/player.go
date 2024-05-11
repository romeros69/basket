package mongo_rp

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/romeros69/basket/internal/apperrors"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	mongodb "github.com/romeros69/basket/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerRepo struct {
	mngCollection *mongo.Collection
}

const (
	name        = "name"
	surname     = "surname"
	middleName  = "middle_name"
	age         = "age"
	height      = "height"
	weight      = "weight"
	team        = "team"
	role        = "role"
	citizenship = "citizenship"
)

func NewPlayerRepo(mng *mongodb.Mongo, collectionName string) *PlayerRepo {
	return &PlayerRepo{
		mngCollection: mng.Db.Collection(collectionName),
	}
}

var _ usecase.PlayerRp = (*PlayerRepo)(nil)

func (p *PlayerRepo) CreatePlayer(ctx context.Context, player *entity.Player) (string, error) {
	res, err := p.mngCollection.InsertOne(ctx, player)
	if err != nil {
		return "", fmt.Errorf("create player: %w", err)
	}
	spew.Dump(res.InsertedID.(primitive.ObjectID).String())
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (p *PlayerRepo) UpdatePlayer(ctx context.Context, playerID string, player *entity.Player) (*entity.Player, error) {
	objID, err := primitive.ObjectIDFromHex(playerID)
	if err != nil {
		return nil, apperrors.ErrInvalidPlayerID
	}

	filter := bson.M{
		"_id": objID,
	}

	updatedPlayer := new(entity.Player)
	res := p.mngCollection.FindOneAndReplace(ctx, filter, player)
	if err = res.Decode(updatedPlayer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrPlayerNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return player, nil
}

func (p *PlayerRepo) GetPlayer(ctx context.Context, playerID string) (*entity.Player, error) {
	objID, err := primitive.ObjectIDFromHex(playerID)
	if err != nil {
		return nil, apperrors.ErrInvalidPlayerID
	}

	filter := bson.M{
		"_id": objID,
	}

	player := new(entity.Player)
	if err := p.mngCollection.FindOne(ctx, filter).Decode(player); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrPlayerNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return player, nil
}

func (p *PlayerRepo) DeletePlayer(ctx context.Context, playerID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerRepo) GetPlayerList(ctx context.Context) ([]*entity.Player, error) {
	//TODO implement me
	panic("implement me")
}
