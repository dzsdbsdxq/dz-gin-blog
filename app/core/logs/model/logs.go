package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysLogs struct {
	global.Model
	Type    string `gorm:"column:type;type:varchar(60);not null;default:'';comment:类型"`
	Key     string `gorm:"column:key;type:varchar(60);not null;default:'';comment:关键词"`
	Content string `gorm:"column:content;type:varchar(255);not null;default:'';comment:内容"`
	Ip      string `gorm:"column:ip;type:varchar(20);not null;default:'';comment:客户端IP"`
	Creator string `gorm:"column:creator;type:varchar(60);not null;default:'';comment:操作者"`
}
