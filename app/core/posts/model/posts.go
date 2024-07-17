package model

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type SysPosts struct {
	global.Model
	Title           string `gorm:"column:title;type:varchar(255);not null;default:'';comment:标题"`
	Body            string `gorm:"column:body;type:text;comment:内容"`
	Extend          string `gorm:"column:extend;type:text;comment:自定义扩展信息"`
	Thumb           string `gorm:"column:thumb;type:varchar(255);not null;default:'';comment:缩略图"`
	CommentAccepted int    `gorm:"column:comment_accepted;type:int(1);not null;default:1;comment:是否开启评论(1=开启,2=禁止评论)"`
	Recommend       int    `gorm:"column:recommend;type:int(1);not null;default:1;comment:是否推荐（1=否，2=推荐）"`
	Status          uint   `gorm:"column:status;type:int(3);not null;default:1;index:key_status;comment:状态（1=已发布，2=草稿，3=回收站）"`
	Top             uint   `gorm:"column:top;type:int(10);not null;default:0;comment:置顶"`
	SortBy          uint   `gorm:"column:sort_by;type:int(10);not null;default:0;comment:排序"`
	Flag            uint   `gorm:"column:flag;type:int(10);not null;default:0;comment:文章类型(1=图文，2=视频)"`
	Views           uint   `gorm:"column:views;type:int(10);not null;default:0;comment:查看数"`
	Likes           uint   `gorm:"column:likes;type:int(10);not null;default:0;comment:点赞数"`
	Password        string `gorm:"column:password;type:varchar(191);not null;default:'';comment:密码"`
	Slug            string `gorm:"column:slug;type:varchar(255);not null;default:'';index:key_slug;comment:别名(可设置成路径访问)"`
	Desc            string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述"`
	SeoKey          string `gorm:"column:seo_key;type:varchar(255);not null;default:'';comment:seo关键词"`
	SeoDesc         string `gorm:"column:seo_desc;type:varchar(255);not null;default:'';comment:seo描述"`
}

func (SysPosts) TableName() string {
	return "sys_posts"
}
