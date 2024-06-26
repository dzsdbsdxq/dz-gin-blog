//go:build wireinject

package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users"
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initializeApp() (*gin.Engine, error) {
	panic(wire.Build(
		initialize.NewGinEngine,
		posts.InitPostModule,
		wire.FieldsOf(new(*posts.Module), "Hdl"),
		users.InitUsersModule,
		wire.FieldsOf(new(*users.Module), "Hdl"),
	))
}
