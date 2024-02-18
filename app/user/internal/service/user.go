package service

import (
	"context"
	v1 "realworld/api/user/v1"
	"realworld/app/user/internal/biz"
)

// UserService is a user service.
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (res *v1.GetCurrentUserReply, err error) {
	r, err := s.uc.CurrentUser(ctx)

	if err != nil {
		return nil, err
	}

	res = &v1.GetCurrentUserReply{
		User: r,
	}

	return
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (res *v1.UpdateUserReply, err error) {
	r, err := s.uc.UpdateUser(ctx, req.User)
	if err != nil {
		return nil, err
	}

	res = &v1.UpdateUserReply{
		User: r,
	}

	return
}
