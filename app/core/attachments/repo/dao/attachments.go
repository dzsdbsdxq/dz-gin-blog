package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/vo"
	"gorm.io/gorm"
)

type IAttachmentsDao interface {
	Create(attachments *model.SysAttachMents) error
	Get(req *vo.AttachmentsGetReq) (attachments []*model.SysAttachMents, total int64, err error)
}

var _ IAttachmentsDao = (*AttachmentsDao)(nil)

func NewAttachmentsDao(db *gorm.DB) *AttachmentsDao {
	return &AttachmentsDao{
		coll: db,
	}
}

type AttachmentsDao struct {
	coll *gorm.DB
}

func (a *AttachmentsDao) Create(attachments *model.SysAttachMents) error {
	return a.coll.Create(attachments).Error
}

func (a *AttachmentsDao) Get(req *vo.AttachmentsGetReq) (attachments []*model.SysAttachMents, total int64, err error) {
	offset := req.PageSize * (req.PageNo - 1)
	db := a.coll.Model(&model.SysAttachMents{})
	if req.Keyword != "" {
		db = db.Where("file_name LIKE ?", "%"+req.Keyword+"%")
	}
	if req.FileType != "" {
		db = db.Where("file_type = ?", req.FileType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return attachments, total, err
	} else {
		db = db.Limit(int(req.PageSize)).Offset(int(offset))
		err = db.Order("created_at DESC").Find(&attachments).Error
	}
	return attachments, total, err
}
