package initialize

import (
	cache "github.com/chenmingyong0423/go-generics-cache"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"time"
)

func InitCache() {
	global.G_DZ_CACHE = cache.NewSimpleCache[string, any](10, time.Minute*10)
	global.G_DZ_LOG.Info("初始化Cache日志完成!")
}
