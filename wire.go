//go:build wireinject

package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users"
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initializeApp() (*gin.Engine, error) {
	panic(wire.Build(
		initialize.NewGinEngine,
		initialize.NewMysql,
		posts.InitPostModule,
		oss.InitOssModule,
		post_category.InitPostCategoriesModule,
		wire.FieldsOf(new(*posts.Module), "Hdl"),
		users.InitUsersModule,
		wire.FieldsOf(new(*users.Module), "Hdl"),
		attachments.InitAttachmentsModule,
		wire.FieldsOf(new(*attachments.Module), "Hdl"),
		categories.InitCategoriesModule,
		wire.FieldsOf(new(*categories.Module), "Hdl"),
		tags.InitTagsModule,
		wire.FieldsOf(new(*tags.Module), "Hdl"),
		comments.InitCommentsModule,
		wire.FieldsOf(new(*comments.Module), "Hdl"),
		logs.InitLogsModule,
		wire.FieldsOf(new(*logs.Module), "Hdl"),
		setting.InitSettingModule,
		wire.FieldsOf(new(*setting.Module), "Hdl"),
	))
}
