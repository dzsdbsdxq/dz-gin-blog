// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users"
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initializeApp() (*gin.Engine, error) {
	module := posts.InitPostModule()
	postHandler := module.Hdl
	usersModule := users.InitUsersModule()
	userHandler := usersModule.Hdl
	ossModule := oss.InitOssModule()
	attachmentsModule := attachments.InitAttachmentsModule(ossModule)
	attachmentsHandler := attachmentsModule.Hdl
	categoriesModule := categories.InitCategoriesModule()
	categoryHandler := categoriesModule.Hdl
	engine, err := initialize.NewGinEngine(postHandler, userHandler, attachmentsHandler, categoryHandler)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
