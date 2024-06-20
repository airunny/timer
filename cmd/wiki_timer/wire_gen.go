// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"timer/api/common"
	"timer/internal/conf"
	"timer/internal/dao"
	"timer/internal/server"
	"timer/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(serverConfig *common.ServerConfig, dataConfig *common.DataConfig, business *conf.Business, logger log.Logger) (*kratos.App, func(), error) {
	db, cleanup, err := dao.NewMySQL(dataConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	client, cleanup2, err := dao.NewRedis(dataConfig, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	user := dao.NewUser(db, client)
	serviceService := service.NewService(business, user)
	grpcServer := server.NewGRPCServer(serverConfig, serviceService, logger)
	httpServer := server.NewHTTPServer(serverConfig, serviceService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
