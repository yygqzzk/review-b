package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// BusinessRepo is a Business repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyEntity) error
}

// BusinessUsecase is a Business usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BusinessUsecase) ReplyReview(ctx context.Context, replyEntity *ReplyEntity) error {
	uc.log.WithContext(ctx).Infof("[biz] ReplyReview: %v", replyEntity)
	return uc.repo.Reply(ctx, replyEntity)
}
