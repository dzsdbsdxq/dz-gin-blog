package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type IInstallService interface {
	RegisterTables() error
}

var _ IInstallService = (*InstallService)(nil)

type InstallService struct {
	coll *gorm.DB
}

func (i *InstallService) RegisterTables() error {
	err := i.coll.AutoMigrate(
		attachments.AttachmentModel{},
		categories.CategoryModel{},
		comments.CommentModel{},

	)
	if err != nil {
		global.G_DZ_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.G_DZ_LOG.Info("register table success")
	return err
}

func NewInstallService() *InstallService {
	return &InstallService{coll: global.G_DZ_DB}
}
