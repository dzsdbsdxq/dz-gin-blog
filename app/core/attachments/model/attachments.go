package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysAttachMents struct {
	global.Model
	FileSize   uint   `gorm:"column:file_size;type:bigint(20);default:0;not null;comment:文件大小"`
	FileWidth  uint   `gorm:"column:file_width;type:int(10);default:0;not null;comment:图片宽度"`
	FileHeight uint   `gorm:"column:file_height;type:int(10);default:0;not null;comment:图片高度"`
	FileName   string `gorm:"column:file_name;type:varchar(255);not null;default:'';comment:文件名"`
	FileMd5    string `gorm:"column:file_md5;type:char(32);not null;default:'';comment:文件md5"`
	FilePath   string `gorm:"column:file_path;type:varchar(255);not null;default:'';comment:文件相对路径"`
	FileExt    string `gorm:"column:file_ext;type:varchar(255);not null;default:'';comment:文件扩展名"`
	FileType   string `gorm:"column:file_type;type:varchar(191);not null;default:'';comment:文件mine类型"`
	Position   string `gorm:"column:position;type:varchar(10);not null;default:'local';comment:上传位置"`
}
