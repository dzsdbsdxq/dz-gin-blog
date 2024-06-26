package model

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type SysUsers struct {
	global.Model
	Enable    int    `json:"enable" gorm:"comment:用户是否被冻结 1正常 2冻结"`
	UserName  string `json:"user_name" gorm:"comment:用户名"`
	Password  string `json:"password" gorm:"comment:密码"`
	Uuid      string `json:"uuid" gorm:"comment:用户UUID"`
	NickName  string `json:"nick_name" gorm:"comment:用户昵称"`
	HeaderImg string `json:"header_img" gorm:"comment:用户头像"`
	Phone     string `json:"phone" gorm:"comment:用户手机号"`
	Email     string `json:"email" gorm:"comment:用户邮箱"`
	WebLink   string `json:"web_link" gorm:"comment:主页链接"`
	Desc      string `json:"desc" gorm:"comment:用户描述"`
}
