package dao

import (
	"errors"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/model"
	"gorm.io/gorm"
)

type ISettingDao interface {
	Create(sts []*model.SysSetting) error
	Update(st *model.SysSetting) error
	Select() (st []*model.SysSetting, err error)
	FindByKey(key string) (st *model.SysSetting, err error)
}
type SettingDao struct {
	coll *gorm.DB
}

func (s *SettingDao) FindByKey(key string) (st *model.SysSetting, err error) {
	err = s.coll.Where("`key` = ?", key).First(&st).Error
	return
}

func (s *SettingDao) Select() (st []*model.SysSetting, err error) {
	err = s.coll.Find(&st).Error
	return
}

func (s *SettingDao) Create(sts []*model.SysSetting) error {
	//查询是否存在相同的Key，如果存在相同的Key，则跳过，不进行创建
	insertSts := make([]*model.SysSetting, 0)
	for _, st := range sts {
		_, err := s.FindByKey(st.Key)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			insertSts = append(insertSts, st)
		}
	}
	if len(insertSts) > 0 {
		return s.coll.Create(insertSts).Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)").Error
	}
	return nil

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
