package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
)

type ICategoriesRepository interface {
	CreateCategory(req *vo.CategoryReq) error
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

func NewCategoriesRepository(dao dao.ICategoriesDao) *CategoriesRepository {
	return &CategoriesRepository{
		dao: dao,
	}
}
