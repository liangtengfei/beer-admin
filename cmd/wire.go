//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/liangtengfei/beer-admin/api"
	"github.com/liangtengfei/beer-admin/internal/biz"
	"github.com/liangtengfei/beer-admin/internal/conf"
	"github.com/liangtengfei/beer-admin/internal/data"
	"github.com/liangtengfei/beer-admin/internal/server"
	"github.com/liangtengfei/beer-admin/internal/service"
	"go.uber.org/zap"
)

func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.App, *zap.SugaredLogger) (*App, func(), error) {
	panic(wire.Build(
		server.ProviderSetServer,
		data.ProviderSetData,
		biz.ProviderSetBiz,
		service.ProviderSetService,
		api.NewSetAPI,
		newApp,
	))
}
