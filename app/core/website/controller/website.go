package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat"
	"github.com/gin-gonic/gin"
)

type WebSiteHandler struct {
	statService stat.Service
}

func NewWebsiteHandler(statService stat.Service) *WebSiteHandler {
	return &WebSiteHandler{
		statService: statService,
	}
}
func (web *WebSiteHandler) RegisterRoutes(engine *gin.Engine) {
	//adminGroup := engine.Group("/admin-api/tags")
	//adminGroup.POST("/create", global.WrapWithBody(t.adminCreateTags))
	//adminGroup.GET("/lists", global.WrapWithBody(t.adminGetTagsLists))
	//adminGroup.PUT("/update/:id", global.WrapWithBody(t.adminUpdateTags))
	//adminGroup.DELETE("/delete/:id", global.Wrap(t.adminDeleteTags))
}
