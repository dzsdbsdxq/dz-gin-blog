package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysComments struct {
	global.Model
	UserType  int    `gorm:"column:user_type;type:int(1);not null;default:1;comment:用户类型（1=普通，2=博主）"`
	CommentId uint   `gorm:"column:comment_id;type:bigint(20);not null;default:0;index:comment_id;comment:评论ID"`
	Status    int    `gorm:"column:status;type:int(1);not null;default:1;comment:状态（1=待审核，2=已发布，3=回收站）"`
	Type      int    `gorm:"column:type;type:int(1);not null;default:1;comment:类型（1=文章，2=页面）"`
	FromId    uint   `gorm:"column:from_id;type:bigint(20);not null;default:0;comment:文章ID或者页面ID"`
	NickName  string `gorm:"column:nick_name;type:varchar(255);not null;default:'';comment:昵称"`
	Email     string `gorm:"column:email;type:varchar(255);not null;default:'';comment:邮箱"`
	Link      string `gorm:"column:link;type:varchar(255);not null;default:'';comment:链接"`
	Body      string `gorm:"column:body;type:text;comment:内容"`
	FromTitle string `gorm:"column:from_title;type:varchar(255);not null;default:'';comment:文章标题或页面标题"`
}
