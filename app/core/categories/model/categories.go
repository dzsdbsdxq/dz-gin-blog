package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysCategories struct {
	global.Model
	ParentId uint   `json:"parent_id" gorm:"comment:上级目录"`
	SortBy   int    `json:"sort_by" gorm:"comment:排序"`
	Name     string `json:"name" gorm:"comment:名称"`
	Slug     string `json:"slug" gorm:"comment:别名"`
	Thumb    string `json:"thumb" gorm:"comment:封面图"`
	Password string `json:"password" gorm:"comment:密码"`
	Desc     string `json:"desc" gorm:"comment:描述"`
}
