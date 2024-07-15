package controller

import (
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	"github.com/gin-gonic/gin"
	"strconv"
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
	adminGroup.POST("/reply", global.WrapWithBody(c.adminCreateComments))
	adminGroup.GET("/lists", global.WrapWithBody(c.adminGetCommentsLists))
	adminGroup.PUT("/update/:id", global.WrapWithBody(c.adminUpdateComments))
	adminGroup.DELETE("/delete", global.WrapWithBody(c.adminDeleteComments))

	publicGroup := engine.Group("/api/comments")
	publicGroup.POST("/create", global.WrapWithBody(c.CreateComments))
}
func (c *CommentsHandler) adminCreateComments(ctx *gin.Context, req vo.CommentsReq) (*global.ResponseBody[any], error) {
	return global.SuccessResponse(), c.serv.AddComment(req)
}
func (c *CommentsHandler) CreateComments(ctx *gin.Context, req vo.CommentsReq) (*global.ResponseBody[any], error) {
	if req.NickName == "" {
		return global.ErrorResponse(409000000, ""), nil
	}
	if !utils.EmailVerify(req.Email) {
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
func (c *CommentsHandler) adminUpdateComments(ctx *gin.Context, req vo.CommentsReq) (*global.ResponseBody[any], error) {
	check := ctx.Query("ck")
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	//查询是否存在
	comment, _ := c.serv.FindCommentById(id)
	if comment == nil {
		return global.ErrorResponse(409000003, ""), nil
	}

	if check == "update" {
		//只通过
		return global.SuccessResponse(), c.serv.UpdateComment(id, req)
	} else if check == "update-reply" {
		//通过并回复
		_ = c.serv.UpdateComment(id, req)
		return global.SuccessResponse(), c.serv.AddComment(vo.CommentsReq{
			UserType:  2, //2 = 博主
			CommentId: req.CommentId,
			Status:    2, //博主回复直接发布
			Type:      comment.Type,
			FromId:    comment.FromId,
			NickName:  "", // TODO :这里需要登录的用户信息
			Email:     "", // TODO :这里需要登录的用户信息
			Link:      "", // TODO :这里需要登录的用户信息
			Body:      req.Body,
			FromTitle: comment.FromTitle,
		})
	}
	return global.ErrorResponse(400001000, ""), nil
}
func (c *CommentsHandler) adminDeleteComments(ctx *gin.Context, ids global.IdsReq) (*global.ResponseBody[any], error) {
	return global.SuccessResponse(), c.serv.DeleteComment(ids)
}
func (c *CommentsHandler) adminGetCommentsLists(ctx *gin.Context, req vo.GetCommentsListReq) (*global.ResponseBody[global.PageRes[[]*vo.CommentsRes]], error) {
	comments, total, err := c.serv.SelectComments(req)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(global.PageRes[[]*vo.CommentsRes]{
		PageInfo: global.PageInfo{
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Category: req.Category,
			Keyword:  req.Keyword,
		},
		TotalCount: total,
		TotalPages: utils.CalcPages(total, req.PageSize),
		List:       comments,
	}), nil
}
