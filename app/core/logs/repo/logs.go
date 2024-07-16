package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/vo"
)

type ILogsRepository interface {
	CreateLogs(log *model.SysLogs) error
	QueryLogsPage(req *vo.LogsGetReq) (logs []*model.SysLogs, total int64, err error)
}

var _ ILogsRepository = (*LogsRepository)(nil)

type LogsRepository struct {
	dao dao.ILogsDao
}

func (l *LogsRepository) CreateLogs(log *model.SysLogs) error {
	return l.dao.Create(log)
}

func (l *LogsRepository) QueryLogsPage(req *vo.LogsGetReq) (logs []*model.SysLogs, total int64, err error) {
	return l.dao.Get(req)
}

func NewLogsRepository(dao dao.ILogsDao) *LogsRepository {
	return &LogsRepository{
		dao: dao,
	}
}
