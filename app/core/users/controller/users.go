package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	serv service.IUserService
}

func NewUserHandler(serv service.IUserService) *UserHandler {
	return &UserHandler{serv: serv}
}
func (h *UserHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/users")
	adminGroup.POST("/create", global.WrapWithBody(h.adminCreateUser))
	adminGroup.PUT("/update", global.WrapWithBody(h.adminUpdateUser))
	adminGroup.GET("/:id", global.Wrap(h.adminGetUserInfoById))
}

func (h *UserHandler) adminCreateUser(ctx *gin.Context, req vo.UserReq) (*global.ResponseBody[any], error) {
	return global.SuccessResponse(), h.serv.AdminCreateUser(&req)
}

func (h *UserHandler) adminUpdateUser(ctx *gin.Context, req vo.UserReq) (*global.ResponseBody[any], error) {
	return global.SuccessResponse(), h.serv.AdminCreateUser(&req)
}

func (h *UserHandler) adminGetUserInfoById(ctx *gin.Context) (*global.ResponseBody[vo.UserRes], error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return nil, err
	}
	user := h.serv.AdminGetUserInfoById(uint(id))
	if user == nil {
		return nil, nil
	}
	retUser := vo.UserRes{
		ID:        user.ID,
		UserName:  user.UserName,
		Password:  "",
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
		WebLink:   user.WebLink,
		Desc:      user.Desc,
	}
	return global.SuccessResponseWithData(retUser), nil
}
