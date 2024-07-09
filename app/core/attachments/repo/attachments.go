package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/vo"
)

type IAttachmentsRepository interface {
	AdminCreateAttachments(attachments *vo.AttachmentsReq) error
	QueryAdminPostsPage(req *vo.AttachmentsGetReq) (attachments []*model.SysAttachMents, total int64, err error)
}

var _ IAttachmentsRepository = (*AttachmentsRepository)(nil)

type AttachmentsRepository struct {
	dao dao.IAttachmentsDao
}

func (a *AttachmentsRepository) AdminCreateAttachments(attachments *vo.AttachmentsReq) error {
	return a.dao.Create(&model.SysAttachMents{
		FileSize:   attachments.FileSize,
		FileWidth:  attachments.FileWidth,
		FileHeight: attachments.FileHeight,
		FileName:   attachments.FileName,
		FileMd5:    attachments.FileMd5,
		FilePath:   attachments.FilePath,
		FileExt:    attachments.FileExt,
		FileType:   attachments.FileType,
		Position:   "local",
	})
}
func (a *AttachmentsRepository) QueryAdminPostsPage(req *vo.AttachmentsGetReq) (attachments []*model.SysAttachMents, total int64, err error) {
	return a.dao.Get(req)
}

func NewAttachmentsRepository(dao dao.IAttachmentsDao) *AttachmentsRepository {
	return &AttachmentsRepository{
		dao: dao,
	}
}
