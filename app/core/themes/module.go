package themes

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/service"
)

type (
	Handler    = controller.ThemesHandler
	Service    = service.IThemesService
	ThemeModel = model.SysThemes
	Module     struct {
		Svc Service
		Hdl *Handler
	}
)
