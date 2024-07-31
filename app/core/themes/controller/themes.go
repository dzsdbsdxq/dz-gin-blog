package controller

import (
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type ThemesHandler struct {
	serv        service.IThemesService
	settingServ setting.Service
}

func NewThemesHandler(serv service.IThemesService, settingServ setting.Service) *ThemesHandler {
	return &ThemesHandler{
		serv:        serv,
		settingServ: settingServ,
	}
}
func (t *ThemesHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/themes")
	//安装主题
	adminGroup.POST("/:name/install", global.Wrap(t.adminInstallThemes))
	//启用主题
	adminGroup.PUT("/:name/enable", global.Wrap(t.adminEnableThemes))
	//删除主题
	adminGroup.DELETE("/:name/delete", global.Wrap(t.adminDeleteThemes))
	//获取主题列表
	adminGroup.GET("/lists", global.Wrap(t.adminGetThemesLists))
	//获取数据看主题配置
	adminGroup.GET("/:name/settings", global.Wrap(t.adminGetThemeSetting))
	//获取主题配置文件的信息
	adminGroup.GET("/:name/configurations", global.Wrap(t.adminGetThemesByName))
	//adminGroup.PUT("/update/:id", global.WrapWithBody(t.adminUpdateTags))
	//adminGroup.DELETE("/delete/:id", global.Wrap(t.adminDeleteTags))
}

func (t *ThemesHandler) adminGetThemesLists(ctx *gin.Context) (*global.ResponseBody[*vo.ThemeSettingConfig], error) {
	return nil, nil
}
func (t *ThemesHandler) adminInstallThemes(ctx *gin.Context) (*global.ResponseBody[any], error) {
	//上传安装包，或在线安装 解压到主题文件夹
	//
	//判断主题是否存在
	//t.settingServ.FindByKey(themeName)
	return global.SuccessResponse(), nil
}
func (t *ThemesHandler) adminEnableThemes(ctx *gin.Context) (*global.ResponseBody[any], error) {
	themeName := ctx.Param("name")
	//操作，enable 启用，disable
	option := ctx.Query("type")
	switch option {
	case "enable":
		file, err := utils.ReadConfigFile(themeName)
		if err != nil {
			return global.ErrorResponse(0, "主题不存在"), err
		}
		var config = new(vo.ThemeSettingConfig)
		//解析JSON
		err = jsoniter.Unmarshal(file, &config)
		if err != nil {
			return global.ErrorResponse(0, "主题配置格式错误"), err
		}
		conf := make([]*vo.ThemeReq, 0)
		//读取配置信息
		fmt.Printf("%+v\n", config)
		if len(config.Pages) > 0 {
			for _, page := range config.Pages {
				conf = append(conf, &vo.ThemeReq{
					Value: page.File,
					Key:   page.Name,
					Type:  themeName,
					Desc:  page.Description,
					Hook:  "pages",
				})
			}
		}
		if len(config.Posts) > 0 {
			for _, post := range config.Posts {
				conf = append(conf, &vo.ThemeReq{
					Value: post.Value,
					Key:   post.Name,
					Type:  themeName,
					Desc:  post.Label,
					Hook:  "posts",
				})
				if len(post.Group) > 0 {

				}
			}

		}

		_ = t.settingServ.UpdateValByKey("BlogEnableTheme", themeName)
	case "disable":
		key, _ := t.settingServ.FindByKey("BlogEnableTheme")
		if key.Val == themeName {
			_ = t.settingServ.UpdateValByKey("BlogEnableTheme", "")
		}
	default:
		return global.ErrorResponse(0, "无效操作"), nil
	}

	return global.SuccessResponse(), nil
}
func (t *ThemesHandler) adminDeleteThemes(ctx *gin.Context) (*global.ResponseBody[any], error) {
	return nil, nil
}

func (t *ThemesHandler) adminGetThemeSetting(ctx *gin.Context) (*global.ResponseBody[[]*vo.ThemeKeyValue], error) {
	themeName := ctx.Param("name")
	themes, err := t.serv.GetThemeSetting(themeName)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(themes), nil
}

func (t *ThemesHandler) adminGetThemesByName(ctx *gin.Context) (*global.ResponseBody[*vo.ThemeSettingConfig], error) {
	themeName := ctx.Param("name")
	config, err := t.serv.GetThemeSettingConfig(themeName)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(config), nil
}
