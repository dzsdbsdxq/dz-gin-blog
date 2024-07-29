//go:build wireinject

package stat

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var StatProviders = wire.NewSet(
	service.NewStatService,
	repo.NewStatRepository,
	dao.NewStatDao,
	wire.Bind(new(service.IStatService), new(*service.StatService)),
	wire.Bind(new(repo.IStatRepository), new(*repo.StatRepository)),
	wire.Bind(new(dao.IStatDao), new(*dao.StatDao)),
)

func InitStatModule(db *gorm.DB) *Module {
	panic(wire.Build(
		StatProviders,
		wire.Struct(new(Module), "Svc"),
	))
}
