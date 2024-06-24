//go:build wireinject

package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initializeApp() (*gin.Engine, error) {
	panic(wire.Build(
		initialize.InitViper,
		initialize.InitZapLogger,
		initialize.InitMysql,
		initialize.InitCache,
	))
}
