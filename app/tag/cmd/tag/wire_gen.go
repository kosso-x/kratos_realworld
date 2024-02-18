// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"realworld/app/tag/internal/biz"
	"realworld/app/tag/internal/conf"
	"realworld/app/tag/internal/data"
	"realworld/app/tag/internal/server"
	"realworld/app/tag/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	tagRepo := data.NewTagRepo(dataData, logger)
	tagUsecase := biz.NewTagUsecase(tagRepo, logger)
	tagService := service.NewTagService(tagUsecase)
	grpcServer := server.NewGRPCServer(confData, confServer, tagService, logger)
	httpServer := server.NewHTTPServer(confData, confServer, tagService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}