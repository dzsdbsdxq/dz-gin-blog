package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type ICommentsRepository interface {
	CreateComments(comment *model.SysComments) error
	SelectComments(req vo.GetCommentsListReq) ([]*model.SysComments, uint64, error)
	DeleteComments(ids global.IdsReq) error
}
type CommentsRepository struct {
	dao dao.ICommentsDao
}

func (c *CommentsRepository) CreateComments(comment *model.SysComments) error {
	return c.dao.Create(comment)
}

func (c *CommentsRepository) SelectComments(req vo.GetCommentsListReq) ([]*model.SysComments, uint64, error) {
	return c.dao.GetCommentsLists(req)
}

func (c *CommentsRepository) DeleteComments(ids global.IdsReq) error {
	return c.dao.Delete(ids)
}

var _ ICommentsRepository = (*CommentsRepository)(nil)

func NewCommentsRepository(dao dao.ICommentsDao) *CommentsRepository {
	return &CommentsRepository{
		dao: dao,
	}
}
