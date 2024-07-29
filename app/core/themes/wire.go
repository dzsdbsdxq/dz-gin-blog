//go:build wireinject

package themes

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ThemeProviders = wire.NewSet(
	controller.NewThemesHandler,
	service.NewThemesService,
	repo.NewThemesRepository,
	dao.NewThemesDao,
	wire.Bind(new(service.IThemesService), new(*service.ThemesService)),
	wire.Bind(new(repo.IThemesRepository), new(*repo.ThemesRepository)),
	wire.Bind(new(dao.IThemesDao), new(*dao.ThemesDao)),
)

func InitThemesModule(db *gorm.DB) *Module {
	panic(wire.Build(
		ThemeProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
