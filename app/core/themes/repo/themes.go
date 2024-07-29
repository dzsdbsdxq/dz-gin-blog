package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/repo/dao"
)

type IThemesRepository interface {
}

var _ IThemesRepository = (*ThemesRepository)(nil)

type ThemesRepository struct {
	dao dao.IThemesDao
}

func NewThemesRepository(dao dao.IThemesDao) *ThemesRepository {
	return &ThemesRepository{
		dao: dao,
	}
}
