package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"path/filepath"
	"time"
)

type IAttachmentsService interface {
	AdminCreateAttachments(post *vo.AttachmentsReq) error
	AdminGetAttachments(req *vo.AttachmentsGetReq) ([]*vo.AttachmentsRes, int64, error)
}

var _ IAttachmentsService = (*AttachmentsService)(nil)

type AttachmentsService struct {
	repo repo.IAttachmentsRepository
	//cfgService website_config.Service
	//eventBus   *eventbus.EventBus
}

func (a *AttachmentsService) AdminCreateAttachments(attach *vo.AttachmentsReq) error {
	return a.repo.AdminCreateAttachments(attach)
}
func (a *AttachmentsService) AdminGetAttachments(req *vo.AttachmentsGetReq) ([]*vo.AttachmentsRes, int64, error) {
	lists, total, err := a.repo.QueryAdminPostsPage(req)
	if err != nil {
		return nil, total, err
	}
	res := make([]*vo.AttachmentsRes, 0)

	for _, list := range lists {
		res = append(res, &vo.AttachmentsRes{
			Url:       filepath.Join(global.G_DZ_CONFIG.Local.BaseUrl, list.FilePath),
			Path:      list.FilePath,
			Driver:    "local",
			CreatedAt: list.CreatedAt.Format(time.DateTime),
			Name:      list.FileName,
			Key:       list.FileMd5,
			FileType:  list.FileType,
		})
	}
	return res, total, nil
}

func NewAttachmentsService(repo repo.IAttachmentsRepository) *AttachmentsService {
	s := &AttachmentsService{
		repo: repo,
		//cfgService: cfgService,
		//eventBus:   eventBus,
	}
	//go s.subscribeCommentEvent()
	return s
}
