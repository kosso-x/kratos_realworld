package server

import (
	v1 "realworld/api/article/v1"
	"realworld/app/article/internal/conf"
	"realworld/app/article/internal/data"
	"realworld/app/article/internal/service"
	rwmiddleware "realworld/package/rwMiddleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(conf *conf.Data, c *conf.Server, article *service.ArticleService, logger log.Logger) *http.Server {
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
	v1.RegisterArticleHTTPServer(srv, article)
	return srv
}
