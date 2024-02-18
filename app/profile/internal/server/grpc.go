package server

import (
	v1 "realworld/api/profile/v1"
	"realworld/app/profile/internal/conf"
	"realworld/app/profile/internal/data"
	"realworld/app/profile/internal/service"
	rwmiddleware "realworld/package/rwMiddleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(conf *conf.Data, c *conf.Server, profile *service.ProfileService, logger log.Logger) *grpc.Server {
	var (
		redisCli = data.NewRedisCli(conf)
		opts     = []grpc.ServerOption{
			grpc.Middleware(
				recovery.Recovery(),
				logging.Server(logger),
				rwmiddleware.AuthTokenHandler(conf.Jwt.Secretkey, redisCli),
			),
		}
	)
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterProfileServer(srv, profile)
	return srv
}