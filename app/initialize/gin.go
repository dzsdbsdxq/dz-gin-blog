package initialize

import (
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
)

func NewGinEngine(postHdl *posts.Handler, usersHdl *users.Handler, attachmentHdl *attachments.Handler, categoriesHdl *categories.Handler, tagsHdl *tags.Handler, commentsHdl *comments.Handler, logsHdl *logs.Handler, settingHdl *setting.Handler) (*gin.Engine, func(), error) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	//ginS.Static()
	engine.Static(global.G_DZ_CONFIG.Local.StorePath, global.G_DZ_CONFIG.Local.StorePath)

	//// 参数校验器注册
	//if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	for k, v := range validators {
	//		err := validate.RegisterValidation(k, v)
	//		if err != nil {
	//			return nil, err
	//		}
	//	}
	//}
	// 中间件注册
	//engine.Use(middleware...)
	postHdl.RegisterRoutes(engine)
	usersHdl.RegisterRoutes(engine)
	attachmentHdl.RegisterRoutes(engine)
	categoriesHdl.RegisterRoutes(engine)
	tagsHdl.RegisterRoutes(engine)
	commentsHdl.RegisterRoutes(engine)
	logsHdl.RegisterRoutes(engine)
	settingHdl.RegisterRoutes(engine)
	return engine, func() {
		fmt.Println("Hello")

	}, nil

}

func InitMiddlewares() {

}
