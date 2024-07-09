package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	serv service.ICategoryService
}

func NewCategoryHandler(serv service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		serv: serv,
		//cfgService:   cfgService,
		//postLikeServ: postLikeServ,
		//eventBus:     eventBus,
	}
}

func (c *CategoryHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/categories")
	adminGroup.POST("/create", global.WrapWithBody(c.adminCreateCategories))
}

func (c *CategoryHandler) adminCreateCategories(ctx *gin.Context, req vo.CategoryReq) (*global.ResponseBody[any], error) {
	return global.SuccessResponse(), c.serv.AdminCreateCategory(&req)
}
