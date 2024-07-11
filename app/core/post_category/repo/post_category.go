package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/repo/dao"
)

type IPostCategoryRepository interface {
	CreatePostCategory(pcs []*model.SysPostCategory) error
	DeletePostCategory(pcs []*model.SysPostCategory) error
	SelectPostCategory(category uint) (pcs []*model.SysPostCategory, err error)
	FindPostCategory(category uint, postId uint) (pc *model.SysPostCategory, err error)
}

var _ IPostCategoryRepository = (*PostCategoryRepository)(nil)

func NewPostCategoryRepository(dao dao.IPostCategoryDao) *PostCategoryRepository {
	return &PostCategoryRepository{
		dao: dao,
	}
}

type PostCategoryRepository struct {
	dao dao.IPostCategoryDao
}

func (p *PostCategoryRepository) DeletePostCategory(pcs []*model.SysPostCategory) error {
	return p.dao.Delete(pcs)
}

func (p *PostCategoryRepository) CreatePostCategory(pcs []*model.SysPostCategory) error {
	return p.dao.Create(pcs)
}

func (p *PostCategoryRepository) SelectPostCategory(category uint) (pcs []*model.SysPostCategory, err error) {
	return p.dao.SelectPostIdsByCid(category)
}

func (p *PostCategoryRepository) FindPostCategory(category uint, postId uint) (pc *model.SysPostCategory, err error) {
	return p.dao.Find(category, postId)
}
