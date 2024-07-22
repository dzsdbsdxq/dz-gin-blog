package setting

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/service"
)

type (
	Handler       = controller.SettingHandler
	Serv          = service.SettingService
	Service       = service.ISettingService
	SettingsModel = model.SysSetting
	Module        struct {
		Svc Service
		Hdl *Handler
	}
)
