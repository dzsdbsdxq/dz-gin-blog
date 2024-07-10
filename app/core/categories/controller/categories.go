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
	adminGroup.GET("/lists", global.WrapWithBody(c.adminGetCategoriesLists))
	//adminGroup.PUT("/update",global.WrapWithBody())
}

func (c *CategoryHandler) adminCreateCategories(ctx *gin.Context, req vo.CategoryReq) (*global.ResponseBody[any], error) {
	if req.Name == "" {
		return global.ErrorResponse("名称不能为空"), nil
	}
	if req.Slug == "" {
		return global.ErrorResponse("别名不能为空"), nil
	}
	return global.SuccessResponse(), c.serv.AdminCreateCategory(&req)
}

func (c *CategoryHandler) adminGetCategoriesLists(ctx *gin.Context, req vo.CategoryListReq) (*global.ResponseBody[[]vo.MenuRes], error) {
	tree, err := c.serv.GetMenuTree(req.ParentId, req.IsTree)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(tree), nil
}
