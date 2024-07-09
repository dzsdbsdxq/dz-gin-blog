package vo

import "github.com/dzsdbsdxq/dz-gin-blog/app/global"

type AttachmentsReq struct {
	FileSize   uint   `json:"file_size"`
	FileWidth  uint   `json:"file_width"`
	FileHeight uint   `json:"file_height"`
	FileName   string `json:"file_name"`
	FileMd5    string `json:"file-md5"`
	FilePath   string `json:"file_path"`
	FileExt    string `json:"file_ext"`
	FileType   string `json:"file_type"`
}

type AttachmentsGetReq struct {
	global.PageInfo
	FileType string `json:"file_type" form:"file_type"`
}
