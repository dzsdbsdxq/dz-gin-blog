//go:build wireinject

package install

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var InstallProviders = wire.NewSet(
	controller.NewInstallHandler,
	service.NewInstallService,
	wire.Bind(new(service.IInstallService), new(*service.InstallService)),
)

func InitInstallModule(db *gorm.DB) *Module {
	panic(wire.Build(
		InstallProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
