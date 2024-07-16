package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/vo"
	"github.com/gin-gonic/gin"
)

type ILogsService interface {
	CreateLogs(ctx *gin.Context, req *vo.LogsReq) error
	SelectLogs(req *vo.LogsGetReq) ([]*vo.LogsRes, int64, error)
}

var _ ILogsService = (*LogService)(nil)

type LogService struct {
	repo repo.ILogsRepository
}

func (a *LogService) CreateLogs(ctx *gin.Context, req *vo.LogsReq) error {
	return a.repo.CreateLogs(&model.SysLogs{
		Type:    req.Type,
		Key:     req.Key,
		Content: req.Content,
		Ip:      ctx.ClientIP(),
		Creator: req.Creator,
	})
}
func (a *LogService) SelectLogs(req *vo.LogsGetReq) ([]*vo.LogsRes, int64, error) {
	lists, total, err := a.repo.QueryLogsPage(req)
	if err != nil {
		return nil, total, err
	}
	res := make([]*vo.LogsRes, 0)
	for _, list := range lists {
		res = append(res, &vo.LogsRes{
			ID:        list.ID,
			CreatedAt: list.CreatedAt,
			Type:      list.Type,
			Key:       list.Key,
			Content:   list.Content,
			Ip:        list.Ip,
			Creator:   list.Creator,
		})
	}
	return res, total, nil
}

func NewLogsService(repo repo.ILogsRepository) *LogService {
	return &LogService{
		repo: repo,
	}
}
