package rwmiddleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Iat      int    `json:"iat"`
	jwt.RegisteredClaims
}

var (
	UserClaims *MyClaims
)

func AuthTokenHandler(secret string, redisCli *redis.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if err = parseToken(ctx, req, secret); err != nil {
				return nil, err
			}

			if err = verifyToken(ctx, redisCli); err != nil {
				return nil, err
			}

			if err = prolongedToken(ctx, redisCli); err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}

func parseToken(ctx context.Context, req interface{}, secret string) (err error) {
	if request, ok := transport.FromServerContext(ctx); ok {
		authorization := strings.SplitN(request.RequestHeader().Get("Authorization"), " ", 2)
		token, _ := jwt.ParseWithClaims(authorization[1], &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secret), nil
		})

		if UserClaims, ok = token.Claims.(*MyClaims); !ok {
			return fmt.Errorf("keyFunc is missing")
		}

	}
	return
}

func verifyToken(ctx context.Context, redisCli *redis.Client) (err error) {
	if redis_username, err := redisCli.Get(ctx, fmt.Sprintf("realworld-%d-username", UserClaims.Id)).Result(); (err != nil) || (redis_username != UserClaims.UserName) {
		return fmt.Errorf("token verification failed")
	}

	return nil
}

func prolongedToken(ctx context.Context, redisCli *redis.Client) (err error) {
	id := UserClaims.Id
	if ok, err := redisCli.Expire(ctx, fmt.Sprintf("realworld-%v-email", id), 15*60*time.Second).Result(); !ok {
		return err
	}
	if ok, err := redisCli.Expire(ctx, fmt.Sprintf("realworld-%v-username", id), 15*60*time.Second).Result(); !ok {
		return err
	}
	if ok, err := redisCli.Expire(ctx, fmt.Sprintf("realworld-%v-token", id), 15*60*time.Second).Result(); !ok {
		return err
	}

	return
}
