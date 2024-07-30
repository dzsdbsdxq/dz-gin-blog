package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/model"
	"gorm.io/gorm"
)

type IThemesDao interface {
	SelectByType(ty string) (themes []*model.SysThemes, err error)
	FindByTypeAndKey(ty string, key string) (theme *model.SysThemes, err error)
}

var _ IThemesDao = (*ThemesDao)(nil)

func NewThemesDao(db *gorm.DB) *ThemesDao {
	return &ThemesDao{
		coll: db,
	}
}

type ThemesDao struct {
	coll *gorm.DB
}

func (t *ThemesDao) SelectByType(ty string) (themes []*model.SysThemes, err error) {
	err = t.coll.Where("type = ?", ty).Find(&themes).Error
	return
}

func (t *ThemesDao) FindByTypeAndKey(ty string, key string) (theme *model.SysThemes, err error) {
	err = t.coll.Where("type = ? and key = ?", ty, key).Limit(1).First(&theme).Error
	return
}
