//go:build wireinject

package oss

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/service"
	"github.com/google/wire"
)

var ossProviders = wire.NewSet(service.NewOssService, repo.NewOssRepository,
	wire.Bind(new(service.IOssService), new(*service.OssService)),
	wire.Bind(new(repo.IOssRepository), new(*repo.OssRepository)),
)

func InitOssModule() *Module {
	panic(wire.Build(
		ossProviders,
		wire.Struct(new(Module), "Svc"),
	))
}
