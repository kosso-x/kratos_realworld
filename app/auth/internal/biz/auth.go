package biz

import (
	"context"
	"fmt"
	v1 "realworld/api/auth/v1"
	"realworld/app/auth/internal/conf"
	pwdoptions "realworld/package/pwdOptions"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRepo is a Greater repo.
type AuthRepo interface {
	FindByEmail(ctx context.Context, email string) (user *User, err error)
	UpdateToken(ctx context.Context, user *User) (err error)
	PreserveToken(ctx context.Context, user *User) (err error)
	CreateUser(ctx context.Context, user *User) (err error)
}

// AuthUsecase is a Auth usecase.
type AuthUsecase struct {
	conf *conf.Data
	repo AuthRepo
	log  *log.Helper
}

// NewAuthUsecase new a Auth usecase.
func NewAuthUsecase(conf *conf.Data, repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{conf: conf, repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest_UserLogin) (user *v1.LoginReply_User, err error) {
	login_user, err := uc.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if login_user == nil {
		return nil, fmt.Errorf("this account has not been registered yet!")
	}

	if err = pwdoptions.ComparePass(login_user.Password, req.Password); err != nil {
		return nil, fmt.Errorf("account password is wrong!")
	}

	if err = uc.generateToken(login_user); err != nil {
		return nil, err
	}

	if err = uc.repo.UpdateToken(ctx, login_user); err != nil {
		return nil, err
	}

	if err = uc.repo.PreserveToken(ctx, login_user); err != nil {
		return nil, err
	}

	if err = gconv.Struct(login_user, &user); err != nil {
		return nil, err
	}

	return
}

func (uc *AuthUsecase) Register(ctx context.Context, req *v1.RegisterRequest_Data) (user *v1.RegisterReply_User, err error) {
	register_user, _ := uc.repo.FindByEmail(ctx, req.Email)

	if register_user != nil {
		return nil, fmt.Errorf("account has been registered!")
	}

	register_user = &User{
		Email:    req.Email,
		UserName: req.Username,
		Password: req.Password,
	}

	if err = uc.generateToken(register_user); err != nil {
		return nil, err
	}

	if register_user.Password, err = pwdoptions.EncryptPassword(register_user.Password); err != nil {
		return nil, err
	}

	if err = uc.repo.CreateUser(ctx, register_user); err != nil {
		return nil, err
	}

	if err = gconv.Struct(register_user, &user); err != nil {
		return nil, err
	}

	return
}

// 生成 token
func (uc *AuthUsecase) generateToken(user *User) (err error) {
	gen_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Model.ID,
		"username": user.UserName,
		"iat":      int(time.Now().Unix()),
	})
	user.Token, err = gen_token.SignedString([]byte(uc.conf.Jwt.Secretkey))

	if err != nil {
		message := fmt.Errorf("sign token error: %s", err)
		return message
	}

	return
}
