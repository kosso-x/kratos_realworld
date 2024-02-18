package data

import (
	"realworld/app/profile/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewProfileRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *profileRepo) FindByName(username string) (user *biz.CelebUser, err error) {
	if err = r.data.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return
}

func (r *profileRepo) FindByEmail(email string) (user *biz.User, err error) {
	if err = r.data.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return
}

func (r *profileRepo) FindFollow(user_id, celeb_user_id int64) (err error) {
	if err = r.data.db.Where("user_id = ? and follow_id = ?", user_id, celeb_user_id).Error; err != nil {
		return
	}

	return
}

func (r *profileRepo) CreateFollow(user_id, celeb_user_id int64) (err error) {
	if err = r.data.db.Create(&biz.Follows{
		UserId:   user_id,
		FollowId: celeb_user_id,
	}).Error; err != nil {
		return
	}

	return
}

func (r *profileRepo) DeleteFollow(user_id, celeb_user_id int64) (err error) {
	if err = r.data.db.Where("user_id = ? and follow_id = ?", user_id, celeb_user_id).Delete(&biz.Follows{}).Error; err != nil {
		return
	}

	return
}

func (r *profileRepo) CreateCelebUser(user *biz.CelebUser) (err error) {
	if err = r.data.db.Create(user).Error; err != nil {
		return
	}

	return
}

func (r *profileRepo) CelebUserFollows(celebUser *biz.CelebUser) (follows []biz.Follows, err error) {
	if err = r.data.db.Where("follow_id = ?", celebUser.Model.ID).Find(&follows).Error; err != nil {
		return
	}

	return
}
