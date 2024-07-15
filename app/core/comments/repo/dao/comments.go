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
	GetCommentsLists(req vo.GetCommentsListReq) (comments []*model.SysComments, total int64, err error)
	FindById(id uint64) (comment *model.SysComments, err error)
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

func (c *CommentsDao) FindById(id uint64) (comment *model.SysComments, err error) {
	err = c.coll.Where("id = ?", id).First(&comment).Error
	return
}

func (c *CommentsDao) GetCommentsLists(req vo.GetCommentsListReq) (comments []*model.SysComments, total int64, err error) {
	offset := req.PageSize * (req.PageNo - 1)
	db := c.coll.Model(&model.SysComments{})
	if req.Keyword != "" {
		db = db.Where("file_name LIKE ?", "%"+req.Keyword+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return comments, total, err
	} else {
		db = db.Limit(int(req.PageSize)).Offset(int(offset))
		err = db.Order("created_at DESC").Find(&comments).Error
	}
	return comments, total, err
}

func (c *CommentsDao) Delete(ids global.IdsReq) error {
	return c.coll.Transaction(func(tx *gorm.DB) error {
		var comments []model.SysComments
		err := tx.Find(&comments, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]model.SysComments{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (c *CommentsDao) Create(comment *model.SysComments) error {
	return c.coll.Create(comment).Error
}

func (c *CommentsDao) Update(comment *model.SysComments) error {
	return c.coll.Updates(comment).Error
}
