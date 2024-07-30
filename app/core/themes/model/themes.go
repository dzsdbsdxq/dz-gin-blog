package model

type SysThemes struct {
	ID   uint64 `gorm:"primaryKey;column:id;type:bigint(20);not null;autoIncrement;" json:"ID"` //主键ID
	Key  string `gorm:"column:key;type:varchar(255);not null;default:'';comment:key"`
	Val  string `gorm:"column:val;type:varchar(255);not null;default:'';comment:val"`
	Type string `gorm:"column:type;type:varchar(255);not null;default:'';comment:主题名称"`
}
