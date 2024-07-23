package vo

import "github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/model"

type StatVO struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func InitializeData() []*model.SysStat {
	return []*model.SysStat{
		{
			Key: "SiteView",
			Val: 0,
		},
	}
}
