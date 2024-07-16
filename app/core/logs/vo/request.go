package vo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type LogsGetReq struct {
	global.PageInfo
	Type string `json:"type" form:"type"`
}
type LogsReq struct {
	Type    string `json:"type"`
	Key     string `json:"key"`
	Content string `json:"content"`
	Ip      string `json:"ip"`
	Creator string `json:"creator"`
}
