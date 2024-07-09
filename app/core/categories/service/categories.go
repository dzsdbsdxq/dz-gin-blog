package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
)

type ICategoryService interface {
	AdminCreateCategory(req *vo.CategoryReq) error
}

var _ ICategoryService = (*CategoryService)(nil)

type CategoryService struct {
	repo repo.ICategoriesRepository
}

func (c *CategoryService) AdminCreateCategory(req *vo.CategoryReq) error {
	return c.repo.CreateCategory(req)
}

func NewCategoryService(repo repo.ICategoriesRepository) *CategoryService {
	return &CategoryService{repo: repo}
}
