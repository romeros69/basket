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

type LeagueRepo struct {
	mngCollection *mongo.Collection
}

func NewLeagueRepo(mng *mongodb.Mongo, collectionName string) *LeagueRepo {
	return &LeagueRepo{
		mngCollection: mng.Db.Collection(collectionName),
	}
}

var _ usecase.LeagueRp = (*LeagueRepo)(nil)

func (l *LeagueRepo) CreateLeague(ctx context.Context, league *entity.League) (string, error) {
	res, err := l.mngCollection.InsertOne(ctx, league)
	if err != nil {
		return "", fmt.Errorf("create league")
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (l *LeagueRepo) UpdateLeague(ctx context.Context, leagueID string, league *entity.League) (*entity.League, error) {
	objID, err := primitive.ObjectIDFromHex(leagueID)
	if err != nil {
		return nil, apperrors.ErrInvalidLeagueID
	}

	filter := bson.M{
		"_id": objID,
	}

	if err = l.mngCollection.FindOneAndReplace(ctx, filter, league).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrLeagueNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return league, nil
}

func (l *LeagueRepo) GetLeague(ctx context.Context, leagueID string) (*entity.League, error) {
	objID, err := primitive.ObjectIDFromHex(leagueID)
	if err != nil {
		return nil, apperrors.ErrInvalidLeagueID
	}

	filter := bson.M{
		"_id": objID,
	}

	league := new(entity.League)
	if err := l.mngCollection.FindOne(ctx, filter).Decode(league); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrLeagueNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return league, nil
}

func (l *LeagueRepo) DeleteLeague(ctx context.Context, leagueID string) error {
	objID, err := primitive.ObjectIDFromHex(leagueID)
	if err != nil {
		return apperrors.ErrInvalidLeagueID
	}

	filter := bson.M{
		"_id": objID,
	}

	if err = l.mngCollection.FindOneAndDelete(ctx, filter).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return apperrors.ErrLeagueNotFound
		}
		return fmt.Errorf("mongo error: %w", err)
	}

	return nil
}

func (l *LeagueRepo) GetLeagueList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.League, error) {
	cursor, err := l.mngCollection.Find(ctx, bson.M{}, options.Find().SetLimit(pageSize).SetSkip((pageNumber-1)*pageSize))
	if err != nil {
		return nil, fmt.Errorf("mongo error: %w", err)
	}
	defer cursor.Close(ctx)

	var leagues []*entity.League
	for cursor.Next(ctx) {
		var league *entity.League
		err := cursor.Decode(&league)
		if err != nil {
			return nil, fmt.Errorf("mongo error: %w", err)
		}
		leagues = append(leagues, league)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return leagues, nil
}
