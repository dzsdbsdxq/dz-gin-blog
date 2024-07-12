package vo

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type GetCommentsListReq struct {
	global.PageInfo
	Status uint `json:"status" form:"status"`
	Type   uint `json:"type" form:"type"`
	FromId uint `json:"from_id" form:"from_id"`
}
