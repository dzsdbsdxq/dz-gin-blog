package controller

import (
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
)

type SettingHandler struct {
	serv service.ISettingService
}

func NewSettingHandler(serv service.ISettingService) *SettingHandler {
	return &SettingHandler{
		serv: serv,
	}
}

func (s *SettingHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/setting")
	adminGroup.POST("/update", global.Wrap(s.AdminUpdate))
}

func (s *SettingHandler) AdminUpdate(ctx *gin.Context) (*global.ResponseBody[any], error) {
	//读取请求体
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return global.ErrorResponse(405000000, ""), err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(ctx.Request.Body)
	var reqs []*vo.SettingReq
	var rawData map[string]string
	err = jsoniter.Unmarshal(body, &rawData)
	if err != nil {
		return global.ErrorResponse(405000000, ""), err
	}
	for key, value := range rawData {
		reqs = append(reqs, &vo.SettingReq{
			Key: key,
			Val: value,
		})
	}
	fmt.Printf("%+v\n", reqs)
	//s.serv.Update()
	return global.SuccessResponse(), s.serv.Update(reqs)
}
