package model

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type SysAttachMents struct {
	global.Model
	FileSize   uint   `json:"file_size" gorm:"comment:文件大小"`
	FileWidth  uint   `json:"file_width" gorm:"comment:图片宽度"`
	FileHeight uint   `json:"file_height" gorm:"comment:图片高度"`
	FileName   string `json:"file_name" gorm:"comment:文件名"`
	FileMd5    string `json:"file-md5" gorm:"comment:文件md5"`
	FilePath   string `json:"file_path" gorm:"comment:文件相对路径"`
	FileExt    string `json:"file_ext" gorm:"comment:文件扩展名"`
	FileType   string `json:"file_type" gorm:"comment:文件mine类型"`
	Position   string `json:"position" gorm:"comment:上传位置"`
}
