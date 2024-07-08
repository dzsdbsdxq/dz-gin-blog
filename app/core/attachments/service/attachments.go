package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/vo"
)

type IAttachmentsService interface {
	AdminCreateAttachments(post *vo.AttachmentsReq) error
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

func NewAttachmentsService(repo repo.IAttachmentsRepository) *AttachmentsService {
	s := &AttachmentsService{
		repo: repo,
		//cfgService: cfgService,
		//eventBus:   eventBus,
	}
	//go s.subscribeCommentEvent()
	return s
}
