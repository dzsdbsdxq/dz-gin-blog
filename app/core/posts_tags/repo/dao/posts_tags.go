package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type IPostTagDao interface {
	Create(pts []*model.SysPostTag) error
	Delete(pts []*model.SysPostTag) error
}

var _ IPostTagDao = (*PostTagDao)(nil)

func NewPostTagDao() *PostTagDao {
	return &PostTagDao{
		coll: global.G_DZ_DB,
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
