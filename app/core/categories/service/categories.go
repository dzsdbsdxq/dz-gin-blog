package service

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type ICategoryService interface {
	AdminCreateCategory(req *vo.CategoryReq) error
	AdminUpdateCategory(id uint64, req *vo.CategoryReq) error
	AdminDeleteCategory(id uint) error
	GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error)
}

var _ ICategoryService = (*CategoryService)(nil)

type CategoryService struct {
	repo      repo.ICategoriesRepository
	pcService post_category.Service
}

func (c *CategoryService) AdminCreateCategory(req *vo.CategoryReq) error {
	return c.repo.CreateCategory(req)
}
func (c *CategoryService) AdminUpdateCategory(id uint64, req *vo.CategoryReq) error {
	return c.repo.UpdateCategory(id, req)
}

func (c *CategoryService) GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error) {
	return c.repo.GetMenuTree(parentId, isTree)
}
func (c *CategoryService) AdminDeleteCategory(id uint) error {
	//查询分类下是否有文章，如果有文章，则不进行删除，如果没有，则删除
	category, _ := c.pcService.SelectPostCategory(id)
	if len(category) > 0 {
		return errors.New(fmt.Sprintf("%s", global.GetErrorMsg(402000013, "")))
	}
	//查询该分类下是否含有子分类，如果存在子分类，则不能进行删除，如果没有则删除
	tree, _ := c.GetMenuTree(id, false)
	if len(tree) > 0 {
		return errors.New(fmt.Sprintf("%s", global.GetErrorMsg(402000014, "")))
	}
	return c.repo.DeleteCategory(uint64(id))
}

func NewCategoryService(repo repo.ICategoriesRepository, pcService post_category.Service) *CategoryService {
	return &CategoryService{repo: repo, pcService: pcService}
}
