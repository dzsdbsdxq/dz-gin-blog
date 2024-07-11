package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TagsHandler struct {
	serv service.ITagService
}

func NewTagsHandler(serv service.ITagService) *TagsHandler {
	return &TagsHandler{
		serv: serv,
		//cfgService:   cfgService,
		//postLikeServ: postLikeServ,
		//eventBus:     eventBus,
	}
}
func (t *TagsHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/tags")
	adminGroup.POST("/create", global.WrapWithBody(t.adminCreateTags))
	adminGroup.GET("/lists", global.WrapWithBody(t.adminGetTagsLists))
	adminGroup.PUT("/update/:id", global.WrapWithBody(t.adminUpdateTags))
	adminGroup.DELETE("/delete/:id", global.Wrap(t.adminDeleteTags))
}
func (t *TagsHandler) adminCreateTags(ctx *gin.Context, req vo.TagsReq) (*global.ResponseBody[any], error) {
	if req.Name == "" {
		return global.ErrorResponse(403000000, ""), nil
	}
	if req.Slug == "" {
		return global.ErrorResponse(403000002, ""), nil
	}
	return global.SuccessResponse(), t.serv.AdminCreateTag(&req)
}
func (t *TagsHandler) adminGetTagsLists(ctx *gin.Context, req global.PageInfo) (*global.ResponseBody[global.PageRes[[]*vo.TagsRes]], error) {
	tags, total, err := t.serv.SelectTags(&req)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(global.PageRes[[]*vo.TagsRes]{
		PageInfo: global.PageInfo{
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Category: req.Category,
			Keyword:  req.Keyword,
		},
		TotalCount: total,
		TotalPages: utils.CalcPages(total, req.PageSize),
		List:       tags,
	}), nil
}

func (t *TagsHandler) adminUpdateTags(ctx *gin.Context, req vo.TagsReq) (*global.ResponseBody[any], error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	return global.SuccessResponse(), t.serv.AdminUpdateTag(id, &req)

}
func (t *TagsHandler) adminDeleteTags(ctx *gin.Context) (*global.ResponseBody[any], error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	return global.SuccessResponse(), t.serv.AdminDeleteTag(uint(id))
}
