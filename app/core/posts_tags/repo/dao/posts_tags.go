package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/model"
	"gorm.io/gorm"
)

type IPostTagDao interface {
	Create(pts []*model.SysPostTag) error
	Delete(pts []*model.SysPostTag) error
}

var _ IPostTagDao = (*PostTagDao)(nil)

func NewPostTagDao(db *gorm.DB) *PostTagDao {
	return &PostTagDao{
		coll: db,
	}
}

type PostTagDao struct {
	coll *gorm.DB
}

func (p *PostTagDao) Create(pts []*model.SysPostTag) error {
	return p.coll.Create(pts).Error
}

func (p *PostTagDao) Delete(pts []*model.SysPostTag) error {
	return p.coll.Unscoped().Delete(pts).Error
}
