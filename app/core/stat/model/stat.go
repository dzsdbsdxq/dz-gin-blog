package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysStat struct {
	global.Model
	Key string `json:"key" gorm:"comment:统计字段"`
	Val uint   `json:"val" gorm:"comment:统计值"`
}
