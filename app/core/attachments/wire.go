//go:build wireinject

package attachments

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var AttachmentsProviders = wire.NewSet(
	controller.NewAttachmentsHandler,
	service.NewAttachmentsService,
	repo.NewAttachmentsRepository,
	dao.NewAttachmentsDao,
	wire.Bind(new(service.IAttachmentsService), new(*service.AttachmentsService)),
	wire.Bind(new(repo.IAttachmentsRepository), new(*repo.AttachmentsRepository)),
	wire.Bind(new(dao.IAttachmentsDao), new(*dao.AttachmentsDao)),
)

func InitAttachmentsModule(db *gorm.DB, ossMdl *oss.Module) *Module {
	panic(wire.Build(
		AttachmentsProviders,
		wire.FieldsOf(new(*oss.Module), "Svc"),
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
