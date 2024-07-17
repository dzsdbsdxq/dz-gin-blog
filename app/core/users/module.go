package users

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/service"
)

type (
	Handler   = controller.UserHandler
	Service   = service.IUserService
	UserModel = model.SysUsers
	Module    struct {
		Svc Service
		Hdl *Handler
	}
)
