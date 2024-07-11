package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/repo/dao"
)

type ITagsRepository interface {
	CreateTags(tag *model.SysTags) error
	UpdateTags(tag *model.SysTags) error
	DeleteTags(id uint64) error
	SelectTags(pageSize int64, pageNo int64) (tags []*model.SysTags, total int64, err error)
}

var _ ITagsRepository = (*TagsRepository)(nil)

type TagsRepository struct {
	dao dao.ITagsDao
}

func NewTagRepository(dao dao.ITagsDao) *TagsRepository {
	return &TagsRepository{
		dao: dao,
	}
}

func (t *TagsRepository) CreateTags(tag *model.SysTags) error {
	return t.dao.Create(tag)
}

func (t *TagsRepository) UpdateTags(tag *model.SysTags) error {
	return t.dao.Update(tag)
}

func (t *TagsRepository) DeleteTags(id uint64) error {
	return t.dao.DeleteById(id)
}
func (t *TagsRepository) SelectTags(pageSize int64, pageNo int64) (tags []*model.SysTags, total int64, err error) {
	return t.dao.Get(pageSize, pageNo)
}
