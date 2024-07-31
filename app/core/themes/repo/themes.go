package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/repo/dao"
	"gorm.io/gorm"
)

type IThemesRepository interface {
	Create(themes []*model.SysThemes) error
	SelectByType(ty string) (themes []*model.SysThemes, err error)
	FindByTypeAndKey(ty string, key string) (theme *model.SysThemes, err error)
}

var _ IThemesRepository = (*ThemesRepository)(nil)

type ThemesRepository struct {
	dao dao.IThemesDao
}

func (t *ThemesRepository) Create(themes []*model.SysThemes) error {
	_, err := t.FindByTypeAndKey()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//如果key不存在，直接返回
		return nil
	}
	return t.dao.Create(themes)
}

func (t *ThemesRepository) SelectByType(ty string) (themes []*model.SysThemes, err error) {
	return t.dao.SelectByType(ty)
}

func (t *ThemesRepository) FindByTypeAndKey(ty string, key string) (theme *model.SysThemes, err error) {
	return t.dao.FindByTypeAndKey(ty, key)
}

func NewThemesRepository(dao dao.IThemesDao) *ThemesRepository {
	return &ThemesRepository{
		dao: dao,
	}
}
