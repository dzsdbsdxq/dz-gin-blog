package initialize

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users"
	"github.com/gin-gonic/gin"
)

func NewGinEngine(postHdl *posts.Handler, usersHdl *users.Handler) (*gin.Engine, error) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

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
	return engine, nil

}

func InitMiddlewares() {

}
