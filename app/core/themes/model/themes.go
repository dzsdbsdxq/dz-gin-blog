package model

type SysThemes struct {
	ID   uint64 `gorm:"primaryKey;column:id;type:bigint(20);not null;autoIncrement;" json:"ID"` //主键ID
	Key  string `gorm:"column:key;type:varchar(255);not null;default:'';comment:key"`
	Val  string `gorm:"column:val;type:varchar(255);not null;default:'';comment:val"`
	Type string `gorm:"column:type;type:varchar(255);not null;default:'';comment:主题名称"`
	Name string `gorm:"column:name;type:varchar(255);not null;default:'';comment:标题"`
	Desc string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述"`
	Hook string `gorm:"column:hook;type:varchar(10);not null;default:'';comment:插入点（文章：post,页面：page，全局：global）"`
}
