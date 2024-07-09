//go:build wireinject

package categories

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/service"
	"github.com/google/wire"
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

func InitCategoriesModule() *Module {
	panic(wire.Build(
		CategoryProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
