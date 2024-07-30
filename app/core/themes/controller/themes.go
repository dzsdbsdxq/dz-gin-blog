package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
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
	adminGroup := engine.Group("/admin-api/themes")
	//adminGroup.POST("/create", global.WrapWithBody(t.adminCreateTags))
	adminGroup.GET("/lists", global.Wrap(t.adminGetThemesLists))
	adminGroup.GET("/:name", global.Wrap(t.adminGetThemesByName))
	//adminGroup.PUT("/update/:id", global.WrapWithBody(t.adminUpdateTags))
	//adminGroup.DELETE("/delete/:id", global.Wrap(t.adminDeleteTags))
}

func (t *ThemesHandler) adminGetThemesLists(ctx *gin.Context) (*global.ResponseBody[*vo.ThemeSettingConfig], error) {
	return nil, nil
}
func (t *ThemesHandler) adminGetThemesByName(ctx *gin.Context) (*global.ResponseBody[*vo.ThemeSettingConfig], error) {
	themeName := ctx.Param("name")
	config, err := t.serv.GetThemeSettingConfig(themeName)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(config), nil
}
