package categories

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/service"
)

type (
	Handler = controller.CategoryHandler
	Service = service.ICategoryService
	Module  struct {
		Svc Service
		Hdl *Handler
	}
)
