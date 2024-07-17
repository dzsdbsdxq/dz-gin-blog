package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysStat struct {
	global.Model
	Key string `gorm:"column:key;type:varchar(255);not null;default:'';comment:统计字段"`
	Val uint   `gorm:"column:val;type:bigint(20);not null;default:0;comment:统计值"`
}
