package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/attachments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/categories"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/logs"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users"
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
		logs.LogModel{},
		post_category.PostCategoryModel{},
		posts.PostModel{},
		posts_tags.PostTagModel{},
		setting.SettingsModel{},
		tags.TagModel{},
		themes.ThemeModel{},
		stat.StatsModel{},
		users.UserModel{},
	)
	if err != nil {
		global.G_DZ_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.G_DZ_LOG.Info("register table success")
	return err
}

func NewInstallService(db *gorm.DB) *InstallService {
	return &InstallService{coll: db}
}
