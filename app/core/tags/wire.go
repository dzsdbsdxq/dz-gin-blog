//go:build wireinject

package tags

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/service"
	"github.com/google/wire"
)

var TagsProviders = wire.NewSet(
	controller.NewTagsHandler,
	service.NewTagService,
	repo.NewTagRepository,
	dao.NewTagsDao,
	wire.Bind(new(service.ITagService), new(*service.TagService)),
	wire.Bind(new(repo.ITagsRepository), new(*repo.TagsRepository)),
	wire.Bind(new(dao.ITagsDao), new(*dao.TagsDao)),
)

func InitTagsModule() *Module {
	panic(wire.Build(
		TagsProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
