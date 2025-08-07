package service

import (
	"context"

	pb "github.com/yygqzzk/review-b/api/business/v1"
	"github.com/yygqzzk/review-b/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewReq) (*pb.ReplyReviewRsp, error) {
	// 商家回复评价
	replyEntity := &biz.ReplyEntity{
		ReviewId:  req.GetReviewId(),
		StoreId:   req.GetStoreId(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	}
	err := s.uc.ReplyReview(ctx, replyEntity)
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewRsp{
		ReplyId: replyEntity.ReplyId,
	}, nil
}
