package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysLogs struct {
	global.Model
	Type    string `json:"type" gorm:"comment:类型"`
	Key     string `json:"key" gorm:"comment:关键词"`
	Content string `json:"content" gorm:"comment:内容"`
	Ip      string `json:"ip" gorm:"comment:客户端IP"`
	Creator string `json:"creator" gorm:"comment:操作者"`
}
