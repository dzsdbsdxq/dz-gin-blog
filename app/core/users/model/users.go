package model

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type SysUsers struct {
	global.Model
	Enable    int    `gorm:"column:enable;type:int(1);not null;default:1;comment:用户是否被冻结 1正常 2冻结"`
	UserName  string `gorm:"column:user_name;type:varchar(255);not null;default:'';comment:用户名"`
	Password  string `gorm:"column:password;type:varchar(255);not null;default:'';comment:密码"`
	Uuid      string `gorm:"column:uuid;type:varchar(255);not null;default:'';comment:用户UUID"`
	NickName  string `gorm:"column:nick_name;type:varchar(255);not null;default:'';comment:用户昵称"`
	HeaderImg string `gorm:"column:header_img;type:varchar(255);not null;default:'';comment:用户头像"`
	Phone     string `gorm:"column:phone;type:char(11);not null;default:'';comment:用户手机号"`
	Email     string `gorm:"column:email;type:varchar(255);not null;default:'';comment:用户邮箱"`
	WebLink   string `gorm:"column:web_link;type:varchar(255);not null;default:'';comment:主页链接"`
	Desc      string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:用户描述"`
}
