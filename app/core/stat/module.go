package stat

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/service"
)

type (
	Service    = service.IStatService
	StatsModel = model.SysStat
	Module     struct {
		Svc Service
	}
)
