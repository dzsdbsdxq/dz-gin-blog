package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"mime/multipart"
)

type Qiniu struct{}

func (*Qiniu) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	return nil, nil
}
func (*Qiniu) DeleteFile(key string) error {
	return nil
}
