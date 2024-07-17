package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

//ALTER TABLE `db_blog`.`sys_setting`
//ADD COLUMN `comp` varchar(20) NOT NULL COMMENT '组件类型';

type SysSetting struct {
	global.Model
	Required int    `gorm:"column:required;type:int(1);not null;default:1;comment:是否必须（1=否，2=必须）"`
	Key      string `gorm:"column:key;type:varchar(255);not null;default:'';comment:key"`
	Val      string `gorm:"column:val;type:varchar(255);not null;default:'';comment:val"`
	Name     string `gorm:"column:name;type:varchar(255);not null;default:'';comment:标题"`
	Desc     string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述"`
	Type     string `gorm:"column:type;type:varchar(10);not null;default:'normal';comment:分组"`
	Comp     string `gorm:"column:comp;type:varchar(20);not null;default:'text';comment:组件类型"`
}
