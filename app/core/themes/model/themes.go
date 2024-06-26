package model

type SysThemes struct {
	ID   uint64 `gorm:"primaryKey" json:"ID"` //主键ID
	Key  string `json:"key" gorm:"comment:key"`
	Val  string `json:"val" gorm:"comment:val"`
	Type string `json:"type" gorm:"comment:主题名称"`
}
