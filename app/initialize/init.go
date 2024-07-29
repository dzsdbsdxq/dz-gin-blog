package initialize

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

func NewInitServer(settingHdl *setting.Handler, installHdl *install.Handler) error {
	err := installHdl.SystemAutoReRegister()
	if err != nil {
		global.G_DZ_LOG.Error("初始化数据库错误:", err.Error())
	}
	err = settingHdl.SystemAutoRegister()
	if err != nil {
		global.G_DZ_LOG.Error("初始化数据库设置表错误:", err.Error())
	}
	return err

}
