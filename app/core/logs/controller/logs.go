package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	"github.com/gin-gonic/gin"
)

type LogsHandler struct {
	serv service.ILogsService
}

func NewLogsHandler(serv service.ILogsService) *LogsHandler {
	return &LogsHandler{
		serv: serv,
	}
}
func (l *LogsHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/logs")
	adminGroup.GET("/lists", global.WrapWithBody(l.AdminGetLogs))
}

func (l *LogsHandler) AdminGetLogs(ctx *gin.Context, req *vo.LogsGetReq) (*global.ResponseBody[global.PageRes[[]*vo.LogsRes]], error) {
	req.ValidateAndSetDefault()
	logs, total, err := l.serv.SelectLogs(req)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(global.PageRes[[]*vo.LogsRes]{
		PageInfo: global.PageInfo{
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Category: req.Category,
			Keyword:  req.Keyword,
		},
		TotalCount: total,
		TotalPages: utils.CalcPages(total, req.PageSize),
		List:       logs,
	}), nil
}
