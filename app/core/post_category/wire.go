//go:build wireinject

package post_category

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var PostCategoryProviders = wire.NewSet(
	service.NewPostCategoryService,
	repo.NewPostCategoryRepository,
	dao.NewPostCategoryDao,
	wire.Bind(new(service.IPostCategoryService), new(*service.PostCategoryService)),
	wire.Bind(new(repo.IPostCategoryRepository), new(*repo.PostCategoryRepository)),
	wire.Bind(new(dao.IPostCategoryDao), new(*dao.PostCategoryDao)),
)

func InitPostCategoriesModule(db *gorm.DB) *Module {
	panic(wire.Build(
		PostCategoryProviders,
		wire.Struct(new(Module), "Svc"),
	))
}
