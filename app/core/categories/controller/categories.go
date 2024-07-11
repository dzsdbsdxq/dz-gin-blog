package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
	"strconv"
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
	adminGroup.PUT("/update/:id", global.WrapWithBody(c.adminUpdateCategories))
	adminGroup.DELETE("/delete/:id", global.Wrap(c.adminDeleteCategories))
}

func (c *CategoryHandler) adminCreateCategories(ctx *gin.Context, req vo.CategoryReq) (*global.ResponseBody[any], error) {
	if req.Name == "" {
		return global.ErrorResponse(402000002, ""), nil
	}
	if req.Slug == "" {
		return global.ErrorResponse(402000003, ""), nil
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

func (c *CategoryHandler) adminUpdateCategories(ctx *gin.Context, req vo.CategoryReq) (*global.ResponseBody[any], error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	return global.SuccessResponse(), c.serv.AdminUpdateCategory(id, &req)

}
func (c *CategoryHandler) adminDeleteCategories(ctx *gin.Context) (*global.ResponseBody[any], error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	return global.SuccessResponse(), c.serv.AdminDeleteCategory(uint(id))
}
