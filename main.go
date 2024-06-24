package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
)

func main() {
	initialize.InitViper()
	initialize.InitZapLogger()
	initialize.InitMysql()
	initialize.InitCache()
	global.G_DZ_LOG.Info("系统完成!")
}
