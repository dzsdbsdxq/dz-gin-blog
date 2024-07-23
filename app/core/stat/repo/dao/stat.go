package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/model"
	"gorm.io/gorm"
)

type IStatDao interface {
	Create(sts []*model.SysStat) error
	Update(st *model.SysStat) error
	Select() (st []*model.SysStat, err error)
	FindByKey(key string) (st *model.SysStat, err error)
}
type StatDao struct {
	coll *gorm.DB
}

func (s *StatDao) Update(st *model.SysStat) error {
	return s.coll.Updates(st).Error
}

func (s *StatDao) Select() (st []*model.SysStat, err error) {
	err = s.coll.Find(&st).Error
	return
}

func (s *StatDao) FindByKey(key string) (st *model.SysStat, err error) {
	err = s.coll.Where("key = ?", key).First(&st).Error
	return
}

func (s *StatDao) Create(sts []*model.SysStat) error {
	return s.coll.Create(sts).Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)").Error
}

var _ IStatDao = (*StatDao)(nil)

func NewStatDao(db *gorm.DB) *StatDao {
	return &StatDao{
		coll: db,
	}
}
