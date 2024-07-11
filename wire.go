//go:build wireinject

package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users"
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initializeApp() (*gin.Engine, error) {
	panic(wire.Build(
		initialize.NewGinEngine,
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
	))
}
