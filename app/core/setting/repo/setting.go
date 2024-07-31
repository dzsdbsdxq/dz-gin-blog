package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/repo/dao"
)

type ISettingRepository interface {
	Create(setting []*model.SysSetting) error
	Update(setting *model.SysSetting) error
	UpdateValByKey(key string, val string) error
	Select() (sts []*model.SysSetting, err error)
	FindByKey(key string) (st *model.SysSetting, err error)
}

var _ ISettingRepository = (*SettingRepository)(nil)

func NewSettingRepository(dao dao.ISettingDao) *SettingRepository {
	return &SettingRepository{
		dao: dao,
	}
}

type SettingRepository struct {
	dao dao.ISettingDao
}

func (s *SettingRepository) UpdateValByKey(key string, val string) error {
	return s.dao.UpdateValByKey(key, val)
}

func (s *SettingRepository) FindByKey(key string) (st *model.SysSetting, err error) {
	return s.dao.FindByKey(key)
}

func (s *SettingRepository) Select() (sts []*model.SysSetting, err error) {
	return s.dao.Select()
}

func (s *SettingRepository) Create(setting []*model.SysSetting) error {
	return s.dao.Create(setting)
}

func (s *SettingRepository) Update(setting *model.SysSetting) error {
	return s.dao.Update(setting)
}
