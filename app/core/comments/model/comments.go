package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysComments struct {
	global.Model
	UserType  int    `json:"user_type" gorm:"comment:用户类型（1=普通，2=博主）"`
	CommentId uint   `json:"comment_id" gorm:"comment:评论ID"`
	Status    int    `json:"status" gorm:"comment:状态（1=待审核，2=已发布，3=回收站）"`
	Type      int    `json:"type" gorm:"comment:类型（1=文章，2=页面）"`
	FromId    uint   `json:"from_id" gorm:"comment:文章ID或者页面ID"`
	NickName  string `json:"nick_name" gorm:"comment:昵称"`
	Email     string `json:"email" gorm:"comment:邮箱"`
	Link      string `json:"link" gorm:"comment:链接"`
	Body      string `json:"body" gorm:"comment:内容"`
	FromTitle string `json:"from_title" gorm:"comment:文章标题或页面标题"`
}
