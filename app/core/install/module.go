package install

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install/service"
)

type (
	Handler = controller.InstallHandler
	Service = service.IInstallService
	Module  struct {
		Hdl *Handler
		Svc Service
	}
)
