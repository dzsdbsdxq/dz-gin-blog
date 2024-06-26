package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/vo"
)

type IPostCategoryRepository interface {
	AdminCreatePostCategory(pcs []*vo.PostCategoryReq) error
	AdminDeletePostCategory(pcs []*vo.PostCategoryReq) error
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

func (p *PostCategoryRepository) AdminDeletePostCategory(pcs []*vo.PostCategoryReq) error {
	sysPostCategorys := make([]*model.SysPostCategory, len(pcs))
	for _, pc := range pcs {
		sysPostCategorys = append(sysPostCategorys, &model.SysPostCategory{
			PostId:     pc.PostId,
			CategoryId: pc.Category,
		})
	}
	return p.dao.Delete(sysPostCategorys)
}

func (p *PostCategoryRepository) AdminCreatePostCategory(pcs []*vo.PostCategoryReq) error {
	sysPostCategorys := make([]*model.SysPostCategory, len(pcs))
	for _, pc := range pcs {
		sysPostCategorys = append(sysPostCategorys, &model.SysPostCategory{
			PostId:     pc.PostId,
			CategoryId: pc.Category,
		})
	}
	return p.dao.Create(sysPostCategorys)
}
