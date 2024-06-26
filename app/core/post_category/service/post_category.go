package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/vo"
)

type IPostCategoryService interface {
	AdminCreatePostCategory(pcs []*vo.PostCategoryReq) error
	AdminDeletePostCategory(pcs []*vo.PostCategoryReq) error
}

var _ IPostCategoryService = (*PostCategoryService)(nil)

type PostCategoryService struct {
	repo repo.IPostCategoryRepository
}

func (p *PostCategoryService) AdminDeletePostCategory(pcs []*vo.PostCategoryReq) error {
	return p.repo.AdminDeletePostCategory(pcs)
}

func (p *PostCategoryService) AdminCreatePostCategory(pcs []*vo.PostCategoryReq) error {
	return p.repo.AdminCreatePostCategory(pcs)
}

func NewPostCategoryService(repo repo.IPostCategoryRepository) *PostCategoryService {
	s := &PostCategoryService{
		repo: repo,
	}
	return s
}
