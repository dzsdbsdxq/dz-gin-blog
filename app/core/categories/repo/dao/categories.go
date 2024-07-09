package dao

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type ICategoriesDao interface {
	Create(category *model.SysCategories) error
	Update(category *model.SysCategories) error
	GetCategoryBySlug(slug string) (sysCategory *model.SysCategories, err error)
	GetCategoryDetailById(id uint) (sysCategory *model.SysCategories, err error)
}

var _ ICategoriesDao = (*CategoriesDao)(nil)

func NewCategoriesDao() *CategoriesDao {
	return &CategoriesDao{
		coll: global.G_DZ_DB,
	}
}

type CategoriesDao struct {
	coll *gorm.DB
}

func (c *CategoriesDao) Create(category *model.SysCategories) error {
	_, err := c.GetCategoryBySlug(category.Slug)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("存在相同的别名:%s", category.Slug))
	}
	return c.coll.Create(category).Error
}

func (c *CategoriesDao) GetCategoryBySlug(slug string) (sysCategory *model.SysCategories, err error) {
	err = c.coll.Where("slug = ?", slug).First(&sysCategory).Error
	return
}
func (c *CategoriesDao) GetCategoryDetailById(id uint) (sysCategory *model.SysCategories, err error) {
	err = c.coll.Where("id = ?", id).First(&sysCategory).Error
	return
}
func (c *CategoriesDao) Update(category *model.SysCategories) error {
	_, err := c.GetCategoryBySlug(category.Slug)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("存在相同的别名:%s", category.Slug))
	}
	return c.coll.Save(category).Error
}

func (c *CategoriesDao) DeleteCategoriesByIds(ids global.IdsReq) error {
	var categories []model.SysCategories
	return c.coll.Find(categories, "id in ?", ids.Ids).Delete(&categories).Error
}
