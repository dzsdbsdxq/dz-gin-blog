package dao

import (
	"errors"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type IPostDao interface {
	Create(post *model.SysPosts) error
}

var _ IPostDao = (*PostDao)(nil)

func NewPostDao() *PostDao {
	return &PostDao{
		coll: global.G_DZ_DB.Model(&model.SysPosts{}),
	}
}

type PostDao struct {
	coll *gorm.DB
}

func (p *PostDao) Create(post *model.SysPosts) error {
	if !errors.Is(p.coll.Where("slug = ?", post.Slug).Find(&struct {
		ID uint
	}{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同别名")
	}
	return p.coll.Create(&post).Error
}
