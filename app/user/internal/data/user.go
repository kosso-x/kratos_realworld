package data

import (
	"context"
	"realworld/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) FindById(ctx context.Context, id int64) (user *biz.User, err error) {
	if err = r.data.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (err error) {
	if err = r.data.db.Save(user).Error; err != nil {
		return
	}

	return
}
