package biz

import (
	"context"
	v1 "realworld/api/user/v1"
	rwmiddleware "realworld/package/rwMiddleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
)

// UserRepo is a Greater repo.
type UserRepo interface {
	FindById(ctx context.Context, id int64) (user *User, err error)
	UpdateUser(ctx context.Context, user *User) (err error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CurrentUser(ctx context.Context) (req *v1.GetCurrentUserReply_User, err error) {
	user, err := uc.repo.FindById(ctx, rwmiddleware.UserClaims.Id)
	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(user, &req); err != nil {
		return nil, err
	}

	return
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest_User) (res *v1.UpdateUserReply_User, err error) {
	user, err := uc.repo.FindById(ctx, rwmiddleware.UserClaims.Id)
	if err != nil {
		return nil, err
	}

	user.Email = req.Email
	user.Bio = req.Bio
	user.Image = req.Image

	err = uc.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(user, &req); err != nil {
		return nil, err
	}

	return
}
