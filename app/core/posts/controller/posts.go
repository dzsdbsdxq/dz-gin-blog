package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/service"
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

func (h *PostHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/posts")
	adminGroup.POST("/create", global.WrapWithBody(h.AdminCreatePost))
}

func (h *PostHandler) AdminCreatePost(ctx *gin.Context) (*global.ResponseBody[any], error) {
	return global.SuccessResponse(), h.serv.AdminCreatePost()
}
