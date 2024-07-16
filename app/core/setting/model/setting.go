package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

//ALTER TABLE `db_blog`.`sys_setting`
//ADD COLUMN `comp` varchar(20) NOT NULL COMMENT '组件类型';

type SysSetting struct {
	global.Model
	Required int    `json:"required" gorm:"comment:是否必须（1=否，2=必须）"`
	Key      string `json:"key" gorm:"comment:key"`
	Val      string `json:"val" gorm:"comment:val"`
	Name     string `json:"name" gorm:"comment:标题"`
	Desc     string `json:"desc" gorm:"comment:描述"`
	Type     string `json:"type" gorm:"comment:分组"`
	Comp     string `json:"comp" gorm:"comment:组件类型"`
}
