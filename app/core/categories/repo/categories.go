package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
)

type ICategoriesRepository interface {
	CreateCategory(category *model.SysCategories) error
	UpdateCategory(category *model.SysCategories) error
	DeleteCategory(id uint64) error
	GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error)
}

var _ ICategoriesRepository = (*CategoriesRepository)(nil)

type CategoriesRepository struct {
	dao dao.ICategoriesDao
}

func (c *CategoriesRepository) CreateCategory(category *model.SysCategories) error {
	return c.dao.Create(category)
}

func (c *CategoriesRepository) GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error) {
	return c.dao.GetMenuTree(parentId, isTree)
}
func (c *CategoriesRepository) UpdateCategory(category *model.SysCategories) error {
	return c.dao.Update(category)
}

func (c *CategoriesRepository) DeleteCategory(id uint64) error {
	return c.dao.DeleteCategoriesById(id)
}

func NewCategoriesRepository(dao dao.ICategoriesDao) *CategoriesRepository {
	return &CategoriesRepository{
		dao: dao,
	}
}
