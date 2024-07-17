package posts

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/service"
)

type (
	Handler   = controller.PostHandler
	Service   = service.IPostService
	PostModel = model.SysPosts
	Module    struct {
		Svc Service
		Hdl *Handler
	}
)
