//go:build wireinject

package website

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/website/controller"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var WebProviders = wire.NewSet(
	controller.NewWebsiteHandler,
)

func InitWebModule(db *gorm.DB, statModule *stat.Module) *Module {
	panic(wire.Build(
		WebProviders,
		wire.Struct(new(Module), "Hdl"),
		wire.FieldsOf(new(*stat.Module), "Svc"),
	))
}
