package controller

import (
	"errors"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

type AttachmentsHandle struct {
	serv service.IAttachmentsService
	oss  oss.Service
}

func NewAttachmentsHandler(serv service.IAttachmentsService, oss oss.Service) *AttachmentsHandle {
	return &AttachmentsHandle{
		serv: serv,
		oss:  oss,
		//cfgService:   cfgService,
		//postLikeServ: postLikeServ,
		//eventBus:     eventBus,
	}
}
func (a *AttachmentsHandle) RegisterRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/attachment")
	adminGroup.POST("/upload", global.Wrap(a.AdminUploadFile))
	adminGroup.GET("/lists", global.Wrap(a.AdminGetFile))
}

func (a *AttachmentsHandle) AdminUploadFile(ctx *gin.Context) (*global.ResponseBody[*vo.AttachmentsUploadSuccessResponse], error) {
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

	return global.SuccessResponseWithData(&vo.AttachmentsUploadSuccessResponse{
		Url:       filepath.Join(global.G_DZ_CONFIG.System.Domain, at.FilePath),
		Path:      at.FilePath,
		Driver:    "local",
		CreatedAt: time.Now(),
		Name:      at.FileName,
		Key:       at.FileMd5,
	}), nil
}

func (a *AttachmentsHandle) AdminGetFile(ctx *gin.Context) {

}
