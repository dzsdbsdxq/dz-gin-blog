//go:build wireinject

package setting

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var SettingProviders = wire.NewSet(
	controller.NewSettingHandler,
	service.NewSettingService,
	repo.NewSettingRepository,
	dao.NewSettingDao,
	wire.Bind(new(service.ISettingService), new(*service.SettingService)),
	wire.Bind(new(repo.ISettingRepository), new(*repo.SettingRepository)),
	wire.Bind(new(dao.ISettingDao), new(*dao.SettingDao)),
)

func InitSettingModule(db *gorm.DB) *Module {
	panic(wire.Build(
		SettingProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
