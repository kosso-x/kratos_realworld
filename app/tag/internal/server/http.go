package server

import (
	v1 "realworld/api/tag/v1"
	"realworld/app/tag/internal/conf"
	"realworld/app/tag/internal/data"
	"realworld/app/tag/internal/service"
	rwmiddleware "realworld/package/rwMiddleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(conf *conf.Data, c *conf.Server, tag *service.TagService, logger log.Logger) *http.Server {
	var (
		redisCli = data.NewRedisCli(conf)
		opts     = []http.ServerOption{
			http.Middleware(
				recovery.Recovery(),
				logging.Server(logger),
				rwmiddleware.AuthTokenHandler(conf.Jwt.Secretkey, redisCli),
			),
		}
	)
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterTagHTTPServer(srv, tag)
	return srv
}
