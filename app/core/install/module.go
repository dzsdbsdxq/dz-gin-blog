package install

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install/service"
)

type (
	Service  = service.InstallService
	IService = service.IInstallService
	Module   struct {
		Svc  *Service
		ISvc IService
	}
)
