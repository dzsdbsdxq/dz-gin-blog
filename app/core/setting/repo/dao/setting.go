package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/model"
	"gorm.io/gorm"
)

type ISettingDao interface {
	Create(st *model.SysSetting) error
	Update(st *model.SysSetting) error
}
type SettingDao struct {
	coll *gorm.DB
}

func (s *SettingDao) Create(st *model.SysSetting) error {
	return s.coll.Create(st).Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)").Error
}

func (s *SettingDao) Update(st *model.SysSetting) error {
	return s.coll.Updates(st).Error
}

var _ ISettingDao = (*SettingDao)(nil)

func NewSettingDao(db *gorm.DB) *SettingDao {
	return &SettingDao{
		coll: db,
	}
}
