package dao

import "gorm.io/gorm"

type IThemesDao interface{}

var _ IThemesDao = (*ThemesDao)(nil)

func NewThemesDao(db *gorm.DB) *ThemesDao {
	return &ThemesDao{
		coll: db,
	}
}

type ThemesDao struct {
	coll *gorm.DB
}
