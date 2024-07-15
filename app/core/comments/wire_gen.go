// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitCommentsModule(db *gorm.DB, postsModule *posts.Module) *Module {
	commentsDao := dao.NewCommentsDao(db)
	commentsRepository := repo.NewCommentsRepository(commentsDao)
	commentService := service.NewCommentService(commentsRepository)
	iPostService := postsModule.Svc
	commentsHandler := controller.NewCommentsHandler(commentService, iPostService)
	module := &Module{
		Svc: commentService,
		Hdl: commentsHandler,
	}
	return module
}

// wire.go:

var CommentProviders = wire.NewSet(controller.NewCommentsHandler, service.NewCommentService, repo.NewCommentsRepository, dao.NewCommentsDao, wire.Bind(new(service.ICommentService), new(*service.CommentService)), wire.Bind(new(repo.ICommentsRepository), new(*repo.CommentsRepository)), wire.Bind(new(dao.ICommentsDao), new(*dao.CommentsDao)))