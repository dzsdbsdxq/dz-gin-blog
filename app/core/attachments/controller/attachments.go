package controller

import (
	"errors"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

type AttachmentsHandler struct {
	serv service.IAttachmentsService
	oss  oss.Service
}

func NewAttachmentsHandler(serv service.IAttachmentsService, oss oss.Service) *AttachmentsHandler {
	return &AttachmentsHandler{
		serv: serv,
		oss:  oss,
		//cfgService:   cfgService,
		//postLikeServ: postLikeServ,
		//eventBus:     eventBus,
	}
}
func (a *AttachmentsHandler) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/attachment")
	adminGroup.POST("/upload", global.Wrap(a.AdminUploadFile))
	adminGroup.GET("/lists", global.WrapWithBody(a.AdminGetFile))
}

func (a *AttachmentsHandler) AdminUploadFile(ctx *gin.Context) (*global.ResponseBody[*vo.AttachmentsRes], error) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.G_DZ_LOG.Error("接收文件失败!", zap.Error(err))
		return nil, errors.New("接收文件失败")
	}
	at, uploadErr := a.oss.UploadFile(header)
	if uploadErr != nil {
		return nil, uploadErr
	}
	saveErr := a.serv.AdminCreateAttachments(&vo.AttachmentsReq{
		FileSize:   at.FileSize,
		FileWidth:  at.FileWidth,
		FileHeight: at.FileHeight,
		FileName:   at.FileName,
		FileMd5:    at.FileMd5,
		FilePath:   at.FilePath,
		FileExt:    at.FileExt,
		FileType:   at.FileType,
	})
	if saveErr != nil {
		_ = a.oss.DeleteFile(at.FilePath)
		return nil, saveErr
	}

	return global.SuccessResponseWithData(&vo.AttachmentsRes{
		Url:       filepath.Join(global.G_DZ_CONFIG.Local.BaseUrl, at.FilePath),
		Path:      at.FilePath,
		Driver:    "local",
		CreatedAt: time.Now().Format(time.DateTime),
		Name:      at.FileName,
		Key:       at.FileMd5,
	}), nil
}

func (a *AttachmentsHandler) AdminGetFile(ctx *gin.Context, req *vo.AttachmentsGetReq) (*global.ResponseBody[global.PageRes[[]*vo.AttachmentsRes]], error) {
	req.ValidateAndSetDefault()
	attachments, total, err := a.serv.AdminGetAttachments(req)
	if err != nil {
		return nil, err
	}
	return global.SuccessResponseWithData(global.PageRes[[]*vo.AttachmentsRes]{
		PageInfo: global.PageInfo{
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Category: req.Category,
			Keyword:  req.Keyword,
		},
		TotalCount: total,
		TotalPages: utils.CalcPages(total, req.PageSize),
		List:       attachments,
	}), nil
}
