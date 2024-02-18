package biz

import (
	"context"
	"fmt"
	v1 "realworld/api/profile/v1"
	pwdoptions "realworld/package/pwdOptions"
	rwmiddleware "realworld/package/rwMiddleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
)

// Profile is a Profile model.
type Profile struct {
	Hello string
}

// ProfileRepo is a Greater repo.
type ProfileRepo interface {
	FindByName(username string) (user *CelebUser, err error)
	FindByEmail(username string) (user *User, err error)
	FindFollow(user_id, follow_id int64) (err error)
	CreateFollow(user_id, celeb_user_id int64) (err error)
	DeleteFollow(user_id, celeb_user_id int64) (err error)
	CreateCelebUser(user *CelebUser) (err error)
	CelebUserFollows(celebUser *CelebUser) (follows []Follows, err error)
}

// ProfileUsecase is a Profile usecase.
type ProfileUsecase struct {
	repo ProfileRepo
	log  *log.Helper
}

// NewProfileUsecase new a Profile usecase.
func NewProfileUsecase(repo ProfileRepo, logger log.Logger) (uc *ProfileUsecase) {
	return &ProfileUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProfileUsecase) RegisterCeleb(ctx context.Context, req *v1.RegisterCelebRequest_User) (res *v1.ProfileReply_Data, err error) {
	register_user, _ := uc.repo.FindByName(req.Username)
	if register_user != nil {
		return nil, fmt.Errorf("account has been registered!")
	}

	register_user = &CelebUser{
		Email:    req.Email,
		UserName: req.Username,
		Password: req.Password,
	}

	if register_user.Password, err = pwdoptions.EncryptPassword(register_user.Password); err != nil {
		return nil, err
	}

	if err = uc.repo.CreateCelebUser(register_user); err != nil {
		return nil, err
	}

	if err = gconv.Struct(register_user, &res); err != nil {
		return nil, err
	}

	return
}

func (uc *ProfileUsecase) GetProfile(ctx context.Context, useranme string) (res *v1.ProfileReply_Data, err error) {
	user, err := uc.repo.FindByName(useranme)
	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(user, &res); err != nil {
		return nil, err
	}

	if err = uc.repo.FindFollow(rwmiddleware.UserClaims.Id, int64(user.Model.ID)); err == nil {
		res.Following = true
		return res, nil
	}

	return
}

func (uc *ProfileUsecase) FollowUser(ctx context.Context, celebUserName, followUserEmail string) (res *v1.ProfileReply_Data, err error) {
	celebUser, err := uc.repo.FindByName(celebUserName)
	if err != nil {
		return nil, err
	}

	user, err := uc.repo.FindByEmail(followUserEmail)
	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(celebUser, &res); err != nil {
		return nil, err
	}

	if err = uc.repo.FindFollow(int64(user.Model.ID), int64(celebUser.Model.ID)); err != nil {
		if err = uc.repo.CreateFollow(int64(user.Model.ID), int64(celebUser.Model.ID)); err != nil {
			return nil, err
		}
	}

	res.Following = true
	return
}

func (uc *ProfileUsecase) UnfollowUser(ctx context.Context, celebUserName string) (res *v1.ProfileReply_Data, err error) {
	celebUser, err := uc.repo.FindByName(celebUserName)
	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(celebUser, &res); err != nil {
		return nil, err
	}

	if err = uc.repo.FindFollow(rwmiddleware.UserClaims.Id, int64(celebUser.Model.ID)); err == nil {
		if err = uc.repo.DeleteFollow(rwmiddleware.UserClaims.Id, int64(celebUser.Model.ID)); err != nil {
			return nil, err
		}
	}

	res.Following = false
	return
}
