package service

import (
	"context"

	v1 "realworld/api/profile/v1"
	"realworld/app/profile/internal/biz"
)

// ProfileService is a profile service.
type ProfileService struct {
	v1.UnimplementedProfileServer

	uc *biz.ProfileUsecase
}

// NewProfileService new a profile service.
func NewProfileService(uc *biz.ProfileUsecase) *ProfileService {
	return &ProfileService{uc: uc}
}

func (s *ProfileService) RegisterCeleb(ctx context.Context, req *v1.RegisterCelebRequest) (res *v1.ProfileReply, err error) {
	r, err := s.uc.RegisterCeleb(ctx, req.User)
	if err != nil {
		return nil, err
	}

	res = &v1.ProfileReply{
		Profile: r,
	}

	return
}

func (s *ProfileService) GetProfile(ctx context.Context, req *v1.ProfileRequest) (res *v1.ProfileReply, err error) {
	r, err := s.uc.GetProfile(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	res = &v1.ProfileReply{
		Profile: r,
	}

	return
}

func (s *ProfileService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (res *v1.ProfileReply, err error) {
	r, err := s.uc.FollowUser(ctx, req.Username, req.User.Email)
	if err != nil {
		return nil, err
	}

	res = &v1.ProfileReply{
		Profile: r,
	}

	return
}

func (s *ProfileService) UnfollowUser(ctx context.Context, req *v1.ProfileRequest) (res *v1.ProfileReply, err error) {
	r, err := s.uc.UnfollowUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	res = &v1.ProfileReply{
		Profile: r,
	}

	return
}
