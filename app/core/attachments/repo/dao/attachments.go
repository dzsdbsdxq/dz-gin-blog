package dao

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type IAttachmentsDao interface {
	Create(attachments *model.SysAttachMents) error
}

var _ IAttachmentsDao = (*AttachmentsDao)(nil)

func NewAttachmentsDao() *AttachmentsDao {
	return &AttachmentsDao{
		coll: global.G_DZ_DB,
	}
}

type AttachmentsDao struct {
	coll *gorm.DB
}

func (a *AttachmentsDao) Create(attachments *model.SysAttachMents) error {
	return a.coll.Create(attachments).Error
}
