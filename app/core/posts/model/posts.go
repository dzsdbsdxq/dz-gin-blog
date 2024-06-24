package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysPosts struct {
	global.G_DZ_MODEL
	Title           string `json:"title" gorm:"comment:标题"`
	Body            string `json:"body" gorm:"comment:内容"`
	Extend          string `json:"extend" gorm:"comment:自定义扩展信息"`
	Thumb           string `json:"thumb" gorm:"comment:缩略图"`
	CommentAccepted int    `json:"comment_accepted" gorm:"comment:是否开启评论(1=开启,2=禁止评论)"`
	Recommend       int    `json:"recommend" gorm:"comment:是否推荐（1=否，2=推荐）"`
	Status          uint   `json:"status" gorm:"comment:状态（1=已发布，2=草稿，3=回收站）"`
	Top             uint   `json:"top" gorm:"comment:置顶"`
	SortBy          uint   `json:"sortby" gorm:"comment:排序"`
	Flag            uint   `json:"flag" gorm:"comment:文章类型(1=图文，2=视频)"`
	Views           uint   `json:"views" gorm:"comment:查看数"`
	Likes           uint   `json:"likes" gorm:"comment:点赞数"`
	Password        string `json:"password" gorm:"comment:密码"`
	Slug            string `json:"slug" gorm:"comment:别名(可设置成路径访问)"`
	Desc            string `json:"desc" gorm:"comment:描述"`
	SeoKey          string `json:"seo_key" gorm:"comment:seo关键词"`
	SeoDesc         string `json:"seo_desc" gorm:"comment:seo描述"`
}

func (SysPosts) TableName() string {
	return "sys_posts"
}
