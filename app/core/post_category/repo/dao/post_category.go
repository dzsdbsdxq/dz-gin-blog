package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type IPostCategoryDao interface {
	Create(pcs []*model.SysPostCategory) error
	Delete(pcs []*model.SysPostCategory) error
	SelectPostIdsByCid(category uint) (pcs []*model.SysPostCategory, err error)
	Find(category uint, postId uint) (pc *model.SysPostCategory, err error)
}

var _ IPostCategoryDao = (*PostCategoryDao)(nil)

func NewPostCategoryDao() *PostCategoryDao {
	return &PostCategoryDao{
		coll: global.G_DZ_DB,
	}
}

type PostCategoryDao struct {
	coll *gorm.DB
}

func (p *PostCategoryDao) Create(pcs []*model.SysPostCategory) error {
	return p.coll.Create(pcs).Error
}

func (p *PostCategoryDao) Delete(pcs []*model.SysPostCategory) error {
	return p.coll.Unscoped().Delete(pcs).Error
}

func (p *PostCategoryDao) SelectPostIdsByCid(category uint) (pcs []*model.SysPostCategory, err error) {
	err = p.coll.Where("category_id = ?", category).Order("post_id DESC").Find(&pcs).Error
	return
}

func (p *PostCategoryDao) Find(category uint, postId uint) (pc *model.SysPostCategory, err error) {
	err = p.coll.Where("category_id = ? AND post_id = ?", category, postId).First(&pc).Error
	return
}
