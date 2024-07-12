package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
	"sync"
)

type PostHandler struct {
	serv service.IPostService
	//cfgService   website_config.Service
	//postLikeServ post_like.Service
	ipMap sync.Map
	//eventBus     *eventbus.EventBus
}

func NewPostHandler(serv service.IPostService) *PostHandler {
	return &PostHandler{
		serv: serv,
		//cfgService:   cfgService,
		//postLikeServ: postLikeServ,
		//eventBus:     eventBus,
	}
}

func (h *PostHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/posts")
	adminGroup.POST("/create", global.WrapWithBody(h.adminCreatePost))
}

func (h *PostHandler) adminCreatePost(ctx *gin.Context, req vo.PostReq) (*global.ResponseBody[any], error) {
	return global.SuccessResponse(), h.serv.CreatePost(&req)
}
