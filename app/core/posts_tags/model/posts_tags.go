package model

type SysPostTag struct {
	PostId uint `json:"post_id" gorm:"comment:文章ID"`
	TagId  uint `json:"tag_id" gorm:"comment:标签ID"`
}
