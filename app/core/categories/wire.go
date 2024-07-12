//go:build wireinject

package categories

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var CategoryProviders = wire.NewSet(
	controller.NewCategoryHandler,
	service.NewCategoryService,
	repo.NewCategoriesRepository,
	dao.NewCategoriesDao,
	wire.Bind(new(service.ICategoryService), new(*service.CategoryService)),
	wire.Bind(new(repo.ICategoriesRepository), new(*repo.CategoriesRepository)),
	wire.Bind(new(dao.ICategoriesDao), new(*dao.CategoriesDao)),
)

func InitCategoriesModule(db *gorm.DB, pcModel *post_category.Module) *Module {
	panic(wire.Build(
		CategoryProviders,
		wire.FieldsOf(new(*post_category.Module), "Svc"),
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
