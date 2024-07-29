package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/service"
	"github.com/gin-gonic/gin"
)

type ThemesHandler struct {
	serv service.IThemesService
}

func NewThemesHandler(serv service.IThemesService) *ThemesHandler {
	return &ThemesHandler{
		serv: serv,
	}
}
func (t *ThemesHandler) RegisterRoutes(engine *gin.Engine) {
	//adminGroup := engine.Group("/admin-api/tags")
	//adminGroup.POST("/create", global.WrapWithBody(t.adminCreateTags))
	//adminGroup.GET("/lists", global.WrapWithBody(t.adminGetTagsLists))
	//adminGroup.PUT("/update/:id", global.WrapWithBody(t.adminUpdateTags))
	//adminGroup.DELETE("/delete/:id", global.Wrap(t.adminDeleteTags))
}
