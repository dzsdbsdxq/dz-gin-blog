package website

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/website/controller"
)

type (
	Handler = controller.WebSiteHandler
	Module  struct {
		Hdl *Handler
	}
)
