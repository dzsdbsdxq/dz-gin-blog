package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type ICommentsDao interface {
	Create(comment *model.SysComments) error
	Update(comment *model.SysComments) error
	Delete(ids global.IdsReq) error
	GetCommentsLists(req vo.GetCommentsListReq) ([]*model.SysComments, uint64, error)
}

var _ ICommentsDao = (*CommentsDao)(nil)

func NewCommentsDao(db *gorm.DB) *CommentsDao {
	return &CommentsDao{
		coll: db,
	}
}

type CommentsDao struct {
	coll *gorm.DB
}

func (c *CommentsDao) GetCommentsLists(req vo.GetCommentsListReq) ([]*model.SysComments, uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentsDao) Delete(ids global.IdsReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CommentsDao) Create(comment *model.SysComments) error {
	return c.coll.Create(comment).Error
}

func (c *CommentsDao) Update(comment *model.SysComments) error {
	//TODO implement me
	panic("implement me")
}
