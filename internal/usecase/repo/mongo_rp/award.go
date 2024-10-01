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

type AwardRepo struct {
	mngCollection *mongo.Collection
}

func NewAwardRepo(mng *mongodb.Mongo, collectionName string) *AwardRepo {
	return &AwardRepo{
		mngCollection: mng.DB.Collection(collectionName),
	}
}

var _ usecase.AwardRp = (*AwardRepo)(nil)

func (a *AwardRepo) CreateAward(ctx context.Context, award *entity.Award) (string, error) {
	res, err := a.mngCollection.InsertOne(ctx, award)
	if err != nil {
		return "", fmt.Errorf("create award: %w", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *AwardRepo) UpdateAward(ctx context.Context, awardID string, award *entity.Award) (*entity.Award, error) {
	objID, err := primitive.ObjectIDFromHex(awardID)
	if err != nil {
		return nil, apperrors.ErrInvalidAwardID
	}

	filter := bson.M{
		"_id": objID,
	}

	if err = a.mngCollection.FindOneAndReplace(ctx, filter, award).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrAwardNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return award, nil
}

func (a *AwardRepo) GetAward(ctx context.Context, awardID string) (*entity.Award, error) {
	objID, err := primitive.ObjectIDFromHex(awardID)
	if err != nil {
		return nil, apperrors.ErrInvalidAwardID
	}

	filter := bson.M{
		"_id": objID,
	}

	award := new(entity.Award)
	if err := a.mngCollection.FindOne(ctx, filter).Decode(award); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.ErrAwardNotFound
		}
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return award, nil
}

func (a *AwardRepo) DeleteAward(ctx context.Context, awardID string) error {
	objID, err := primitive.ObjectIDFromHex(awardID)
	if err != nil {
		return apperrors.ErrInvalidAwardID
	}

	filter := bson.M{
		"_id": objID,
	}

	if err = a.mngCollection.FindOneAndDelete(ctx, filter).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return apperrors.ErrAwardNotFound
		}
		return fmt.Errorf("mongo error: %w", err)
	}

	return nil
}

func (a *AwardRepo) GetAwardList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Award, error) {
	cursor, err := a.mngCollection.Find(ctx, bson.M{}, options.Find().SetLimit(pageSize).SetSkip((pageNumber-1)*pageSize))
	if err != nil {
		return nil, fmt.Errorf("mongo error: %w", err)
	}
	defer cursor.Close(ctx)

	var awards []*entity.Award
	for cursor.Next(ctx) {
		var award *entity.Award
		err := cursor.Decode(&award)
		if err != nil {
			return nil, fmt.Errorf("mongo error: %w", err)
		}
		awards = append(awards, award)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("mongo error: %w", err)
	}

	return awards, nil
}
