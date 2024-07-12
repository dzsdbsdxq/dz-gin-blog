//go:build wireinject

package posts

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var PostProviders = wire.NewSet(
	controller.NewPostHandler,
	service.NewPostService,
	repo.NewPostRepository,
	dao.NewPostDao,
	wire.Bind(new(service.IPostService), new(*service.PostService)),
	wire.Bind(new(repo.IPostRepository), new(*repo.PostRepository)),
	wire.Bind(new(dao.IPostDao), new(*dao.PostDao)),
)

func InitPostModule(db *gorm.DB) *Module {
	panic(wire.Build(
		PostProviders,
		//wire.FieldsOf(new(*website_config.Module), "Svc"),
		//wire.FieldsOf(new(*post_like.Module), "Svc"),
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
