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

type Model struct {
	ID        uint64         `gorm:"primaryKey" json:"ID"` //主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type PageInfo struct {
	PageNo   int64  `json:"pageNo" form:"pageNo"`     // 当前页
	PageSize int64  `json:"pageSize" form:"pageSize"` // 每页数量
	Category int64  `json:"category" form:"category"` //栏目ID
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
type PageRes[T any] struct {
	PageInfo
	// 总页数
	TotalPages int64 `json:"total"`
	// 总数量
	TotalCount int64 `json:"count"`
	List       T     `json:"list"`
}
