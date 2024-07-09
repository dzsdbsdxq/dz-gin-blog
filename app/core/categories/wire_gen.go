// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package categories

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitCategoriesModule() *Module {
	categoriesDao := dao.NewCategoriesDao()
	categoriesRepository := repo.NewCategoriesRepository(categoriesDao)
	categoryService := service.NewCategoryService(categoriesRepository)
	categoryHandler := controller.NewCategoryHandler(categoryService)
	module := &Module{
		Svc: categoryService,
		Hdl: categoryHandler,
	}
	return module
}

// wire.go:

var CategoryProviders = wire.NewSet(controller.NewCategoryHandler, service.NewCategoryService, repo.NewCategoriesRepository, dao.NewCategoriesDao, wire.Bind(new(service.ICategoryService), new(*service.CategoryService)), wire.Bind(new(repo.ICategoriesRepository), new(*repo.CategoriesRepository)), wire.Bind(new(dao.ICategoriesDao), new(*dao.CategoriesDao)))
