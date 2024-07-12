package controller

import (
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	"github.com/gin-gonic/gin"
)

type CommentsHandler struct {
	serv        service.ICommentService
	postService posts.Service
}

func NewCommentsHandler(serv service.ICommentService, postService posts.Service) *CommentsHandler {
	return &CommentsHandler{serv: serv, postService: postService}
}
func (c *CommentsHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/comments")
	adminGroup.POST("/reply", global.WrapWithBody(c.adminCreateCategories))
	//adminGroup.GET("/lists", global.WrapWithBody(c.adminGetCategoriesLists))
	//adminGroup.PUT("/update/:id", global.WrapWithBody(c.adminUpdateCategories))
	//adminGroup.DELETE("/delete/:id", global.Wrap(c.adminDeleteCategories))

	publicGroup := engine.Group("/api/comments")
	publicGroup.POST("/create", global.WrapWithBody(c.CreateComments))
}
func (c *CommentsHandler) adminCreateCategories(ctx *gin.Context, req vo.CommentsReq) (*global.ResponseBody[any], error) {
	return nil, nil
}
func (c *CommentsHandler) CreateComments(ctx *gin.Context, req vo.CommentsReq) (*global.ResponseBody[any], error) {
	if req.NickName == "" {
		return global.ErrorResponse(409000000, ""), nil
	}
	if utils.EmailVerify(req.Email) {
		return global.ErrorResponse(409000001, ""), nil
	}
	//查询文章信息
	post, err := c.postService.FindPostByPostId(req.FromId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", post)
	return global.SuccessResponse(), c.serv.AddComment(req)
}
