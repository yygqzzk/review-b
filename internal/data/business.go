package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/yygqzzk/review-b/api/review/v1"
	"github.com/yygqzzk/review-b/internal/biz"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (b *businessRepo) Reply(ctx context.Context, replyEntity *biz.ReplyEntity) error {
	b.log.WithContext(ctx).Infof("[data] Reply: %v", replyEntity)
	rsp, err := b.data.reviewClient.ReplyReview(ctx, &pb.ReplyReviewReq{
		ReviewID:  replyEntity.ReviewID,
		StoreID:   replyEntity.StoreID,
		Content:   replyEntity.Content,
		PicInfo:   replyEntity.PicInfo,
		VideoInfo: replyEntity.VideoInfo,
	})
	if err != nil {
		return err
	}
	replyEntity.ReplyID = rsp.ReplyID
	return nil
}

func (b *businessRepo) SaveAppeal(ctx context.Context, appealEntity *biz.AppealEntity) error {
	b.log.WithContext(ctx).Infof("[data] SaveAppeal: %v", appealEntity)
	rsp, err := b.data.reviewClient.AppealReview(ctx, &pb.AppealReviewReq{
		ReviewID:  appealEntity.ReviewID,
		StoreID:   appealEntity.StoreID,
		Reason:    appealEntity.Reason,
		Content:   appealEntity.Content,
		PicInfo:   appealEntity.PicInfo,
		VideoInfo: appealEntity.VideoInfo,
	})
	if err != nil {
		return err
	}
	appealEntity.AppealID = rsp.AppealID
	return nil
}
