package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type ICategoriesRepository interface {
	CreateCategory(req *vo.CategoryReq) error
	UpdateCategory(id uint64, req *vo.CategoryReq) error
	DeleteCategory(id uint64) error
	GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error)
}

var _ ICategoriesRepository = (*CategoriesRepository)(nil)

type CategoriesRepository struct {
	dao dao.ICategoriesDao
}

func (c *CategoriesRepository) CreateCategory(req *vo.CategoryReq) error {
	return c.dao.Create(&model.SysCategories{
		ParentId: req.ParentId,
		SortBy:   req.SortBy,
		Name:     req.Name,
		Slug:     req.Slug,
		Thumb:    req.Thumb,
		Password: req.Password,
		Desc:     req.Desc,
	})
}

func (c *CategoriesRepository) GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error) {
	return c.dao.GetMenuTree(parentId, isTree)
}
func (c *CategoriesRepository) UpdateCategory(id uint64, req *vo.CategoryReq) error {
	return c.dao.Update(&model.SysCategories{
		Model: global.Model{
			ID: id,
		},
		ParentId: req.ParentId,
		SortBy:   req.SortBy,
		Name:     req.Name,
		Slug:     req.Slug,
		Thumb:    req.Thumb,
		Password: req.Password,
		Desc:     req.Desc,
	})
}

func (c *CategoriesRepository) DeleteCategory(id uint64) error {
	return c.dao.DeleteCategoriesById(id)
}

func NewCategoriesRepository(dao dao.ICategoriesDao) *CategoriesRepository {
	return &CategoriesRepository{
		dao: dao,
	}
}
