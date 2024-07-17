package model

type SysPostCategory struct {
	PostId     uint `gorm:"column:post_id;primaryKey;type:bigint(20);not null;default:0;comment:文章ID"`
	CategoryId uint `gorm:"column:category_id;primaryKey;type:bigint(20);not null;default:0;comment:栏目ID"`
}
