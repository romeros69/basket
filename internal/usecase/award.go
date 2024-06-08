package usecase

import (
	"context"

	"github.com/romeros69/basket/internal/entity"
)

type AwardUC struct {
	awardRp AwardRp
}

func NewAwardUC(awardRp AwardRp) *AwardUC {
	return &AwardUC{
		awardRp: awardRp,
	}
}

var _ Award = (*AwardUC)(nil)

func (a *AwardUC) CreateAward(ctx context.Context, award *entity.Award) (string, error) {
	return a.awardRp.CreateAward(ctx, award)
}

func (a *AwardUC) UpdateAward(ctx context.Context, awardID string, award *entity.Award) (*entity.Award, error) {
	return a.awardRp.UpdateAward(ctx, awardID, award)
}

func (a *AwardUC) GetAward(ctx context.Context, awardID string) (*entity.Award, error) {
	return a.awardRp.GetAward(ctx, awardID)
}

func (a *AwardUC) DeleteAward(ctx context.Context, awardID string) error {
	return a.awardRp.DeleteAward(ctx, awardID)
}

func (a *AwardUC) GetAwardList(ctx context.Context, pageSize, pageNumber int64) ([]*entity.Award, error) {
	return a.awardRp.GetAwardList(ctx, pageSize, pageNumber)
}
