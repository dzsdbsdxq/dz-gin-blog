package comments

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/service"
)

type (
	Handler      = controller.CommentsHandler
	Service      = service.ICommentService
	CommentModel = model.SysComments
	Module       struct {
		Svc Service
		Hdl *Handler
	}
)
