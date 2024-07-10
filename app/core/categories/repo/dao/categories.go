package dao

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type ICategoriesDao interface {
	Create(category *model.SysCategories) error
	Update(category *model.SysCategories) error
	GetCategoryBySlug(slug string) (sysCategory *model.SysCategories, err error)
	GetCategoryDetailById(id uint) (sysCategory *model.SysCategories, err error)
	DeleteCategoriesByIds(ids global.IdsReq) error
	getChildrenList(menu *vo.MenuRes, treeMap vo.MenuTreeMap) (err error)
	getMenuTreeMap() (treeMap vo.MenuTreeMap, err error)
	GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error)
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

func (c *CategoriesDao) getChildrenList(menu *vo.MenuRes, treeMap vo.MenuTreeMap) (err error) {
	idx := fmt.Sprintf("%d", menu.ID)
	if value, ok := treeMap[idx]; ok {
		menu.Children = value
		for i := 0; i < len(menu.Children); i++ {
			err = c.getChildrenList(&menu.Children[i], treeMap)
		}
	}
	return err
}

func (c *CategoriesDao) getMenuTreeMap() (treeMap vo.MenuTreeMap, err error) {
	var menuTree []*model.SysCategories
	treeMap = make(vo.MenuTreeMap, 0)
	err = c.coll.Order("sort_by DESC").Find(&menuTree).Error
	if err != nil {
		return nil, err
	}
	for _, v := range menuTree {
		pit := fmt.Sprintf("%d", v.ParentId)
		treeMap[pit] = append(treeMap[pit], vo.MenuRes{
			ID:       v.ID,
			ParentId: v.ParentId,
			SortBy:   v.SortBy,
			Name:     v.Name,
			Slug:     v.Slug,
			Thumb:    v.Thumb,
			Password: v.Password,
			Desc:     v.Desc,
			Url:      fmt.Sprintf("%s/categories/%s%s", global.G_DZ_CONFIG.System.Domain, v.Slug, global.G_DZ_CONFIG.System.UrlSuffix),
			Children: nil,
		})
	}
	return
}

func (c *CategoriesDao) GetMenuTree(parentId uint, isTree bool) (menus []vo.MenuRes, err error) {
	idx := fmt.Sprintf("%d", parentId)
	menuTree, err := c.getMenuTreeMap()
	menus = menuTree[idx]
	if isTree {
		for i := 0; i < len(menus); i++ {
			err = c.getChildrenList(&menus[i], menuTree)
		}
	}
	return menus, err
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
