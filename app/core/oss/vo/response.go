package vo

type OssFile struct {
	FileSize   uint   `json:"file_size"`
	FileWidth  uint   `json:"file_width"`
	FileHeight uint   `json:"file_height"`
	FileName   string `json:"file_name"`
	FileMd5    string `json:"file-md5"`
	FilePath   string `json:"file_path"`
	FileExt    string `json:"file_ext"`
	FileType   string `json:"file_type"`
}
