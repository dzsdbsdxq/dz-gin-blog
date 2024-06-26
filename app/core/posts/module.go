package posts

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/service"
)

type (
	Handler = controller.PostHandler
	Service = service.IPostService
	Module  struct {
		Svc Service
		Hdl *Handler
	}
)
