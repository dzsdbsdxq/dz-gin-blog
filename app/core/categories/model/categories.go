package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysCategories struct {
	global.Model
	ParentId uint   `gorm:"column:parent_id;type:bigint(20);not null;default:0;comment:上级目录"`
	SortBy   int    `gorm:"column:sort_by;type:int(10);not null;default:0;comment:排序"`
	Name     string `gorm:"column:name;type:varchar(255);not null;default:'';comment:名称"`
	Slug     string `gorm:"column:slug;type:varchar(255);not null;default:'';index:key_categories_slug;comment:别名"`
	Thumb    string `gorm:"column:thumb;type:varchar(255);not null;default:'';comment:封面图"`
	Password string `gorm:"column:password;type:varchar(255);not null;default:'';comment:密码"`
	Desc     string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述"`
}
