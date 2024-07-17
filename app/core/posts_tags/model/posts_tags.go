package model

type SysPostTag struct {
	PostId uint `json:"post_id" gorm:"column:post_id;primaryKey;type:bigint(20);not null;default:0;comment:文章ID"`
	TagId  uint `json:"tag_id" gorm:"column:tag_id;primaryKey;type:bigint(20);not null;default:0;comment:标签ID"`
}
