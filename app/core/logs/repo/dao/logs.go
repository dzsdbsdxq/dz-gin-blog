package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs/vo"
	"gorm.io/gorm"
)

type ILogsDao interface {
	Create(log *model.SysLogs) error
	Get(req *vo.LogsGetReq) (logs []*model.SysLogs, total int64, err error)
}

var _ ILogsDao = (*LogsDao)(nil)

type LogsDao struct {
	coll *gorm.DB
}

func (l *LogsDao) Get(req *vo.LogsGetReq) (logs []*model.SysLogs, total int64, err error) {
	offset := req.PageSize * (req.PageNo - 1)
	db := l.coll.Model(&model.SysLogs{})
	if req.Keyword != "" {
		db = db.Where("file_name LIKE ?", "%"+req.Keyword+"%")
	}
	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return logs, total, err
	} else {
		db = db.Limit(int(req.PageSize)).Offset(int(offset))
		err = db.Order("created_at DESC").Find(&logs).Error
	}
	return logs, total, err
}

func (l *LogsDao) Create(log *model.SysLogs) error {
	return l.coll.Create(log).Error
}

func NewLogsDao(db *gorm.DB) *LogsDao {
	return &LogsDao{
		coll: db,
	}
}
