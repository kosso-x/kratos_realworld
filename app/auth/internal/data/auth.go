package data

import (
	"context"
	"fmt"
	"realworld/app/auth/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) FindByEmail(ctx context.Context, email string) (user *biz.User, err error) {
	if err = r.data.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return
}

func (r *authRepo) UpdateToken(ctx context.Context, user *biz.User) (err error) {
	if err = r.data.db.Save(user).Error; err != nil {
		return
	}

	return
}

// redis 存 token
func (r *authRepo) PreserveToken(ctx context.Context, user *biz.User) (err error) {
	id := user.Model.ID
	if err = r.data.redisCli.Set(ctx, fmt.Sprintf("realworld-%v-email", id), user.Email, 15*60*time.Second).Err(); err != nil {
		return
	}
	if err = r.data.redisCli.Set(ctx, fmt.Sprintf("realworld-%v-username", id), user.UserName, 15*60*time.Second).Err(); err != nil {
		return
	}
	if err = r.data.redisCli.Set(ctx, fmt.Sprintf("realworld-%v-token", id), user.Token, 15*60*time.Second).Err(); err != nil {
		return
	}

	return
}

func (r *authRepo) CreateUser(ctx context.Context, user *biz.User) (err error) {
	if err = r.data.db.Create(user).Error; err != nil {
		return
	}

	return
}
