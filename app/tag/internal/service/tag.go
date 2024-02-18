package service

import (
	"context"

	v1 "realworld/api/tag/v1"
	"realworld/app/tag/internal/biz"
)

// TagService is a tag service.
type TagService struct {
	v1.UnimplementedTagServer

	uc *biz.TagUsecase
}

// NewTagService new a tag service.
func NewTagService(uc *biz.TagUsecase) *TagService {
	return &TagService{uc: uc}
}

func (s *TagService) GetTags(ctx context.Context, req *v1.GetTagRequest) (res *v1.GetTagReply, err error) {
	r, err := s.uc.GetTags(ctx)

	if err != nil {
		return nil, err
	}

	res = &v1.GetTagReply{
		Tags: r,
	}

	return
}
