//go:build wireinject

package logs

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var LogsProviders = wire.NewSet(
	controller.NewLogsHandler,
	service.NewLogsService,
	repo.NewLogsRepository,
	dao.NewLogsDao,
	wire.Bind(new(service.ILogsService), new(*service.LogService)),
	wire.Bind(new(repo.ILogsRepository), new(*repo.LogsRepository)),
	wire.Bind(new(dao.ILogsDao), new(*dao.LogsDao)),
)

func InitLogsModule(db *gorm.DB) *Module {
	panic(wire.Build(
		LogsProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
