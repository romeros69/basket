package mongo_rp

import (
	"context"
	"fmt"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	mongodb "github.com/romeros69/basket/pkg/mongo"
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
	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (p *PlayerRepo) UpdatePlayer(ctx context.Context, player *entity.Player) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerRepo) GetPlayer(ctx context.Context, playerID string) (*entity.Player, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerRepo) DeletePlayer(ctx context.Context, playerID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerRepo) GetPlayerList(ctx context.Context) ([]*entity.Player, error) {
	//TODO implement me
	panic("implement me")
}
