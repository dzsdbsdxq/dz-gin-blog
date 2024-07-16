package logs

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/vo"
)

type (
	Handler  = controller.LogsHandler
	Service  = service.ILogsService
	LogVO    = vo.LogsReq
	LogModel = model.SysLogs
	Module   struct {
		Svc Service
		Hdl *Handler
	}
)
