package attachments

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/vo"
)

type (
	Handle   = controller.AttachmentsHandle
	Service  = service.IAttachmentsService
	AttachVO = vo.AttachmentsReq
	Module   struct {
		Svc Service
		Hdl *Handle
	}
)
