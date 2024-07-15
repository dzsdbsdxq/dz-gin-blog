//go:build wireinject

package comments

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var CommentProviders = wire.NewSet(
	controller.NewCommentsHandler,
	service.NewCommentService,
	repo.NewCommentsRepository,
	dao.NewCommentsDao,
	wire.Bind(new(service.ICommentService), new(*service.CommentService)),
	wire.Bind(new(repo.ICommentsRepository), new(*repo.CommentsRepository)),
	wire.Bind(new(dao.ICommentsDao), new(*dao.CommentsDao)),
)

func InitCommentsModule(db *gorm.DB, postsModule *posts.Module) *Module {
	panic(wire.Build(
		CommentProviders,
		wire.FieldsOf(new(*posts.Module), "Svc"),
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
