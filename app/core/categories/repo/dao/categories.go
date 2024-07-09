package dao

import "github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/model"

type ICategoriesDao interface {
	Create(post *model.SysCategories) error
}
