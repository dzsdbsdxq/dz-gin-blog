package initialize

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

// installMdl *install.Service,
func NewInitServer(settingServ *setting.Serv) error {
	//err := installMdl.RegisterTables()
	//if err != nil {
	//	global.G_DZ_LOG.Error("初始化数据库错误:", err.Error())
	//}
	err := settingServ.Create()
	if err != nil {
		global.G_DZ_LOG.Error("初始化数据库设置表错误:", err.Error())
	}
	return err

}
