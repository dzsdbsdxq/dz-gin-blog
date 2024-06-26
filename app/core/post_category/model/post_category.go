package model

type SysPostCategory struct {
	PostId     uint `json:"post_id" gorm:"comment:文章ID"`
	CategoryId uint `json:"category_id" gorm:"comment:栏目ID"`
}
