package dao

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type ITagsDao interface {
	Create(tag *model.SysTags) error
	Update(tag *model.SysTags) error
	DeleteById(id uint64) error
	GetTagDetailById(id uint) (tag *model.SysTags, err error)
	GetTagBySlug(slug string) (tag *model.SysTags, err error)
	Get(pageSize int64, pageNo int64) (tags []*model.SysTags, total int64, err error)
}

var _ ITagsDao = (*TagsDao)(nil)

func NewTagsDao(db *gorm.DB) *TagsDao {
	return &TagsDao{
		coll: db,
	}
}

type TagsDao struct {
	coll *gorm.DB
}

func (t *TagsDao) DeleteById(id uint64) error {
	var tag model.SysTags
	return t.coll.Find(&tag, "id = ?", id).Delete(&tag).Error
}

func (t *TagsDao) Create(tag *model.SysTags) error {
	_, err := t.GetTagBySlug(tag.Slug)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("%s:%s", global.GetErrorMsg(403000009, ""), tag.Slug))
	}
	return t.coll.Create(tag).Error
}
func (t *TagsDao) Update(tag *model.SysTags) error {
	dict, err := t.GetTagDetailById(uint(tag.ID))

	if dict.Slug != tag.Slug {
		_, err = t.GetTagBySlug(tag.Slug)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("%s:%s", global.GetErrorMsg(403000009, ""), tag.Slug))
		}
	}

	return t.coll.Updates(tag).Error
}
func (t *TagsDao) GetTagBySlug(slug string) (tag *model.SysTags, err error) {
	err = t.coll.Where("slug = ?", slug).First(&tag).Error
	return
}
func (t *TagsDao) GetTagDetailById(id uint) (tag *model.SysTags, err error) {
	err = t.coll.Where("id = ?", id).First(&tag).Error
	return
}
func (t *TagsDao) Get(pageSize int64, pageNo int64) (tags []*model.SysTags, total int64, err error) {
	offset := pageSize * (pageNo - 1)
	db := t.coll.Model(&model.SysTags{})
	err = db.Count(&total).Error
	if err != nil {
		return tags, total, err
	} else {
		db = db.Limit(int(pageSize)).Offset(int(offset))
		err = db.Order("created_at DESC").Find(&tags).Error
	}
	return tags, total, err
}
