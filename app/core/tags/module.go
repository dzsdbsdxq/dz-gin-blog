package tags

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/service"
)

type (
	Handler = controller.TagsHandler
	Service = service.ITagService
	Module  struct {
		Svc Service
		Hdl *Handler
	}
)
