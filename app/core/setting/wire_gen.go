// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package setting

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitSettingModule(db *gorm.DB) *Module {
	settingDao := dao.NewSettingDao(db)
	settingRepository := repo.NewSettingRepository(settingDao)
	settingService := service.NewSettingService(settingRepository)
	settingHandler := controller.NewSettingHandler(settingService)
	module := &Module{
		Svc: settingService,
		Hdl: settingHandler,
	}
	return module
}

// wire.go:

var SettingProviders = wire.NewSet(controller.NewSettingHandler, service.NewSettingService, repo.NewSettingRepository, dao.NewSettingDao, wire.Bind(new(service.ISettingService), new(*service.SettingService)), wire.Bind(new(repo.ISettingRepository), new(*repo.SettingRepository)), wire.Bind(new(dao.ISettingDao), new(*dao.SettingDao)))