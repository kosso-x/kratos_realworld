package service

import (
	"context"
	v1 "realworld/api/auth/v1"
	"realworld/app/auth/internal/biz"
)

// AuthService is a auth service.
type AuthService struct {
	v1.UnimplementedAuthServer

	uc *biz.AuthUsecase
}

// NewAuthService new a auth service.
func NewAuthService(uc *biz.AuthUsecase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, req *v1.LoginRequest) (rep *v1.LoginReply, err error) {
	/*
		1. 通过登录信息找 user
		2. 验证密码
		3. 生成token
		4. 保存token
	*/

	r, err := s.uc.Login(ctx, req.User)
	if err != nil {
		return nil, err
	}

	rep = &v1.LoginReply{
		User: r,
	}

	return
}

func (s *AuthService) Register(ctx context.Context, req *v1.RegisterRequest) (rep *v1.RegisterReply, err error) {
	/*
		1. 查找有没有邮箱账号
		2. 有就报错
		3. 没有就创建(生成token, 加密密码)
		4. 创建成功
	*/
	r, err := s.uc.Register(ctx, req.User)
	if err != nil {
		return nil, err
	}

	rep = &v1.RegisterReply{
		User: r,
	}

	return
}
