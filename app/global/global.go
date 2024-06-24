package global

import (
	cache "github.com/chenmingyong0423/go-generics-cache"
	"github.com/dzsdbsdxq/dz-gin-blog/app/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	G_DZ_DB     *gorm.DB
	G_DZ_VP     *viper.Viper
	G_DZ_LOG    *zap.SugaredLogger
	G_DZ_CONFIG = new(config.Config)
	G_DZ_CACHE  *cache.Cache[string, any]
)

type G_DZ_MODEL struct {
	ID        uint64         `gorm:"primaryKey" json:"ID"` //主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
