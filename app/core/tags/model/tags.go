package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysTags struct {
	global.Model
	Name  string `gorm:"column:name;type:varchar(255);not null;default:'';comment:名称"`
	Slug  string `gorm:"column:slug;type:varchar(255);not null;default:'';index:key_slug;comment:别名"`
	Thumb string `gorm:"column:thumb;type:varchar(255);not null;default:'';comment:缩略图"`
	Nums  int    `gorm:"column:nums;type:int(10);not null;default:0;comment:文章数"`
}
