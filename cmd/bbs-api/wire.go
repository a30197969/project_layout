// +build wireinject
package main

import (
	"github.com/google/wire"
	"project_layout/internal/biz"
	"project_layout/internal/conf"
	"project_layout/internal/data"
	"project_layout/internal/pkg"
	"project_layout/internal/server"
)

func initApp(cfg *conf.Config) *pkg.App {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, newApp))
}
