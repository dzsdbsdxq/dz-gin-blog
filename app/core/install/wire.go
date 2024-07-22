//go:build wireinject

package install

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var InstallProviders = wire.NewSet(
	service.NewInstallService,
	wire.Bind(new(service.IInstallService), new(*service.InstallService)),
)

func InitInstallModule(db *gorm.DB) *Module {
	panic(wire.Build(
		InstallProviders,
		wire.Struct(new(Module), "Svc", "ISvc"),
	))
}
