package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysTags struct {
	global.Model
	Name  string `json:"name" gorm:"comment:名称"`
	Slug  string `json:"slug" gorm:"comment:别名"`
	Thumb string `json:"thumb" gorm:"comment:缩略图"`
	Nums  int    `json:"nums" gorm:"comment:文章数"`
}
