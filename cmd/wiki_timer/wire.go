//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"timer/api/common"
	"timer/internal/conf"
	"timer/internal/dao"
	"timer/internal/server"
	"timer/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*common.ServerConfig, *common.DataConfig, *conf.Business, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, service.ProviderSet, newApp))
}
